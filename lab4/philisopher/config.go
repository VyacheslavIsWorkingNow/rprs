package philisopher

import "sync"

const (
	NumPhilosophers = 3
	MaxRandTime     = 500
)

type philosopher struct {
	id                  int
	leftFork, rightFork *sync.Mutex
}

func newPhilosopher(forks []*sync.Mutex, index int) *philosopher {
	return &philosopher{
		id:        index + 1,
		leftFork:  forks[index],
		rightFork: forks[(index+1)%NumPhilosophers],
	}
}

type CelebratoryMeal struct {
	philosophers []*philosopher
	forks        []*sync.Mutex
	wg           sync.WaitGroup
}

func NewCelebratoryMeal() *CelebratoryMeal {
	return &CelebratoryMeal{
		forks:        make([]*sync.Mutex, NumPhilosophers),
		philosophers: make([]*philosopher, NumPhilosophers),
		wg:           sync.WaitGroup{},
	}
}

func (cm *CelebratoryMeal) InitCelebratoryMeal() {
	for i := 0; i < NumPhilosophers; i++ {
		cm.forks[i] = &sync.Mutex{}
	}

	for i := 0; i < NumPhilosophers; i++ {
		cm.philosophers[i] = newPhilosopher(cm.forks, i)
	}
}
