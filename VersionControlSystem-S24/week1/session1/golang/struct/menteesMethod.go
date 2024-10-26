package main

import (
	"fmt"
	"log"
)

type Error struct {
  Message string
  code int
}

func (m *Mentees) print() {
  // for i, v := range m {
    // s := fmt.Sprintf("%s is very intelligent and smart", v)
    // fmt.Println(i, s)
  // }
  fmt.Println(*m)
}

func (m *Mentees) TellIfAudit() (bool, *Error) {
  if m.Audit {
    return true, nil
  } else {
    err := Error{"Yeh kya kardiya Bhai ?", 404}
    return false, &err
  }
}

func (v VCS) PrintAudits() {
  for i, m := range v {
    fmt.Println(i)
    _, err := m.TellIfAudit()
    if err == nil {
      s := fmt.Sprintf("%s is good", m.Name)
      fmt.Println(s)
    } else {
      log.Fatal(err)
    }
  }
  fmt.Println("Fatal error nhi hua bhai")
}
