package main

import "fmt"

type Student struct {
	name string
	age  int

	grade string
	class string
}

func (s *Student) PrintGrade() {
	fmt.Println(s.class, s.grade)
}

func (s *Student) InputGrade(class string, grade string) {
	s.class = class
	s.grade = grade
}

func main() {
	var s Student
	s.class = "math"
	s.grade = "A"

	s.PrintGrade()
	s.InputGrade("history", "B")
	s.PrintGrade()
}
