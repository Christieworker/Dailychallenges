package main

import "fmt"

func Calculator(Age, Poids, Taille float32) {

	var IMC float32

	IMC = Poids / (Taille * Taille)

	fmt.Println("Vous avez ", Age)

	fmt.Println("Votre indice de masse corporelle est :", IMC)

	if IMC < 18.5 {
		fmt.Println("Votre poids est trop faible.")
	} else if IMC > 18.5 && IMC <= 25 {
		fmt.Println("Vous avez un poids normal.")
	} else {
		fmt.Println("Vous êtes en surpoids.")
	}
}

func main() {
	// var then variable name then variable type
	var age float32
	age = -1
	var weight float32
	var height float32
	var shouldstop bool
	var stopvalue string

	for shouldstop = false; shouldstop == false; {

		for isAgevalid := false; isAgevalid == false; isAgevalid = age >= 1 && age <= 150 {

			// Println function is used to display output in the next line
			fmt.Println("Quel est votre âge :")

			// Taking input from user
			fmt.Scanln(&age)

			if age < 1 || age > 150 {

				fmt.Println("Votre âge est invalide. Entrez un âge compris entre 1 et 150")
			}
		}

		for correctweight := false; correctweight == false; correctweight = weight >= 5 && weight <= 1000 {

			fmt.Println("Entrez votre poids :")

			fmt.Scanln(&weight)

			if weight < 5 || weight > 1000 {
				fmt.Println("Votre poids est invalide. Entrez un poids compris entre 5 et 1000 kg")
			}

		}

		for incorrectheight := true; incorrectheight; incorrectheight = height < 0.5 || height > 3 {
			fmt.Println("Entrez votre taille : ")
			fmt.Scanln(&height)

			if height < 0.5 || height > 3 {
				fmt.Println("Votre taille est invalide. Entrez une taille comprise entre 0.5 et 3 m")
			}
		}

		Calculator(age, weight, height)

		fmt.Println("Souhaitez-vous continuer ? Oui ou Non")
		fmt.Scanln(&stopvalue)

		if stopvalue == "Non" {
			shouldstop = true
		} else {
			shouldstop = false
		}

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
