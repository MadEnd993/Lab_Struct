package main

import (
	"fmt"
	"strings"
)

func simulate(coins []bool, S, K int) int {
	n := len(coins)
	state := make([]bool, n)
	copy(state, coins)

	idx := 0
	startIdx := 0
	for i, c := range state {
		if c { // герб вверх
			startIdx = i
			break
		}
	}
	idx = startIdx

	for step := 0; step < K; step++ {
		idx = (idx + S - 1) % n
		state[idx] = !state[idx]

		idx = (idx + 1) % n
	}

	count := 0
	for _, c := range state {
		if c {
			count++
		}
	}
	return count
}

func isValidInitial(coins []bool, N, M int) bool {
	heads := 0
	tails := 0
	for _, c := range coins {
		if c {
			heads++
		} else {
			tails++
		}
	}
	return heads == N && tails == M
}

func search(pos int, coins []bool, N, M, S, K, L int) bool {
	if pos == len(coins) {
		if !isValidInitial(coins, N, M) {
			return false
		}
		return simulate(coins, S, K) == L
	}

	coins[pos] = true
	if search(pos+1, coins, N, M, S, K, L) {
		return true
	}

	coins[pos] = false
	if search(pos+1, coins, N, M, S, K, L) {
		return true
	}

	return false
}

func formatCoins(coins []bool) string {
	symbols := make([]string, len(coins))
	for i, c := range coins {
		if c {
			symbols[i] = "O" // орел (вверх)
		} else {
			symbols[i] = "Р" // решка (вниз)
		}
	}
	return strings.Join(symbols, " ")
}

func main() {
	var N, M, S, K, L int

	fmt.Println("Введите N M S K L (через пробел):")
	fmt.Scan(&N, &M, &S, &K, &L)

	total := N + M
	if total == 0 {
		fmt.Println("Нет монет")
		return
	}

	coins := make([]bool, total)

	if search(0, coins, N, M, S, K, L) {
		fmt.Println("Расстановка найдена:")
		fmt.Println(formatCoins(coins))

		finalHeads := simulate(coins, S, K)
		fmt.Printf("После %d ходов гербов вверх: %d (ожидалось %d)\n", K, finalHeads, L)
	} else {
		fmt.Println("Невозможно найти расстановку.")
	}
}
