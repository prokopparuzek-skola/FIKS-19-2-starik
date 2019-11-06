package main

import "fmt"

func main() {
	var T int
	fmt.Scanf("%d", &T)
	for ; T > 0; T-- {
		var M, N uint
		fmt.Scanf("%d%d", &M, &N)
		if N%2 == 0 {
			fmt.Println("Vyhraje starik")
		} else {
			if M%2 == 1 {
				fmt.Println("Vyhraju ja")
			} else {
				fmt.Println("Vyhraje starik")
			}
		}
	}
}
