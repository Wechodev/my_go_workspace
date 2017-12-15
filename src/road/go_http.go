package main

import (
	"net"
	"net/http"
	"time"
	"fmt"
	"crypto/x509"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"encoding/pem"
	"errors"
	"io/ioutil"
)

const SERVER_PORT = 8081
const SERVER_DOMIAN = "localhost"
const RESPONSE_TEMPLATE = "hello"

func rootHandler(w http.ResponseWriter, req *http.Request)  {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Content-Length", fmt.Sprint(len(RESPONSE_TEMPLATE)))
	w.Write([]byte(RESPONSE_TEMPLATE))
}

func YourListenAndServeTLS(addr string, certFile string, keyFile string, handler http.Handler) error {
	config := &tls.Config{
		Rand: rand.Reader,
		Time: time.Now,
		NextProtos: []string{"http/1.1"},
	}

	var err error
	config.Certificates = make([]tls.Certificate, 1)
	config.Certificates[0], err = YourLoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return  err
	}

	conn, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	tlsListener := tls.NewListener(conn, config)
	return http.Serve(tlsListener, handler)
}

func YourLoadX509KeyPair(certFile string, keyFile string) (cert tls.Certificate, err error) {
	certPEMBlock, err := ioutil.ReadFile(certFile)
	if err != nil {
		return
	}

	certDERBlock, restPEMBlock := pem.Decode(certPEMBlock)
	if certPEMBlock == nil {
		err = errors.New("crypto/tls: failed to parse certificate PEM data")
		return
	}

	certDERBlockChain, _ := pem.Decode(restPEMBlock)
	if certDERBlockChain == nil {
		cert.Certificate = [][]byte{certDERBlock.Bytes}
	} else {
		cert.Certificate = [][]byte{certDERBlock.Bytes, certDERBlockChain.Bytes}
	}

	keyPEMBlock, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return
	}

	keyDERBlock, _ := pem.Decode(keyPEMBlock)
	if keyDERBlock == nil {
		err = errors.New("crypto/tls:failed to parse key PEM data")
		return
	}

	key, err := x509.ParsePKCS1PrivateKey(keyDERBlock.Bytes)
	if err != nil {
		err = errors.New("crypto/tls: failed to parse key")
		return
	}

	cert.PrivateKey = key

	x509Cert, err := x509.ParseCertificate(certDERBlock.Bytes)
	if err != nil {
		return
	}

	if x509Cert.PublicKeyAlgorithm != x509.RSA ||
		x509Cert.PublicKey.(*rsa.PublicKey).N.Cmp(key.PublicKey.N) != 0{
		err = errors.New("crypto/tls: private key does not match public key")
		return
	}

	return
}

func main()  {
	http.HandleFunc(fmt.Sprintf("%s:%d/", SERVER_DOMIAN, SERVER_PORT), rootHandler)
	YourListenAndServeTLS(fmt.Sprintf(":%d", SERVER_PORT), "cert.pem", "key.pem", nil)
}

/*
在本例子中，用到了crypto中的一些包
rand,伪随机函数发生器，用于产生基于时间和CPU时钟的伪随机数;rsa,非对称加密算法，rsa是三个发明者首字母拼接而成的，tls是传输层安全协议，
x509一种常见的数字证书格式;pem在非对称体系下，一般存放公钥和私钥的文件

*/