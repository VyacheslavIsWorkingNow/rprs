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
	fmt.Printf("left fork for p: %d\n", p.id)
	p.leftFork.Lock()
	timer := time.NewTimer(time.Millisecond * time.Duration(rand.Intn(1000)+10))
	select {
	case <-timer.C:
		p.leftFork.Unlock()
		fmt.Printf("left put fork for p: %d\n", p.id)
		return
	case <-time.After(time.Millisecond * time.Duration(rand.Intn(10))):
		p.rightFork.Lock()
		fmt.Printf("right fork for p: %d\n", p.id)
		timer.Stop()
		return
	}
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
