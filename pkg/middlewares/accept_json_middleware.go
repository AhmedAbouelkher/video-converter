package middlewares

import "net/http"

func ContentTypeApplicationJsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request)  {
		res.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(res, req)
	})
}