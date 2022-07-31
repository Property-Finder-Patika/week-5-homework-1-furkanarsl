package main

import "fmt"

type IUse interface {
	Use()
}

type Program struct {
	Name string
}

func (p *Program) Use() {
	fmt.Printf("%v: Using the program...\n", p.Name)
}

type ProgramProxy struct {
	program      *Program
	usrCount     int
	LicenseLimit int
}

func (p *ProgramProxy) Use() {
	// If we dont have a program create it.
	if p.program == nil {
		p.program = &Program{Name: "Program1"}
	}
	// Check the current user count. This would be a bit more complicated if we actually wanted to implement concurrent users.
	// To make the example simpler we are going to increment the user count by one every time the program is used and assume that they keep using it.
	if p.usrCount >= p.LicenseLimit {
		fmt.Println("Looks like you have reached your license limit. Please try using the program later.")
	} else {
		p.program.Use()
		p.usrCount++
	}
}

func main() {
	p := ProgramProxy{LicenseLimit: 5}

	for i := 0; i < 10; i++ {
		p.Use()
	}
}
