package main

import "fmt"

func readi(Ccity string, newmap map[string][]int64) int64 {
	pop, exists := newmap[Ccity]
	//err := ""
	if exists {
		return pop[0]
	}
	fmt.Println("City not Found")
	return 0
}

func ToDel(Ccity string, newmap map[string][]int64) int64 {
	_, exists := newmap[Ccity]
	if exists {
		delete(newmap, Ccity)
	}
	return 0
}
func main() {
	//creating a map
	newmap := map[string][]int64{
		"Hoshiarpur":  {210938129},
		"Bhubneshwar": {21093456},
		"KOTA":        {21093324},
		"PUNE":        {21093432},
	}

	{
		if newmap == nil {
			fmt.Println("Empty")
		}
		fmt.Println("Not Empty")
	}
	fmt.Println(newmap)

	Ccity := "PUNE"
	fmt.Println(readi(Ccity, newmap))
	fmt.Println(ToDel(Ccity, newmap))
	fmt.Println(readi(Ccity, newmap))
}
