package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the Time App")

	currentTime := time.Now()
	fmt.Println("Current Time is:", currentTime)
	fmt.Println("===================================")

	fmt.Println("Formatted time is : ", currentTime.Format("02-01-2026 15:04:05 Monday"))
	fmt.Println("====================================")

	createdDate := time.Date(2024, time.March, 25, 10, 30, 0, 0, time.UTC)
	fmt.Println("Created Date is:", createdDate)
	fmt.Println("Formatted Created Date is : ", createdDate.Format("02-01-2006 15:04:05 Monday"))
}
