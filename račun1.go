package main

import (
	"fmt"
	"math"
)

func izracunajHipotenuzu(stranicaA, stranicaB float64) float64 {
	hipotenuza := math.Sqrt(stranicaA*stranicaA + stranicaB*stranicaB)
	return hipotenuza
}

func izracunajKuteve(stranicaA, stranicaB, hipotenuza float64) []float64 {
	sinA := stranicaA / hipotenuza
	sinB := stranicaB / hipotenuza

	cosA := stranicaB / hipotenuza
	cosB := stranicaA / hipotenuza

	tanA := stranicaA / stranicaB
	tanB := stranicaB / stranicaA

	cotA := stranicaB / stranicaA
	cotB := stranicaA / stranicaB

	kutevi := []float64{sinA, sinB, cosA, cosB, tanA, tanB, cotA, cotB}
	return kutevi
}

func main() {
	var stranicaA, stranicaB float64
	fmt.Print("Unesite duljinu prve stranice: ")
	fmt.Scan(&stranicaA)
	fmt.Print("Unesite duljinu druge stranice: ")
	fmt.Scan(&stranicaB)

	hipotenuza := izracunajHipotenuzu(stranicaA, stranicaB)
	kutevi := izracunajKuteve(stranicaA, stranicaB, hipotenuza)

	fmt.Println("Hipotenuza: ", hipotenuza)
	fmt.Println("Sinusi kuteva (A i B): ", kutevi[0:2])
	fmt.Println("Kosinusi kuteva (A i B): ", kutevi[2:4])
	fmt.Println("Tangensi kuteva (A i B): ", kutevi[4:6])
	fmt.Println("Kotangensi kuteva (A i B): ", kutevi[6:8])
}
