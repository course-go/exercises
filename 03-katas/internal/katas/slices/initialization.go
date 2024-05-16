package slices

type Director string

type Movie struct {
	Director Director
}

func GetMovies() []Movie {
	return []Movie{
		{
			Director: Director("Christopher Nolan"),
		},
		{
			Director: Director("Steven Spielberg"),
		},
		{
			Director: Director("Quentin Tarantino"),
		},
		{
			Director: Director("Martin Scorsese"),
		},
	}
}

func ProcessDirectors(directors []Director) {}
