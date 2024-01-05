package tickets

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	ID                 uint32
	Name               string
	Email              string
	DestinationCountry string
	FlightTime         string
	Price              uint32
}

type TimePeriod int

const (
	Madrugada TimePeriod = iota
	Mañana
	Tarde
	Noche
)

var TicketCache []Ticket
var ErrorBadFormat = errors.New("the csv file has format errors")
var ErrorBadTimePeriod = errors.New("the time period used is not known")

// Dumps the CSV data on a slice for quick access
func dumpData() ([]Ticket, error) {
	var data []Ticket
	absolutePath := "/Users/jdoffo/Desktop/Practica Bootcamp/Bootcamp-GO/Clase 06/desafio-go-bases-main/tickets.csv"
	file, err := os.Open(absolutePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		newTicket, err := parseStringToTicket(scanner.Text())
		if err != nil {
			return nil, err
		}
		data = append(data, newTicket)
	}
	return data, nil
}

// Convert the given string into a Ticket
func parseStringToTicket(str string) (Ticket, error) {
	fields := strings.Split(str, ",")
	if len(fields) != 6 {
		return Ticket{}, ErrorBadFormat
	}
	id, err := strconv.Atoi(fields[0])
	if err != nil {
		return Ticket{}, ErrorBadFormat
	}
	name := fields[1]
	email := fields[2]
	dstCountry := fields[3]
	time := fields[4]
	price, err := strconv.Atoi(fields[5])
	if err != nil {
		return Ticket{}, ErrorBadFormat
	}
	newTicket := Ticket{uint32(id), name, email, dstCountry, time, uint32(price)}
	return newTicket, nil
}

func InitializeCache() {
	var err error
	TicketCache, err = dumpData()
	if err != nil {
		panic(err)
	}
}

// ejemplo 1
func GetTotalTickets(destination string) int {
	var counter int
	for _, ticket := range TicketCache {
		if destination == ticket.DestinationCountry {
			counter++
		}
	}
	return int(counter)
}

func TicketFlyTimePeriod(t Ticket) (TimePeriod, error) {
	ticketFlyHour, err := strconv.Atoi(strings.Split(t.FlightTime, ":")[0])
	if err != nil {
		return -1, err
	}
	switch {
	case 0 <= ticketFlyHour && ticketFlyHour <= 6:
		return Madrugada, nil
	case 7 <= ticketFlyHour && ticketFlyHour <= 12:
		return Mañana, nil
	case 13 <= ticketFlyHour && ticketFlyHour <= 19:
		return Tarde, nil
	case 20 <= ticketFlyHour && ticketFlyHour <= 23:
		return Noche, nil
	default:
		return -1, ErrorBadTimePeriod
	}
}

// ejemplo 2
func FlyCountByTimePeriod(time TimePeriod) (int, error) {
	var count int = 0
	for _, ticket := range TicketCache {
		ticketPeriod, err := TicketFlyTimePeriod(ticket)
		if err != nil {
			return 0, err
		}
		if ticketPeriod == time {
			count++
		}
	}
	return count, nil
}

// ejemplo 3
func AverageDestination(destination string) float64 {
	destinationFlyCount := GetTotalTickets(destination)
	return float64(destinationFlyCount) / float64(len(TicketCache)) * 100.0
}
