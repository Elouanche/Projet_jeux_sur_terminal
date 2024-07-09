package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

//Permets au joueur de choisir son pseudo
func choixNom() string { 
	for {
		isValidNom()

		_, err := fmt.Scanln() 
		if err != nil {
			fmt.Println("Le Pseudo ne doit contenir que des lettres, avec la première en maj et les autres en minuscule. Réessayez :", err)
			fmt.Println("Réessayez :")
			continue
		}
	}
}

//Permets de savoir si le pseudo est valide
func isValidNom() string {
	scanner := bufio.NewScanner(os.Stdin)
	usernamePattern := regexp.MustCompile("^[a-zA-Z]+$")
	clearTerminal()
	for {
		fmt.Println("==================================")
		fmt.Println("Quel est votre Pseudo?")
		fmt.Println("==================================")
		scanner.Scan()
		playerName := scanner.Text()
		if !usernamePattern.MatchString(playerName) {
			clearTerminal()
			fmt.Printf("Vous ne pouvez mettre que des caractères dans votre pseudo, réessayer.\n")
		} else {
			return strings.Title(strings.ToLower(playerName))
		}
	}
}

//permet au joueur de choisir sa classe 
func choixClasse() (string, int) {
	var playerClasse string
	var PvMax int
	fmt.Println("==================================")
	fmt.Println("Quel est votre classe? (Sorcière, Paysan ou Chevalier)")
	fmt.Println("==================================")
	for {
		_, err := fmt.Scanln(&playerClasse)
		if err != nil {
			fmt.Println("Tu ne peux pas choisir cette classe !", err)
			continue
		}
		switch playerClasse {
		case "Chevalier":
			PvMax = 120
		case "Paysan":
			PvMax = 100
		case "Sorcière":
			PvMax = 80
		default:
			fmt.Println("Tu ne peux pas choisir cette classe !")
			continue
		}
		clearTerminal()
		fmt.Printf("Vous avez choisi la classe : %s\n", playerClasse)
		time.Sleep(2 * time.Second)
		clearTerminal()
		return playerClasse, PvMax
	}
}
