package main
import "fmt"

func main() {
	sahil := User{"Sahil", "shl@mail.com", true, 21}
	fmt.Println(sahil)
	fmt.Println(sahil.Email)
	sahil.Score()
    sahil.Verified()
}

type User struct {
	Name 	string
	Email 	string
	Status 	bool
	Age 	int
}

func (u User) Score() {
	fmt.Println("Sahil has scored 100  marks")
}

func (u User) Verified() {
    fmt.Println("Sahil has successfully verfied")
}