package main

import "fmt"

func main() {

	arr := [5]int{10, 20, 10, 40, 20}

	var count int
	for i := 0; i < len(arr); i++ {
		//fmt.Println(arr[i])
		count = 0
		for j := 0; j < len(arr); j++ {
			if i != j {
				if arr[i] == arr[j] {
					count = count + 1
				}
			}

		}

		if count == 0 {
			fmt.Println(arr[i])
		}

	}

}
