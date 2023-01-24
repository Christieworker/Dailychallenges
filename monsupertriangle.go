package main

import "fmt"

func superTriangle(carac string, nblines int) {
	fmt.Println("Quel caractere utiliser")
	fmt.Scanln(&carac)
	fmt.Println(carac)

	/*for index, rune := range carac {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}*/

	fmt.Println("Combien de lignes")
	fmt.Scanln(&nblines)
	fmt.Println(nblines)

	if carac == " " {
		fmt.Println("Saisissez un seul caractère et un caractère différent du caractère espace")
		fmt.Scanln(&carac)
		fmt.Println(carac)
	}

	/*if len(carac) > 1 || carac == " " {
		fmt.Println("Saisissez un seul caractère et un caractère différent du caractère espace")
		fmt.Scanln(&carac)
		fmt.Println(carac)
	}*/

	if nblines > 16 {
		fmt.Println("Vous ne devez pas dépasser 16 lignes. Entrez à nouveau le nombre de ligne")
		fmt.Scanln(&nblines)
		fmt.Println(nblines)
	}

}

func main() {

	var caractere string
	var nbrlignes int

	superTriangle(caractere, nbrlignes)

}
