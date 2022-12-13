package middlewares

import (
	"fmt"
	"net/http"
)

func AllowOnlyGET(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Only GET is allowed!", http.StatusBadRequest)
			return
		}
		fmt.Printf("%s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	})
}

func AllowOnlyPOST(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Only POST is allowed!", http.StatusBadRequest)
			return
		}
		fmt.Printf("%s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	})
}

func AllowOnlyPUT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			http.Error(w, "Only PUT is allowed!", http.StatusBadRequest)
			return
		}
		fmt.Printf("%s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	})
}

func AllowOnlyDELETE(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			http.Error(w, "Only DELETE is allowed!", http.StatusBadRequest)
			return
		}
		fmt.Printf("%s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	})
}