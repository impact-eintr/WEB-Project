package cache

func (this *Stat) add(key string, val []byte) {
	this.Count++
	this.KeySize += int64(len(key))
	this.ValueSize += int64(len(val))

}

func (this *Stat) del(key string, val []byte) {
	this.Count--
	this.KeySize -= int64(len(key))
	this.ValueSize -= int64(len(val))

}
