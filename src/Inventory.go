package main

import (
	"fmt"
	"time"
)

//Permets de voir l'inventaire
func displayInventory(p *Personnage) {
	clearTerminal()
	fmt.Println("\nContenu de l'inventaire :")
	for i, item := range p.PersoInv {
		fmt.Printf("%d. %s (x%d)\n", i+1, item.VarInventaire.nomItem, item.VarInventaire.quantity)
	}
	fmt.Println("\nVotre équipement:")
	// Accédez aux champs Tete, Torse et Pieds à partir de Equipement
	fmt.Println("Tete: ", p.PersoEquipement.head)
	fmt.Println("Torse: ", p.PersoEquipement.body)
	fmt.Println("Pieds: ", p.PersoEquipement.boots)

}

//Permet d'enlever les items de l'inventaire du joueur une fois utilisé
func removeItem(p *Personnage, itemName string, quantity int) {
	for i, item := range p.PersoInv {
		if item.VarInventaire.nomItem == itemName {
			p.PersoInv[i].VarInventaire.quantity -= quantity
			if p.PersoInv[i].VarInventaire.quantity <= 0 {
				p.PersoInv = append(p.PersoInv[:i], p.PersoInv[i+1:]...)
			}
			return
		}
	}
}

// Fonction pour accéder à l'inventaire
func accessInventory(p *Personnage, mobGobelinEntrainement *Monstre) {
	for {
		displayInventory(p)
		fmt.Print("Choisissez un item à utiliser (ou 0 pour retourner au menu précédent) : ")
		var choix2 int
		fmt.Scan(&choix2)
		if choix2 == 0 {
			return
		}
		if choix2 > 0 && choix2 <= len(p.PersoInv) {
			item := &p.PersoInv[choix2-1].VarInventaire
			switch item.nomItem {
			case "Heal Potion":
				useHealPotion(p, item)
			case "Poison Potion":
				usePoisonPotion(p, item, mobGobelinEntrainement)
			case "SpellBook":
				SpellBook(p)
			default:
				fmt.Println("Cet item n'est pas utilisable.")
			}
			removeItem(p, item.nomItem, 1) 
		}
	}
}
//fonction pour utiliser les Heal Potion
func useHealPotion(p *Personnage, h *Item) { 
	p.PvActuels += 50
	if p.PvActuels > p.PvMax {
		p.PvActuels = p.PvMax
	}
}

//fonction pour utiliser une potion de poison
func usePoisonPotion(p *Personnage, po *Item, mobGobelinEntrainement *Monstre) { 
	for x := 0; x < 3; x++ {
		mobGobelinEntrainement.Pv_Actuels -= 10
		time.Sleep(1 * time.Second)
	}
}

 //fonction permetant d'ajouter des items à l'inventaire
func addInventaire(p *Personnage, ItemNom string, Effet string, Quantity int) {
	if len(p.PersoInv) >= p.maxSize {
		clearTerminal()
		fmt.Printf("==================================\nVotre inventaire est plein. Vous ne pouvez plus prendre d'objet.\n==================================\n")
		time.Sleep(2 * time.Second)
		return
	} else {
		for i, inv := range p.PersoInv {
			if inv.VarInventaire.nomItem == ItemNom && p.PersoInv[i].VarInventaire.quantity < 10 {
				p.PersoInv[i].VarInventaire.quantity += 1
				return
			}
		}
		NewItem := inventaire{VarInventaire: Item{nomItem: ItemNom, effet: Effet, quantity: Quantity}} 
		p.PersoInv = append(p.PersoInv, NewItem)                                                      
		return
	}
}
