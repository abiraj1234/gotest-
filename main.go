package main

import (
	"bufio"
	"fmt"
	"os"
	"sam_elite/customers"
	"strconv"
	"strings"
)

func main() {
	var userinput int
	var exit string
	userlist := customers.Users_Details{}
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("*****>WELCOME TO SAM_ELITE<****")

	for ok := true; ok; ok = exit == "y" {
		fmt.Println("Choose the option")
		fmt.Println("1.LOGIN 2.REGISTER 3.EXIT")
		fmt.Scanln(&userinput)
		switch userinput {
		case 1:

			login(&userlist, reader)
			fmt.Println("Do You Want to Contiune(y/n) ")
			fmt.Scanln(&exit)
		case 2:

			Signup(&userlist)
			fmt.Println("Do You Want to Contiune(y/n) ")
			fmt.Scanln(&exit)
		case 3:
			fmt.Println("Thank you visit again ")
			exit = "n"
		default:
			fmt.Println("Choose the vaild option ")
			fmt.Println("Do You Want to Contiune(y/n) ")
			fmt.Scanln(&exit)
		}

	}

}

func Signup(userlist *customers.Users_Details) {
	//var username, useremail, userpassword, userconfpassword, useraddress string
	var userage int
	var userphone int64
	var Confirmpassword bool

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Your Name : ")
	username, _ := reader.ReadString('\n')
	//fmt.Scanln(&username)

	fmt.Print("Enter Your Age (Age should be Above 18) : ")
	fmt.Scanln(&userage)
	if userage < 18 {
		fmt.Println("You are not eligible to use this app BYE! ...")
		return
	}

	fmt.Print("Enter your address : ")
	useraddress, _ := reader.ReadString('\n')

	fmt.Print("Enter Your Phone number : ")
	fmt.Scanln(&userphone)

	fmt.Print("Enter Your Email : ")
	useremail, _ := reader.ReadString('\n')

	for ok := true; ok; ok = Confirmpassword {
		fmt.Print("Enter Your Password : ")
		userpassword, _ := reader.ReadString('\n')

		fmt.Print("Confirm your password : ")
		userconfpassword, _ := reader.ReadString('\n')

		if userpassword == userconfpassword {
			fmt.Println("Sucessfully registered ")
			Confirmpassword = false
			userlist.Add_user(strings.TrimSpace(username), userage, strings.TrimSpace(useraddress), userphone, strings.TrimSpace(useremail), strings.TrimSpace(userpassword))
			userlist.Show_users()
		} else {
			Confirmpassword = true
			fmt.Println("Your Confirm Password Does Not  Matching ...!")
		}

	}

}

func login(userlist *customers.Users_Details, reader *bufio.Reader) {

	fmt.Print("Enter the Name : ")
	u_name, _ := reader.ReadString('\n')

	fmt.Print("Enter Your Password: ")
	u_password, _ := reader.ReadString('\n')

	fmt.Println()

	successMsg := strings.Split(userlist.Check_User(strings.TrimSpace(u_name), strings.TrimSpace(u_password)), " ")
	id, _ := strconv.Atoi(successMsg[0])
	//fmt.Println(successMsg, id)
	if successMsg[1] == "true" {
		fmt.Println("successfully Logined ")
		Loginsuccess(userlist, id)
	} else {
		fmt.Println("Entered wrong credentials !Please login again ")

	}

}

func Loginsuccess(userlist *customers.Users_Details, id int) {

	fmt.Println("please pick any option")
	var after_login int
	var amount int
	var Back string

	for ok := true; ok; ok = Back == "y" {

		fmt.Println("1.MENU CARD 2.BALANCE 3.RECHARGE 4.ORDERS 5.HISTORY 6.EDIT PROFILE 7.LOGOUT")
		fmt.Scanln(&after_login)
		switch after_login {
		case 1:
			fmt.Println("YOUR MENU ")
			menu_card()
		case 2:
			balance := userlist.Balance(id)
			fmt.Println("Your Balance Is : ", balance)
		case 3:
			fmt.Println("How Much Amount You Want Recharge : ")
			fmt.Scanln(&amount)

			userlist.Recharge(id, amount)
			fmt.Println("Recharged successfull")
		case 4:
			fmt.Println("Your order : ")
		case 5:
			fmt.Println("its Your history : ")
		case 6:
			fmt.Println("edit your profile :  ")
		case 7:
			fmt.Println("logout")

		default:
			fmt.Println("please enter the correct option ")
		}
		fmt.Println("do u want to  go back (y/n) ")
		fmt.Scanln(&Back)
	}
}

func menu_card() {
	var menudraft int
	var b string

	for ok := true; ok; ok = b == "y" {

		fmt.Printf("1.BEERS \n2.WHISHKEY BOURBON")
		fmt.Print("CHOOSE THE OPTION ")
		fmt.Scanln(&menudraft)
		switch menudraft {
		case 1:
			BEERS()
		case 2:
			WHISHKEY_BOURBON()
		default:
			fmt.Println("WRONG OPTION ")
		}
		fmt.Println("go back to menu (y/n) ")
		fmt.Scanln(&b)
	}
}

func Billcal(amount int, quantity int) int {
	return amount * quantity

}
func BEERS() {

	var menudraft int
	var quantity int

	fmt.Println("1.CoranaBeer ------ 400Rs\n2.BUDWEISER------- 350Rs\n3.HENINEKEN--------450Rs")
	fmt.Println("CHOOSE THE OPTION : ")
	fmt.Scanln(&menudraft)
	fmt.Println("NO OF QUANTITY : ")
	fmt.Scanln(&quantity)
	switch menudraft {
	case 1:
		bill := Billcal(400, quantity)
		fmt.Println(bill)
	case 2:
		bill := Billcal(350, quantity)
		fmt.Println(bill)
	case 3:
		bill := Billcal(450, quantity)
		fmt.Println(bill)
	default:
		fmt.Println("WRONG OPTION ")

	}

}

func WHISHKEY_BOURBON() {
	var menudraft int

	var quantity int

	fmt.Println("1.JACK DANILES ------5000\n2.HENNESSY --------5500\n3.JIMBEAM----------4000")
	fmt.Println("CHOOSE THE BEST OPTION: ")
	fmt.Scanln(&menudraft)
	fmt.Println("NO OF QUANTITY: ")
	fmt.Scanln(&quantity)
	switch menudraft {
	case 1:
		bill := Billcal(5000, quantity)
		fmt.Println(bill)
	case 2:
		bill := Billcal(5500, quantity)
		fmt.Println(bill)

	case 3:
		bill := Billcal(4500, quantity)
		fmt.Println(bill)
	default:
		fmt.Println("WRONG OPTION")

	}

}
