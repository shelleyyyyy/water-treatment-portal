package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type BrokerConfig struct {
	Broker   string
	Port     string
	Pass     string
	User     string
	ClientID string
}

type ResponseData struct {
	Items []Item `json:"items"`
}

type Item struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Value string `json:"value"`
	Topic string `json:"topic"`
	Chart bool   `json:"chart"`
}

type Historical struct {
	Topic string `json:"topic"`
	Value string `json:"value"`
}

type Msg struct {
	Topic   string `json:"topic"`
	Message string `json:"message"`
}

func main() {

	// Load Configuration
	err := godotenv.Load("app.env")

	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	brokerConf := BrokerConfig{
		Broker:   os.Getenv("BROKER"),
		Port:     os.Getenv("BROKER_PORT"),
		Pass:     os.Getenv("BROKER_PASS"),
		User:     os.Getenv("BROKER_USER"),
		ClientID: os.Getenv("MQTT_CLIENT_ID"),
	}

	pb := Pocket{
		Username:       os.Getenv("POCKETBASE_USERNAME"),
		Password:       os.Getenv("POCKETBASE_PASSWORD"),
		Pocketbase:     os.Getenv("POCKETBASE"),
		PocketbasePort: os.Getenv("POCKETBASE_PORT"),
	}

	for i := 0; i < 10; i++ {
		fmt.Println("Connecting to Pocketbase...")
		err := pb.Auth()

		if err != nil {
			fmt.Println("Failed Pocketbase Connection. Retrying in 5s")
		} else {
			break
		}

		time.Sleep(5 * time.Second)
	}

	fmt.Println("Connected to Pocketbase!")

	res, err := pb.GetList("subsciptions")

	if err != nil {
		return
	}

	var subs ResponseData

	err = json.Unmarshal([]byte(res), &subs)

	if err != nil {
		fmt.Println(err)
		return
	}

	// MQTT Handlers
	var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		// fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())

		var i Item

		for _, s := range subs.Items {
			if s.Topic == msg.Topic() {
				i = s

				clone := s
				clone.Value = string(msg.Payload())
				jsonData, err := json.Marshal(clone)

				if err != nil {
					return
				}

				pb.Update(s.Id, "subsciptions", jsonData)

			}

		}

		if i.Chart {
			hist := Historical{
				Topic: msg.Topic(),
				Value: string(msg.Payload()),
			}

			// fmt.Println("REACHING THIS")
			// fmt.Println(hist)

			jsonData, err := json.Marshal(hist)

			if err != nil {
				return
			}

			pb.Create("historical", jsonData)
		}
	}

	var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
		fmt.Println("Connected to mqtt Broker!")
	}

	var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
		fmt.Printf("Connect lost: %v", err)

		if token := client.Connect(); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", brokerConf.Broker, brokerConf.Port))
	opts.SetClientID(brokerConf.ClientID)
	opts.SetUsername(brokerConf.User)
	opts.SetPassword(brokerConf.Pass)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for _, s := range subs.Items {
		mqttSub(client, s)
	}

	status := "on"

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/publishMessage", func(c echo.Context) error {
		u := new(Msg)

		if err := c.Bind(u); err != nil {
			return err
		}

		mqttPub(client, u.Topic, u.Message)

		return c.JSON(http.StatusCreated, "success: "+status)
	})
	e.GET("/start", func(c echo.Context) error {

		status = "run"

		for _, s := range subs.Items {
			mqttSub(client, s)
		}

		return c.JSON(http.StatusFound, "Starting Experiment... Status: "+status)

	})
	e.GET("/kill", func(c echo.Context) error {

		status = "off"

		err = pb.RemoveCollection("historical")
		if err != nil {
			return c.JSON(http.StatusNotFound, "IDK SOMETHING HEPPEND")
		}

		for _, s := range subs.Items {
			client.Unsubscribe(s.Topic)
		}

		return c.JSON(http.StatusFound, "Destroying Experiment..., Status: "+status)

	})
	e.GET("/pause", func(c echo.Context) error {

		status = "pause"

		for _, s := range subs.Items {
			client.Unsubscribe(s.Topic)
		}

		status = "pause"

		return c.JSON(http.StatusFound, "Pausing Experiment..., Status: "+status)

	})
	e.GET("/status", func(c echo.Context) error {

		type Status struct {
			Status string `json:"status"`
		}

		s := Status{status}

		return c.JSON(http.StatusFound, s)

	})
	e.Logger.Fatal(e.Start(":1323"))
}

func mqttSub(client mqtt.Client, item Item) {
	token := client.Subscribe(item.Topic, 1, nil)
	token.Wait()
}

func mqttPub(client mqtt.Client, topic string, msg string) {
	text := fmt.Sprintf(msg)
	token := client.Publish(topic, 0, false, text)
	token.Wait()
}
