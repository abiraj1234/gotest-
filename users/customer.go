package users

import (
	"fmt"
	"strconv"
)

type User struct {
	id       int
	name     string
	password string
	phone    int64
	age      int
	email    string
	balance  int
	next     *User
}

type Users_Details struct {
	head *User
	len  int
}

func (u *Users_Details) Add_user(Name string, Password string, Phone int64, Age int, Email string) {

	n := User{
		name: Name, password: Password, phone: Phone, age: Age, email: Email}
	if u.len == 0 {
		u.head = &n
		u.len++
		n.id = u.len
		return
	}
	temp := u.head
	for i := 0; i < u.len; i++ {
		if temp.next == nil {
			temp.next = &n
			u.len++
			n.id = u.len
			return
		}
		temp = temp.next
	}
}

func (u *Users_Details) Show_users() {
	if u.len == 0 {
		fmt.Println("No users found")
		return
	}
	temp := u.head
	for i := 0; i < u.len; i++ {
		fmt.Println("-------------------------------")
		fmt.Printf("NAME : %s,\nID : %d,\nPASSWORD : *************,\nPHONE : %d,\nAGE : %d,\nEMAIL : %s\n", temp.name, temp.id, temp.phone, temp.age, temp.email)
		temp = temp.next
	}
	fmt.Println()
}

func (u *Users_Details) Check_User(name string, password string) string {
	if u.len == 0 {
		fmt.Println("no user found")
		return "0 false "
	}
	temp := u.head
	for i := 0; i < u.len; i++ {
		if temp.name == name && temp.password == password {

			return strconv.Itoa(temp.id) + " true " + temp.name

		}
		temp = temp.next
	}
	return "0 false "
}

func (u *Users_Details) Balance(id int) int {
	temp := u.head
	for i := 0; i < u.len; i++ {
		if id == temp.id {
			return temp.balance
		}
		temp = temp.next
	}
	return 0
}

func (u *Users_Details) Recharge(id int, amount int) {
	if u.len == 0 {
		fmt.Println("NO USERS FOUND")
		return
	}
	temp := u.head
	for i := 0; i < u.len; i++ {
		if id == temp.id {
			temp.balance = temp.balance + amount

		}
		temp = temp.next
	}

}
