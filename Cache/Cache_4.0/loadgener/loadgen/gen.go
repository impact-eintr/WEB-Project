package loadgen

import (
	"bytes"
	"context"
	"fmt"
	"loadgener/loadgen/lib"
	"loadgener/loadgen/log"
	"math"

	"time"
)

// 日志记录器。
var logger = log.DLogger()

type myGenerator struct {
	timeoutNS  time.Duration //响应超时时间
	lps        uint32        //每秒荷载量
	durationNS time.Duration //持续负载时间

	concurrency uint32        //荷载并发量
	tickets     lib.Gotickets //gotoutine票池

	ctx        context.Context    //上下文
	cancelFunc context.CancelFunc //取消函数

	caller    lib.Caller //调用器
	callCount int64      //调用计数

	status uint32 // 荷载发生器的状态

	resultCh chan *lib.CallResult //结果调用通道
}

func NewGenerator(pset ParamSet) (lib.Generator, error) {
	logger.Infoln("New a load generator...")
	if err := pset.Check(); err != nil {
		return nil, err
	}

	gen := &myGenerator{
		caller:     pset.Caller,
		timeoutNS:  pset.TimeoutNS,
		lps:        pset.LPS,
		durationNS: pset.DurationNS,
		status:     lib.STATUS_ORIGINAL,
		resultCh:   pset.ResultCh,
	}
	if err := gen.init(); err != nil {
		return nil, err
	}
	return gen, nil
}

//初始化荷载发生器
func (gen *myGenerator) init() error {
	var buf bytes.Buffer
	buf.WriteString("初始化荷载发生器中...")

	// 载荷的并发量 ≈ 载荷的响应超时时间 / 载荷的发送间隔时间
	var total64 = int64(gen.timeoutNS) / int64(1e9/gen.lps)
	if total64 > math.MaxInt32 {
		total64 = math.MaxInt32
	}
	//荷载并发量
	gen.concurrency = uint32(total64)

	//初始化票仓
	tickets, err := lib.NewGoTickets(gen.concurrency)
	if err != nil {
		return err
	}
	gen.tickets = tickets

	buf.WriteString(fmt.Sprintf("初始化完成(concurrency=%d)", gen.concurrency))
	logger.Infoln(buf.String())
	return nil
}

func (gen *myGenerator) Start() bool {

}

func (gen *myGenerator) Stop() bool {

}

func (gen *myGenerator) Status() uint32 {

}

func (gen *myGenerator) CallCount() int64 {

}
