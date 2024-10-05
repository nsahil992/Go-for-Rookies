package main

import (
	"fmt"
)

func main() {
	languages := map[string]string{"en" : "English", "hi" : "Hindi", "san" : "Sanskrit", "mar" : "Marathi"}
	fmt.Println(languages)  // map[en:English hi:Hindi mar:Marathi san:Sanskrit]

	// Remember (en, hi, san, mar) are the keys and (English, Hindi, Sanskrit, Marathi) are the values

	fmt.Println(len(languages))

	// To print values of the keys
	fmt.Println(languages["mar"]) // Marathi
	fmt.Println(languages["en"]) // English

	// Getting a key
	value, found := languages["en"]
	fmt.Println(value, found)  	// It will print English and true as English value exists in the map

	value, found = languages["hh"]  //  It will print false as hh doesn't exist in the map
	fmt.Println(value, found)

	// Adding key value pair
	languages["jap"] = "Japanese"
	fmt.Println(languages)

	languages["hi"] = "Hindi language"
	fmt.Println(languages)

	// Deleting key value pair
	delete(languages, "jap")
	fmt.Println(languages)

	// Iterating over a map
	for key, value := range languages {
		fmt.Println(key, "=>", value)
	}

	studentGrades := map[string]string{
        "Alice": "A",
        "Bob":   "B",
        "Eve":   "A",
    }

    fmt.Println("Alice's grade:", studentGrades["Alice"])

	for student, grade := range studentGrades {
        fmt.Printf("%s has grade %s\n", student, grade)
    }
}