package piscine

import (
	"github.com/01-edu/z01"
)

func QuadA(x, y int) {
	if x <= 0 || y <= 0 { // Test si les paramètre sont negatif ou null
		return
	}
	for j := 1; j <= y; j++ { //boucle pour les lignes du carré
		for i := 1; i <= x; i++ { //boucle pour les colones du carré
			switch {
			case j == 1 && i == 1: // pour l'angle en haut a gauche
				z01.PrintRune(rune('o'))
			case j == 1 && i == x: // pour l'angle en haut a droit
				z01.PrintRune(rune('o'))
			case j == y && i == 1: // pour l'angle en bas a gauche
				z01.PrintRune(rune('o'))
			case j == y && i == x: // pour l'angle en bas a droit
				z01.PrintRune(rune('o'))

			case j == 1 && i > 1 && i < x: // pour la ligne du haut
				z01.PrintRune(rune('-'))
			case j == y && i > 1 && i < x: // pour la ligne du bas
				z01.PrintRune(rune('-'))

			case i == 1 && j > 1 && j < y: // pour la colone de gauche
				z01.PrintRune(rune('|'))
			case i == x && j > 1 && j < y: // pour la colone de droit
				z01.PrintRune(rune('|'))
			default:
				z01.PrintRune(rune(' ')) // si c'est aucun de c'est ca on met un espace
			}
		}
		z01.PrintRune(rune('\n'))
	}
}
func QuadB(x, y int) {
	if x <= 0 || y <= 0 { // Test si les paramètre sont negatif ou null
		return
	}
	for j := 1; j <= y; j++ { //boucle pour les lignes du carré
		for i := 1; i <= x; i++ { //boucle pour les colones du carré
			switch {
			case j == 1 && i == 1: // pour l'angle en haut a gauche
				z01.PrintRune(rune('/'))
			case j == 1 && i == x: // pour l'angle en haut a droit
				z01.PrintRune(rune(92))
			case j == y && i == 1: // pour l'angle en bas a gauche
				z01.PrintRune(rune(92))
			case j == y && i == x: // pour l'angle en bas a droit
				z01.PrintRune(rune('/'))

			case j == 1 && i > 1 && i < x: // pour la ligne du haut
				z01.PrintRune(rune('*'))
			case j == y && i > 1 && i < x: // pour la ligne du bas
				z01.PrintRune(rune('*'))

			case i == 1 && j > 1 && j < y: // pour la colone de gauche
				z01.PrintRune(rune('*'))
			case i == x && j > 1 && j < y: // pour la colone de droit
				z01.PrintRune(rune('*'))
			default:
				z01.PrintRune(rune(' ')) // si c'est aucun de c'est ca on met un espace
			}
		}
		z01.PrintRune(rune('\n'))
	}
}
func QuadC(x, y int) {
	if x <= 0 || y <= 0 { // Test si les paramètre sont negatif ou null
		return
	}
	for j := 1; j <= y; j++ { //boucle pour les lignes du carré
		for i := 1; i <= x; i++ { //boucle pour les colones du carré
			switch {
			case j == 1 && i == 1: // pour l'angle en haut a gauche
				z01.PrintRune(rune('A'))
			case j == 1 && i == x: // pour l'angle en haut a droit
				z01.PrintRune(rune('A'))
			case j == y && i == 1: // pour l'angle en bas a gauche
				z01.PrintRune(rune('C'))
			case j == y && i == x: // pour l'angle en bas a droit
				z01.PrintRune(rune('C'))

			case j == 1 && i > 1 && i < x: // pour la ligne du haut
				z01.PrintRune(rune('B'))
			case j == y && i > 1 && i < x: // pour la ligne du bas
				z01.PrintRune(rune('B'))

			case i == 1 && j > 1 && j < y: // pour la colone de gauche
				z01.PrintRune(rune('B'))
			case i == x && j > 1 && j < y: // pour la colone de droit
				z01.PrintRune(rune('B'))
			default:
				z01.PrintRune(rune(' ')) // si c'est aucun de c'est ca on met un espace
			}
		}
		z01.PrintRune(rune('\n'))
	}
}
func QuadD(x, y int) {
	if x <= 0 || y <= 0 { // Test si les paramètre sont negatif ou null
		return
	}
	for j := 1; j <= y; j++ { //boucle pour les lignes du carré
		for i := 1; i <= x; i++ { //boucle pour les colones du carré
			switch {
			case j == 1 && i == 1: // pour l'angle en haut a gauche
				z01.PrintRune(rune('A'))
			case j == 1 && i == x: // pour l'angle en haut a droit
				z01.PrintRune(rune('C'))
			case j == y && i == 1: // pour l'angle en bas a gauche
				z01.PrintRune(rune('A'))
			case j == y && i == x: // pour l'angle en bas a droit
				z01.PrintRune(rune('C'))

			case j == 1 && i > 1 && i < x: // pour la ligne du haut
				z01.PrintRune(rune('B'))
			case j == y && i > 1 && i < x: // pour la ligne du bas
				z01.PrintRune(rune('B'))

			case i == 1 && j > 1 && j < y: // pour la colone de gauche
				z01.PrintRune(rune('B'))
			case i == x && j > 1 && j < y: // pour la colone de droit
				z01.PrintRune(rune('B'))
			default:
				z01.PrintRune(rune(' ')) // si c'est aucun de c'est ca on met un espace
			}
		}
		z01.PrintRune(rune('\n'))
	}
}
func QuadE(x, y int) {
	if x <= 0 || y <= 0 { // Test si les paramètre sont negatif ou null
		return
	}
	for j := 1; j <= y; j++ { //boucle pour les lignes du carré
		for i := 1; i <= x; i++ { //boucle pour les colones du carré
			switch {
			case j == 1 && i == 1: // pour l'angle en haut a gauche
				z01.PrintRune(rune('A'))
			case j == 1 && i == x: // pour l'angle en haut a droit
				z01.PrintRune(rune('C'))
			case j == y && i == 1: // pour l'angle en bas a gauche
				z01.PrintRune(rune('C'))
			case j == y && i == x: // pour l'angle en bas a droit
				z01.PrintRune(rune('A'))

			case j == 1 && i > 1 && i < x: // pour la ligne du haut
				z01.PrintRune(rune('B'))
			case j == y && i > 1 && i < x: // pour la ligne du bas
				z01.PrintRune(rune('B'))

			case i == 1 && j > 1 && j < y: // pour la colone de gauche
				z01.PrintRune(rune('B'))
			case i == x && j > 1 && j < y: // pour la colone de droit
				z01.PrintRune(rune('B'))
			default:
				z01.PrintRune(rune(' ')) // si c'est aucun de c'est ca on met un espace
			}
		}
		z01.PrintRune(rune('\n'))
	}
}
