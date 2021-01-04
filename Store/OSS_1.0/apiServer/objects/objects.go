package objects

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func put(w http.ResponseWriter, r *http.Request) {
	object := strings.Split(r.URL.EscapedPath(), "/")[2]
	c, err := storeObject(r.Body, object)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(c)
}

func storeObject(r io.Reader, object string) (int, error) {
	stream, err := putStrea
}
