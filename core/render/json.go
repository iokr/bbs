package render

import (
	"encoding/json"
	"net/http"
)

var jsonContentType = []string{"application/json; charset=utf-8"}

type JSON struct {
	Data interface{}
}

func (r JSON) Render(w http.ResponseWriter) (err error) {
	return WriteJSON(w, r.Data)
}

func WriteJSON(w http.ResponseWriter, obj interface{}) error {
	writeContentType(w, jsonContentType)
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = w.Write(jsonBytes)
	return err
}

func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}
