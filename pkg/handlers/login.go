package handlers

import (
	"bytes"
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-openapi/runtime/middleware"
	"github.com/martinlindhe/gogost/gost34112012512"
	"time"
	"univers/pkg/config"
	"univers/pkg/httperrors"
	"univers/pkg/models"
	"univers/pkg/restapi/operations/login"
	"univers/pkg/store"
)

var loginHandler *LoginHandler

type LoginHandler struct {
	cfg *config.Config
	st *store.Store
}

func GetLoginHandler() *LoginHandler {
	if nil == loginHandler {
		loginHandler = &LoginHandler{}
		loginHandler.init()
	}

	return loginHandler
}

func (h *LoginHandler) init() {
	h.cfg = config.GetConfig()
	h.st = store.GetStore()
}

func (h *LoginHandler) Login(params login.PostLoginParams) middleware.Responder {
	hash := gost34112012512.New()
	_, _ = hash.Write([]byte(params.Body.APIKey))
	apiKey, _ := base64.StdEncoding.DecodeString(h.cfg.ApiKey)
	if 0 != bytes.Compare(hash.Sum(nil), apiKey) {
		h.st.IncrLimit(params.HTTPRequest.RemoteAddr)
		return login.NewPostLoginUnauthorized().WithPayload(httperrors.WrongApiKeyError)
	}

	return login.NewPostLoginOK().WithPayload(&models.PostLoginOKBody{Token: genToken(h.cfg.Secret)})
}

func genToken(secret string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
		"iss": "univer",
	})
	tokenString, _ := token.SignedString([]byte(secret))
	return tokenString
}