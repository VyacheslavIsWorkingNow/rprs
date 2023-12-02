package main

import (
	"fmt"
	"github.com/VyacheslavIsWorkingNow/rprs/lab4/philisopher"
)

func main() {

	cm := philisopher.NewCelebratoryMeal()
	cm.InitCelebratoryMeal()

	fmt.Println("Starting for 3 times")
	cm.RunCelebratoryMealNTimes(3)
	cm.Wait()

	//fmt.Println("Starting for 3 seconds")
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	//defer cancel()
	//cm.RunCelebratoryWithTimeout(ctx)
	//cm.Wait()
}
