package dish

type Dish struct {
	DishID           int64        `db:"dish_id" json:"dish_id"`
	DishName         string       `db:"name" json:"name"`
	Price            float64      `db:"price" json:"price"`
	IngredientsBytes []byte       `db:"ingredients"`
	Ingredients      []Ingredient `json:"ingredients"`
}

type RecieveDish struct {
	DishID   int64   `db:"dish_id" json:"dish_id"`
	DishName string  `db:"dish_name" json:"name"`
	Price    float64 `db:"price" json:"price"`
}

type Ingredient struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}
