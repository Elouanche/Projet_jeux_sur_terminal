package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// Fonction pour clear le terminal et avoir que la partie qui nous intéresse
func clearTerminal() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

//function pour le texte du menu
func textMenu() { 
	fmt.Println("\n==========================================")
	fmt.Println("Menu :")
	fmt.Printf("==========================================\n")
	fmt.Printf("1. Afficher les informations du personnage\n")
	fmt.Printf("2. Accéder au contenu de l'inventaire\n")
	fmt.Printf("3. Marchand\n")
	fmt.Printf("4. Skill\n")
	fmt.Printf("5. Forgeron\n")
	fmt.Printf("6. Combat\n")
	fmt.Printf("7. Quitter\n")
	fmt.Printf("==========================================\n")
	fmt.Printf("Choix (1-7) :\n")
	fmt.Printf("==========================================\n")
}

//Affichage du menu
func main() {
	clearTerminal()
	playerName := isValidNom()
	playerClasse, PvMax := choixClasse()
	p1 := Init(playerName, playerClasse, 1, 0, 100, PvMax, PvMax/2, 100, 10, 1,100, 100, []inventaire{}, []Skill{}) 
	addInventaire(&p1, "Heal Potion", "Ajoute 50 pv", 1)                                        
	addInventaire(&p1, "Poison Potion", "Retire 10 pv par seconde", 1)
	addSkill(&p1, "Coup de poing", "Permet au joueur de mettre un coup de poing", 10)
	mobGobelinEntrainement := initMonstre("Gobelin d'entrainement", 40, 40, 5, 1)

	var choix int

	for {
		textMenu()
		fmt.Scan(&choix)

		switch choix {
		case 1:
			displayInfo(p1)
		case 2:
			accessInventory(&p1, &mobGobelinEntrainement)
		case 3:
			merchantMenu(&p1, merchantItems)
		case 4:
			displaySkills(&p1)
		case 5:
			forgeronMenu(&p1)
		case 6:
			areneCombat(&p1, &mobGobelinEntrainement)
		case 7:
			clearTerminal()
			fmt.Println("Au revoir !")
			os.Exit(0)
		default:
			fmt.Println("Choix invalide. Veuillez choisir entre 1 et 7.")
		}
	}
}
