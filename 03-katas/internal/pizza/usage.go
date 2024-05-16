package pizza

// makeHawaiPizza creates a hawai pizza.
// Its sole purpose is to demonstrate the traditional builder usage.
// But man, hawai? Seriously? From all the pizzas you chose hawai?
func makeHawaiPizza() (pizza *Pizza, err error) {
	return NewBuilder().
		SetBase(Sugo).
		SetSize(Regular).
		AddTopping(Ham).
		AddTopping(Pineapple).
		Build()
}
