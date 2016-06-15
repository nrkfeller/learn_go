package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
)

func messageReceived(client *MQTT.Client, msg MQTT.Message) {
	topics := strings.Split(msg.Topic(), "/")
	msgFrom := topics[len(topics)-1]
	fmt.Print(msgFrom + ": " + string(msg.Payload()))
}

func main() {
	stdin := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().Unix())

	server := flag.String("server", "tcp://iot.eclipse.org:1883", "The MQTT server to connect to")
	room := flag.String("room", "gochat", "The chat room to enter. default 'gochat'")
	name := flag.String("name", "user"+strconv.Itoa(rand.Intn(1000)), "Username to be displayed")
	flag.Parse()

	subTopic := strings.Join([]string{"/gochat/", *room, "/+"}, "")
	pubTopic := strings.Join([]string{"/gochat/", *room, "/", *name}, "")

	opts := MQTT.NewClientOptions().AddBroker(*server).SetClientID(*name).SetCleanSession(true)

	opts.OnConnect = func(c *MQTT.Client) {
		if token := c.Subscribe(subTopic, 1, messageReceived); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}

	client := MQTT.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Printf("Connected as %s to %s\n", *name, *server)

	for {
		message, err := stdin.ReadString('\n')
		if err == io.EOF {
			os.Exit(0)
		}
		if token := client.Publish(pubTopic, 1, false, message); token.Wait() && token.Error() != nil {
			fmt.Println("Failed to send message")
		}
	}
}

// //define a function for the default message handler
// var f MQTT.MessageHandler = func(client *MQTT.Client, msg MQTT.Message) {
// 	fmt.Printf("TOPIC: %s\n", msg.Topic())
// 	fmt.Printf("MSG: %s\n", msg.Payload())
// }
//
// func main() {
// 	//create a ClientOptions struct setting the broker address, clientid, turn
// 	//off trace output and set the default message handler
// 	opts := MQTT.NewClientOptions().AddBroker("tcp://iot.eclipse.org:1883")
// 	opts.SetClientID("go-simple")
// 	opts.SetDefaultPublishHandler(f)
//
// 	//create and start a client using the above ClientOptions
// 	c := MQTT.NewClient(opts)
// 	if token := c.Connect(); token.Wait() && token.Error() != nil {
// 		panic(token.Error())
// 	}
//
// 	//subscribe to the topic /go-mqtt/sample and request messages to be delivered
// 	//at a maximum qos of zero, wait for the receipt to confirm the subscription
// 	if token := c.Subscribe("go-mqtt/sample", 0, nil); token.Wait() && token.Error() != nil {
// 		fmt.Println(token.Error())
// 		os.Exit(1)
// 	}
//
// 	//Publish 5 messages to /go-mqtt/sample at qos 1 and wait for the receipt
// 	//from the server after sending each message
// 	for i := 0; i < 5; i++ {
// 		text := fmt.Sprintf("this is msg #%d!", i)
// 		token := c.Publish("go-mqtt/sample", 0, false, text)
// 		token.Wait()
// 	}
//
// 	time.Sleep(3 * time.Second)
//
// 	//unsubscribe from /go-mqtt/sample
// 	if token := c.Unsubscribe("go-mqtt/sample"); token.Wait() && token.Error() != nil {
// 		fmt.Println(token.Error())
// 		os.Exit(1)
// 	}
//
// 	c.Disconnect(250)
// }
