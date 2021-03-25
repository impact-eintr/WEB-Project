package cache

/*
#include <stdlib.h>
#include <rocksdb/c.h>
#cgo LDFLAGS: -lpthread -lstdc++ -lrocksdb -ldl -lm -O3
*/
import "C"
import "runtime"

func newRocksdbCache() *rocksdbCache {

}
