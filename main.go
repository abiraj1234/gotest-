package main

import (
	"Cafeteria/users"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var useroption int
	var exit string
	userlist := users.Users_Details{}
	historyList := users.HistoryList{}
	fmt.Println("Welcome fam")

	for ok := true; ok; ok = exit == "y" {
		fmt.Println("1.Login 2.Signup 3.Exit")
		fmt.Scanln(&useroption)
		switch useroption {
		case 1:
			login(&userlist, &historyList)

		case 2:
			sign_up(&userlist)

		case 3:
			fmt.Println("Thank you..come again!")
		default:
			fmt.Println("You Entered wrong option")

		}
		fmt.Println("Do You Want to Contiune(y/n) ")
		fmt.Scanln(&exit)
	}

}
func sign_up(userlist *users.Users_Details) {
	var username, useremail, userpassword string
	var userage int
	var userphone int64

	fmt.Println("Enter Your Name : ")
	fmt.Scanln(&username)

	fmt.Println("Enter Your Password : ")
	fmt.Scanln(&userpassword)

	fmt.Println("Enter Your age : ")
	fmt.Scanln(&userage)

	fmt.Println("Enter Your Email : ")
	fmt.Scanln(&useremail)

	fmt.Println("Enter Your Phone : ")
	fmt.Scanln(&userphone)

	userlist.Add_user(username, userpassword, userphone, userage, useremail)
	userlist.Show_users()

}
func login(userlist *users.Users_Details, historylist *users.HistoryList) {
	var namecheck, passwordcheck string
	fmt.Println("Enter Your Name : ")
	fmt.Scanln(&namecheck)
	fmt.Println("Enter Your Password : ")
	fmt.Scanln(&passwordcheck)
	successMsg := strings.Split(userlist.Check_User(namecheck, passwordcheck), " ")
	ids, _ := strconv.Atoi(successMsg[0])
	if successMsg[1] == "true" {
		name := successMsg[2]
		Loginsuccess(ids, historylist, name, userlist)
	} else {
		fmt.Println("wrong input")
	}

}

func Loginsuccess(id int, historylist *users.HistoryList, name string, userlist *users.Users_Details) {
	var login_option int
	var amount int
	var Back string
	for ok := true; ok; ok = Back == "y" {
		fmt.Println("1.Menu Card 2.Balance 3.Recharge 4.Your Order 5.History ")
		fmt.Scanln(&login_option)
		switch login_option {
		case 1:
			menulist(id, historylist, name)
		case 2:
			balance := userlist.Balance(id)
			fmt.Println("Your Balance Is : ", balance)

		case 3:
			fmt.Println("How Much Amount You Want Recharge : ")
			fmt.Scanln(&amount)

			userlist.Recharge(id, amount)
			fmt.Println("Recharged successfull")

		case 4:
			fmt.Println("YOUR ORDER : ")
			historylist.Current_History(id, userlist)

		case 5:
			historylist.Show_History(id)

		default:
			fmt.Println("YOU ENTER THE WRONG OPTION")
		}
		fmt.Println("Do You Want go back (y/n) ")
		fmt.Scanln(&Back)

	}

}

func menulist(id int, history *users.HistoryList, name string) {
	var menuoption int
	var quantity int
	fmt.Printf("1.TEA ---> 20Rs\n2.COFFEE ---> 25Rs\n3.MILK ---> 15Rs\n4.BADAMMILK --->35Rs\n5.LEMONTEA ---->Rs20\n")
	fmt.Print("CHOOSE THE OPTION : ")
	fmt.Scanln(&menuoption)
	fmt.Println("NO OF QUANTITY : ")
	fmt.Scanln(&quantity)
	switch menuoption {
	case 1:
		bill := Billcal(20, quantity)
		history.Add_History(id, name, "TEA", bill)
	case 2:
		bill := Billcal(25, quantity)
		history.Add_History(id, name, "COFFEE", bill)

	case 3:
		bill := Billcal(15, quantity)
		history.Add_History(id, name, "MILK", bill)

	case 4:
		bill := Billcal(35, quantity)
		history.Add_History(id, name, "BADAM MILK", bill)

	case 5:
		bill := Billcal(20, quantity)
		history.Add_History(id, name, "LEMONTEA", bill)
	default:
		fmt.Println("worng option")

	}
	fmt.Println("Your order is booked ..goto yourOrder section to check your order..")

}
func Billcal(amount int, quantity int) int {
	return amount * quantity
}
