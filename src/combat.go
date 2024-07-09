package main

import (
	"fmt"
	"math/rand"
	"time"
)
//Fonction pour les textes de combats
func textCombatEntrainement() {
	clearTerminal()
	fmt.Println("=====================================")
	fmt.Println("Bienvenue dans l'arène d'entrainement")
	fmt.Println("=====================================")
	time.Sleep(2 * time.Second)
	clearTerminal()
	fmt.Printf("C'est l'heure du-du-du-du-duel\n")
	time.Sleep(1 * time.Second)
}

//Fonction pour les textes de combats
func textCombat() {
	clearTerminal()
	fmt.Println("=====================================")
	fmt.Println("Bienvenue en combat !")
	fmt.Println("=====================================")
	time.Sleep(2 * time.Second)
	clearTerminal()
	fmt.Printf("C'est l'heure du-du-du-du-duel\n")
	time.Sleep(1 * time.Second)
}

//Fonction pour les textes de combats
func textArrene() {
	clearTerminal()
	fmt.Println("=====================================")
	fmt.Println("Bienvenue dans le hub des combats !")
	fmt.Println("=====================================")
	fmt.Println("Que veux-tu faire ? 1.Entrainement/2.Mission/3.Quitter")
}

//Permets de définir l'initiative du joueur
func Initiative() int {
	return rand.Intn(20) + 1
}
func initiativeStart(p *Personnage, mobGobelinEntrainement *Monstre) (int, int) {
	p.initiative = Initiative()
	mobGobelinEntrainement.initiative = Initiative()
	return p.initiative, mobGobelinEntrainement.initiative
}

//Choix du joueur
func areneCombat(p *Personnage, mobGobelinEntrainement *Monstre) {
	var choix5 int

	for {
		textArrene()
		fmt.Scan(&choix5)

		switch choix5 {
		case 1:
			trainingFight(p, mobGobelinEntrainement)
		case 2:
			mission(p, mobGobelinEntrainement)
		case 3:
			return
		default:
			fmt.Println("Choix invalide. Veuillez choisir entre 1 et 3.")
		}
	}
}

//Fonction pour les combats d'entrainement
func trainingFight(p *Personnage, mobGobelinEntrainement *Monstre) {
outerLoop:
	for {
		x := 1
		mobGobelinEntrainement.Pv_Actuels = mobGobelinEntrainement.Pv_Max
		PvEssaie := p.PvActuels
		initiativeStart(p, mobGobelinEntrainement)
		if mobGobelinEntrainement.initiative > p.initiative {
			fmt.Println("Le gobelin a plus d'initiave que toi, il ta donc infligé des dégats en premier.")
			PvEssaie -= 5
			fmt.Printf("Tu commence donc avec %d pv", PvEssaie)
			time.Sleep(5 * time.Second)
			clearTerminal()
		} else {
			fmt.Println("Tu as plus d'initiative que le gobelin, tu attaque donc en premier")
			time.Sleep(5 * time.Second)
			clearTerminal()
		}

		textCombatEntrainement()

		for {
			clearTerminal()
			fmt.Printf("Tu es manche %d\n", x)
			fmt.Printf("Le gobelin a %d pv sur %d pv\n", mobGobelinEntrainement.Pv_Actuels, mobGobelinEntrainement.Pv_Max)
			fmt.Printf("Tu as %d pv sur %d pv\n", PvEssaie, p.PvMax)
			fmt.Printf("Que souhaitez-vous faire ? (1.Attaque / 2.Regarde l'inventaire / 3. Utiliser un skill / 4. Arrêter)\n")
			var choix4 int
			fmt.Scan(&choix4)

			switch choix4 {
			case 1:
				PvEssaie, mobGobelinEntrainement.Pv_Actuels = combatEntrainement(p, mobGobelinEntrainement, PvEssaie, x)
				x += 1

				if PvEssaie <= 0 || mobGobelinEntrainement.Pv_Actuels <= 0 {
					x = 1
					PvEssaie = p.PvActuels
					mobGobelinEntrainement.Pv_Actuels = mobGobelinEntrainement.Pv_Max
					break outerLoop
				}
			case 2:
				accessInventory(p, mobGobelinEntrainement)
			case 3:
				AccessSkills(p, mobGobelinEntrainement)
			case 4:
				return
			}

		}
	}
}

