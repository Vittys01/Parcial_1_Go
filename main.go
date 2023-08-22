package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	total, err := tickets.GetTotalTickets("Brazil")
	fmt.Println(total, err)
	total2, err := tickets.GetMornings("tarde")
	fmt.Println(total2, err)
}
