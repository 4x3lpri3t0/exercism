package lasagna

func PreparationTime(layers []string, layerPrepTimeMins int) int {
	if layerPrepTimeMins == 0 {
		layerPrepTimeMins = 2
	}

	return len(layers) * layerPrepTimeMins
}

func Quantities(layers []string) (noodles int, sauce float64) {
	noodles = 0
	sauce = 0.0
	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		} else if layer == "sauce" {
			sauce += 0.2
		}
	}
	return
}

func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, numberOfPortions int) []float64 {
	var quantitiesCopy = make([]float64, len(quantities))
	copy(quantitiesCopy, quantities)

	// Scale down to one portion + Scale up
	for i := 0; i < len(quantitiesCopy); i++ {
		quantitiesCopy[i] /= float64(2)
		quantitiesCopy[i] *= float64(numberOfPortions)
	}

	return quantitiesCopy
}
