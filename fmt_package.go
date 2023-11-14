package main

import "fmt"

func main() {
	nome := "Carlos"
	ano_nascimento := 1989
	fmt.Print(nome, ano_nascimento)
	fmt.Println(nome, ano_nascimento)
	fmt.Printf("Nome: %s\nAno de Nascimento: %d\n", nome, ano_nascimento)

	dados_sprint := fmt.Sprint(nome, ano_nascimento)
	dados_sprintln := fmt.Sprintln(nome, ano_nascimento)
	dados_sprintf := fmt.Sprintf("Nome: %s\nAno de Nascimento: %d\n", nome, ano_nascimento)
	fmt.Print(dados_sprint)
	fmt.Print(dados_sprintln)
	fmt.Print(dados_sprintf)
}
