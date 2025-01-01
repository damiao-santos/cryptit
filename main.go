package main

import (
	"fmt"

	"github.com/damiao-santos/cryptit/decrypt" //para importar o pacote decrypt devemos usar o nome do modulo + o package
	"github.com/damiao-santos/cryptit/encrypt" //para importar o pacote encrypt devemos usar o nome do modulo + o package
)

func main() {
	encrypted := encrypt.Convert("Damião")
	fmt.Println(encrypted)
	fmt.Println(decrypt.Revert(encrypted))
}

// parei 05:14
