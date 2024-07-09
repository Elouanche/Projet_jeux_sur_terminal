package main

//Structure et fonction pour initialiser un monstre
type Monstre struct {
	Nom           string
	Pv_Max        int
	Pv_Actuels    int
	Point_attaque int
	initiative    int
}

func initMonstre(Nom string, Pv_Max int, Pv_Actuels int, Point_attaque int, initiative int) Monstre {
	return Monstre{
		Nom:           Nom,
		Pv_Max:        Pv_Max,
		Pv_Actuels:    Pv_Actuels,
		Point_attaque: Point_attaque,
		initiative:    initiative,
	}
}
