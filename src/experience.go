package main

//Fonction d'experience du joueur
func experience(p *Personnage) {
	x := p.experienceMax
	if p.experience >= x {
		p.Niveau += 1
		p.experience -= p.experienceMax
		p.experienceMax = x * 2
	}
}
