package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Matrix struct {
	data [][]int
	size int
}

func NewMatrix(size int) *Matrix {
	matrix := &Matrix{
		data: make([][]int, size),
		size: size,
	}
	for i := range matrix.data {
		matrix.data[i] = make([]int, size)
		for j := range matrix.data[i] {
			if i == j {
				matrix.data[i][j] = int(^uint(0) >> 1) // символ бесконечности
			} else {
				matrix.data[i][j] = 0 // изначально 0
			}
		}
	}
	return matrix
}

// Автоматическое заполнение матрицы
func (m *Matrix) FillRandom() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < m.size; i++ {
		for j := 0; j < m.size; j++ {
			if m.data[i][j] != int(^uint(0)>>1) { // если не бесконечность
				m.data[i][j] = rand.Intn(100) + 1
			}
		}
	}
}

// Ручное заполнение матрицы
func (m *Matrix) FillManual() {
	fmt.Println("Введите элементы матрицы:")
	for i := 0; i < m.size; i++ {
		for j := 0; j < m.size; j++ {
			fmt.Printf("Элемент [%d][%d]: ", i, j)
			fmt.Scan(&m.data[i][j])
		}
	}
}

// Проверка нормализации матрицы
func (m *Matrix) IsNormalized() bool {
	for i := 0; i < m.size; i++ {
		hasZero := false
		for j := 0; j < m.size; j++ {
			if m.data[i][j] == 0 {
				hasZero = true
				break
			}
		}
		if !hasZero {
			return false
		}
	}
	return true
}

// Нормализация матрицы
func (m *Matrix) Normalize() {
	for i := 0; i < m.size; i++ {
		minValue := int(^uint(0) >> 1)
		for j := 0; j < m.size; j++ {
			if m.data[i][j] < minValue && m.data[i][j] != int(^uint(0)>>1) {
				minValue = m.data[i][j]
			}
		}
		for j := 0; j < m.size; j++ {
			if m.data[i][j] != int(^uint(0)>>1) { // Если не бесконечность
				m.data[i][j] -= minValue
			}
		}
	}
}

// Печать матрицы
func (m *Matrix) Print() {
	for _, row := range m.data {
		for _, value := range row {
			if value == int(^uint(0)>>1) {
				fmt.Print("∞ ")
			} else {
				fmt.Printf("%d ", value)
			}
		}
		fmt.Println()
	}
}