package main

import (
	"fmt"
)

func main() {
	var n int // количество городов
	fmt.Print("Введите количество городов от 5 до 20")
	fmt.Scan(&n)

	if n < 5 || n > 20 {
		fmt.Println("Неверное значение. Пожалуйста, введите число от 5 до 20.")
		return
	}

	matrix := NewMatrix(n)

	var fill_choice string // способ заполнения матрицы
	fmt.Print("Выберите способ заполнения матрицы (random or manual)")
	fmt.Scan(&fill_choice)

	if fill_choice == "random" {
		matrix.FillRandom()
	} else if fill_choice == "manual" {
		matrix.FillManual()
	} else {
		fmt.Println("Метод заполнения не определён! Используется случайное заполнение")
		matrix.FillRandom()
	}

	fmt.Println("Исходная матрица:")
	matrix.Print()

	if !matrix.IsNormalized(){
		fmt.Println("Матрица не приведена. Запущен процесс нормализации...")
		matrix.Normalize()
		fmt.Println("Приведённая матрица")
		matrix.Print()
	} else {
		fmt.Println("Матрица уже приведена.")
	}

	//запуск алгоритма Литтла
	LittleAlgorithm(matrix)
}