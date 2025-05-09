package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func oddFilter(args []string) ([]int, []string) {
	oddNumbers := []int{}
	invalidNumbers := []string{}
	for _, i := range args {
		num, err := strconv.Atoi(i)
		if err != nil {
			invalidNumbers = append(invalidNumbers, i)
			continue
		}
		if num%2 != 0 {
			oddNumbers = append(oddNumbers, num)
		}
	}
	return oddNumbers, invalidNumbers
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Digite os números separados por espaço (ex: 1 2 3 4) e pressione Enter:")
	fmt.Print("> ")

	userInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Erro ao ler entrada:", err)
		os.Exit(1)
	}

	userInput = strings.TrimSpace(userInput)
	if userInput == "" {
		fmt.Println("Nenhum número foi digitado.")
		os.Exit(0)
	}

	stringNumbers := strings.Fields(userInput)

	result, invalidResult := oddFilter(stringNumbers)

	if len(result) > 0 {
		fmt.Println("Números ímpares encontrados:")
		for _, i := range result {
			fmt.Print(i, " ")
		}
	} else {
		fmt.Println("Nenhum número ímpar encontrado (ou todos os inputs válidos eram pares).")
	}

	if len(invalidResult) > 0 {
		fmt.Println("\nValores invalidos que foram ignorados:")
		for _, i := range invalidResult {
			fmt.Print(i, " ")
		}
	}
}
