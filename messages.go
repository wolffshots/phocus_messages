// Package phocus_messages contains the
// various queries and commands that
// that can be sent with phocus
package phocus_messages

import (
	"github.com/google/uuid"
	"log"
)

// Message is the shape of a message for phocus to interpret and handle queuing of
type Message struct {
	ID      uuid.UUID `json:"id"`
	Command string    `json:"command"`
	Payload string    `json:"payload"`
}

// Interpret converts the generic `phocus` message into a specific inverter message
// TODO add even more generalisation and separated implementation details here
func Interpret(input Message) error {
	switch input.Command {
	case "QPGSn":
		// TODO pass in inverter number
		errOne := HandleQPGS(1)
		if errOne != nil {
			log.Printf("Failed to handle QPGS1 :%v\n", errOne)
		}
		errTwo := HandleQPGS(2)
		if errTwo != nil {
			log.Printf("Failed to handle QPGS2 :%v\n", errTwo)
        }
		if errOne != nil {
            return errOne
		} else {
            return errTwo
		}
	case "QID":
		log.Println("TODO send QID")
	default:
		log.Println("Unexpected message on queue")
	}
	return nil
}

// Command interface is a WIP
type Command interface {
	New()
	Print()
}
