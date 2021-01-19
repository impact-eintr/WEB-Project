package objectstream

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type TempPutStream struct {
	Server string
	Uuid   string
}

func NewTempPutStream(server, object string, size int64) (*TempPutStream, error) {
	log.Println(object)
	request, err := http.NewRequest("POST", "http://"+server+"/temp/"+object, nil)
	if err != nil {
		return nil, err
	}

	//构造发往dataserver的请求
	request.Header.Set("size", strconv.FormatInt(size, 10))
	client := http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	uuid, err := ioutil.ReadAll(response.Body) //返回dataserver返回的uuid值
	if err != nil {
		return nil, err
	}

	return &TempPutStream{
		server,
		string(uuid),
	}, nil
}

//io.TeeReader() 调用 Write() 递送 PATCH请求
func (w *TempPutStream) Write(p []byte) (n int, err error) {
	request, err := http.NewRequest("PATCH", "http://"+w.Server+"/temp/"+w.Uuid, strings.NewReader(string(p)))
	if err != nil {
		return 0, err
	}
	//递送请求
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return 0, err
	}
	if response.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("dataServer return http code : %d", response.StatusCode)
	}

	log.Println("p[]:", p)
	return len(p), nil
}

//commit 根据传入参数决定项数据节点发送put还是delete
func (w *TempPutStream) Commit(good bool) {
	method := "DELETE"
	if good {
		method = "PUT"

	}
	request, _ := http.NewRequest(method, "http://"+w.Server+"/temp/"+w.Uuid, nil)
	client := http.Client{}
	client.Do(request)

}

func NewTempGetStream(server, uuid string) (*GetStream, error) {
	return newGetStream("http://" + server + "/temp/" + uuid)

}
