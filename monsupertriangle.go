package main

import "fmt"

func superTriangle(carac string, nblines int) {

	fmt.Println("Quel caractere utiliser")
	fmt.Scanln(&carac)

	for len(carac) > 1 {
		fmt.Println("Saisissez un seul caractère et un caractère différent du caractère espace")
		fmt.Scanln(&carac)
	}

	fmt.Println("Combien de lignes")
	fmt.Scanln(&nblines)

	for nblines > 16 {
		fmt.Println("Vous ne devez pas dépasser 16 lignes. Entrez à nouveau le nombre de lignes souhaité")
		fmt.Scanln(&nblines)
		fmt.Println("Vous avez demandé un triangle formé de ", carac, " alignés sur ", nblines, "")

	}

	/*if len(carac) > 1 || carac == " " {
		fmt.Println("Saisissez un seul caractère et un caractère différent du caractère espace")
		fmt.Scanln(&carac)
		fmt.Println(carac)

	}*/

	//----------------------------------Impression du triangle-------------------------------------

	var a, b int

	fmt.Println("Ci-dessous votre triangle")

	// La boucle for imprime le nombre de lignes
	for a = 1; a <= nblines; a++ {

		// La boucle for imprime le nombre de colonnes
		for b = 1; b <= a; b++ {

			// imprimer le caractère entré par l'utilisateur suivi d'un espace en utilisant fmt.Print()
			fmt.Print(carac, " ")
		}
		// imprimer une nouvelle ligne en utilisant la fonction fmt.Println()
		fmt.Println()
	}

}

func main() {

	var caractere string
	var nbrlignes int
	var arret bool
	var valeurdarret string

	for arret = false; arret == false; {

		superTriangle(caractere, nbrlignes)

		fmt.Println("Voulez-vous afficher un nouveau triangle ? Oui ou Non")
		fmt.Scanln(&valeurdarret)

		if valeurdarret != "Oui" && valeurdarret != "Non" {

			fmt.Println("Souhaitez-vous continuer ? Entrez 'Oui' ou 'Non'")
			fmt.Scanln(&valeurdarret)

			if valeurdarret == "Non" {
				arret = true
			} else if valeurdarret == "Oui" {
				arret = false
			}

		}

	}

}
