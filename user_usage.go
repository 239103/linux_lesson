package main

import "fmt"

func math() {
	var result, i, n, m int64
	for i = 1; i < 100000; i++ {
		for n = 1; n < 100000; n++ {
			for m = 1; m < 100000; m++ {
				result += m * n * i
			}
		}
	}
	fmt.Println("result=", result)
}

func webaccess() {

}

func main() {
	go math()
}
