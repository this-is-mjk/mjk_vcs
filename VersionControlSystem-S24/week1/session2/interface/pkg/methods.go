package pkg

import "fmt"

func (m *Mentees) Print() {
	fmt.Println(m.Name)
}

func (M *Mentor) Print() {
	fmt.Println(M.Name)
}

func PrintNameOnly(s *Student) {
  (*s).Print()
}

func (c *circle) area() float64 {
  return 3.14 * c.radius * c.radius
}

func (r *rectangle) area() float64 {
  return r.width * float64(r.breadth) 
}

func printArea(s *shape) {
  fmt.Println((*s).area())
}
