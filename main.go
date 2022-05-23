package main

import (
	"bitbucket.org/service-ekspedisi/application"
	"fmt"
	"sort"
)

func main() {
	application.StartApp()
	//fmt.Println(BinaryRepresentation(3, 7))
	//fmt.Println(Number1(5))
}

func CheckMissNumber(a []int) int {
	sort.Ints(a)

	var result int
	missing := 1

	for _, i := range a {
		if i == missing {
			missing += 1
		}
		if i > missing {
			result = missing
			break
		}

		result = missing

	}

	return result
}

func BinaryRepresentation(a, b int) int {
	var n = a * b
	var count = 0

	for i := 0; i < n; i++ {
		n = n & (n - 1)
		count++
	}

	return count
}

/*

SELECT DISTINCT web_page_ipv4, COUNT(distinct(user_ipv4)) AS users_cnt from visits GROUP BY web_page_ipv4;
*/

var Colors = []string{"R", "G", "B"}

func Number1(n int) int {
	var result int
	var arrTemp []int
	arrTemp = []int{1, 2, 3, 4}
	for i := 0; i < n; i++ {
		if i < len(arrTemp) {
			result += arrTemp[i]
		}
	}

	arrTemp = append(arrTemp, result)

	fmt.Println(arrTemp)
	fmt.Println(len(arrTemp))

	return result
}

func TotalInconvenience() {
}
