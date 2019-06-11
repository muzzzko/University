package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"univers/pkg/config"
	"univers/pkg/handlers"
	"univers/pkg/httperrors"
	"univers/pkg/restapi/operations/login"
	"univers/pkg/restapi/operations/region"
	"univers/pkg/restapi/operations/university"
	"univers/pkg/store"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"univers/pkg/restapi/operations"
)

func configureFlags(api *operations.UniversAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.UniversAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	universityHandler := handlers.GetUniversityHandler()
	loginHandler := handlers.GetLoginHandler()
	regionHandler := handlers.GetRegionHandler()

	cfg := config.GetConfig()
	api.JWTAuth = func(tokenString string) (i interface{}, e error) {
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(cfg.Secret), nil
		})

		if nil != err {
			return nil, errors.New(401, "Unauthorized")
		}

		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			return token, nil
		} else {
			return nil, errors.New(401, "Unauthorized")
		}
	}
	
	if api.GetPingHandler == nil {
		api.GetPingHandler = operations.GetPingHandlerFunc(func(params operations.GetPingParams) middleware.Responder {
			return middleware.NotImplemented("operation .GetPing has not yet been implemented")
		})
	}
	api.UniversityGetUniversityHandler = university.GetUniversityHandlerFunc(universityHandler.GetUniversity)
	api.UniversityGetUniversitiesHandler = university.GetUniversitiesHandlerFunc(universityHandler.GetUniversities)
	api.UniversityPostUniversityHandler = university.PostUniversityHandlerFunc(universityHandler.PostUniversity)
	api.UniversityPatchUniversityHandler = university.PatchUniversityHandlerFunc(universityHandler.PatchUniversity)
	api.LoginPostLoginHandler = login.PostLoginHandlerFunc(loginHandler.Login)
	api.RegionGetRegionsHandler = region.GetRegionsHandlerFunc(regionHandler.GetRegions)
	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return recoveryMiddleware(corsMiddleware(rateLimitMeddleware(handler)))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w, r)
		if (*r).Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				spew.Dump(err)
				res, _ := json.Marshal(httperrors.ServiceUnavailable)
				http.Error(w, string(res) , 400)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func rateLimitMeddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if "/login" == r.URL.Path {
			st := store.GetStore()
			if limit, _ := st.GetLimit(r.RemoteAddr); limit > 2 {
				err := struct {
					Code    string `json:"code"`
					Message string `json:"message"`
				}{
					"002-001",
					"Too many request",
				}

				body, _ := json.Marshal(&err)

				w.WriteHeader(429)
				_, _ = w.Write(body)
				return
			}
		}
 		next.ServeHTTP(w,r)
	})
}