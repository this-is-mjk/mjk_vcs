package main

import (
	"fmt"
	"session2/pkg"
)

func main() {
	m := pkg.Mentees{"Rushabh", false, "Batti"}
	M := pkg.Mentor{"Yash", "Y22"}
	m.Print()
	M.Print()

	p := pkg.Project{pkg.Mentor{}, pkg.Mentees{}}
	fmt.Println(p)
}
