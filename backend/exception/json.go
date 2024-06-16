package exception

import (
	"encoding/json"
	"fullstack_toko/backend/model/web"
	"net/http"
)

func ParseJson(r *http.Request, stx any) error {
	return json.NewDecoder(r.Body).Decode(stx)
}

func WriteJson(w http.ResponseWriter, code int, status, message string, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	wr := &web.WebResponse{
		Code:    code,
		Status:  status,
		Message: message,
		Data:    v,
	}
	return json.NewEncoder(w).Encode(wr)
}

func JsonForbidden(w http.ResponseWriter, message string, v any) error {
	return WriteJson(w, http.StatusForbidden, "status forbidden", message, v)
}

func JsonBadRequest(w http.ResponseWriter, message string, v any) error {
	return WriteJson(w, http.StatusBadRequest, "status bad reqeust", message, v)
}

func JsonInternalError(w http.ResponseWriter, message string, v any) error {
	return WriteJson(w, http.StatusInternalServerError, "status internal server error", message, v)
}

func JsonUnauthorized(w http.ResponseWriter, message string, v any) error {
	return WriteJson(w, http.StatusUnauthorized, "status unauthorized", message, v)
}
