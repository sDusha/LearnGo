package main

import (
	"fmt"
	"math/rand"
)

/*
При планировании поездки на Марс будет удобно собрать расценки различных космических станций в одном месте.
Есть множество сайтов для авиалиний, но не для космических. Для нас это не будет проблемой. При умелом руководстве,
Go сможет решить проблемы подобного рода.

В таблице четыре столбца:

Космическая станция (Spaceline), что предоставляет услуги;
Продолжительность (Duration) в днях поездки на Марс в один конец;
Покрывает ли цена поездку туда и обратно (Trip type);
Цена (Price) в миллионах долларов.
Для каждого билета случайным образом выбирается космическая станция: Space Adventures, SpaceX или Virgin Galactic.

Датой отправления на каждом билете значится 13 Октября 2020 года.
В этот день Марс будет на расстоянии 62 100 000 км от Земли.

Скорость космического корабля будет выбрана случайным образом из диапазона от 16 до 30 км/ч.
Это определит продолжительность поездки на Марс, а также цену билета. Более быстрые корабли намного дороже.
Цены на билеты варьируются от $36 до $50 миллионов. Цена для поездки туда-обратно удваивается.
*/

const secondsPerDay = 86400

func main() {
	distance := 62100000
	company := ""
	trip := ""

	fmt.Println("Spaceline        Days Trip type  Price")
	fmt.Println("======================================")

	for count := 0; count < 10; count++ {
		switch rand.Intn(3) {
		case 0:
			company = "Space Adventures"
		case 1:
			company = "SpaceX"
		case 2:
			company = "Virgin Galactic"
		}

		speed := rand.Intn(15) + 16                  // 16-30 km/s
		duration := distance / speed / secondsPerDay // days
		price := 20.0 + speed                        // millions

		if rand.Intn(2) == 1 {
			trip = "Round-trip"
			price = price * 2
		} else {
			trip = "One-way"
		}

		fmt.Printf("%-16v %4v %-10v $%4v\n", company, duration, trip, price)
	}
}
