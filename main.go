package main

import "fmt"

func main() {
	order := map[string][]string{
		"batch1": {"brownies", "nastar", "muffin", "muffin", "brownies", "pie", "nastar", "muffin"},
		"batch2": {"brownies", "pie", "pie", "nastar"},
	}

	counts := make(map[string]int)

	for batch, cakes := range order {
		fmt.Println("Processing", batch)
		printCakes(cakes, counts)
	}
}

func printCakes(cakes []string, counts map[string]int) {
	for _, cake := range cakes {
		counts[cake]++
		fmt.Printf("%d %s siap!\n", counts[cake], cake)
	}
}
