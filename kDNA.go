package kiblin

import "math/rand"

const mutationChance = 0.05

func mate(gene1 string, gene2 string) string {
	d := ""

	for i := 0; i < len(gene1); i++ {
		chance := rand.Float64()
		if chance > 0.5 {
			d += string(gene1[i])
		}
	}
	return d
}
