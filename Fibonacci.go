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

// Première méthode, utilisant une fonction récursive (Fibonacci())
func genererSequenceRecursive(lim int) []int {
	defer chrono("Génération par récursion")()
	var res []int
	for i := 0; i < lim; i++ {
		res = append(res, Fibonacci(i))
	}
	return res
}

// Génère le nombre n de la séquence de Fibonacci de façon récursive
func Fibonacci(n int) int {
	if n <= 1 {
		return int(n)
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// Deuxième méthode, utilisant une boucle
func genererSequenceBoucle(lim int) []int {
	defer chrono("Génération par boucle")()
	var res []int
	if lim <= 1 {
		res = append(res, lim)
		return res
	}
	n1, n2 := 1, 0
	for i := 2; i < lim; i++ {
		n2, n1 = n1, n1+n2
	}
	res = append(res, n2+n1)
	return res
}

// Affiche un prompt puis attrape un entier du standard input
// Ça passe ou ça casse. À améliorer.
func readInt(msg string, i int) int {
	fmt.Println(msg)
	fmt.Print("(Entrez un entier de 1-50 - Attention: ça pourrait être long!) :")
	_, err := fmt.Scanf("%d", &i)
	if err != nil || i > 50 {
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
