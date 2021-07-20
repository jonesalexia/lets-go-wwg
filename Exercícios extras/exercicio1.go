// Exercício 1

// package main

// import (
// "fmt"
//)

// func main() {
//	var quilometros int8
//	quilometros = 150

//	fmt.Println(quilometros)
//}

// 1 - A aplicação não compila, porque o valor da variavel ultrapassa o tipo do inteiro (int8)

// 2 - O erro indica que que o valor 150 não está correto, pois excede o valor máximo permitido pelo tipo int8

// Correção:

package main

import (
	"fmt"
)

func main() {
	var quilometros int16
	quilometros = 150

	fmt.Println(quilometros)
}
