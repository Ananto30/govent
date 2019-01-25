# govent
Event driven architecture for Go using channel

## Usage
Currently usage can be found in the example folder. Here is the easy way - 

Register new event types - 
```go
const (
	MyMessage govent.EventType = iota
)
```

Make new events -
```go
type MessageEvent struct {
	Message string
}
```

Make a consumer function for that event -
```go
func ShowMessage(e govent.Event) {
	if e, ok := e.(MessageEvent); ok {
		fmt.Println(e.Message)
	}
}
```

Start a new Event publisher - 
```go
publisher := govent.NewEventPublisher()
```

Subscribe events to the Event publisher - 
```go
govent.Subscribe(ShowMessage, MyMessage)
```

Publish events - 
```go
publisher <- govent.EventObject{
		EventType: MyMessage,
		Event:     MessageEvent{Message: "Hello World"},
	}
```

## Run
```
cd example
go run *.go
```
