package helpers

import "fmt"

func Check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}