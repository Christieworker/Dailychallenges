package main

import "fmt"

func main() {
	fmt.Println("Hello, World!") // afficher du texte

	var Poids, Taille, IMC float32

	Poids = 80
	Taille = 1.75

	IMC = Poids / (Taille * Taille)

	fmt.Println("L'indice de masse corporelle est :", IMC)

	if IMC < 18.5 {
		fmt.Println("Votre poids est trop faible.")
	} else if IMC > 18.5 && IMC <= 25 {
		fmt.Println("Vous avez un poids normal.")
	} else {
		fmt.Println("Vous êtes en surpoids.")
	}
}

// renseigner le poids en kilogramme et la taille en metre
/*func calculator(Poids, Taille float32) {
	Poids = 80
	Taille = 1.75
	var IMC float32
	IMC = Poids / (Taille * Taille)

	fmt.Println("L'indice de masse corporelle est :", IMC)

	if IMC < 18.5 {
		fmt.Println("Votre poids est trop faible.")
	} else if IMC > 18.5 && IMC <= 25 {
		fmt.Println("Vous avez un poids normal.")
	} else {
		fmt.Println("Vous êtes en surpoids.")
	}

	//return IMC

}*/
