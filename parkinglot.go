package main

import (
	"fmt"
	"time"
)

type space struct {
	ADA         bool
	reservation reservation
}

type reservation struct {
	customerName string
	startDate    time.Time
	endDate      time.Time
}

type lot map[int]space

func main() {
	fmt.Println("Creating new lot with three regular spots, and three ADA spots...")
	l1, rsp := newLot(3, 3)
	fmt.Println(rsp)
	l1.print()

}

func newLot(regular int, ADA int) (lot, string) {
	lot := lot{}
	lot.addSpaces(regular, ADA)
	return lot, fmt.Sprintf("Successfully created lot with %v spaces", len(lot))
}

func (l lot) print() {
	for id, space := range l {
		fmt.Printf("Space with id # %v\n", id)
		fmt.Printf("ADA compliant: %v\n", space.ADA)
		if (reservation{}) == space.reservation {
			fmt.Println("NOT currently Reserved")
		} else {
			fmt.Printf("Currently Reserved from %v to %v", space.reservation.startDate, space.reservation.endDate)
		}
	}
}

func (l lot) addSpaces(regular int, ADA int) string {
	added := 0

	for i := 0; i < regular; i++ {
		nextID := len(l) + 1
		newSpace := space{}
		l[nextID] = newSpace
	}

	for j := 0; j < ADA; j++ {
		nextID := len(l) + 1
		newSpace := space{ADA: true}
		l[nextID] = newSpace
	}
	return fmt.Sprintf("Successfully added %v spaces.  Lot now has %v total spaces", added, len(l))
}
