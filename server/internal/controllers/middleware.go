package controllers

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/bcinnovationlabs/Apps/Chatbot-AI/server/pkg/utils"
)

// Middleware will split the full URL path into slash-sperated parts and save to
// the context to flow downstream in the app for this particular request.
func URLProcessorMiddleware(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Split path into slash-separated parts, for example, path "/foo/bar"
		// gives p==["foo", "bar"] and path "/" gives p==[""]. Our API starts with
		// "/api/v1", as a result we will start the array slice at "3".
		p := strings.Split(r.URL.Path, "/")[1:]
		n := len(p)

		// Open our program's context based on the request and save the
		// slash-seperated array from our URL path.
		ctx := r.Context()
		ctx = context.WithValue(ctx, "url_split", p)
		ctx = context.WithValue(ctx, "length", n)

		// Flow to the next middleware.
		fn(w, r.WithContext(ctx))
	}
}

func JWTProcessorMiddleware(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// Read our application's signing key and attach it to the application
		// context so it can flow downstream in all our applications.

		accessToken := r.Header.Get("Authorization")

		if accessToken != "" {
			// Special thanks to "poise" via https://stackoverflow.com/a/44700761
			splitToken := strings.Split(accessToken, "JWT ")
			accessToken = splitToken[1]
			ctx = context.WithValue(ctx, "access_token", accessToken)
			// log.Println(reqToken) // For debugging purposes only.
			secretKey, err := ioutil.ReadFile(".env")
			sessionUUID, err := utils.ProcessJWTToken(secretKey, accessToken)
			if err == nil {
				ctx = context.WithValue(ctx, "is_authorized", true)
				ctx = context.WithValue(ctx, "session_uuid", sessionUUID)

				// Flow to the next middleware with our JWT token saved.
				fn(w, r.WithContext(ctx))
				return
			}
			log.Println("JWTProcessorMiddleware | ProcessJWT | err", err)
		}

		// Flow to the next middleware without anything done.
		ctx = context.WithValue(ctx, "is_authorized", false)
		fn(w, r.WithContext(ctx))
	}
}

func ChainMiddleware(fn http.HandlerFunc) http.HandlerFunc {
	// Attach our middleware
	fn = URLProcessorMiddleware(fn)
	fn = JWTProcessorMiddleware(fn)
	return func(w http.ResponseWriter, r *http.Request) {
		// Flow to the next middleware.
		fn(w, r)
	}
}
