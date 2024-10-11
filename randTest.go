package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	var randInt = rand.Intn(100) + 22
	fmt.Println(randInt)

	var command = "выйти наружу"
	var exit = strings.Contains(command, "наружу")

	fmt.Println("Вы покидаете пещеру:", exit)

	var array_ = [...]string{"идти на восток", "зайти внутрь", "another", "прочитать знак", "зайти в пещеру"}
	command = array_[rand.Intn(3)]

	if command == "идти на восток" { // если command равен "идти на восток"
		fmt.Println("Вы направляетесь к горе.")
	} else if command == "зайти внутрь" { // в противном случае, если command равен "зайти внутрь"
		fmt.Println("Вы заходите в пещеру, где будете жить до конца своей жизни.")
	} else { // или что-то другое
		fmt.Println("Пока не совсем понятно.")
	}

	fmt.Println("Здесь вход в пещеру и путь на восток.")

	switch command { // Сравнивает case с command
	case "идти на восток":
		fmt.Println("Вы направляетесь к горе.")
	case "зайти в пещеру", "зайти внутрь": // Запятая разделяет список возможных значений
		fmt.Println("Вы находитесь в тускло освещенной пещере.")
		fallthrough // переход на следующий случай
	case "прочитать знак":
		fmt.Println("На знаке написано 'Несовершеннолетним вход запрещен'.")
	default:
		fmt.Println("Пока не совсем понятно.")
	}

	var count = 2 // Объявление и инициализация

	for count > 0 { // Условие
		fmt.Println(count)
		time.Sleep(time.Second)
		count-- // Обратный отсчет; в противном случае цикл будет длиться вечно
	}
	fmt.Println("Запуск!")

	var today = time.Now().Weekday()
	fmt.Println(today)
	for i := 1; i < 10; i++ {
		fmt.Println(i * i)
	}

	// краткое обявление
	if num := rand.Intn(3); num == 0 {
		fmt.Println("Space Adventures")
	} else if num == 1 {
		fmt.Println("SpaceX")
	} else {
		fmt.Println("Virgin Galactic")
	} // num больше не в области видимости
}
