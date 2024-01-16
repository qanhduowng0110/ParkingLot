package main

import (
	"fmt"
	"time"
)

type ParkingLot struct {
	Size       int
	Occupied   map[int]Car
	NextSlot   int
	ChargeRate int
}

type Car struct {
	Number     string
	EntryTime  time.Time
	ParkingSlot int
}

func NewParkingLot(size, chargeRate int) *ParkingLot {
	return &ParkingLot{
		Size:       size,
		Occupied:   make(map[int]Car),
		NextSlot:   1,
		ChargeRate: chargeRate,
	}
}

func (pl *ParkingLot) Park(carNumber string) {
	if len(pl.Occupied) == pl.Size {
		fmt.Println("Sorry, parking lot is full.")
		return
	}

	entryTime := time.Now()
	slot := pl.NextSlot
	pl.Occupied[slot] = Car{Number: carNumber, EntryTime: entryTime, ParkingSlot: slot}
	pl.NextSlot++

	fmt.Printf("Allocated slot number: %d\n", slot)
}

func (pl *ParkingLot) Leave(carNumber string) {
	for slot, car := range pl.Occupied {
		if car.Number == carNumber {
			elapsedHours := time.Since(car.EntryTime).Hours()
			charge := pl.calculateCharge(elapsedHours)

			fmt.Printf("Registration number %s with Slot Number %d is free with Charge $%d\n", car.Number, slot, charge)

			delete(pl.Occupied, slot)
			return
		}
	}

	fmt.Printf("Registration number %s not found in the parking lot.\n", carNumber)
}

func (pl *ParkingLot) Status(carNumber string) {
	for _, car := range pl.Occupied {
		if car.Number == carNumber {
			elapsedHours := time.Since(car.EntryTime).Hours()
			charge := pl.calculateCharge(elapsedHours)

			fmt.Printf("Slot number: %d\n", car.ParkingSlot)
			fmt.Printf("Registration number: %s\n", car.Number)
			fmt.Printf("Parking time: %.2f hours\n", elapsedHours)
			fmt.Printf("Charges: $%d\n", charge)
			return
		}
	}

	fmt.Printf("Registration number %s not found in the parking lot.\n", carNumber)
}

func (pl *ParkingLot) calculateCharge(elapsedHours float64) int {
	// Minimum charge for 2 hours
	charge := pl.ChargeRate
	if elapsedHours > 2 {
		// Additional charge for each additional hour
		charge += pl.ChargeRate * int(elapsedHours-2)
	}

	return charge
}

func main() {
	var parkingLot *ParkingLot

	for {
		var command string
		fmt.Print("Enter command: ")
		fmt.Scan(&command)

		switch command {
		case "create":
			var size int
			fmt.Print("Enter parking lot size: ")
			fmt.Scan(&size)
			parkingLot = NewParkingLot(size, 10)
			fmt.Printf("Parking lot with size %d created.\n", size)

		case "park":
			var carNumber string
			fmt.Print("Enter car number: ")
			fmt.Scan(&carNumber)
			if parkingLot == nil {
				fmt.Println("Please create a parking lot first.")
			} else {
				parkingLot.Park(carNumber)
			}

		case "leave":
			var carNumber string
			fmt.Print("Enter car number: ")
			fmt.Scan(&carNumber)
			if parkingLot == nil {
				fmt.Println("Please create a parking lot first.")
			} else {
				parkingLot.Leave(carNumber)
			}

		case "status":
			var carNumber string
			fmt.Print("Enter car number: ")
			fmt.Scan(&carNumber)
			if parkingLot == nil {
				fmt.Println("Please create a parking lot first.")
			} else {
				parkingLot.Status(carNumber)
			}

		case "exit":
			fmt.Println("Exiting program.")
			return

		default:
			fmt.Println("Invalid command. Try again.")
		}
	}
}
