package philisopher

import (
	"context"
	"log"
	"sync"
)

func PhilosopherRoutineNTimes(ph *philosopher, wg *sync.WaitGroup, N int) {
	defer wg.Done()
	for i := 0; i < N; i++ {
		ph.doAll()
	}
}

func PhilosopherRoutineWithTimeout(ctx context.Context, ph *philosopher, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			log.Println("time out")
			return
		default:
		}
		ph.doAll()
	}
}

func (cm *CelebratoryMeal) RunCelebratoryMealNTimes(N int) {
	for i := 0; i < NumPhilosophers; i++ {
		cm.wg.Add(1)
		go PhilosopherRoutineNTimes(cm.philosophers[i], &cm.wg, N)
	}
}

func (cm *CelebratoryMeal) RunCelebratoryWithTimeout(ctx context.Context) {
	for i := 0; i < NumPhilosophers; i++ {
		cm.wg.Add(1)
		go PhilosopherRoutineWithTimeout(ctx, cm.philosophers[i], &cm.wg)
	}
}

func (cm *CelebratoryMeal) Wait() {
	cm.wg.Wait()
}
