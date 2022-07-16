package log

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"rvnx_doener_service/ent"
	"rvnx_doener_service/ent/event"
	"testing"
	"time"
)

type TestEventLogger struct {
	t      *testing.T
	events chan *ent.Event
	closed bool
}

func (l *TestEventLogger) Handle(event *ent.Event) {
	if !l.closed {
		go func() {
			defer func() {
				// recover send on closed channel panics
				recovered := recover()

				if recovered != nil {
					// tag channel as closed
					l.closed = true
				}
			}()

			l.events <- event
		}()
	}
}

func (l *TestEventLogger) WaitUntil(
	eventType event.EventType,
	timeout time.Duration,
	assertion func(t *testing.T, event ent.Event)) {
	l.t.Helper()

	for {
		select {
		case <-time.After(timeout):
			assert.Fail(l.t, fmt.Sprintf("timed out before receiving %s", eventType.String()))
			return
		case e := <-l.events:
			if e == nil {
				assert.Fail(l.t, fmt.Sprintf("did not receive %s", eventType.String()))
				return
			}

			if e.EventType != eventType {
				continue
			}

			if assertion != nil {
				assertion(l.t, *e)
			}

			return
		}
	}
}

func (l *TestEventLogger) Close() {
	l.closed = true
	close(l.events)
}

func NewTestEventLogger(t *testing.T) *TestEventLogger {
	return &TestEventLogger{
		t:      t,
		events: make(chan *ent.Event),
		closed: false,
	}
}
