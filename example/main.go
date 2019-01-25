package main

import (
	"fmt"
	"github/ananto/govent"
)

func main() {

	publisher := govent.NewEventPublisher()

	govent.Subscribe(ShowMessage, MyMessage)
	govent.Subscribe(ShowWeather, MyWeather)

	publisher <- govent.EventObject{
		EventType: MyMessage,
		Event:     MessageEvent{Message: "Hello World"},
	}
	publisher <- govent.EventObject{
		EventType: MyWeather,
		Event:     WeatherEvent{Condition: "Sunny"},
	}
	publisher <- govent.EventObject{
		EventType: MyWeather,
		Event:     MessageEvent{Message: "Hello World"},
	}
	publisher <- govent.EventObject{
		EventType: MyMessage,
		Event:     WeatherEvent{Condition: "Rainy"},
	}

	for {

	}

}


// ShowMessage prints the message
func ShowMessage(e govent.Event) {
	if e, ok := e.(MessageEvent); ok {
		fmt.Println(e.Message)
	}
}


// ShowWeather prints the weather condition
func ShowWeather(e govent.Event) {
	if e, ok := e.(WeatherEvent); ok {
		fmt.Println(e.Condition)
	}
}