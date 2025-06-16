package main

func main() {

	mapList := make(map[string](func(int, int) int))

	mapList["add"] = func(a, b int) int {
		return a + b
	}

	var sub func(int, int) int = func(a, b int) int {
		return a - b
	}

	mapList["sub"] = sub

	mapList["mul"] = func(a, b int) int {
		return a * b
	}

	mapList["div"] = func(a, b int) int {
		return a / b
	}

	for key, value := range mapList {
		result := value(10, 20)
		println(key, ":", result)
	}

}
