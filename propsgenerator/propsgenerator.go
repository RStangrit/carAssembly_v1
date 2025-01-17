package propsgenerator

import (
	"time"

	"golang.org/x/exp/rand"
)

func GenerateRandomCarProperties() (int, int, int, string, string, int, int, string, string, int, bool) {
	rand.Seed(uint64(time.Now().UnixNano()))

	horsePowers := rand.Intn(301) + 100  // 100-400
	volume := rand.Intn(3001) + 1000     // 1000-4000
	torqueMoment := rand.Intn(501) + 200 // 200-700
	gearboxTypes := []string{"manual", "automatic", "CVT"}
	clutchTypes := []string{"dry", "wet", "double"}
	gearboxType := gearboxTypes[rand.Intn(len(gearboxTypes))]
	clutchType := clutchTypes[rand.Intn(len(clutchTypes))]
	rimSize := rand.Intn(11) + 15   // 15-25
	tireSize := rand.Intn(11) + 195 // 195-205
	bodyMaterials := []string{"steel", "aluminum", "carbon fiber"}
	bodyMaterial := bodyMaterials[rand.Intn(len(bodyMaterials))]
	bodyColors := []string{"red", "blue", "black", "white", "silver"}
	bodyColor := bodyColors[rand.Intn(len(bodyColors))]
	bodyTrunkVolume := rand.Intn(501) + 300 // 300-800
	bodyRoof := rand.Intn(2) == 1           // true/false

	return horsePowers, volume, torqueMoment, gearboxType, clutchType, rimSize, tireSize, bodyMaterial, bodyColor, bodyTrunkVolume, bodyRoof
}
