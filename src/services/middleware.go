package services

import (
	"gomscode/src/logger"
	"net/http"
	"time"

)

// Middleware exported
func Middleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
	/*	if (r.RequestURI != "/"){	
			userName := utility.GetUserName(r)
			if utility.IsEmpty(userName) {
				http.Redirect(w, r, "/", 302)
			}
		}
	*/	f(w, r)
		requesterIP := r.RemoteAddr
		logger.LogOutInfo(r.Method, r.RequestURI, requesterIP, r.Host, time.Since(start).String(), r.UserAgent())	
		}
}
