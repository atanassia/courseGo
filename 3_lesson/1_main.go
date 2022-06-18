package main

import (
	"fmt"
)

const PassStatus string = "pass"
const FailStatus string = "fail"

type HealthCheck struct{
	ServiceID int
	status string
}

func GenerateCheck() []HealthCheck{
	var sl []HealthCheck
	for id := 0; id <= 5; id++{
		if id%2 == 0 {
			sl = append(sl, HealthCheck{id, PassStatus})
		}else{
			sl = append(sl, HealthCheck{id, FailStatus})
		}
	}

	return sl
}

func main() {
	slice_values := GenerateCheck()
	for _, value :=range slice_values{
		if value.status == PassStatus{
			fmt.Printf("id = %d, value = %s\n", value.ServiceID, value.status)
		}
	}

	fmt.Println("Конец программы")
}