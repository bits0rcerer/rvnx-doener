package log

import (
	"log"
	"rvnx_doener_service/ent"
)

type EventLogger interface {
	// Handle handles a new Event
	Handle(event *ent.Event)
}

type ConsoleEventLogger struct {
}

func (c ConsoleEventLogger) Handle(event *ent.Event) {
	log.Printf("[Event: %s] %s", event.EventType.String(), event.String())
}
