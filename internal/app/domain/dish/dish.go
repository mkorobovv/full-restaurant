package dish

type Dish struct {
	DishID      int64        `db:"dish_id"`
	DishName    string       `db:"dish_name"`
	Price       float64      `db:"price"`
	Ingredients []Ingredient `db:"ingredients"`
}

type Ingredient struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}
