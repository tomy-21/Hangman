package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var hangmanStages = []string{
	`
  +---+
  |   |
      |
      |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
      |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
  |   |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
 /|   |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
 /|\  |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
 /|\  |
 /    |
      |
=========`,
	`
  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
=========`,
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file>")
		return
	}

	// Lecture du fichier de mots
	words, err := readWordsFromFile(os.Args[1])
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}

	// Choix aléatoire d'un mot
	rand.Seed(time.Now().UnixNano())
	wordToGuess := words[rand.Intn(len(words))]

	// Révélation de lettres aléatoires
	revealedWord := revealRandomLetters(wordToGuess)

	fmt.Println("Mot à deviner:", revealedWord)

	// Début du jeu
	playHangman(wordToGuess, revealedWord)
}

// readWordsFromFile lit les mots du fichier fourni en paramètre
func readWordsFromFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

// revealRandomLetters révèle quelques lettres aléatoires dans le mot
func revealRandomLetters(word string) string {
	wordRunes := []rune(word)
	revealed := make([]rune, len(wordRunes))

	// Initialiser tous les caractères à '_'
	for i := range revealed {
		revealed[i] = '_'
	}

	// Révéler aléatoirement quelques lettres
	numRevealed := len(wordRunes) / 3
	revealedIndices := rand.Perm(len(wordRunes))[:numRevealed]
	for _, index := range revealedIndices {
		revealed[index] = wordRunes[index]
	}

	return string(revealed)
}

// playHangman démarre le jeu du pendu
func playHangman(wordToGuess, revealedWord string) {
	attempts := 0
	maxAttempts := len(hangmanStages) - 1
	guessedLetters := make(map[rune]bool)
	incorrectLetters := []rune{} // Stocker les lettres incorrectes
	wordRunes := []rune(wordToGuess)
	revealedRunes := []rune(revealedWord)

	for attempts <= maxAttempts {
		fmt.Println(hangmanStages[attempts])
		fmt.Printf("Tentatives restantes: %d\n", maxAttempts-attempts)
		fmt.Printf("Mot actuel: %s\n", string(revealedRunes))

		// Afficher les lettres incorrectes déjà devinées
		showIncorrectGuesses(incorrectLetters)

		fmt.Print("Entrez une lettre: ")

		var guess string
		fmt.Scanf("%s\n", &guess)

		if len(guess) != 1 {
			fmt.Println("Veuillez entrer une seule lettre.")
			continue
		}

		letter := rune(guess[0])
		if guessedLetters[letter] {
			fmt.Println("Vous avez déjà deviné cette lettre.")
			continue
		}

		guessedLetters[letter] = true

		found := false
		for i, r := range wordRunes {
			if r == letter {
				revealedRunes[i] = letter
				found = true
			}
		}

		if !found {
			attempts++
			incorrectLetters = append(incorrectLetters, letter) // Ajouter à la liste des lettres incorrectes
			fmt.Println("Lettre incorrecte!")
		} else {
			fmt.Println("Lettre correcte!")
		}

		if string(revealedRunes) == wordToGuess {
			fmt.Printf("Félicitations! Vous avez deviné le mot: %s\n", wordToGuess)
			return
		}

		if attempts > maxAttempts {
			fmt.Println(hangmanStages[maxAttempts])
			fmt.Printf("Désolé, vous avez perdu. Le mot était: %s\n", wordToGuess)
			return
		}
	}
}

// showIncorrectGuesses affiche les lettres incorrectes déjà essayées
func showIncorrectGuesses(incorrectLetters []rune) {
	if len(incorrectLetters) > 0 {
		fmt.Printf("Lettres incorrectes déjà essayées: %s\n", string(incorrectLetters))
	}
}
