package lib

import (
	"time"
)

const ()

// 声明代表载荷发生器状态的常量。
const (
	// STATUS_ORIGINAL 代表原始。
	STATUS_ORIGINAL uint32 = 0
	// STATUS_STARTING 代表正在启动。
	STATUS_STARTING uint32 = 1
	// STATUS_STARTED 代表已启动。
	STATUS_STARTED uint32 = 2
	// STATUS_STOPPING 代表正在停止。
	STATUS_STOPPING uint32 = 3
	// STATUS_STOPPED 代表已停止。
	STATUS_STOPPED uint32 = 4
)

//原生请求
type RawReq struct {
	ID  int64
	req []byte
}

//原生响应
type RawRsp struct {
	ID     int64
	Resp   []byte
	Err    error
	Elapse time.Duration
}

//结果代码
type RetCode int

type CallResult struct {
	ID     int64         //ID
	Req    RawReq        //原生请求
	Resp   RawRsp        //原生响应
	Code   RetCode       //响应代码
	Msg    string        //结果成因的简述
	Elapse time.Duration //耗时
}

// Generator 荷载发生器接口
type Generator interface {
	// 启动载荷发生器。
	// 结果值代表是否已成功启动。
	Start() bool
	// 停止载荷发生器。
	// 结果值代表是否已成功停止。
	Stop() bool
	// 获取状态。
	Status() uint32
	// 获取调用计数。每次启动会重置该计数。
	CallCount() int64
}
