package mux

import (
	"encoding/json"
	"net/http"
)


func writeJson(w http.ResponseWriter, code int, data map[string]any) {
    encoder := json.NewEncoder(w)
    encoder.SetIndent("", "    ")
    w.WriteHeader(code)
    encoder.Encode(data)
}
