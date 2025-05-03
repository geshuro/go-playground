package main

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

func main() {
	course := Course{
		Name:    "Go Programming",
		Price:   29.99,
		IsFree:  false,
		UserIDs: []uint{1, 2, 3},
		Classes: map[uint]string{
			1: "Introduction to Go",
			2: "Advanced Go",
			3: "Go for Web Development",
		},
	}

	println("Course Name:", course.Name)
	println("Course Price:", course.Price)
	println("Is Course Free:", course.IsFree)
	println("User IDs enrolled in the course:", course.UserIDs)
	for id, class := range course.Classes {
		println("Class ID:", id, "Class Name:", class)
	}
}
