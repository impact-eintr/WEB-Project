package rabbitmq

import (
	"encoding/json"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	channel  *amqp.Channel
	conn     *amqp.Connection
	Name     string
	exchange string
}

//创建一个新的rabbitmq结构体
func New(s string) *RabbitMQ {
	conn, err := amqp.Dial(s)
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	//申请新的队列
	q, err := ch.QueueDeclare(
		"",    //name
		false, //durable
		true,  //delete when unused
		false, //exclusive
		false, //no-wait
		nil,   //argument
	)
	if err != nil {
		panic(err)
	}
	mq := new(RabbitMQ)
	mq.channel = ch
	mq.conn = conn
	mq.Name = q.Name
	return mq
}

//将自己的消息队列与一个exchange绑定，所有发往该exchange的消息都能在自己的消息队列中哦个被接收到
func (q *RabbitMQ) Bind(exchange string) {
	e := q.channel.QueueBind(
		q.Name,   // queue name
		"",       // routing key
		exchange, // exchange
		false,
		nil,
	)
	if e != nil {
		panic(e)
	}
	q.exchange = exchange

}

//send可以往某个消息队列发送消息
func (q *RabbitMQ) Send(queue string, body interface{}) {
	str, e := json.Marshal(body)
	if e != nil {
		panic(e)
	}
	e = q.channel.Publish("",
		queue,
		false,
		false,
		amqp.Publishing{
			ReplyTo: q.Name,
			Body:    []byte(str),
		})
	if e != nil {
		panic(e)
	}

}

//publish 可以向某个exchange发送消息
func (q *RabbitMQ) Publish(exchange string, body interface{}) {
	str, e := json.Marshal(body)
	if e != nil {
		panic(e)
	}
	e = q.channel.Publish(exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ReplyTo: q.Name,
			Body:    []byte(str),
		})
	if e != nil {
		panic(e)
	}

}

//Consume方法用于生成一个接收消息的go channel,使客户端可以通过go语言的原生机制接收队列中的消息
func (q *RabbitMQ) Consume() <-chan amqp.Delivery {
	c, e := q.channel.Consume(q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if e != nil {
		panic(e)
	}
	return c

}

//关闭消息队列
func (q *RabbitMQ) Close() {
	q.channel.Close()
	q.conn.Close()
}
