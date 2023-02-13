package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	population := os.Args[1]
	data, err := os.ReadFile(population)
	for index, value := range data {
		if (value >= 65 && value <= 90) || (value >= 97 && value <= 122) {
			data = append(data[:index], data[index+1:]...)
		}
	}
	if err != nil {
		panic(err)
	}
	//fmt.Println(data)
	tranche := strings.Fields(string(data))
	//fmt.Println(tranche)
	tab := make([]float64, len(tranche))
	for i, x := range tranche {

		tab[i], _ = strconv.ParseFloat(x, 64)
	}
	sort.Float64s(tab) // tri le tableau de façon croissant
	//fmt.Println(tab)
	var sum float64
	// La Moyenne
	for _, value := range tab {
		sum += value
	}
	a := sum / float64(len(tab))
	fmt.Println("Average:", int(math.Round(a)))
	// La médiane
	if len(tab)%2 == 0 {
		middle := len(tab) / 2
		middle1 := (tab[middle-1] + tab[middle]) / 2
		fmt.Println("Median:", int(math.Round(middle1)))
	} else {
		middle := len(tab) / 2
		middle1 := tab[middle]
		fmt.Println("Median:", int(math.Round(middle1)))
	}
	// La variance
	var variance float64
	for _, value := range tab {
		variance += math.Pow((value-a), 2) / float64(len(tab))
	}
	//b := variance / float64(len(tab))
	fmt.Println("Variance:", int(math.Round(variance)))

	// L'écart-type
	Standard_Deviation := math.Sqrt(variance)
	fmt.Println("Standard Deviation:", int(math.Round(Standard_Deviation)))
}
