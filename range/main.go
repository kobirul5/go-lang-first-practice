package main

func main() {
	myMap := map[string]string{
		"Alice": "Engineer",
		"Bob":   "Designer",
		"success": "Developer",
	}

	for key, value := range myMap {
		println(key, ":", value)
	}

	myArr := []int{1, 2, 3, 4, 5}
	for index, value := range myArr {
		println("Index:", index, "Value:", value)
	}

	// Using range with a slice
	for i, v := range myArr {
		println("Index:", i, "Value:", v)
	}

}