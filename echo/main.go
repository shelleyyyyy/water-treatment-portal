package main

import (
	"log"
	"net/http"
	"os"

	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"github.com/joho/godotenv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Config struct {
	Broker         string
	BrokerPort     string
	BrokerPass     string
	BrokerUser     string
	Pocketbase     string
	PocketbasePort string
}

// func LoadConfig(path string) (config Config, err error) {
// 	viper.AddConfigPath(path)
// 	viper.SetConfigName("app")
// 	viper.SetConfigType("env")

// 	viper.AutomaticEnv()

// 	err = viper.ReadInConfig()
// 	if err != nil {
// 		return
// 	}

// 	err = viper.Unmarshal(&config)
// 	return
// }

func main() {

	err := godotenv.Load("app.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	broker := os.Getenv("BROKER")
	brokerPort := os.Getenv("BROKER_PORT")
	brokerPass := os.Getenv("BROKER_PASS")
	brokerUser := os.Getenv("BROKER_USER")
	pocketbase := os.Getenv("POCKETBASE")
	pocketbasePort := os.Getenv("POCKETBASE_PORT")

	fig := Config{broker, brokerPort, brokerPass, brokerUser, pocketbase, pocketbasePort}

	fmt.Println(fig)

	go fig.StartSubs()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/publishMessage", fig.publishMessage)
	e.Logger.Fatal(e.Start(":1323"))
}

type Msg struct {
	Topic   string `json:"topic"`
	Message string `json:"message"`
}

func (fig Config) publishMessage(c echo.Context) error {
	u := new(Msg)
	if err := c.Bind(u); err != nil {
		return err
	}

	fig.startMqtt(u.Topic, u.Message)

	return c.JSON(http.StatusCreated, "success")
}

func (fig Config) startMqtt(topic string, msg string) {
	// var broker = "broker.emqx.io"
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", fig.Broker, fig.BrokerPort))
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername(fig.BrokerUser)
	opts.SetPassword(fig.BrokerPass)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// sub(client)
	publish(client, topic, msg)

	client.Disconnect(250)
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func publish(client mqtt.Client, topic string, msg string) {
	text := fmt.Sprintf(msg)
	token := client.Publish(topic, 0, false, text)
	token.Wait()
}

func sub(client mqtt.Client) {
	topic := "topic/test"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", topic)
}
