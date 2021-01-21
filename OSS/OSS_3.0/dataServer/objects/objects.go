package objects

import (
	"OSS/dataServer/conf"
	"OSS/dataServer/locate"
	"OSS/dataServer/utils"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodGet {
		log.Println(strings.Split(r.URL.EscapedPath(), "/")[2])
		get(w, r)
		return
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	file := getFile(strings.Split(r.URL.EscapedPath(), "/")[2])
	if file == "" {
		w.WriteHeader(http.StatusNotFound)
		return

	}
	sendFile(w, file)

}

func getFile(hash string) string {
	file := conf.Conf.Dir + "/objects/" + hash
	f, _ := os.Open(file)
	d := url.PathEscape(utils.CalculateHash(f))
	f.Close()
	if d != hash {
		log.Println("object hash mismatch, remove", file)
		locate.Del(hash)
		os.Remove(file)
		return ""

	}
	return file

}

func sendFile(w io.Writer, file string) {
	f, _ := os.Open(file)
	defer f.Close()
	io.Copy(w, f)
}
