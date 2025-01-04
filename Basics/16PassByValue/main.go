/*
1. Passing by Value (Alice’s Book Summary)
In this scenario, Alice writes her summary in her own notebook and then gives Bob a copy of her summary.

Alice is the original author of the summary.
Bob gets a copy of the summary, so if Bob changes anything in the summary, it doesn't affect Alice’s original notebook.
This is like passing by value — Bob gets a copy of Alice's summary, but any changes Bob makes won't change Alice’s original summary.
*/

package main
import "fmt"

func main() {
	// Alice's original book summary
	bookSummary := "The book is about a detective solving a mystery."
	fmt.Println("Before passing the summary:", bookSummary)
	
	// Pass by value (copy of the book summary)
	modifySummary(bookSummary)

	// Alice's summary remains unchanged
	fmt.Println("After Bob tries to modify the summary:", bookSummary)
}

func modifySummary(summary string) {
	// Bob tries to modify the copy of the summary (pass by value)
	summary = "The detective solves a big mystery and catches the criminal."
	fmt.Println("Inside the function (modified copy):", summary)
}