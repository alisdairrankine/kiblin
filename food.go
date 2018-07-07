package kiblin

type Food interface {
	SaltnessFactor() int
	SweetnessFactor() int
	SournessFactor() int

	Vitamins() int
	Calories() int
}
