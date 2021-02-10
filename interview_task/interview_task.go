//Package interview_task for ituniver18
package interview_task

import "fmt"

//ShowFizzBuzz returns filtered number
func ShowFizzBuzz(numbers []int) {
	for tIndex := range numbers {
		if numbers[tIndex]%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if numbers[tIndex]%5 == 0 {
			fmt.Println("Buzz")
		} else if numbers[tIndex]%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(numbers[tIndex])
		}
	}
}
