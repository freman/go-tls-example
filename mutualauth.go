package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	peerCert := r.TLS.PeerCertificates[0]

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w,
		"Issued by %s\nIssued to %s\n",
		peerCert.Issuer.CommonName,
		peerCert.Subject.CommonName,
	)
}

func main() {
	// CA Certificate that is signing your client certificates
	caCert, err := ioutil.ReadFile("ssl/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		ClientCAs:  caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}

	tlsConfig.BuildNameToCertificate()

	http.HandleFunc("/", hello)

	server := &http.Server{
		Addr:      ":8443",
		TLSConfig: tlsConfig,
	}

	// Certificate and key from any CA that the client knows about (could be yours,
	// could be letsencrypt, could be anyone)
	if err := server.ListenAndServeTLS("ssl/bob.crt", "ssl/bob.key"); err != nil {
		log.Fatal(err)
	}
}
