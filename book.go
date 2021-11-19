package restBook

type Book struct {
	Id            int     `json:"-"`
	Title         string  `json:"title" binding:"required"`
	Author        string  `json:"author" binding:"required"`
	Price         float32 `json:"price" binding:"required"`
	YearOfPublish int     `json:"year" binding:"required" db:"year_of_publish"`
	Count         int     `json:"count" binding:"required"`
}
