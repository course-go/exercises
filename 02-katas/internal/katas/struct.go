package katas

type Person struct {
	Name string
}

// Rename renames a person.
// However, it does not seem to work.
// Fix it.
func (p Person) Rename(name string) {
	p.Name = name
}
