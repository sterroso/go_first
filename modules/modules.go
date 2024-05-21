package modules

import "fmt"

type Ticket struct {
	ID    int
	Event string
}

func (t *Ticket) PrintEvent() {
	fmt.Printf("Wellcome to '%s'!\n", t.Event)
}
