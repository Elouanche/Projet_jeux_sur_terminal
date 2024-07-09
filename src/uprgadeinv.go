package main

import "fmt"

//Permets d'ajouter de la place à l'inventaire
func UpgradeInventorySlot(p *Personnage, choix2 int) {
	item := merchantItems[choix2-1]
	clearTerminal()
	if p.Money < 30 {
		fmt.Println("Tu n'a plus d'argent")
		return
	} else if len(p.PersoInv) <= p.maxSize {
		p.Money -= 30
		fmt.Printf("Vous avez acheté : %s vous avez : %d pièces\n", item.nomItem, p.Money)
		p.maxSize += 10
		addInventaire(p, "Sac à dos magique", "Permet d'augmenter le nombre de slot d'inventaire", 1)
		return
	}
	return
}
