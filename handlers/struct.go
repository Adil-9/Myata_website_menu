package handlers

type Position struct {
	Id          int
	Dish        string
	Description string
	Price       int
	Category    string
}

type Items struct {
	Category string
	Dishes   []Position
}
