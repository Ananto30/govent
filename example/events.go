package main

import "github/ananto/govent"

// my event types, registering to govent
const (
	MyMessage govent.EventType = iota
	MyWeather govent.EventType = iota
)

// MessageEvent is an event type for messaging
type MessageEvent struct {
	Message string
}

// WeatherEvent is an event type for weather conditions
type WeatherEvent struct {
	Condition string
}
