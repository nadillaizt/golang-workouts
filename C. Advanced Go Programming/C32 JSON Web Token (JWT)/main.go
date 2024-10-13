package main

import (
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "path/filepath"
    "strings"

    jwt "github.com/golang-jwt/jwt/v4"
    gubrak "github.com/novalagung/gubrak/v2"
)

type M map[string]interface{}

var APPLICATION_NAME = "My Simple JWT App"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("the secret of kalimdor")

func main() {
    mux := new(CustomMux)
    mux.RegisterMiddleware(MiddlewareJWTAuthorization)

    mux.HandleFunc("/login", HandlerLogin)
    mux.HandleFunc("/index", HandlerIndex)

    server := new(http.Server)
    server.Handler = mux
    server.Addr = ":8080"

    fmt.Println("Starting server at", server.Addr)
    server.ListenAndServe()
}

func HandlerLogin(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Unsupported http method", http.StatusBadRequest)
        return
    }

    username, password, ok := r.BasicAuth()
    if !ok {
        http.Error(w, "Invalid username or password", http.StatusBadRequest)
        return
    }

	type MyClaims struct {
		jwt.StandardClaims
		Username string `json:"Username"`
		Email    string `json:"Email"`
		Group    string `json:"Group"`
	}

	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
		Username: userInfo["username"].(string),
		Email:    userInfo["email"].(string),
		Group:    userInfo["group"].(string),
	}
}

