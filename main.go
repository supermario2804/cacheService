package main

import (
	"cacheDataService/handlers"
	"cacheDataService/utils"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/streadway/amqp"
	"net/http"
	"strings"
	"time"
)

func main() {

	url := "amqp://guest:guest@amq:5672"
	http.HandleFunc("/api/setPage", handlers.SetPageCache)
	http.HandleFunc("/api/set", handlers.SetTableCache)
	http.HandleFunc("/api/get", handlers.GetTableCache)
	http.HandleFunc("/api/getPage", handlers.GetPageCache)
	http.HandleFunc("/api/healthcheck", handlers.HealthCheck)

	rabbitmqSetupQueue(url)
	go pollForDBChanges(url)
	go backupDB()
	fmt.Println("The server started...")
	http.ListenAndServe(":8090", nil)
}

func backupDB() {

	time.Sleep(10 * time.Second)

	for {
		utils.PrintInfo("Taking backup of redis DB")
		client := redis.NewClient(&redis.Options{
			Addr:     "redis:6379",
			Password: "",
			DB:       0,
		})

		defer client.Close()

		_, connErr := client.Ping().Result()
		if connErr != nil {
			utils.HandleError(connErr, "Redis DB backup failed!!")
		}
		saveErr := client.BgSave().Err()
		if saveErr != nil {
			utils.HandleError(saveErr, "Redis DB backup has failed!")
		}

		utils.PrintInfo("Redis backup completed successfully!!")
		time.Sleep(30 * time.Minute)
	}

}

func rabbitmqSetupQueue(url string) {
	time.Sleep(20 * time.Second)
	// Connect to the rabbitMQ instance
	connection, err := amqp.Dial(url)
	defer connection.Close()
	if err != nil {
		panic("could not establish connection with RabbitMQ:" + err.Error())

	}
	channel, err := connection.Channel()

	fmt.Printf("Channel has been declared\n")
	if err != nil {
		panic("could not open RabbitMQ channel:" + err.Error())
	}
	err = channel.ExchangeDeclare("events", "topic", true, false, false, false, nil)

	fmt.Printf("Exchange has been declared\n")
	if err != nil {
		panic(err)
	}
	// We create a queue named Test
	_, err = channel.QueueDeclare("test", true, false, false, false, nil)

	if err != nil {
		panic("error declaring the queue: " + err.Error())
	}

	fmt.Printf("Queue has been declared\n")
	// We bind the queue to the exchange to send and receive data from the queue
	err = channel.QueueBind("test", "#", "events", false, nil)

	fmt.Printf("Queue has been binded\n")
	if err != nil {
		panic("error binding to the queue: " + err.Error())
	}
}

func pollForDBChanges(url string) {
	connection, err := amqp.Dial(url)
	if err != nil {
		panic("could not establish connection with RabbitMQ:" + err.Error())
	}
	channel, err := connection.Channel()

	if err != nil {
		panic("could not open RabbitMQ channel:" + err.Error())
	}

	for {
		fmt.Printf("Polling has been started\n")
		// We consume data in the queue named test using the channel we created in go.
		msgs, err := channel.Consume("test", "", false, false, false, false, nil)

		if err != nil {
			panic("error consuming the queue: " + err.Error())
		}

		// We loop through the messages in the queue and print them to the console.
		// The msgs will be a go channel, not an amqp channel
		for msg := range msgs {
			//print the message to the console
			fmt.Println("message received: " + string(msg.Body))
			data := strings.Split(string(msg.Body), "_")
			fmt.Println("==> Notification received from rabbitmq")
			fmt.Printf("==> Row from Table : %s with Primary Key : %s has been changed.\n", data[0], data[1])
			fmt.Println("==> Reloading Data from DB.")
			// Acknowledge that we have received the message so it can be removed from the queue
			msg.Ack(false)
		}
	}
}
