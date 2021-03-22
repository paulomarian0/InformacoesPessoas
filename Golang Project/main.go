package main

import (
	"fmt"
)

type Pessoa struct {
	Nome  string
	Idade int
}

func main() {

	var nomePessoa string
	var idadePessoa int
	var pessoaStruct Pessoa
	var entrada int
	slicePessoas := []Pessoa{}

	for {

		fmt.Println("Digite 1 para adicionar informações de alguma pessoa")
		fmt.Println("Digite 2 para exibir as informações atuais")
		fmt.Println("Digite 3 para finalizar o programa")
		fmt.Scan(&entrada)

		switch entrada {
		case 1:
			fmt.Println("Digite o nome da pessoa que deseja adicionar")
			fmt.Scan(&nomePessoa)
			fmt.Println("Digite a idade da pessoa que deseja adicionar")
			fmt.Scan(&idadePessoa)

			pessoaStruct = Pessoa{nomePessoa, idadePessoa}
			slicePessoas = append(slicePessoas, pessoaStruct)

			fmt.Println("Adicionado com sucesso")

			fmt.Print("-----------------------------------------------------\n")

			break

		case 2:
			fmt.Println(slicePessoas)

			fmt.Print("-----------------------------------------------------\n")

			break
		case 3:

			goto Final
		}

	}

Final:
	fmt.Println("Programa finalizado")
}
