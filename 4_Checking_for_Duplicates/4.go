package main

import "fmt"

type Developer struct {
	Name string
	Age  int
}

func FilterUnique(developers []Developer) []string {
	// TODO Implement
	uniqueMap := make(map[string]bool)
	var uniqueNames []string
	for _, dev := range developers {
		if _, exist := uniqueMap[dev.Name]; !exist {
			uniqueMap[dev.Name] = true
			uniqueNames = append(uniqueNames, dev.Name)
		}
	}
	return uniqueNames

}

func main() {
	fmt.Println("Filter Unique Challenge")
	dev := []Developer{
		{Name: "Elliot"},
		{Name: "Alan"},
		{Name: "Jennifer"},
		{Name: "Graham"},
		{Name: "Paul"},
		{Name: "Alan"},
	}
	FilterUnique(dev)

}
