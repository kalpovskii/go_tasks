package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// Написать двумерный массив и заполнить его случайными числами

func main() {
	var n, m uint32 // кол-во столбцов и строк
	if _, err := fmt.Scanf("%d %d", &m, &n); err != nil {
		log.Fatal(err)
	}

	// инициализируем массив слайсов
	arr := make([][]int32, m)
	for i := range arr {
		arr[i] = make([]int32, n)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	mem := map[int32]struct{}{}
	for i := range arr {
		for j := range arr[i] {
			num := r.Int31n(900) + 100
			for _, ok := mem[num]; ok; {
				num = r.Int31n(900) + 100
			}
			mem[num] = struct{}{}
			arr[i][j] = num
		}
	}

	for i := range arr {
		for j := range arr[i] {
			fmt.Printf("\t%d", arr[i][j])
		}
		fmt.Println()
	}

	// Вывести строку с наибольшим числом
	maxNum := int32(100)
	maxRowIdx := 0
	for i := range arr {
		for j := range arr[i] {
			if arr[i][j] > maxNum {
				maxRowIdx = i
				maxNum = arr[i][j]
			}
		}
	}

	fmt.Println()
	for j := range arr[maxRowIdx] {
		fmt.Printf("\t%d", arr[maxRowIdx][j])
	}
	fmt.Println()
}
