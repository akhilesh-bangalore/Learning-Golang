package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	plants := []PowerPlant{
		PowerPlant{Hydro, 300, Active},
		PowerPlant{Wind, 30, Active},
		PowerPlant{Wind, 25, Inactive},
		PowerPlant{Wind, 35, Active},
		PowerPlant{Solar, 45, Unavailable},
		PowerPlant{Solar, 40, Inactive},
	}

	grid := PowerGrid{300, plants}

	if option, err := requestOption(); err == nil {
		fmt.Println("")

		switch option {
		case "1":
			grid.generatePlantSummaryReport()
		case "2":
			grid.generateGridCapacityReport()
		}
	} else {
		fmt.Println(err.Error())
	}
}

func requestOption() (option string, err error) {
	fmt.Println("1) Generate Power Plant Report")
	fmt.Println("2) Generate Power Grid Report")
	fmt.Print("Please choose an option: ")

	fmt.Scanln(&option)

	if option != "1" && option != "2" {
		err = errors.New("Invalid option selected")
	}
	return
}

type PlantType string

const (
	Hydro    PlantType = "Hydro"
	Thermal  PlantType = "Thermal"
	Wind     PlantType = "Wind"
	Solar    PlantType = "Solar"
	Chemical PlantType = "Chemical"
)

type PlantStatus string

const (
	Active      PlantStatus = "Active"
	Inactive    PlantStatus = "Inactive"
	Unavailable PlantStatus = "Unavailable"
)

type PowerPlant struct {
	Type     PlantType
	Capacity float64
	Status   PlantStatus
}

type PowerGrid struct {
	Load   float64
	Plants []PowerPlant
}

func (pg *PowerGrid) generatePlantSummaryReport() {
	for idx, p := range pg.Plants {
		heading := fmt.Sprintf("%s%d", "Plant #", idx)
		fmt.Println(heading)
		fmt.Println(strings.Repeat("-", len(heading)))

		fmt.Printf("%-20s%s\n", "Type:", p.Type)
		fmt.Printf("%-20s%.2f\n", "Capacity:", p.Capacity)
		fmt.Printf("%-20s%s\n", "Status:", p.Status)

		fmt.Println("")
	}
}

func (pg *PowerGrid) generateGridCapacityReport() {
	capacity := 0.

	for _, p := range pg.Plants {
		if p.Status == Active {
			capacity += p.Capacity
		}
	}

	heading := fmt.Sprintf("%s", "Power Grid Capcity")
	fmt.Println(heading)
	fmt.Println(strings.Repeat("-", len(heading)))

	fmt.Printf("%-20s%.2f\n", "Capacity:", capacity)
	fmt.Printf("%-20s%.2f\n", "Load:", pg.Load)
	fmt.Printf("%-20s%.2f%%\n", "Utilization:", pg.Load/capacity*100)
}
