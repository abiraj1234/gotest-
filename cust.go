package customers

import (
	"fmt"
	"strconv"
)

type User struct {
	id       int
	name     string
	age      int
	address  string
	phone    int64
	email    string
	password string
	balance  int
	next     *User
}

type Users_Details struct {
	head *User
	len  int
}

func (u *Users_Details) Add_user(Name string, Age int, Address string, Phone int64, Email string, Password string) {
	x := User{
		name: Name, age: Age, address: Address, phone: Phone, email: Email, password: Password}

	if u.len == 0 {
		u.head = &x
		u.len++
		x.id = u.len
		return
	}
	temp := u.head
	for i := 0; i < u.len; i++ {
		if temp.next == nil {
			temp.next = &x
			u.len++
			x.id = u.len
			return
		}
		temp = temp.next
	}
}

func (u *Users_Details) Show_users() {
	if u.len == 0 {
		fmt.Println("No User found")
		return
	}
	temp := u.head
	for i := 0; i < u.len; i++ {
		fmt.Println("----------------------------")
		fmt.Printf("Username : %s\nUser_id : %d\nUserage : %d\nUseremail :%s\nUserpassword:*********\nUseraddress : %s\nUserphone :%d\n", temp.name, temp.id, temp.age, temp.email, temp.address, temp.phone)
		temp = temp.next
	}
	fmt.Println()

}

func (u *Users_Details) Check_User(name string, password string) string {
	if u.len == 0 {
		fmt.Println("NO User Found")
		return "0 false none"
	}
	temp := u.head
	for i := 0; i < u.len; i++ {
		if temp.name == name && temp.password == password {
			return strconv.Itoa(temp.id) + " true " + temp.name
		}
		temp = temp.next
	}
	return "0 false none"
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
		fmt.Println("No User Found ")
	}
	temp := u.head
	for i := 0; i < u.len; i++ {
		if id == temp.id {
			temp.balance = temp.balance + amount
		}
		temp = temp.next
	}
}
