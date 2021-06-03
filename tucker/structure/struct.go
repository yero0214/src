package main

import "fmt"

type Student struct {
	name  string
	class int

	grade Grade
}

type Grade struct {
	name  string
	grade string
}

func (s Student) ViewGrade() {
	fmt.Println(s.grade)
}

func main() {
	var s Student
	s.name = "Dave"
	s.class = 1

	s.grade.name = "math"
	s.grade.grade = "c"

	s.ViewGrade()
}
