package philisopher

import (
	"fmt"
	"math/rand"
	"time"
)

func (p *philosopher) think() {
	fmt.Printf("Философ %d размышляет\n", p.id)
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(MaxRandTime)+1))
}

func (p *philosopher) eat() {
	fmt.Printf("Философ %d ест\n", p.id)
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(MaxRandTime)+1))
}

func (p *philosopher) getForks() {
	p.leftFork.Lock()
	p.rightFork.Lock()
}

func (p *philosopher) putForks() {
	p.leftFork.Unlock()
	p.rightFork.Unlock()
}

func (p *philosopher) doAll() {
	p.think()
	p.getForks()
	p.eat()
	p.putForks()
}
