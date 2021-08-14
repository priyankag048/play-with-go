package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// adds CSRF protection to all post requests
func checkCSRF(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// loads session on every request
func sessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
