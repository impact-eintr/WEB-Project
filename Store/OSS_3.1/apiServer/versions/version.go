package versions

import (
	"OSS/apiServer/es"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Get(c *gin.Context) {
	from := 0
	size := 1000
	name := c.Param("file")
	for {
		metas, err := es.SearchAllVersions(name, from, size)
		if err != nil {
			log.Println(err)
			c.Status(http.StatusInternalServerError)
			return
		}
		for i := range metas {
			b, _ := json.Marshal(metas[i])
			b = append(b, '\n')
			c.Data(http.StatusOK, "application/octet-stream", b)
		}
		if len(metas) != size {
			return
		}
		from += size
	}

}

//	func Handler(w http.ResponseWriter, r *http.Request) {
//		m := r.Method
//		if m != http.MethodGet {
//			w.WriteHeader(http.StatusMethodNotAllowed)
//			return
//		}
//		from := 0
//		size := 1000
//		name := strings.Split(r.URL.EscapedPath(), "/")[2]
//		for {
//			metas, err := es.SearchAllVersions(name, from, size)
//			if err != nil {
//				log.Println(err)
//				w.WriteHeader(http.StatusInternalServerError)
//				return
//			}
//			for i := range metas {
//				b, _ := json.Marshal(metas[i])
//				w.Write(b)
//				w.Write([]byte("\n"))
//			}
//			if len(metas) != size {
//				return
//			}
//			from += size
//		}
//	}
