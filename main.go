package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Jogo da Advinhação")
	fmt.Println("Um numero aleatorio será sorteado, tente acertar. O numero é um inteiro entre 0 e 100")

	x := rand.Int64N(101)
	scanner := bufio.NewScanner(os.Stdin)
	chutes := [10]int64{}

	for i := range chutes {
		fmt.Println("Qual seu chute?")
		scanner.Scan()
		chute := scanner.Text()
		chute = strings.TrimSpace(chute)

		chuteInt, err := strconv.ParseInt(chute, 10, 64)
		if err != nil {
			fmt.Println("O seu chute tem que ser um inteiro")
			return
		}

		switch {
		case chuteInt < x:
			fmt.Println("Você errou, o numero sorteado é maior que ", chuteInt)
		case chuteInt > x:
			fmt.Println("Você errou, o numero sorteado é menor que ", chuteInt)
		case chuteInt == x:
			fmt.Printf("Parabens, você acertou o numero! O numero era %d\n"+
				"Você acertou em %d tentativas.\n"+
				"Essas foram as suas tentativas: %v\n", x, i+1, chutes[:i])
			return
		}
		chutes[i] = chuteInt
	}
	fmt.Printf("Infelizmente você nao acertou o numero, que era %d\n"+
		"Você teve 10 tentativas.\n"+
		"Essas foram as suas tentativas: %v\n", x, chutes)
}
