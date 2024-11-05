package main

import "fmt"

type Node struct {

	level int
	pathCost int
	visited []bool
	lastVisited int
}

// Алгоритм Литтла
func LittleAlgorithm(m *Matrix) {
	fmt.Println("Запуск алгоритма Литтла...")
	m.Print()

	if m.size == 0 {
		fmt.Println("Матрица пуста.")
		return
	}

	// Инициализация
	minCost := int(^uint(0) >> 1) // бесконечность
	visited := make([]bool, m.size)
	visited[0] = true // Начинаем с первой вершины

	// Рекурсивный поиск
	var tsp func(node Node)
	tsp = func(node Node) {
		if node.level == m.size-1 {
			// Добавляем возврат в начальную точку
			totalCost := node.pathCost + m.data[node.lastVisited][0]
			if totalCost < minCost {
				minCost = totalCost
			}
			return
		}

		for i := 0; i < m.size; i++ {
			if !node.visited[i] && m.data[node.lastVisited][i] != int(^uint(0)>>1) {
				// Создаем новый узел
				newNode := Node{
					level:       node.level + 1,
					pathCost:    node.pathCost + m.data[node.lastVisited][i],
					visited:     make([]bool, m.size),
					lastVisited: i,
				}

				// Копируем состояние посещенных вершин
				copy(newNode.visited, node.visited)
				newNode.visited[i] = true

				// Рекурсивный вызов
				tsp(newNode)
			}
		}
	}

	// Начальный узел
	initialNode := Node{
		level:       0,
		pathCost:    0,
		visited:     visited,
		lastVisited: 0,
	}

	// Запуск рекурсивного поиска
	tsp(initialNode)

	// Вывод минимальной стоимости
	fmt.Printf("Минимальная стоимость маршрута: %d\n", minCost)
}