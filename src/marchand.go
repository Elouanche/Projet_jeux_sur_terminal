package main

import (
	"fmt"
)

//Permet d'afficher ce qu'il y a dans le shop
func displayMerchantItems(items []Item) { 
	fmt.Println("Objets disponibles chez le marchand :")
	i := 0
	for _, item := range items {
		fmt.Printf("%d. %s - %s\n", i+1, item.nomItem, item.effet)
		i++
	}
}

//Menu du marchand
func merchantMenu(p *Personnage, merchantItems []Item) { 
	clearTerminal()
	fmt.Printf("\nBienvenue chez le marchand ! Voici ce que je vends :\n")
	displayMerchantItems(merchantItems)
	fmt.Print("\nChoisissez un objet à acheter (ou 0 pour retourner au menu précédent) : ")
	var choix2 int
	fmt.Scan(&choix2)

	if choix2 == 0 {
		return
	}

	if choix2 > 0 && choix2 <= len(merchantItems) && choix2 < 2 { 
		healPotion(p, choix2)
	}
	if choix2 > 1 && choix2 <= len(merchantItems) && choix2 < 3 { 
		poisonPotion(p, choix2)
	}
	if choix2 > 2 && choix2 <= len(merchantItems) && choix2 < 4 { 
		fourrureLoup(p, choix2)
	}
	if choix2 > 3 && choix2 <= len(merchantItems) && choix2 < 5 { 
		peauTroll(p, choix2)
	}
	if choix2 > 4 && choix2 <= len(merchantItems) && choix2 < 6 { 
		cuireSanglier(p, choix2)
	}
	if choix2 > 5 && choix2 <= len(merchantItems) && choix2 < 7 {
		plumeCorbeau(p, choix2)
	}
	if choix2 > 6 && choix2 <= len(merchantItems) && choix2 < 8 {
		UpgradeInventorySlot(p, choix2)
	}
	if choix2 > 7 && choix2 < 9 { 
		clearTerminal()
		if p.Money < 25 {
			fmt.Println("Tu n'a plus assez d'argent")
			return
		} else {
			p.Money -= 25
			addInventaire(p, "SpellBook", "Permet d'apprendre un sort", 1)
		}
	}
	if choix2 > 8 {
		fmt.Println("Désolé, mais tu ne peux pas faire ça ! Pars de là !")
	}
}

//Permets de remplire le tableau d'item du shop
var merchantItems = []Item{ 
	{nomItem: "Potion de vie ", effet: "Restaure 50 PV", quantity: 1},
	{nomItem: "Potion de poison", effet: "Enlève 10 pv par seconde", quantity: 1},
	{nomItem: "Fourrure de Loup", effet: "Permet de créer de l'équipement", quantity: 1},
	{nomItem: "Peau de Troll", effet: "Permet de créer de l'équipement", quantity: 1},
	{nomItem: "Cuire de Sanglier", effet: "Permet de créer de l'équipement", quantity: 1},
	{nomItem: "Plume de Corbeau", effet: "Permet de créer de l'équipement", quantity: 1},
	{nomItem: "Sac à dos magique", effet: "Permet d'augmenter son stockage", quantity: 1},
	{nomItem: "SpellBook", effet: "Permet d'apprendre un sort", quantity: 1},
}

func healPotion(p *Personnage, choix2 int) {
	item := merchantItems[choix2-1]
	clearTerminal()
	if p.Money < 3 {
		fmt.Println("Tu n'a plus assez d'argent")
		return
	} else if len(p.PersoInv) <= 10 {
		p.Money -= 3
		fmt.Printf("Vous avez acheté : %s vous avez: %d pièces\n", item.nomItem, p.Money)
		addInventaire(p, "Heal Potion", "Ajoute 50 pv", 1)
	}
}

func poisonPotion(p *Personnage, choix2 int) {
	item := merchantItems[choix2-1]
	clearTerminal()
	if p.Money < 6 {
		fmt.Println("Tu n'a plus assez d'argent")
		return
	} else if len(p.PersoInv) <= 10 {
		p.Money -= 6
		fmt.Printf("Vous avez acheté : %s vous avez : %d pièces\n", item.nomItem, p.Money)
		addInventaire(p, "Poison Potion", "Enlève 10 pv par seconde", 1)
	}
}

func fourrureLoup(p *Personnage, choix2 int) {
	item := merchantItems[choix2-1]
	clearTerminal()
	if p.Money < 4 {
		fmt.Println("Tu n'a plus assez d'argent")
		return
	} else if len(p.PersoInv) <= 10 {
		p.Money -= 4
		fmt.Printf("Vous avez acheté : %s vous avez : %d pièces\n", item.nomItem, p.Money)
		addInventaire(p, "Fourrure de Loup", "Permet de créer de l'équipement", 1)
	}
}

func peauTroll(p *Personnage, choix2 int) {
	item := merchantItems[choix2-1]
	clearTerminal()
	if p.Money < 7 {
		fmt.Println("Tu n'a plus assez d'argent")
		return
	} else if len(p.PersoInv) <= 10 {
		p.Money -= 7
		fmt.Printf("Vous avez acheté : %s vous avez : %d pièces\n", item.nomItem, p.Money)
		addInventaire(p, "Peau de Troll", "Permet de créer de l'équipement", 1)
	}
}

func cuireSanglier(p *Personnage, choix2 int) {
	item := merchantItems[choix2-1]
	clearTerminal()
	if p.Money < 3 {
		fmt.Println("Tu n'a plus assez d'argent")
		return
	} else if len(p.PersoInv) <= 10 {
		p.Money -= 3
		fmt.Printf("Vous avez acheté : %s vous avez : %d pièces\n", item.nomItem, p.Money)
		addInventaire(p, "Cuire de Sanglier", "Permet de créer de l'équipement", 1)
		return
	}
}

func plumeCorbeau(p *Personnage, choix2 int) {
	item := merchantItems[choix2-1]
	clearTerminal()
	if p.Money < 1 {
		fmt.Println("Tu n'a plus d'argent")
		return
	} else if len(p.PersoInv) <= 10 {
		p.Money -= 1
		fmt.Printf("Vous avez acheté : %s vous avez : %d pièces\n", item.nomItem, p.Money)
		addInventaire(p, "Plume de Corbeau", "Permet de créer de l'équipement", 1)
		return
	}
}
