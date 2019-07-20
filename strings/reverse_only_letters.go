package main

import "fmt"

func isLetter(asciiChar rune) bool {
	return asciiChar >= 'a' && asciiChar <= 'z' || asciiChar >= 'A' && asciiChar <= 'Z'
}

func reverseOnlyLetters(S string) string {
	i, j, str := 0, len(S)-1, []rune(S)

	for i < j {
		if !isLetter(str[i]) {
			i++
			continue
		}

		if !isLetter(str[j]) {
			j--
			continue
		}

		str[i], str[j] = str[j], str[i]
		i++
		j--
	}

	return string(str)
}

func main() {
	result := reverseOnlyLetters("Test1ng-Leet=code-Q!")

	fmt.Println(result) //Qedo1ct-eeLg=ntse-T!
}
