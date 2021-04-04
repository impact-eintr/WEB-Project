package lib

import (
	"errors"
	"fmt"
)

//令牌桶的接口
type Gotickets interface {
	Take()
	Return()
	Active() bool      //票池是否激活
	Total() uint32     //总票数
	Remainder() uint32 //剩余
}

//令牌桶的实现类型
type myGoTickets struct {
	total    uint32        //总票数
	ticketCh chan struct{} //票的容器
	active   bool          //票池是否激活
}

func (gt *myGoTickets) init(total uint32) bool {
	if gt.active {
		return false //已经激活无需初始化
	}

	if total == 0 {
		return false //别逗了
	}

	ch := make(chan struct{}, total)
	n := int(total)
	for i := 0; i < n; i++ {
		ch <- struct{}{}
	}
	gt.ticketCh = ch
	gt.total = total
	gt.active = true
	return true
}

func NewGoTickets(total uint32) (Gotickets, error) {
	gt := myGoTickets{}
	if !gt.init(total) {
		errMsg := fmt.Sprintf("令牌桶无法初始化(total=%d)\n", total)
		return nil, errors.New(errMsg)
	}
	return &gt, nil
}

func (gt *myGoTickets) Take() {
	<-gt.ticketCh
}

func (gt *myGoTickets) Return() {
	gt.ticketCh <- struct{}{}
}

func (gt *myGoTickets) Active() bool {
	return gt.active
}

func (gt *myGoTickets) Total() uint32 {
	return gt.total
}

func (gt *myGoTickets) Remainder() uint32 {
	return uint32(len(gt.ticketCh))
}
