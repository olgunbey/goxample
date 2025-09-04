package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "username123"
	password = "password123"
	dbname   = "testdb"
)

func main() {

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	row, err := db.Query("SELECT * FROM users WHERE ID = $1", 1)
	if err != nil {
		fmt.Println("Değer gelmedi")
	}
	defer row.Close()

	var persons []Person

	for row.Next() {
		var person Person
		if err := row.Scan(&person.Id, &person.Name, &person.Age); err != nil {
			fmt.Println("Scan error:", err)
			continue
		}
		persons = append(persons, person)
	}

	for _, person := range persons {
		fmt.Println(person.Id, person.Name)
	}

	// _, err = db.Exec("CREATE DATABASE " + "testDb")
	// if err != nil {
	// 	fmt.Println("Database may already exist:", err)
	// } else {
	// 	fmt.Println("Database created!")
	// }

	// _, err = db.Exec(`CREATE TABLE users(
	//   Id SERIAL PRIMARY KEY,
	//   Name VARCHAR(15) NOT NULL,
	//   Age INTEGER NOT NULL
	// ) `)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Create Table")
	// }

	//Api'den gelen data bu olsun
	// p := Person{Name: "Olgun", Age: 12, Id: 1}

	// _, err = db.Exec(`INSERT INTO users VALUES
	// ($1,$2,$3)`, p.Id, p.Name, p.Age)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// var a [2]string

	// a[0] = "Olgun"
	// a[1] = "Şahin"

	// for q := 0; q < len(a); q++ {

	// 	fmt.Println(a[q])
	// }

	// var repo = NewInMemoryPersonRepo()

	// fmt.Println(repo)
	// p2 := Person{Name: "Ahmet", Id: 1, Age: 10}

	// p3 := Person{Name: "Mehmet", Id: 2, Age: 20}

	// repo.Save(p2)

	// repo.Save(p3)

	// fmt.Println("Number of persons in repo:", len(repo.persons))

	// person, err := repo.GetById(2)
	// if err != nil {
	// 	fmt.Println("Hata:", err)
	// } else {
	// 	fmt.Println("Kullanıcı:", person.Name)
	// }

	// p := Person{
	// 	Id:   2,
	// 	Name: "Olgun",
	// 	Age:  24,
	// 	Addresses: []Address{
	// 		{Street: "İstiklal Cad", City: "İstanbul"},
	// 		{Street: "Atatürk Cad", City: "Ankara"},
	// 	},
	// }

	// for _, addr := range p.Addresses {
	// 	fmt.Println("Address", addr.City, addr.Street)
	// }

	// for _, addr := range p.Addresses {
	// 	if addr.City == "Ankara" {
	// 		fmt.Println("Adresim ankara")
	// 	} else {
	// 		fmt.Println("Adresim Ankara Değil")
	// 	}
	// }

}

// func (p *Person) IncreateAge() {
// 	p.Age += 1
// }

type Person struct {
	Id   int
	Name string
	Age  int
}

// type InMemoryPersonRepo struct {
// 	persons map[int]Person
// }

// // New-Constuctor
// func NewInMemoryPersonRepo() *InMemoryPersonRepo {
// 	return &InMemoryPersonRepo{
// 		persons: make(map[int]Person),
// 	}
// }

// func (r *InMemoryPersonRepo) GetById(id int) (*Person, error) {
// 	person, ok := r.persons[id]
// 	if !ok {
// 		return nil, fmt.Errorf("User Not Found")
// 	}

// 	return &person, nil
// }

// func (r *InMemoryPersonRepo) Save(p Person) error {
// 	r.persons[p.Id] = p
// 	return nil
// }
