package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Применяем рекурсивную функцию для удаления пустых списков и словарей
func prettify(v interface{}) interface{} {
	switch value := v.(type) {
	case map[string]interface{}:
		// Для каждого ключа в словаре
		for k, v := range value {
			// Рекурсивно обрабатываем значение
			value[k] = prettify(v)
			// Если значение после обработки пустое, удаляем ключ
			if isEmpty(value[k]) {
				delete(value, k)
			}
		}
	case []interface{}:
		// Обрабатываем элементы списка
		var newSlice []interface{}
		for _, item := range value {
			item = prettify(item)
			// Если элемент не пустой, добавляем его в новый список
			if !isEmpty(item) {
				newSlice = append(newSlice, item)
			}
		}
		return newSlice
	}
	return v
}

// Проверка, является ли объект пустым (пустой словарь или список)
func isEmpty(v interface{}) bool {
	switch value := v.(type) {
	case map[string]interface{}:
		return len(value) == 0
	case []interface{}:
		return len(value) == 0
	}
	return false
}

func main() {
	var t int
	// Чтение количества наборов входных данных
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n int
		// Чтение количества строк JSON-объекта
		fmt.Scan(&n)

		// Чтение JSON-объекта
		var jsonData string
		for j := 0; j < n; j++ {
			var line string
			fmt.Scan(&line)
			jsonData += line
		}

		// Декодирование JSON
		var obj interface{}
		decoder := json.NewDecoder(bytes.NewBufferString(jsonData))
		decoder.UseNumber() // Используем для правильного декодирования чисел
		err := decoder.Decode(&obj)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}

		// Применяем операцию prettify
		result := prettify(obj)

		// Кодируем результат в JSON и выводим
		output, err := json.Marshal(result)
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}

		fmt.Println(string(output))
	}
}
