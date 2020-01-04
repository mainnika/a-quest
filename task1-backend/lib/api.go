package lib

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	routing "github.com/jackwhelpton/fasthttp-routing/v2"
	"github.com/jackwhelpton/fasthttp-routing/v2/access"
	"github.com/jackwhelpton/fasthttp-routing/v2/cors"
	"github.com/jackwhelpton/fasthttp-routing/v2/fault"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

const (
	URL        = "/check"
	URLHealthz = "/healthz"
)

type Api struct {
	Base   string
	Alg    string
	Pub    *ecdsa.PublicKey
	Priv   *ecdsa.PrivateKey
	Answer []byte
}

type Answer struct {
	Answer json.RawMessage `json:"answer"`
	Name   string          `json:"name"`
}

func (a *Api) GetHandler() fasthttp.RequestHandler {

	crs := cors.Options{
		AllowOrigins:     "*",
		AllowHeaders:     "*",
		AllowMethods:     "*",
		AllowCredentials: true,
	}

	router := routing.New()
	router.Use(
		access.Logger(log.Debugf),
		cors.Handler(crs),
		fault.PanicHandler(log.Warnf),
	)

	base := strings.TrimSuffix(a.Base, "/")
	api := router.Group(base)

	api.Post(URL, a.checkAnswer)
	api.Get(URLHealthz, a.healthCheck)

	return router.HandleRequest
}

func (a *Api) healthCheck(ctx *routing.Context) (err error) {
	ctx.Response.Header.Set("Content-Type", "application/json")
	_, err = ctx.WriteString("{\"health\":\"ok\"}")
	if err != nil {
		return routing.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return
}

func (a *Api) checkAnswer(ctx *routing.Context) (err error) {

	now := time.Now()
	id := uuid.New()
	answer := &Answer{}

	err = json.Unmarshal(ctx.Request.Body(), answer)
	if err != nil {
		return routing.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	hash := sha256.New()
	hash.Write(answer.Answer)

	hexed := make([]byte, hex.EncodedLen(sha256.Size))

	hex.Encode(hexed, hash.Sum(nil))

	log.Debugf("req answer, %v", answer)
	log.Debugf("req hashed, %s", hexed)

	valid := subtle.ConstantTimeCompare(a.Answer, hexed)
	if valid == 0 {
		return routing.NewHTTPError(http.StatusForbidden)
	}

	claims := &jwtgo.StandardClaims{
		Id:       id.String(),
		Subject:  answer.Name,
		Issuer:   "Anonymous Moroz Grandfather",
		IssuedAt: now.Unix(),
	}

	method := jwtgo.GetSigningMethod(a.Alg)
	token := jwtgo.NewWithClaims(method, claims)

	jwt, err := token.SignedString(a.Priv)
	if err != nil {
		return routing.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	ctx.Response.Header.Set("Content-Type", "application/json")
	_, err = ctx.WriteString(fmt.Sprintf(`
['переход на следующий уровень ищи в альтер имени безопасности узла',
'используй ключ чтобы открыть следующий уровень',
'%s']
`, jwt))
	if err != nil {
		return routing.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return
}

func (a *Api) getPrivKey(t *jwtgo.Token) (interface{}, error) {
	return a.Priv, nil
}

func (a *Api) getKey(t *jwtgo.Token) (interface{}, error) {
	return a.Pub, nil
}