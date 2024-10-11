package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// findRoot находит корень дерева на основе кодировки
func findRoot(treeCode []int) int {
	allVertices := make(map[int]bool)
	children := make(map[int]bool)

	index := 0
	for index < len(treeCode) {
		vertex := treeCode[index]
		allVertices[vertex] = true

		degree := treeCode[index+1]
		for i := 0; i < degree; i++ {
			child := treeCode[index+2+i]
			children[child] = true
		}

		// Переход к следующему блоку
		index += 2 + degree
	}

	// Найти вершину, которая не является ребенком
	for vertex := range allVertices {
		if !children[vertex] {
			return vertex
		}
	}

	return -1 // Если корень не найден
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	testCount, _ := strconv.Atoi(line)

	for i := 0; i < testCount; i++ {
		// Чтение длины кодировки (можно игнорировать)
		_, _ = reader.ReadString('\n')

		// Чтение кодировки
		line, _ = reader.ReadString('\n')
		line = strings.TrimSpace(line)

		// Преобразование строки в массив чисел
		strs := strings.Split(line, " ")
		treeCode := make([]int, len(strs))
		for j, s := range strs {
			treeCode[j], _ = strconv.Atoi(s)
		}

		// Нахождение корня и вывод результата
		root := findRoot(treeCode)
		fmt.Println(root)
	}
}
