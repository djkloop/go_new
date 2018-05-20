package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "cc",
		"course":  "golang",
		"site":    "djkloop",
		"quality": "height",
	}

	m2 := make(map[string]int)

	var m3 map[string]int

	fmt.Println(m, m2, m3)

	fmt.Println("map for")

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName, hasValue := m["course"]
	fmt.Println(courseName, hasValue)
	fmt.Println("Map Zero Value")
	causeName, hasValue := m["cause"]
	fmt.Println(causeName, hasValue)
	if causeName, hasValue := m["cause"]; hasValue {
		fmt.Println(causeName)
	} else {
		fmt.Println("key does not exist.")
	}

	fmt.Println("Deleting values")
	name, hasValue := m["name"]
	fmt.Println(name, hasValue)

	delete(m, "name")
	name, hasValue = m["name"]
	fmt.Println(name, hasValue)
}
