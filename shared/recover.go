package shared

import (
	"log"
)

func Recover() {
	if r := recover(); r != nil {
		// Log the panic or handle it as needed,
		// For example, you can log it to a file or send it to an error tracking service
		log.Println("Recovered from panic:", r)
	}
}
