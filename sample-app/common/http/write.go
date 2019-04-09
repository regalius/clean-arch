package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func WriteStatusCode(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	w.Write([]byte(fmt.Sprintf("%s %s", http.StatusText(status), message)))
}

func WriteJSON(w http.ResponseWriter, data interface{}, status int) (err error) {
	w.WriteHeader(status)
	var dataByte []byte
	dataByte, err = json.Marshal(data)
	if err != nil {
		log.Println("[Common/http] ReturnJSON error unmarshal", err)
		w.Write([]byte("{}"))
		return
	}
	w.Write(dataByte)
	return
}
