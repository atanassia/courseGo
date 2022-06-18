package main

import "math"
import "fmt"

func main(){
	fmt.Println(AreaOfPolygonInsideCircle(3,3))
}

const PI = 3.141592653589793;

func AreaOfPolygonInsideCircle(circleRadius float64, numberOfSides int) float64 {
	var numberOfSides2 = float64(numberOfSides)
	return math.Round((0.5 * math.Pow(circleRadius,2) * numberOfSides2 * math.Sin(360/numberOfSides2*(PI/180)))*1000)/1000
}