package main

import "go-playgrond/course"

func main() {
	cse := course.New("Go Programming", 0, false)
	cse.UserIDs = append(cse.UserIDs, 1, 2, 3)
	cse.Classes = map[uint]string{
		1: "Introduction to Go",
		2: "Advanced Go",
		3: "Go for Web Development",
	}
	cse.String()
}
