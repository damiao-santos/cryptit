package decrypt

func Revert(str string) string {
	decryptedStr := ""
	for _, c := range str {
		asciiCode := int(c)
		character := string(rune(asciiCode - 3))
		decryptedStr += character
	}
	
	return decryptedStr
}
