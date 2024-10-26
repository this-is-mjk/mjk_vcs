package main

type Mentees struct {
	Name   string
	Branch string
	Audit  bool // True means auditting
	Age    int
}

type VCS []Mentees

func main() {
  m1 := Mentees{"Atharv", "CSE", true, 100}
  m2 := Mentees{"Ravi", "CSE", true, 0}
  m3 := Mentees{"Medha", "CHM", false, 50}
  m1.print()
  m2.print()
  m3.print()
  project := VCS{m1, m2, m3}
  project.PrintAudits()
}
