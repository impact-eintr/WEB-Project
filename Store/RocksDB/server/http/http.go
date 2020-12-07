package http

import (
	"RocksDB/server/cache"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Server struct {
	cache.Cache //嵌套（继承）
}

type cacheHandler struct {
	*Server
}

type statusHandler struct {
	*Server
}

func New(c cache.Cache) *Server {
	return &Server{c}
}

func (s *Server) Listen() {
	http.Handle("/cache/", s.cacheHandler())
	http.Handle("/status", s.statusHandler())
	http.ListenAndServe(":12345", nil)
}

func (s *Server) cacheHandler() http.Handler {
	return &cacheHandler{s}
}

func (s *Server) statusHandler() http.Handler {
	return &statusHandler{s}
}

func (h *cacheHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := strings.Split(r.URL.EscapedPath(), "/")[2] //escapedPath returns a path encoded by Unicode
	if len(key) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	m := r.Method
	if m == http.MethodPut { //设置新值
		b, _ := ioutil.ReadAll(r.Body)
		if len(b) != 0 {
			err := h.Set(key, b)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError) //remote server error
			}
		}
		return
	}
	if m == http.MethodGet {
		b, err := h.Get(key)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if len(b) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Write(b)
		return
	}

	if m == http.MethodDelete {
		err := h.Del(key)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusNotFound)
		}
		return
	}

	//Method Not Found
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (h *statusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	b, err := json.Marshal(h.GetStat())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b = append(b, '\n')
	w.Write(b)
}
