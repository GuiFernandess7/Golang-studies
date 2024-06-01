package main

import "fmt"

type Person struct {
	name string
	age int
	email string
	address Address
}

type Address struct {
	street string
	city string
	postal_code string
}

type User struct {
	Person
	likes int
}

func compare_users(user1 Person, user2 Person) bool {
	return user1 == user2
}

func getInceptionSum(n int) func() int {
	fn := 1
	return func() int {
		return fn + n
	}
}

func main(){
	n_result := getInceptionSum(5)
	fmt.Println(n_result())

	p1 := Person{
		name: "Guilherme",
		age: 22,
		email: "guilherme@email.com",
		address: Address{
			street: "Rua John Doe",
			city: "São Paulo",
			postal_code: "0321930",
		},
	}

	p2 := Person{
		name: "Guilherme",
		age: 22,
		email: "guilherme@email.com",
	}

	main_user := User{
        Person: Person{
            name:  "Gabriel",
            age:   20,
            email: "gabriel_2@gmail.com",
            address: Address{
                street:     "Rua 2",
                city:       "Cidade Exemplo",
                postal_code: "12345-678",
            },
        },
        likes: 20,
    }

	println(compare_users(p1, p2))
	fmt.Println(p1.address.city)
	fmt.Print(main_user)
}