package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Subscribtion struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Value string `json:"value"`
	Topic string `json:"topic"`
}

type UpdateDataStruct struct {
	Title float64 `json:"title"`
	Value float64 `json:"value"`
	Topic float64 `json:"topic"`
}

func StartSubs() {

	_, subs := GetData("subsciptions")

	// fmt.Println(subs.Items)

	// subscribtions := []*Subscribtion{}

	// s := Subscribtion{
	// 	Id:    "zlwzmqwbyctrgql",
	// 	Title: "Out Pump Status",
	// 	Value: "off",
	// 	Topic: "outpump",
	// }

	// s1 := Subscribtion{
	// 	Id:    "z1t3vgo4jhg23d5",
	// 	Title: "In Pump Status",
	// 	Value: "tmp",
	// 	Topic: "outpump",
	// }

	// subscribtions = append(subscribtions, &s, &s1)

	var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Message %s received on topic %s\n", msg.Payload(), msg.Topic())

		for _, x := range subs.Items {
			if msg.Topic() == x.Topic {

				fmt.Println(x.Value)
				x.Value = string(msg.Payload())
				fmt.Println(x.Value)

				marsh, err := json.Marshal(x)
				if err != nil {
					fmt.Println(err)
					return
				}

				UpdateData("subsciptions", x.Id, marsh)
			}
		}
	}

	var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
		fmt.Println("Connected")
	}

	var connectionLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
		fmt.Printf("Connection Lost: %s\n", err.Error())
	}

	var broker = "tcp://192.168.1.72:1883"
	options := mqtt.NewClientOptions()
	options.AddBroker(broker)
	options.SetClientID("go_mqtt_example")
	options.SetDefaultPublishHandler(messagePubHandler)
	options.OnConnect = connectHandler
	options.OnConnectionLost = connectionLostHandler

	client := mqtt.NewClient(options)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// letters := []string{"a", "b", "c", "d"}

	for _, x := range subs.Items {
		topic := x.Topic
		token = client.Subscribe(topic, 1, nil)
		token.Wait()
		fmt.Printf("Subscribed to topic %s\n", topic)
	}

	for true {

	}

	// num := 10
	// for i := 0; i < num; i++ {
	// 	text := fmt.Sprintf("%d", i)
	// 	token = client.Publish(topic, 0, false, text)
	// 	token.Wait()
	// 	time.Sleep(time.Second)
	// }

	client.Disconnect(100)
}

func UpdateData(collName string, recordID string, jsonData []byte) error {
	httpposturl := fmt.Sprintf("http://192.168.1.72:8080/api/collections/%s/records/%s", collName, recordID)

	request, error := http.NewRequest("PATCH", httpposturl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		return error
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)

	return nil
}

type ResponseData struct {
	Page       int
	PerPage    int
	TotalItems int
	TotalPages int
	Items      []Subscribtion
}

type Item struct {
	Id    string
	Title string
	Value string
	Topic string
}

func GetData(collID string) (error, ResponseData) {
	requestURL := fmt.Sprintf("http://192.168.1.72:8080/api/collections/%s/records", collID)
	res, err := http.Get(requestURL)

	var jsonData ResponseData

	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		return err, jsonData
	}

	fmt.Printf("client: status code: %d\n", res.StatusCode)
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return err, jsonData
	}

	err = json.Unmarshal([]byte(resBody), &jsonData)

	if err != nil {
		fmt.Println("HERE")
		panic(err)
	}

	return nil, jsonData
}
