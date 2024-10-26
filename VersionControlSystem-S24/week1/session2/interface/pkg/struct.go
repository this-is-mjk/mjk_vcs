package pkg

type Mentees struct {
	Name   string
	Audit  bool
	Branch string
}

type Mentor struct {
	Name  string
	Batch string
}

type Project struct {
  Mentor
  Mentees
}

type Student interface {
  Print()  
}

type circle struct {
  radius float64
}

type rectangle struct {
  width float64
  breadth int
}

// type InterfaceX interface {
  // random() bool
// }


type shape interface {
  area() float64
  // InterefaceX  => Embedded interface
}


