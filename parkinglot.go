package main

import (
	"fmt"
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

type lot []space

func main() {
	newlot := lot{}
	fmt.Println(newlot)
}
