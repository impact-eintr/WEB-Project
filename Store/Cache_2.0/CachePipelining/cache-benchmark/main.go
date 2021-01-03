package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var typ, server, operation string
var total, valueSize, threads, keyspacelen, pipelen int

//初始化参数
func init() {
	flag.StringVar(&typ, "type", "redis", "cache server type")
	flag.StringVar(&server, "h", "localhsot", "cache server address")
	flag.IntVar(&total, "n", 1000, "total number of request")
	flag.IntVar(&valueSize, "d", 1000, "data size of SET/GET value in bytes")
	flag.IntVar(&threads, "c", 1, "number of parallel connections")
	flag.StringVar(&operation, "t", "set", "default set , could be set/get/mixed")
	flag.IntVar(&keyspacelen, "r", 0, "keyspacelen, use random keys from 0 to keyspacelen-1")
	flag.IntVar(&pipelen, "P", 1, "pipeline length")
	flag.Parse()
	fmt.Println("type is", typ)
	fmt.Println("server is", server)
	fmt.Println("total", total, "requests")
	fmt.Println("data size is", valueSize)
	fmt.Println("we have", threads, "connections")
	fmt.Println("operation is", operation)
	fmt.Println("keyspacelen is", keyspacelen)
	fmt.Println("pipeline length is", pipelen)

	rand.Seed(time.Now().UnixNano())

}

//statistic统计值
type statistic struct {
	count int           //同一时间级下操作计数
	time  time.Duration //持续时间
}

type result struct {
	getCount    int         //记录一共进行了多少次Get
	missCount   int         //记录一共进行了多少次Set
	setCount    int         //没有hit到的次数
	statBuckets []statistic //记录操作花费的时间
}

//增加新的统计值
//如果是在现有的"桶"中存在，就直接累加，否则先创建新的桶
func (r *result) addStatistic(bucket int, stat statistic) {
	if bucket > len(r.statBuckets)-1 {
		newStatBuckets := make([]statistic, bucket+1)
		copy(newStatBuckets, r.statBuckets)
		r.statBuckets = newStatBuckets
	}
	s := r.statBuckets[bucket]
	s.count += stat.count
	s.time += stat.time
	r.statBuckets[bucket] = s
}

func (r *result) addDuration(d time.Duration, typ string) {
	bucket := int(d / time.Millisecond)
	r.addStatistic(bucket, statistic{1, d})
	if typ == "get" {
		r.getCount++
	} else if typ == "set" {
		r.setCount++
	} else {
		r.missCount++
	}
}

func (r *result) addResult(src *result) {
	for b, s := range src.statBuckets {
		r.addStatistic(b, s)
	}
	r.getCount += src.getCount
	r.setCount += src.setCount
	r.missCount += src.missCount
}

//operate
func operate() {

}

func main() {
	//goroutine
	ch := make(chan *result, threads)
	res := &result{0, 0, 0, make([]static, 0)}
	start := time.Now()
	for i := 0; i < threads; i++ {
		go operate(i, total/threads, ch)
	}

}
