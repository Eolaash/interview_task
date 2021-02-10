package main

import "github.com/interview_task/interview_task"

func main() {
	tEndNum := 20
	tNumbers := make([]int, tEndNum, tEndNum)

	// Generate array
	for tIndex := range tNumbers {
		tNumbers[tIndex] = tIndex + 1
	}

	// Call task func
	interview_task.ShowFizzBuzz(tNumbers)
}
