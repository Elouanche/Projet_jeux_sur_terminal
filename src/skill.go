package main

import (
	"fmt"
	"time"
)

//Structure pour les compétences
type Skill struct {
	Name        string
	Description string
	Damage      int
}


//Ajout de compétences
func addSkill(p *Personnage, name string, description string, damage int) {
	newSkill := Skill{Name: name, Description: description, Damage: damage}
	p.playerSkill = append(p.playerSkill, newSkill)
}

//Permets d'afficher les compétences
func displaySkills(p *Personnage) {
	clearTerminal()
	fmt.Println("\nVos compétences :")
	for i, skill := range p.playerSkill {
		fmt.Printf("%d. %s - %s\n", i+1, skill.Name, skill.Description)
	}
}

func AccessSkills(p *Personnage, mobGobelinEntrainement *Monstre) {
	for {
		displaySkills(p)
		fmt.Print("Choisissez une compétence à utiliser (ou 0 pour retourner au menu précédent) : ")
		var choix3 int
		fmt.Scan(&choix3)
		if choix3 == 0 {
			return
		}
		if choix3 > 0 && choix3 <= len(p.playerSkill) {

			for skillIndex, skill := range p.playerSkill {

				if skillIndex == choix3-1 {

					if skill.Damage > p.Mana {
						fmt.Println("Pas assez de mana pour utiliser cette compétence.")
						return
					}

					p.Mana -= skill.Damage
					fmt.Printf("Vous utilisez %s et infligez %d dégâts.\n", skill.Name, skill.Damage)
					time.Sleep(2 * time.Second)
					mobGobelinEntrainement.Pv_Actuels -= skill.Damage
					return

				}
			}
		}
	}
}

func skillPlayer() {
	player := Personnage{}

	// Ajoutez des compétences de base
	addSkill(&player, "Coup de poing", "Permet au joueur de mettre un coup de poing", 10)
	fmt.Println("Vous avez utilisé Coup de poing et infligé 10 points de dégâts.")

	// Utilisez le SpellBook pour ajouter d'autres compétences
	SpellBook(&player)

	// Accédez aux compétences
	displaySkills(&player)
}

func SpellBook(p *Personnage) {
	addSkill(p, "Boule de feu", "Permet au joueur de jeter une boule de feu", 20)
}
