package main

//structure information item/ début item
type Item struct { 
	nomItem     string
	effet       string
	quantity    int
	point       int
	description string
}

var forgeItems = []Item{
	{nomItem: "Chapeau de l'aventurier", description: "Chapeau pour les aventuriers", quantity: 1},
	{nomItem: "Tunique de l'aventurier", description: "Tunique pour les aventuriers", quantity: 1},
	{nomItem: "Bottes de l'aventurier", description: "Bottes pour les aventuriers", quantity: 1},
}

var chapeauAventurierRecipe = Recipe{
    ItemName: "Chapeau de l'aventurier",
    Resources: []struct {
        ResourceName string
        Quantity     int
    }{
        {"Plume de Corbeau", 1},
        {"Cuir de Sanglier", 1},
    },
}

var tuniqueAventurierRecipe = Recipe{
    ItemName: "Tunique de l'aventurier",
    Resources: []struct {
        ResourceName string
        Quantity     int
    }{
        {"Fourrure de Loup", 2},
        {"Peau de Troll", 1},
    },
}

var bottesAventurierRecipe = Recipe{
    ItemName: "Bottes de l'aventurier",
    Resources: []struct {
        ResourceName string
        Quantity     int
    }{
        {"Fourrure de Loup", 1},
        {"Cuir de Sanglier", 1},
    },
}

var chapeauAventurierStats = Equipment{
    nomEquipement: "Chapeau de l'aventurier",
    stats: struct {
        attaque int
        defense int
    }{
        attaque: 0,    // Vous pouvez définir les statistiques appropriées
        defense: 5,    // pour chaque équipement
    },
}

var tuniqueAventurierStats = Equipment{
    nomEquipement: "Tunique de l'aventurier",
    stats: struct {
        attaque int
        defense int
    }{
        attaque: 0,
        defense: 10,
    },
}

var bottesAventurierStats = Equipment{
    nomEquipement: "Bottes de l'aventurier",
    stats: struct {
        attaque int
        defense int
    }{
        attaque: 0,
        defense: 3,
    },
}

type Recipe struct {
	ItemName  string
	Resources []struct {
		ResourceName string
		Quantity     int
	}
}

var forgeRecipes = []Recipe{
	{
		ItemName: "Chapeau de l'aventurier",
		Resources: []struct {
			ResourceName string
			Quantity     int
		}{
			{"Plume de Corbeau", 1},
			{"Cuir de Sanglier", 1},
		},
	},
	{
		ItemName: "Tunique de l'aventurier",
		Resources: []struct {
			ResourceName string
			Quantity     int
		}{
			{"Fourrure de Loup", 2},
			{"Peau de Troll", 1},
		},
	},
	{
		ItemName: "Bottes de l'aventurier",
		Resources: []struct {
			ResourceName string
			Quantity     int
		}{
			{"Fourrure de Loup", 1},
			{"Cuir de Sanglier", 1},
		},
	},
}

