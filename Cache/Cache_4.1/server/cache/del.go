package cache

import (
	"errors"
	"unsafe"
)

/*
#include <stdlib.h>
#include <rocksdb/c.h>
#cgo LDFLAGS: -lpthread -lstdc++ -lrocksdb -ldl -lm -O3
*/
import "C"

func (c *rocksdbCache) Del(key string) error {
	k := C.CString(key)
	defer C.free(unsafe.Pointer(k))
	C.rocksdb_delete(c.db, c.wo, k, C.size_t(len(key)), &c.e)
	if c.e != nil {
		return errors.New(C.GoString(c.e))

	}
	return nil

}
