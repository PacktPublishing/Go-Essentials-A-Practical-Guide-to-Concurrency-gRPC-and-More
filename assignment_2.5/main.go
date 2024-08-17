// In this assignment, you will work with custom defined types, interfaces, and error handling in Go by simulating a simple vehicle management system. You will create three vehicle types: Car, Boat, and Motorcycle. Each vehicle will have different properties and methods. Additionally, you will implement an interface to manage vehicles, and handle errors gracefully.

// Step 1: Define Custom Types
// 1.	Create a Vehicle struct with common properties like Brand, Model, Year, and Color.
// 2.	Define three custom types: Car, Boat, and Motorcycle, each embedding the Vehicle struct. Add specific properties to each type. For example:
// •	Car: NumDoors, EngineType.
// •	Boat: Length, PropulsionType.
// •	Motorcycle: NumWheels, HasSidecar.

// Step 2: Create an Interface
// 1.	Declare an interface VehicleInterface with methods that all vehicle types must implement such as Start(), Stop(), Steer()

// Step 3: Implement Error Handling
// 1.	Define a custom error type VehicleError that implements the error interface and includes additional information about the error.
// 2.	Handle potential errors in adding a vehicle by returning the VehicleError when applicable

// Step 4: Create a Management System
// 1.	Implement a function to add new vehicles to the system and return the newly added vehicle.
// 2.	Write a function to start a vehicle and handle any errors that may occur.
// 3.	Implement a function to stop a vehicle.

// Step 5: Test Your Implementation
// 1.	Create a few vehicles of each type and add them to the vehicle management system.
// 2.	Test the functionality of starting, stopping.

package main

import (
	"errors"
	"fmt"
)

type Vehicle struct {
	Brand string
	Model string
	Year int
	Color string
}

type VehicleError struct {
	VehicleType string
	Err error
}

func (ve VehicleError) Error() string {
	return fmt.Sprintf("%s vehicle error: %v", ve.VehicleType, ve.Err)
}

type VehicleInterface interface {
	Start()
	Stop()
	Steer()
}

type Car struct {
	Vehicle
	NumDoors int
	EngineType string
}

type Boat struct {
	Vehicle
	Length int
	PropulsionType string
}

type Motorcycle struct {
	Vehicle
	NumWheels int
	HasSidecar bool
}

func (c Car) Start() {
    fmt.Printf("Starting the %s %s with %d doors\n", c.Color, c.Brand, c.NumDoors)
}

func (c Car) Stop() {
    fmt.Printf("Stopping the %s %s with %d doors\n", c.Color, c.Brand, c.NumDoors)
}

func (c Car) Steer() {
    fmt.Printf("Steering the %s %s with %d doors\n", c.Color, c.Brand, c.NumDoors)
}

func (b Boat) Start() {
    fmt.Printf("Starting the %s boat with length %d\n", b.Brand, b.Length)
}

func (b Boat) Stop() {
    fmt.Printf("Stopping the %s boat with length %d\n", b.Brand, b.Length)
}

func (b Boat) Steer() {
    fmt.Printf("Steering the %s boat with length %d\n", b.Brand, b.Length)
}

func (m Motorcycle) Start() {
    fmt.Printf("Starting the %s motorcycle with %d wheels\n", m.Brand, m.NumWheels)
    if m.NumWheels < 0 {
        panic("negative number of wheels")
    }
}

func (m Motorcycle) Stop() {
    fmt.Printf("Stopping the %s motorcycle with %d wheels\n", m.Brand, m.NumWheels)
}

func (m Motorcycle) Steer() {
    fmt.Printf("Steering the %s motorcycle with %d wheels\n", m.Brand, m.NumWheels)
}

func AddVehicle(vehicleType string, brand string, model string, year int, color string) (VehicleInterface, error){
	switch vehicleType {
	case "car":
		return Car{Vehicle{Brand: brand, Model: model, Year: year, Color: color}, 4, "gasoline"}, nil
	case "boat":
        return Boat{Vehicle{Brand: brand, Model: model, Year: year, Color: color}, 20, "motor"}, nil
    // case "motorcycle":
    //  return Motorcycle{Vehicle{Brand: brand, Model: model, Year: year, Color: color}, 2, false}, nil
    case "motorcycle":
        return Motorcycle{Vehicle{Brand: brand, Model: model, Year: year, Color: color}, -1, false}, nil //simulates an error
    default:
        return nil, VehicleError{VehicleType: "Unknown", Err: errors.New("invalid vehicle type")}
    }
}

func VehicleAbilities(v VehicleInterface){
	v.Start()
	v.Steer()
	v.Stop()
}

func doRecover(){
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

func main() {
	defer doRecover()

	car, err := AddVehicle("car", "Toyota", "Camry", 2020, "blue")
    if err != nil {
        fmt.Println("Error adding car:", err)
        return
    }

    boat, err := AddVehicle("boat", "Bayliner", "Element", 2018, "white")
    if err != nil {
        fmt.Println("Error adding boat:", err)
        return
    }

    motorcycle, err := AddVehicle("motorcycle", "Harley-Davidson", "Street Glide", 2022, "black")
    if err != nil {
        fmt.Println("Error adding motorcycle:", err)
        return
    }

    VehicleAbilities(car)
    VehicleAbilities(boat)
    VehicleAbilities(motorcycle)

}
