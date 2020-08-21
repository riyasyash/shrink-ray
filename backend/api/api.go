package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(v); err != nil {
		return err
	}
	return nil
}

func Respond(statusCode int, w http.ResponseWriter, resp interface{}) {
	b, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%v", err)
	}
	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json")
	w.Write(b)
}

func Redirect(statusCode int, w http.ResponseWriter, resp string) {
	w.Header().Add("Location", resp)
	b, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%v", err)
	}
	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json")
	w.Write(b)
}
