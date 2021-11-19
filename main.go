package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Country string `json:"country"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func main() {
	http.HandleFunc("/api/v1/get-data", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)

			return
		}
		w.Header().Set("Content-Type", "application/json")		

		name := GenerateName()
		age := GenerateAge()

		json.NewEncoder(w).Encode(Person{
			Name: name,
			Age: age,
			Country: GenerateCountry(),
			Email: GenerateEmail(name, age),
			Password: GeneratePassword(),
		})
	})

	fmt.Println("The server is running on https://127.0.0.1:8080")
	fmt.Println("The api is running on https://127.0.0.1:8080/api/v1/get-data")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func GenerateAge() int {
	age := rand.Intn(100)

	return age
}

func GenerateEmail(name string, age int) string {
	var email string 
	emailExtensions := []string{"@hotmail.com", "@gmail.com", "@outlook.com"}
	age2 := strconv.Itoa(age)
	index := rand.Intn(len(emailExtensions))
	email += strings.Replace(name, " ", ".", -1) + "." + age2 + emailExtensions[index]

	return email
}

func GenerateName() string {
	names := []string{"John", "John Smith", "Steve", "Abbey", "Ellie", "Robert", "Mia"}

	return names[rand.Intn(len(names))]
}

func GenerateCountry() string {
	countries := []string{"Russia", "USA", "Canada", "China", "Brazil", "India", "Portugal", "Spain", "Germany", "France", "Denmark", "Mexico"}

	return countries[rand.Intn(len(countries))]
}

func GeneratePassword() string {
	chars := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "1", "2", "3", "4", "5", "6", "7", "8", "9", "!", "@", "#", "$", "%", "¨", "&", "*", "(", ")", "[", "]", "{", "}", ";", ":", "/", "|", ",", ".", "<", ">", "¹", "²", "³", "£", "¢", "¬", "§", "´", "`", "~", "^", "ª", "º", "°", "?", "'", "-", "+", "=", "_"}
	password := ""

	for i := 0; i < 8; i++ {
		password += chars[rand.Intn(len(chars))]
	}

	return password
}
