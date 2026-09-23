package main

import "fmt"

type DayOfWeek int

const (
	Monday DayOfWeek = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func getWorkDayStatus(day DayOfWeek) string {
	switch day {
	case Monday, Tuesday, Wednesday, Thursday, Friday:
		return "Workday"
	case Saturday, Sunday:
		return "Weekend"
	default:
		return "Invalid day"
	}

}


type officeStatus string

const (
	Open   officeStatus = "Open"
	Closed officeStatus = "Closed"
)


func main() {
	fmt.Println(getWorkDayStatus(Monday)) // Output: Workday
}
