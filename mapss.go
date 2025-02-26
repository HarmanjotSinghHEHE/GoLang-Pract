package main

import "fmt"

func main() {
	totalwins := map[string]int{}
	totalwins["Harman"] = 1
	totalwins["Lions"] = 2
	totalwins["Aradhya"] = 0
	fmt.Println(totalwins["Harman"])
	fmt.Println(totalwins["Lions"])
	fmt.Println(totalwins["Aradhya"])
	totalwins["Aradhya"]++
	fmt.Println("Updates", totalwins["Aradhya"])
	totalwins["Harman"] = 3
	fmt.Println(totalwins["Lions"])

}
