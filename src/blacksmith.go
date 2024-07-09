package main

import "fmt"

//Affiche le menu du forgeron
func forgeronMenu(p *Personnage) {
	clearTerminal()
	fmt.Println("Bienvenu chez le forgeron, voici ce que vous pouvez fabriquer :\n ")
	displayForgeItems() 

	fmt.Println("\nChoisisser un objet à fabrique (ou 0 pour retourner au menu précédent): ")
	var choix2 int
	fmt.Scan(&choix2)

	if choix2 == 0 {
		return
	}

	if choix2 > 0 && choix2 <= len(forgeItems) {
		item := &forgeItems[choix2-1]

		if canCraft(p, item) {
			craftItem(p, item)
		} else {
			fmt.Println("Choix invalide")
		}
	} else {
		fmt.Println("Choix invalide")
	}
}

//Affichage des objets du forgeron
func displayForgeItems() {
	fmt.Println("Objet fabriquable par le forgeron")
	for _, item := range forgeItems {
		fmt.Printf("%s\n", item.nomItem)
	}
}

func canCraft(p *Personnage, item *Item) bool {
	for _, recipe := range forgeRecipes {
		if recipe.ItemName == item.nomItem {
			for _, resource := range recipe.Resources {
				if !hasItem(p, resource.ResourceName, resource.Quantity) {
					return false
				}
			}
			return true
		}
	}
	return false
}

//Fonction qui permets de créer un item
func craftItem(p *Personnage, item *Item) {
	for _, recipe := range forgeRecipes {
		if recipe.ItemName == item.nomItem {
			if p.Money < 5 {
				fmt.Println("Vous n'avez pas assez d'argent pour fabriquer cela.")
				return
			}

			// Vérification des ressources
			for _, resource := range recipe.Resources {
				if !hasItem(p, resource.ResourceName, resource.Quantity) {
					fmt.Printf("Vous n'avez pas suffisamment de %s pour fabriquer cela.\n", resource.ResourceName)
					return
				}
			}

			// Retirer les ressources de l'inventaire
			for _, resource := range recipe.Resources {
				removeItem(p, resource.ResourceName, resource.Quantity)
			}

			// Création de l'équipement
			switch item.nomItem {
			case "Chapeau de l'aventurier":
				p.EquipementTete = Equipment{
					nomEquipement: "Chapeau de l'aventurier",
					stats: struct {
							attaque int
							defense int
					}{
						attaque: 0, 
						defense: 10, 
					},
				}
			case "Tunique de l'aventurier":
				p.EquipmentTorse = Equipment{
					nomEquipement: "Tunique de l'aventurier",
					stats: struct {
						attaque int
						defense int
					}{
						attaque: 0,  
						defense: 20, 
					},
				}
			case "Bottes de l'aventurier":
				p.EquipementPieds = Equipment{
					nomEquipement: "Bottes de l'aventurier",
					stats: struct {
						attaque int
						defense int
					}{
						attaque: 0,  
						defense: 15, 
					},
				}
			}

			p.Money -= 5
			fmt.Printf("Vous avez fabriqué un %s. Vous avez maintenant %d pièces d'or.\n", item.nomItem, p.Money)
			return
		}
	}
	fmt.Println("La recette n'a pas été trouvée.")
}

//Permet de savoir si le joueur à l'item
func hasItem(p *Personnage, itemName string, quantity int) bool {
	for _, item := range p.PersoInv {
		if item.VarInventaire.nomItem == itemName && item.VarInventaire.quantity >= quantity {
			return true
		}
	}
	return false
}
