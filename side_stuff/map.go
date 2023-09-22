package main

import "fmt"

func main() {
	var m map[string]string
	fmt.Println(m == nil)
	m = map[string]string{}
	fmt.Println(m == nil)
	fmt.Println(m)
	fmt.Println(len(m))
	m = make(map[string]string, 5)
	m = map[string]string{"Sam": "Programmer to be"}
	fmt.Println(m)
	m["Johnny"] = "Programmer"
	fmt.Println(m)

	delete(m, "Johnny")
	fmt.Println(m)

	for name, title := range m {
		fmt.Println(name, title)
	}

	title, ok := m["Johnny"]
	if ok {
		fmt.Println(title)
	} else {
		fmt.Println("Doesn't exist")
	}
}
