package service

import "net/http"

func applyCORS(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept")

		if r.Method == "OPTIONS" {
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if r.Method == "HEAD" {
			return
		}
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
