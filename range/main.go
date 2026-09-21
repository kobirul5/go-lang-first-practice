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
}