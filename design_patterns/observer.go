package main

import "fmt"

//============Topic
type Topic interface {
	register(observer Observer)
	broadcast()
}

type Item struct {
	observers []Observer
	name      string
	available bool
}

func (i *Item) broadcast() {
	for _, observer := range i.observers {
		observer.updateValue(i.name)
		fmt.Printf("The observer %s has been notified\n", observer.getId())
	}
}

func (i *Item) register(observer Observer) {
	i.observers = append(i.observers, observer)
}

func (i *Item) UpdateAvailable() {
	fmt.Printf("The item %s change\n", i.name)
	i.available = true
	i.broadcast()
}

func NewItem(name string) *Item {
	return &Item{
		name:      name,
		available: false,
	}
}

//=========== Observer
type Observer interface {
	getId() string
	updateValue(string)
}

type EmailClient struct {
	id string
}

func (ec *EmailClient) getId() string {
	return ec.id
}

func (ec *EmailClient) updateValue(itemName string) {
	fmt.Printf("The item %s was modified\n", itemName)
}

//=======main
func main() {
	graphicCard := NewItem("RTX 60000")

	firstObserver := &EmailClient{
		id: "123A",
	}

	secondObserver := &EmailClient{
		id: "456B",
	}

	graphicCard.register(firstObserver)
	graphicCard.register(secondObserver)

	graphicCard.UpdateAvailable()
}