//Fonction pour l'entrainement
func combatEntrainement(p *Personnage, mobGobelinEntrainement *Monstre, PvEssaie int, x int) (int, int) {
	playerTurn(mobGobelinEntrainement)
	if mobGobelinEntrainement.initiative > p.initiative && x == 1 {
		return PvEssaie, mobGobelinEntrainement.Pv_Actuels
	} else {
		gobelinPatern(mobGobelinEntrainement, x, &PvEssaie)
	}
	if PvEssaie <= 0 {
		fmt.Println("Bien essayé mais tu as perdu ! Reviens quand tu seras plus fort !")
		time.Sleep(3 * time.Second)
		return PvEssaie, mobGobelinEntrainement.Pv_Actuels
	} else if mobGobelinEntrainement.Pv_Actuels <= 0 {
		fmt.Printf("Tu as gagné ! Bravo ! (même si c'était vraiment facile...)\n")
		time.Sleep(3 * time.Second)
		return PvEssaie, mobGobelinEntrainement.Pv_Actuels
	} else {
		fmt.Printf("Prépare-toi pour le prochain round\n")
		time.Sleep(2 * time.Second)
		return PvEssaie, mobGobelinEntrainement.Pv_Actuels
	}

}
func playerTurn(mobGobelinEntrainement *Monstre) {
	fmt.Println("Tu as infligé 5 dégats à l'ennemi")
	mobGobelinEntrainement.Pv_Actuels -= 5
	time.Sleep(1 * time.Second)
	clearTerminal()
}
func gobelinPatern(mobGobelinEntrainement *Monstre, x int, PvEssaie *int) {
	if x%3 == 0 {
		fmt.Println("Le monstre t'a infligé", mobGobelinEntrainement.Point_attaque*2, "dégats ")
		*PvEssaie -= (mobGobelinEntrainement.Point_attaque * 2)
		time.Sleep(1 * time.Second)
		clearTerminal()
		return
	} else {
		fmt.Printf("Le monstre t'a infligé %d dégats \n", mobGobelinEntrainement.Point_attaque)
		*PvEssaie -= mobGobelinEntrainement.Point_attaque
		time.Sleep(1 * time.Second)
		clearTerminal()
		return
	}
}

//Fonction pour les missions
func mission(p *Personnage, mobGobelinEntrainement *Monstre) {
	mobGobelinEntrainement.Pv_Actuels = mobGobelinEntrainement.Pv_Max
	initiativeStart(p, mobGobelinEntrainement)
	if mobGobelinEntrainement.initiative > p.initiative {
		fmt.Println("Le gobelin a plus d'initiave que toi, il ta donc infligé des dégats en premier.")
		p.PvActuels -= 5
		fmt.Printf("Tu commence donc avec %d pv", p.PvActuels)
		time.Sleep(2 * time.Second)
		clearTerminal()
	} else {
		fmt.Println("Tu as plus d'initiative que le gobelin, tu attaque donc en premier")
		time.Sleep(2 * time.Second)
		clearTerminal()
	}

	textCombat()

	for x := 1; mobGobelinEntrainement.Pv_Actuels > 0 && p.PvActuels > 0; x++ {
		clearTerminal()
		fmt.Printf("Tu es manche %d\n", x)
		fmt.Printf("Le gobelin a %d pv sur %d pv\n", mobGobelinEntrainement.Pv_Actuels, mobGobelinEntrainement.Pv_Max)
		fmt.Printf("Tu as %d pv sur %d pv\n", p.PvActuels, p.PvMax)
		fmt.Printf("Que souhaitez-vous faire ? (1.Attaque / 2.Regarde l'inventaire / 3. Utiliser un skill / 4. Arrêter)\n")
		var choix6 int
		fmt.Scan(&choix6)

		switch choix6 {
		case 1:
			playerTurn(mobGobelinEntrainement)
			p.PvActuels, mobGobelinEntrainement.Pv_Actuels = combat(p, mobGobelinEntrainement, x)
		case 2:
			accessInventory(p, mobGobelinEntrainement)
			p.PvActuels, mobGobelinEntrainement.Pv_Actuels = combat(p, mobGobelinEntrainement, x)
		case 3:
			AccessSkills(p, mobGobelinEntrainement)
			p.PvActuels, mobGobelinEntrainement.Pv_Actuels = combat(p, mobGobelinEntrainement, x)
		case 4:
			return
		default:
			fmt.Println("Veuillez choisir une option valide")
		}

	}
}

//Fonction de combat
func combat(p *Personnage, mobGobelinEntrainement *Monstre, x int) (int, int) {
	if mobGobelinEntrainement.initiative > p.initiative && x == 1 {
		return p.PvActuels, mobGobelinEntrainement.Pv_Actuels
	} else if mobGobelinEntrainement.Pv_Actuels > 0 {
		gobelinPatern(mobGobelinEntrainement, x, &p.PvActuels)
	}
	if p.PvActuels <= 0 {
		return p.PvActuels, mobGobelinEntrainement.Pv_Actuels
	} else if mobGobelinEntrainement.Pv_Actuels <= 0 {
		fmt.Printf("Tu as gagné ! Bravo ! (même si c'était vraiment facile...)\n")
		p.experience += 120
		experience(p)
		time.Sleep(3 * time.Second)
		return p.PvActuels, mobGobelinEntrainement.Pv_Actuels
	} else {
		fmt.Printf("Prépare-toi pour le prochain round\n")
		time.Sleep(2 * time.Second)
		return p.PvActuels, mobGobelinEntrainement.Pv_Actuels
	}

}
