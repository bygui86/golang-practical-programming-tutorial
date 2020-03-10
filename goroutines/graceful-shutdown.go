package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*
	Example that shows how to gracefully terminate goroutines:
		. producer: flag in a normal for-loop
		. consumer: specific channel
*/
func main() {
	messageChannel := make(chan string, 1000)

	consumer := &Consumer{
		messageChannel: messageChannel,
		quitChannel:    make(chan bool),
	}
	go consumer.startConsumer()
	defer consumer.stopConsumer()

	producer := &Producer{
		messageChannel: messageChannel,
		produceFlag:    true,
	}
	go producer.startProducer()
	defer producer.stopProducer()

	syscallCh := make(chan os.Signal, 3)
	signal.Notify(syscallCh, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-syscallCh
	fmt.Println("GLOBAL TERMINATION SIGNAL")
	time.Sleep(5 * time.Second)
}

type Producer struct {
	messageChannel   chan string
	produceFlag      bool
	produceFlagMutex sync.Mutex
}

func (p *Producer) startProducer() {
	fmt.Println("START PRODUCER")
	counter := 1
	for {
		message := fmt.Sprintf("Message number %d", counter)
		fmt.Printf("producer sending message to channel: %s \n", message)
		p.messageChannel <- message
		counter++
		time.Sleep(time.Second)

		p.produceFlagMutex.Lock()
		if !p.produceFlag {
			fmt.Println("PROCUCER TEMRINATION SIGNAL")
			break
		}
		p.produceFlagMutex.Unlock()
	}
}

func (p *Producer) stopProducer() {
	fmt.Println("STOP PROCUCER")
	p.produceFlagMutex.Lock()
	defer p.produceFlagMutex.Unlock()
	p.produceFlag = false
}

type Consumer struct {
	messageChannel chan string
	quitChannel    chan bool
}

func (c *Consumer) startConsumer() {
	fmt.Println("START CONSUMER")
	for {
		select {
		case message := <-c.messageChannel:
			fmt.Printf("consumer received message: %s\n", message)
		case <-c.quitChannel:
			fmt.Println("CONSUMER TERMINATION SIGNAL")
			break
		}
	}
}

func (c *Consumer) stopConsumer() {
	fmt.Println("STOP CONSUMER")
	c.quitChannel <- true
}
