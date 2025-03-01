package hellope

import "core:fmt"
import "core:time"
import "core:slice"
import "core:os"
import "helpers"
import "models"

main :: proc() {
	fmt.println("Hellope!")
	fmt.printf("Number: %d\n", helpers.incr_number(42))

	u := models.new_user()
	fmt.printf("User: %v\n", u)

	now := time.now()
	fmt.println(now)

	somestr := "Testing for bytes"
	somebyte := transmute([]u8)somestr
	fmt.println(somebyte)

	list1 := [6]int{3,4,2,1,3,3}
	list2 := [6]int{4,3,5,3,9,3}

	slice.sort(list1[:])
	slice.sort(list2[:])

	distance := 0
	for i in 0..<6 {
		distance += abs(list1[i] - list2[i])
	}
	fmt.println("Distance :", distance)
}

