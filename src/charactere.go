package main

import "fmt"

// Structure pour le personnage
type Personnage struct {
	Nom             string
	Classe          string
	Niveau          int
	experience      int
	experienceMax   int
	PvMax           int
	PvActuels       int
	Money           int
	maxSize         int
	initiative      int
	Mana            int
	ManaMax         int
	PersoInv        []inventaire
	PersoEquipement struct {
		head  Equipment
		body  Equipment
		boots Equipment
	}
	playerSkill     []Skill
	EquipementTete  Equipment
	EquipmentTorse  Equipment
	EquipementPieds Equipment
}

type Equipment struct {
	nomEquipement string
	stats         struct {
		attaque int
		defense int
	}
}

type inventaire struct {
	VarInventaire Item
}

type equipement struct {
	head  Item
	body  Item
	boots Item
}

// Fonction pour initialiser un personnage
func Init(nom string, classe string, niveau int, experience int, experienceMax int, pvMax int, pvActuels int, Money int, maxSize int, initiative int,Mana int,ManaMax int, PersoInv []inventaire , playerSkill []Skill) Personnage {
	return Personnage{
		Nom:           nom,
		Classe:        classe,
		Niveau:        niveau,
		experience:    experience,
		experienceMax: experienceMax,
		PvMax:         pvMax,
		PvActuels:     pvActuels,
		Money:         Money,
		maxSize:       maxSize,
		initiative:    initiative,
		Mana:		   Mana,
		ManaMax : ManaMax,
		PersoInv:      []inventaire{},
		playerSkill:   []Skill{},
	}
}

// Fonction pour afficher les informations du personnage
func displayInfo(p Personnage) {
	clearTerminal()
	fmt.Printf("==========================\nInformations du personnage \n==========================\n")
	fmt.Printf("Nom : %s\nClasse : %s\nNiveau : %d\nExperience : %d sur %d\nPoints de vie actuels : %d sur %d pv\nArgent : %d\nTaille Inventaire :%d\n", p.Nom, p.Classe, p.Niveau, p.experience, p.experienceMax, p.PvActuels, p.PvMax, p.Money, p.maxSize)
	fmt.Println("Équipement équipé:")
	fmt.Println("Tête:", p.EquipementTete.nomEquipement)
	fmt.Println("Attaque:", p.EquipementTete.stats.attaque)
	fmt.Println("Défense:", p.EquipementTete.stats.defense)
}
