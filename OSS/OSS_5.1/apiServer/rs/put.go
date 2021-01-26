package rs

import (
	"OSS/apiServer/objectstream"
	"fmt"
	"io"
)

type RSPutStream struct {
	*encoder
}

func NewRSPutStream(dataServers []string, hash string, size int64) (*RSPutStream, error) {
	if len(dataServers) != ALL_SHARDS {
		return nil, fmt.Errorf("dataServers number mismatch")

	}

	perShard := (size + DATA_SHARDS - 1) / DATA_SHARDS //每个数据分片的数据量
	writers := make([]io.Writer, ALL_SHARDS)           //写入流队列
	var e error

	for i := range writers {
		writers[i], e = objectstream.NewTempPutStream(dataServers[i],
			fmt.Sprintf("%s_%d", hash, i), perShard) //发送给某个数据节点：编号后的hash值,每份数据的大小
		if e != nil {
			return nil, e

		}

	}
	enc := NewEncoder(writers)

	return &RSPutStream{enc}, nil

}

func (s *RSPutStream) Commit(success bool) {
	s.Flush()
	for i := range s.writers {
		s.writers[i].(*objectstream.TempPutStream).Commit(success)

	}

}
