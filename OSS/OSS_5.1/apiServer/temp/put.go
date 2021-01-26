package temp

import (
	"OSS/apiServer/es"
	"OSS/apiServer/locate"
	"OSS/apiServer/rs"
	"OSS/apiServer/utils"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"net/url"
)

func Put(c *gin.Context) {
	token := url.PathEscape(c.Param("file")[1:])
	stream, err := rs.NewRSResumablePutStreamFromToken(token)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusForbidden)
		return
	}

	current := stream.CurrentSize()
	color.Green("current : %v\n", current)
	if current == -1 {
		c.Status(http.StatusNotFound)
		return
	}

	offset := utils.GetOffsetFromHeader(c.Request.Header)
	color.Green("offset : %v\n", offset)

	if current != offset {
		c.Status(http.StatusRequestedRangeNotSatisfiable)
		return
	}

	bytes := make([]byte, rs.BLOCK_SIZE)

	for {
		n, err := io.ReadFull(c.Request.Body, bytes)
		if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
			log.Println(err)
			c.Status(http.StatusInternalServerError)
			return
		}

		if n != rs.BLOCK_SIZE && current != stream.Size {
			return
		}

		stream.Write(bytes[:n])
		if current == stream.Size {
			stream.Flush()
			getStream, e := rs.NewRSResumableGetStream(stream.Servers, stream.Uuids, stream.Size)
			hash := url.PathEscape(utils.CalculateHash(getStream))
			if hash != stream.Hash {
				stream.Commit(false)
				log.Println("resumable put done but hash mismatch")
				c.Status(http.StatusForbidden)
				return

			}
			if locate.Exist(url.PathEscape(hash)) {
				stream.Commit(false)

			} else {
				stream.Commit(true)

			}
			e = es.AddVersion(stream.Name, stream.Hash, stream.Size)
			if e != nil {
				log.Println(e)
				c.Status(http.StatusInternalServerError)

			}
			return

		}
	}
}
