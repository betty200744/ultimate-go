package main

// user represents an user in the system.
type user struct {
	name  string
	email string
}

func main() {
	// -------------
	// Pass by value
	// -------------

	// Declare variable of type int with a value of 10.
	// This value is put on a stack with a value of 10.
	count := 10

	// To get the address of a value, we use &.
	println("count:\tValue Of[", count, "], \tAddr Of[", &count, "]")

	// Pass the "value of" count.
	increment1(count)

	// Printing out the result of count. Nothing has changed.
	println("count:\tValue Of[", count, "], \tAddr Of[", &count, "]")

	// Pass the "address of" count.
	// This is still considered pass by value, not by reference because the address itself is a value.
	increment2(&count)

	// Printing out the result of count. count is updated.
	println("count:\tValue Of[", count, "], \tAddr Of[", &count, "]")

	// ---------------
	// Escape analysis
	// ---------------

	stayOnStack()
	escapeToHeap()
}

func increment1(inc int) {
	// Increment the "value of" inc.
	inc++
	println("inc1:\tValue Of[", inc, "], \tAddr Of[", &inc, "]")
}

// increment2 declares count as a pointer variable whose value is always an address and points to
// values of type int.
// The * here is not an operator. It is part of the type name.
// Every type that is declared, whether you declare or it is predeclared, you get for free a pointer.
func increment2(inc *int) {
	// Increment the "value of" count that the "pointer points to".
	// The * is an operator. It tells us the value of the pointer points to.
	*inc++
	println("inc2:\tValue Of[", inc, "], \t Value Points To[", *inc, "], \t  Addr Of [", &inc, "]")
}

// stayOnStack shows how the variable does not escape.
func stayOnStack() user {
	// In the stayOnStack stack frame, create a value and initialize it.
	u := user{
		name:  "Hoanh An",
		email: "hoanhan@bennington.edu",
	}

	// Take the value and return it, pass back up to main stack frame.
	return u
}

// escapeToHeap shows how the variable escape.
func escapeToHeap() *user {
	u := user{
		name:  "Hoanh An",
		email: "hoanhan@bennington.edu",
	}

	return &u
}
