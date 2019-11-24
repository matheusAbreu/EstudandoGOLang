package main

import (
	"bufio"
	"fmt"
	"os"

	pali "./EncontrandoPalindromo"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	testePali := pali.CriaTexto()
	fmt.Printf("\nManual:\n")
	fmt.Printf("1 - Não use letras maiúsculas\n")
	fmt.Printf("2 - Não utilize acentos pontuações, parênteses, colchetes, chaves, arroba ou sharp\n")
	fmt.Printf("3 - Espaços serão ignorados na comparação palíndrica\n")
	fmt.Printf("\nEscreva seu texto:\n")
	scanner.Scan()

	testePali.AddTexto(scanner.Text())
	fmt.Println(testePali.ImprimeMeuTexto())
	fmt.Println("resultado do teste de palindromo: ", testePali.VerificandoPalindromo())
}
