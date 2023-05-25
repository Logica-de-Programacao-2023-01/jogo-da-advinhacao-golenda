package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("|-------------------------------------|")
	fmt.Println("|--Bem-vindo ao jogo da adivinhação!--|")
	fmt.Println("|-------------------------------------|")

	//Declaracao de variaveis
	var resposta int
	var contador int
	var jogar bool
	jogar = true
	var x int
	var slice []int
	for jogar == true {
		rand.Seed(time.Now().UnixNano())
		numeroAleatorio := rand.Intn(100) + 1
		for resposta != numeroAleatorio {
			contador++
			fmt.Println(contador, "º tentativa")
			fmt.Println("Digite um numero aleatorio entre 1 e 100: ")
			fmt.Scan(&resposta)
			fmt.Println("**************************")
			if resposta == numeroAleatorio {
				fmt.Println("Parabens voce acertou!", numeroAleatorio, "era o numero aleatorio. :)")
				fmt.Println("Voce usou", contador, "tentativas!")
				break
			} else if resposta > 100 || resposta < 1 {
				fmt.Println("O numero deve estar entre 1 e 100")
			} else if resposta > numeroAleatorio {
				fmt.Println("O numero é MENOR que", resposta)
				fmt.Println("**************************")
			} else if resposta < numeroAleatorio {
				fmt.Println("O numero é MAIOR que ", resposta)
				fmt.Println("**************************")
			}
		}
		for i := 0; i < 1; i++ {
			slice = append(slice, contador)
		}

		fmt.Println("Quer jogar novamente? (1)Sim (2)Não")
		fmt.Scan(&x)
		contador = 0
		if x == 1 {
			fmt.Println("**************************")
			fmt.Println("Certo! Sorteando um novo numero...")
		} else if x == 2 {
			fmt.Println("**************************")
			fmt.Println("Obrigado por jogar!")
			for i := 0; i < len(slice); i++ {
				fmt.Println("Na", i+1, "º rodada, voce utilizou de", slice[i])
			}
			jogar = false

		}
	}
}
