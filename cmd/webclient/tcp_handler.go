package main

import (
	"crypto"
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/spf13/viper"
	"software.sslmate.com/src/go-pkcs12"
)

func (app *App) connect() (net.Conn, error) {
	addr := fmt.Sprintf("%s:%s", app.host, app.tcpPort)
	if app.tls {
		app.Logger.Infof("connecting with SSL to %s...", addr)
		conn, err := tls.Dial("tcp", addr, app.getTlsConfig())
		if err != nil {
			return nil, err
		}
		app.Logger.Debugf("handshake...")

		if err := conn.Handshake(); err != nil {
			return conn, err
		}
		cs := conn.ConnectionState()

		app.Logger.Infof("Handshake complete: %t", cs.HandshakeComplete)
		app.Logger.Infof("version: %d", cs.Version)
		for i, cert := range cs.PeerCertificates {
			app.Logger.Infof("cert #%d subject: %s", i, cert.Subject.String())
			app.Logger.Infof("cert #%d issuer: %s", i, cert.Issuer.String())
			app.Logger.Infof("cert #%d dns_names: %s", i, strings.Join(cert.DNSNames, ","))
		}
		return conn, nil
	} else {
		app.Logger.Infof("connecting to %s...", addr)
		return net.DialTimeout("tcp", addr, app.dialTimeout)
	}
}

func (app *App) getTlsConfig() *tls.Config {
	p12Data, err := os.ReadFile(viper.GetString("ssl.cert"))
	if err != nil {
		app.Logger.Fatal(err)
	}

	key, cert, _, err := pkcs12.DecodeChain(p12Data, viper.GetString("ssl.password"))
	if err != nil {
		app.Logger.Fatal(err)
	}

	tlsCert := tls.Certificate{
		Certificate: [][]byte{cert.Raw},
		PrivateKey:  key.(crypto.PrivateKey),
		Leaf:        cert,
	}

	return &tls.Config{Certificates: []tls.Certificate{tlsCert}, InsecureSkipVerify: true}
}
