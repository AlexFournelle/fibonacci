package main

import (
	"fmt"
	"os"
	"time"
)

// Fibonacci V1
// Execute et mesure le temps d'éxécution de deux méthodes différentes
// pour générer une série de Fibonacci s'étendant jusqu'à un nombre n
// entré au clavier par l'utilisateur
// A faire: comparer les temps
func main() {
	var limite int
	var sequence []int

	fmt.Println("SÉQUENCE DE FIBONACCI")
	limite = readInt("Quelle longueur voulez vous pour la série?", limite)

	sequence = genererSequenceRecursive(limite)
	fmt.Println(sequence)
	sequence = genererSequenceBoucle(limite)
	fmt.Println(sequence)
}

// Première méthode, utilisant une fonction récursive f()
func genererSequenceRecursive(lim int) []int {
	defer chrono("Génération par récursion")()
	var res []int
	var f func(n int) int
	f = func(n int) int {
		if n <= 1 {
			return int(n)
		}
		return f(n-1) + f(n-2)
	}
	for i := 0; i < lim; i++ {
		res = append(res, f(i))
	}
	return res
}

// Deuxième méthode, utilisant une boucle
func genererSequenceBoucle(lim int) []int {
	defer chrono("Génération par boucle")()
	var res []int
	res = append(res, 0, 1)
	for i := 2; i < lim; i++ {
		res = append(res, res[i-1]+res[i-2])
	}
	return res
}

// Affiche un prompt puis attrape un entier du standard input
// Ça passe ou ça casse. À améliorer.
func readInt(msg string, i int) int {
	fmt.Println(msg)
	fmt.Print("(Entrez un entier de 1-50 - Attention: ça pourrait être long!) :")
	_, err := fmt.Scanf("%d", &i)
	if err != nil || i > 50 || i < 1 {
		fmt.Println(err)
		os.Exit(1)
	}
	return i
}

// Calcule le temps d'éxécution d'une fonction
// Voodoo de fonction déférée que je comprend pas parfaitement bien
func chrono(quoi string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("La %s a demandé %v de temps CPU\n", quoi, time.Since(start))
	}
}
