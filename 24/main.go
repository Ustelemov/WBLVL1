/*Разработать программу нахождения расстояния междудвумя точками,
которые представлены в виде структуры
Point с инкапсулированными параметрами x,y и конструктором. */

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

//Геттер для инкапсулированного поля х
func (p Point) getX() float64 {
	return p.x
}

//Геттер для инкапсулированного поля y
func (p Point) getY() float64 {
	return p.y
}

//Реализуем интерфейс Stringer для красивого вывода
func (p Point) String() string {
	return fmt.Sprintf("{x: %0.1f, y: %0.1f}", p.x, p.y)
}

func findDist(p1 Point, p2 Point) float64 {
	x1, y1 := p1.getX(), p1.getY()
	x2, y2 := p2.getX(), p2.getY()

	return math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))
}

func main() {
	p1 := NewPoint(3, 4)
	p2 := NewPoint(0, 0)

	fmt.Printf("p1: %v, p2: %v, dist(p1,p2): %.1f\n", p1, p2, findDist(*p1, *p2))

}
