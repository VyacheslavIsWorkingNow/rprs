package philisopher

import (
	"log"
	"math/rand"
	"time"
)

func (p *philosopher) think() {
	log.Printf("Phylosoph %d think\n", p.id)
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(MaxRandTime)+1))
}

func (p *philosopher) eat() {
	log.Printf("Phylosoph %d eat\n", p.id)
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(MaxRandTime)+1))
}

func (p *philosopher) getForks() {
	log.Printf("Phylosoph %d get left fork\n", p.id)
	p.leftFork.Lock()
	timer := time.NewTimer(time.Millisecond * time.Duration(rand.Intn(1000)+10))
	select {
	case <-timer.C:
		p.leftFork.Unlock()
		log.Printf("Phylosoph %d put right fork\n", p.id)
		return
	case <-time.After(time.Millisecond * time.Duration(rand.Intn(10))):
		p.rightFork.Lock()
		log.Printf("Phylosoph %d get right fork\n", p.id)
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
