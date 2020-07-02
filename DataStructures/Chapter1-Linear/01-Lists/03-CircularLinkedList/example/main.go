// Package main provides ...
package main

import "fmt"

type People struct {
	number     int
	nextpeople *People
}

type PeopleCircle struct {
	len        int
	headPeople *People
}

func NewPeopleCircle() *PeopleCircle {
	return &PeopleCircle{}
}

func (p *PeopleCircle) AddPeople(number int) {
	people := &People{number: number}
	if p.len == 0 {
		p.headPeople = people
		p.headPeople.nextpeople = people
		p.len++
	} else {
		lastPeople := new(People)
		for lastPeople = p.headPeople; ; lastPeople = lastPeople.nextpeople {
			if lastPeople.nextpeople == p.headPeople {
				break
			}
		}
		lastPeople.nextpeople = people
		people.nextpeople = p.headPeople
		p.len++
	}
}

func GenerateCircle(number int) *PeopleCircle {
	p := NewPeopleCircle()
	if number < 1 {
		fmt.Println("人数不够怎么玩")
		return nil
	} else {
		for i := 1; i < number+1; i++ {
			p.AddPeople(i)
		}
		return p
	}
}

func (p *PeopleCircle) FindPrevPeople(number int) *People {
	node := new(People)
	for node = p.headPeople; node.nextpeople != p.headPeople; node = node.nextpeople {
		if node.nextpeople.number == number {
			break
		}
	}
	return node
}

func (p *PeopleCircle) FindPeople(number int) *People {
	node := new(People)
	for node = p.headPeople; node.nextpeople != p.headPeople; node = node.nextpeople {
		if node.number == number {
			break
		}
	}
	return node
}

// 6 5 4 3 2 1  5 out
// 6 4 3 2 1    4 out
// 6 3 2 1		6 out
// 3 2 1		2 out
// 3 1			3 out
// 1			1 out

func GameStart(size, times, startnumber int) {
	p := GenerateCircle(size)
	pending := new(People)
	startpeople := p.FindPeople(startnumber)
	pending = startpeople
	for {
		for i := 1; i <= times; i++ {
			if i == 1 {
				continue
			} else {
				pending = pending.nextpeople
			}
		}
		fmt.Printf("第%d号出局\n", pending.number)
		if p.len == 1 {
			break
		}
		prevPeople := p.FindPrevPeople(pending.number)
		prevPeople.nextpeople = pending.nextpeople
		p.len--
		pending = pending.nextpeople
	}
}

func main() {
	GameStart(6, 5, 1)
}
