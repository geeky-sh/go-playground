package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Server struct {
	topics      []Topic
	subscribers []Subscriber
}

type Topic struct {
	name     string
	messages []string
}

type ServerInterface interface {
	CreateTopic(name string) error
	Publish(topicName string, msg string) error
	Subscribe(topicName string) (SubscriberInterface, error)
	Offset(sub SubscriberInterface, val int) error
}

func NewServer() ServerInterface {
	return &Server{topics: []Topic{}}
}

func (r *Server) CreateTopic(n string) error {
	r.topics = append(r.topics, Topic{name: n, messages: []string{}})
	return nil
}

func (r *Server) Publish(topicName string, msg string) error {
	for i, t := range r.topics {
		if t.name == topicName {
			r.topics[i].messages = append(r.topics[i].messages, msg)
			fmt.Println(r.topics[i].messages)
		}
	}

	for _, s := range r.subscribers {
		s.putOnChan()
	}

	return nil
}

func (r *Server) Subscribe(topicName string) (SubscriberInterface, error) {
	t := r.findTopic(topicName)

	s := Subscriber{topic: t, ID: uuid.NewString(), offset: 0, consumer: make(chan string, 10)}
	s.putOnChan()
	r.subscribers = append(r.subscribers, s)
	return &s, nil
}

func (r *Server) Offset(sub SubscriberInterface, val int) error {
	sub.SetOffset(val)
	sub.putOnChan()
	return nil
}

func (r *Server) findTopic(n string) *Topic {
	for _, t := range r.topics {
		if t.name == n {
			return &t
		}
	}
	return nil
}

type Subscriber struct {
	ID       string
	topic    *Topic
	offset   int
	consumer chan string
}

func NewSubscriber(topic *Topic) SubscriberInterface {
	return &Subscriber{offset: 0, ID: uuid.NewString(), consumer: make(chan string, 10), topic: topic}
}

type SubscriberInterface interface {
	Consume()
	putOnChan()
	SetOffset(v int)
}

func (r *Subscriber) Consume() {
	go func() {
		for {
			select {
			case msg := <-r.consumer:
				fmt.Printf("Consumer %s consumed msg %s", r.ID, msg)
				r.offset += 1
			}
		}
	}()
}
func (r *Subscriber) putOnChan() {
	fmt.Println(r.offset)
	fmt.Println("msg")
	fmt.Println(len(r.topic.messages))
	if r.offset < len(r.topic.messages) {
		newMsgs := r.topic.messages[r.offset:]
		for _, m := range newMsgs {
			r.consumer <- m
		}
	}
}

func (r *Subscriber) SetOffset(v int) {
	r.offset = v
}

func main() {
	s := NewServer()
	s.CreateTopic("payment")
	s.CreateTopic("sales")

	sb1, _ := s.Subscribe("payment")
	sb1.Consume()

	sb2, _ := s.Subscribe("sales")
	sb2.Consume()

	sb3, _ := s.Subscribe("payment")
	sb3.Consume()

	s.Publish("payment", "p1")
	s.Publish("payment", "p2")
	s.Publish("payment", "p3")
	s.Publish("payment", "p4")
	s.Publish("payment", "p5")
	s.Publish("payment", "p6")
	s.Publish("payment", "p7")

	s.Publish("sales", "p1")
	s.Publish("sales", "p2")
	s.Publish("sales", "p3")
	s.Publish("sales", "p4")
	s.Publish("sales", "p5")
	s.Publish("sales", "p6")
	s.Publish("sales", "p7")

	s.Offset(sb1, 2)
	s.Offset(sb3, 2)

	time.Sleep(5 * time.Second)
}
