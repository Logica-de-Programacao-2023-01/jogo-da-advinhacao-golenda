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

	rand.Seed(time.Now().UnixNano())
	numeroAleatorio := rand.Intn(100) + 1

	//Declaracao de variaveis
	var resposta int
	var contador int

	for resposta != numeroAleatorio {
		contador++
		fmt.Println(numeroAleatorio)
		fmt.Println("Tentativa de numero", contador)
		fmt.Println("Digite um numero aleatorio entre 1 e 100: ")
		fmt.Scan(&resposta)
		fmt.Println("**************************")
		if resposta == numeroAleatorio {
			fmt.Println("Parabens voce acertou!", numeroAleatorio, "era o numero aleatorio. :)")
			fmt.Printf("Voce usou %d tentativas", contador)
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
}
