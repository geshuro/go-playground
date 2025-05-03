package course

type course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

func (c course) String() {
	println("Course Name:", c.Name)
	println("Course Price:", c.Price)
	println("Is Course Free:", c.IsFree)
	println("User IDs enrolled in the course:", c.UserIDs)
	for id, class := range c.Classes {
		println("Class ID:", id, "Class Name:", class)
	}
}

func (c *course) changePrice(price float64) {
	c.Price = price
}
func New(name string, price float64, isFree bool) *course {
	if price == 0 {
		price = 24.5
	}
	return &course{
		Name:   name,
		Price:  price,
		IsFree: isFree,
	}
}
