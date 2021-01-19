package temp

import (
	"OSS/dataServer/conf"
	"encoding/json"
	"io/ioutil"
	"os"
)

type tempInfo struct {
	Uuid string
	Name string
	Size int64
}

func readFromFile(uuid string) (*tempInfo, error) {
	f, err := os.Open(conf.Conf.Dir + "/temp/" + uuid)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	b, _ := ioutil.ReadAll(f)
	var info tempInfo
	json.Unmarshal(b, &info)
	return &info, nil
}
