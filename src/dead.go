package main

import (
	"fmt"
	"time"
)

//Fonction qui vérifie si le joueur est mort
func dead(p1 *Personnage) { // Fonction qui vérifie si le joueur est mort
	if p1.PvActuels <= 0 {
		clearTerminal()
		fmt.Printf("VOUS ÊTES MORTS !")
		time.Sleep(5 * time.Second)
		clearTerminal()
		p1.PvActuels = p1.PvMax / 2
		fmt.Printf("Points de vie actuels : %d sur %d pv ", p1.PvActuels, p1.PvMax) 
	}
}
