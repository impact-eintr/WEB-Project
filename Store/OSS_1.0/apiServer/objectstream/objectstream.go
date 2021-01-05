//本包是对http包的封装，将一些http函数的调用转换为读写流的形式
package objectstream

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type PutStream struct {
	writer *io.PipeWriter
	c      chan error
}

//object 请求对象的名称
//返回 一个阻塞写管道 一个errorchannel
func NewPutStream(server, object string) *PutStream {
	reader, writer := io.Pipe() //注意管道的读写是阻塞的
	c := make(chan error)
	go func() {
		request, _ := http.NewRequest("PUT", "http://"+server+"/objects/"+object, reader)
		log.Println("stream 接收的dataNode:", request.URL)
		client := http.Client{}
		r, e := client.Do(request)
		//r, e = http.Get("http://" + server + "/objects/+object")
		log.Println(r.Body)
		if e == nil && r.StatusCode != http.StatusOK {
			e = fmt.Errorf("dataServer return http code %d", r.StatusCode)

		}
		c <- e
	}()
	return &PutStream{writer, c}
}

func (w *PutStream) Write(p []byte) (n int, err error) {
	return w.writer.Write(p)
}

func (w *PutStream) Close() error {
	w.writer.Close()
	return <-w.c
}

type GetStream struct {
	reader io.Reader
}

//url 目标dataNode 的地址 newGetStream 隐藏了真实的url
func newGetStream(url string) (*GetStream, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("dataServer return http code %d", r.StatusCode)
	}
	return &GetStream{r.Body}, nil
}

func NewGetStream(server, object string) (*GetStream, error) {
	if server == "" || object == "" {
		return nil, fmt.Errorf("非法的节点%s对象%s", server, object)
	}
	return newGetStream("http://" + server + "/objects/" + object)
}
func (r *GetStream) Read(p []byte) (n int, err error) {
	return r.reader.Read(p)
}
