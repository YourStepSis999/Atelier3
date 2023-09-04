package Atelier3

import (
	"math/rand"
)

// le lien du github du package est github.com/YourStepSis999/Atelier3

func ArrayGenerate( nb int) []int{
	array := []int{}
	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(10000 - 1) + 10000
	}
	return array
}
