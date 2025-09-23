package events

type EventObserver interface {
	OnNewEvent(event, took string)
}

type EventStore struct {
	Events         []string
	EventObservers []EventObserver
}

func NewEventStore() *EventStore {
	return &EventStore{}
}

func (es *EventStore) NewEvent(event, took string) {
	es.Events = append(es.Events, event+" "+took)

	for _, v := range es.EventObservers {
		v.OnNewEvent(event, took)
	}
}

func (es *EventStore) AddNewEventObserver(o EventObserver) {
	es.EventObservers = append(es.EventObservers, o)
}
