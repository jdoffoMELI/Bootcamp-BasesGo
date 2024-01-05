package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	tickets.InitializeCache()
	fmt.Println(tickets.TicketCache[0])
	total := tickets.GetTotalTickets("Brazil")
	fmt.Println(total)

	vuelosMañana, err := tickets.FlyCountByTimePeriod(tickets.Noche)
	if err != nil {
		panic(err)
	}
	fmt.Println("Vuelos de la mañana:", vuelosMañana)

}
