package app

import (
	"net/http"
	"strings"

	u "github.com/sudhanshu-chauhan/bpforgo/utils"
)

// JwtAuthentication middleware function to handle auth ...
var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		notAuth := []string{}
		requestPath := r.URL.Path

		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}
		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			response = u.Message(false, "missing auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}
		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			response = u.Message(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}
	})
}
