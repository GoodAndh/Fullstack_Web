package exception

import "net/http"

func NotFoundHandler() http.Handler {
	return http.HandlerFunc(notFound)
}

func notFound(w http.ResponseWriter,r *http.Request)  {
	WriteJson(w,http.StatusNotFound,"not found","sorry but this route is unavailable",nil)
}