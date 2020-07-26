package main

import "fmt"

// user defines a user in the program.
type user struct {
	name    string
	surname string
}

func main() {
	// ----------------------
	// Declare and initialize
	// ----------------------

	// Declare and make a map that stores values of type user with a key of type string.
	users1 := make(map[string]user)

	// Add key/value pairs to the map.
	users1["Roy"] = user{"Rob", "Roy"}
	users1["Ford"] = user{"Henry", "Ford"}
	users1["Mouse"] = user{"Mickey", "Mouse"}
	users1["Jackson"] = user{"Michael", "Jackson"}

	// ----------------
	// Iterate over map with range
	// ----------------

	fmt.Printf("\n=> Iterate over map with range\n")
	for key, value := range users1 {
		fmt.Println(key, value)
	}

	// ------------
	// Map literals
	// ------------

	// Declare and initialize the map with values.
	users2 := map[string]user{
		"Roy":     {"Rob", "Roy"},
		"Ford":    {"Henry", "Ford"},
		"Mouse":   {"Mickey", "Mouse"},
		"Jackson": {"Michael", "Jackson"},
	}

	// Iterate over the map.
	fmt.Printf("\n=> Map literals\n")
	for key, value := range users2 {
		fmt.Println(key, value)
	}

	// ----------
	// Delete key
	// ----------

	delete(users2, "Roy")

	// --------
	// Find key
	// --------

	// Find the Roy key.
	// If found is True, we will get a copy value of that type.
	// if found is False, u is still a value of type user but is set to its zero value.
	u1, found1 := users2["Roy"]
	u2, found2 := users2["Ford"]

	// Display the value and found flag.
	fmt.Printf("\n=> Find key\n")
	fmt.Println("Roy", found1, u1)
	fmt.Println("Ford", found2, u2)
}
