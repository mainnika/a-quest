package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	jwtgo "github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"

	"github.com/mainnika/a-quest/keys"
	"github.com/mainnika/a-quest/task1-backend/lib"
	. "github.com/mainnika/a-quest/task1-backend/lib/configure"
	. "github.com/mainnika/a-quest/task1-backend/lib/env"
)

var version = "dev"

var (
	getVersion bool
)

func init() {

	if IsDevelopment {
		log.SetLevel(log.DebugLevel)
		log.Debug("debug mode")
	}
}

func httpStart(apiserv *lib.Api) (httpserv *fasthttp.Server) {

	httpserv = &fasthttp.Server{
		Logger:           log.StandardLogger(),
		Handler:          apiserv.GetHandler(),
		DisableKeepalive: true,
	}

	lis, err := net.Listen("tcp", Config.HttpAPI.Addr)
	if err != nil {
		log.Fatalf("http listen error: %v", err)
	}

	err = httpserv.Serve(lis)
	if err != nil {
		log.Fatalf("http serve error: %v", err)
	}

	return
}

func main() {

	flag.BoolVar(&getVersion, "v", false, "version")
	flag.Parse()

	if getVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	log.Debugf("version: %s", version)
	log.Debugf("cfg: %v", Config)

	pubKey, err := jwtgo.ParseECPublicKeyFromPEM(keys.PublicKey)
	if err != nil {
		log.Fatalf("can not parse jwt key: %s", err)
	}

	privKey, err := jwtgo.ParseECPrivateKeyFromPEM(keys.PrivateKey)
	if err != nil {
		log.Fatalf("can not parse jwt key: %s", err)
	}

	apiserv := &lib.Api{
		Base:   Config.HttpAPI.Base,
		Pub:    pubKey,
		Priv:   privKey,
		Alg:    keys.Alg,
		Answer: keys.Answer,
	}

	httpStart(apiserv)
}
