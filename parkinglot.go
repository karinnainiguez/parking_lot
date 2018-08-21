package main

import (
	"time"
)

type space struct {
	id          int
	special     string
	reservation reservation
}

type reservation struct {
	customerName string
	startDate    time.Time
	endDate      time.Time
}

func main() {

}
