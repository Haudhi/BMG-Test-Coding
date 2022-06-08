package main

import "fmt"

func input(N int) []string {
	var result []string
	for i := 1; i <= N; i++ {
		if i%3 == 0 {
			result = append(result, `"`+"Frontend"+`",`)
			// fmt.Println("Frontend")
		} else if i%5 == 0 {
			result = append(result, `"`+"Backend"+`",`)
			// fmt.Println("Backend")
		} else if i%3 == 0 && i%5 == 0 {
			result = append(result, `"`+"Frontend Backend"+`",`)
			// fmt.Println("Frontend Backend")
		} else {
			result = append(result, `"`+fmt.Sprint(i)+`",`)
			// fmt.Println(i)
		}

	}
	return result
}

func main() {
	fmt.Println(input(50))
}
