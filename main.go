package main

import (
	"errors"
	"fmt"
)

func maxArray(array []int) (int, error) {
	if len(array) == 0 {
		return 0, errors.New("len of array in 0")
	}
	max := array[0]
	iMax := 0
	for i, num := range array {
		if num >= max {
			max = num
			iMax = i
		}
	}

	return iMax, nil
}

func getMaxCandies(pinatas []int) int {
	iMax, err := maxArray(pinatas)
	if err != nil {
		return 0
	}
	var result int

	for len(pinatas) != 0 {
		mostLeft := 1
		mostRight := 1

		if iMax-1 >= 0 {
			if iMax-2 >= 0 {
				mostLeft = pinatas[iMax-2]
			}
			result += mostLeft * pinatas[iMax-1] * pinatas[iMax]
			pinatas = append(pinatas[:iMax-1], pinatas[iMax:]...)
			iMax--
		} else if iMax+1 < len(pinatas) {
			if iMax+2 < len(pinatas) {
				mostRight = pinatas[iMax+2]
			}
			result += pinatas[iMax] * pinatas[iMax+1] * mostRight
			pinatas = append(pinatas[:iMax+1], pinatas[iMax+2:]...)
		} else {
			return result + pinatas[iMax]
		}
	}

	return result
}

func main() {
	pinatas := []int{1, 1, 4, 5, 1}
	fmt.Println(getMaxCandies(pinatas))
}
