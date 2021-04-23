package internal

type Broker interface {
	// 消息推送 主题 消息
	publish(topic string, msg interface{}) error

	// 消息订阅
	subscribe(topic string) (<-chan interface{}, error)
	unsubscribe(topic string, sub <-chan interface{}) error

	close() // 关闭消息队列

	// 服务端广播消息
	broadcast(msg interface{}, subscribers []chan interface{})
	setCondtions(capacity int)
}