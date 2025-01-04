/*
Passing by Reference (Alice Shares Her Notebook)
Now, in this case, Alice shares her actual notebook with Bob. So when Bob writes in the notebook, Alice will immediately see the changes.

Alice still owns the original notebook, but now she is letting Bob modify it directly. If Bob changes something in the notebook, it affects Alice’s notebook too.
This is like passing by reference — Bob is modifying the original notebook directly, so changes are visible to Alice as well.
*/

package main

import "fmt"

func main() {
	// Alice's original book summary
	bookSummary := "The book is about a detective solving a mystery."
	fmt.Println("Before passing the summary:", bookSummary)

	// Pass by reference (Alice shares the actual notebook with Bob)
	modifySummaryByReference(&bookSummary)

	// Alice's summary is changed after Bob modifies it
	fmt.Println("After Bob modifies the summary:", bookSummary)
}

func modifySummaryByReference(summary *string) {
	// Bob modifies the original summary directly by dereferencing the pointer
	*summary = "The detective solves a big mystery and catches the criminal."
	fmt.Println("Inside the function (modified original):", *summary)
}
