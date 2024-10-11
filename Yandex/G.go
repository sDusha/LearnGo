package main

import (
	"fmt"
	"math"
)

type LogLevel string

const (
	Error LogLevel = "ERROR"
	Info  LogLevel = "INFO"
)

type Logger interface {
	Log(message string)
}

type Log struct {
	Level LogLevel
}

func (l Log) Log(message string) {
	switch l.Level {
	case Error:
		fmt.Printf("ERROR: %s\n", message)
	case Info:
		fmt.Printf("INFO: %s\n", message)
	default:
		fmt.Printf("UNKNOWN LEVEL: %s\n", message)
	}
}

type Animal interface {
	MakeSound()
}

func MakeSound(a Animal) {
	a.MakeSound()
}

type Dog struct {
}

func (Dog) MakeSound() {
	fmt.Println("Гав!")
}

type Cat struct {
}

func (Cat) MakeSound() {
	fmt.Println("Мяу!")
}

type Shape interface {
	Area() float64
}

func Area(s Shape) float64 {
	return s.Area()
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	// интерфейсы
	cat1 := Cat{}
	cat1.MakeSound()
	dog1 := Dog{}
	dog1.MakeSound()
}
