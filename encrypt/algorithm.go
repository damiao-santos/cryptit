package encrypt

func Convert(str string) string { //para acessar a função Convert fora do pacote encrypt, ela deve começar com letra maiúscula
	encrytedStr := ""
	for _, c := range str {
		asciiCode := int(c)                      //int converte o valor de c para inteiro
		character := string(rune(asciiCode + 3)) //rune garante que o valor seja convertido para string
		encrytedStr += character
	}
	return encrytedStr
}
