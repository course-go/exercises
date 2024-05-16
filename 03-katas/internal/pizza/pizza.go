package pizza

type Topping string

const (
	Prosciutto Topping = "prosciutto"
	Salami     Topping = "salami"
	Ham        Topping = "ham"
	Pineapple  Topping = "pineapple" // Do not even think about it
)

type Size string

const (
	Regular Size = "32cm"
	Large   Size = "40cm"
)

type Base string

const (
	Sugo  Base = "sugo"
	Cream Base = "cream"
)

type Pizza struct {
	Base     Base
	Size     Size
	Toppings []Topping
}
