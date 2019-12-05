package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func calculateModuleFuel(moduleFuel int, mass int) int {
	tmpFuel := mass/3 - 2
	if tmpFuel <= 0 {
		return moduleFuel
	}
	moduleFuel += tmpFuel
	return calculateModuleFuel(moduleFuel, tmpFuel)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fuel := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		fuel += calculateModuleFuel(0, mass)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Fuel required: ", fuel)
}
