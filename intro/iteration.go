package main

import "fmt"

func main() {
	// go has no other loops, only a for that can be used as a while and also for

	// while i <= 3
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// for couning to 3
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// for range count
	for i := range 3 {
		fmt.Println("range", i)
	}

	// while true
	for {
		fmt.Println("loop")
		break
	}

	// for range count with continue
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
