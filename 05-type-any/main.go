package main

func main() {

	var myList = []any{"sujan", 3, 6, true, 4.5, 5.6, "hello", 7, 8, 9, 10}

	processList(myList)

}

// processList iterates over a slice of values of any type and processes each element based on its type.
// It counts the number of string and boolean elements, sums all integer values, and sums all float32 and float64 values (converted to float64).
// The function then prints the counts and sums for each type.
// This function is useful for analyzing heterogeneous lists where elements may be of different types.
func processList(list []any) {

	var strCount, boolCount, intSum int
	var floatSum float64 = 0.0

	for _, item := range list {
		switch v := item.(type) {
		case string:
			strCount++
		case int:
			intSum += v
		case bool:
			boolCount++
		case float32:
			floatSum += float64(v)
		case float64:
			floatSum += v
		}
	}

	println("String count:", strCount)
	println("Boolean count:", boolCount)
	println("Sum of ints:", intSum)
	println("Sum of float64s:", floatSum)

}
