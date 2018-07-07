package kiblin

import (
	"errors"
	"math"
)

type Kiblin struct {
	hunger uint
	health uint
	mood   uint

	gender string
	dna    string

	mouth chan string
	ear   chan string
}

func Hatch(dna string) (*Kiblin, error) {
	return nil, errors.New("genetic defect")
}

func (k *Kiblin) tick() {
	if k.hunger > 100 {
		k.health -= uint(math.Ceil(float64(k.hunger) / 100))
	}
}

func (k *Kiblin) Feed(food Food) {

	

	k.health -= food.Vitamins()
}

func (k *Kiblin) Pet() {
	k.mood++
}
