package users

import (
	"fmt"
)

type History struct {
	id        int
	name      string
	orderfood string
	bill      int
	next      *History
}

type HistoryList struct {
	head *History
	len  int
}

func (H *HistoryList) Add_History(Id int, Name string, Orderfood string, Bill int) {

	na := History{
		id: Id, name: Name, orderfood: Orderfood, bill: Bill}
	if H.len == 0 {
		H.head = &na
		H.len++
		return
	}
	temp := H.head
	for i := 0; i < H.len; i++ {
		if temp.next == nil {
			temp.next = &na
			H.len++
			return
		}
		temp = temp.next

	}

}
func (H *HistoryList) Show_History(id int) {
	if H.len == 0 {
		fmt.Println("No History found")
		return
	}
	temp := H.head
	for i := 0; i < H.len; i++ {
		if id == temp.id {
			fmt.Printf("Id : %d\nName : %s\nOrder Food : %s\nBill : %d\n", temp.id, temp.name, temp.orderfood, temp.bill)
			fmt.Println()

		}
		temp = temp.next
	}
}
func (H *HistoryList) Current_History(id int, userlist *Users_Details) {
	var confirmation string
	if H.len == 0 {
		fmt.Println("No History found")
		return
	}
	temp := H.head
	current := H.head
	for i := 0; i < H.len; i++ {
		if id == temp.id {
			current = temp

		}
		temp = temp.next
	}
	fmt.Printf("Id : %d\nName : %s\nOrder Food : %s\nBill : %d\n", current.id, current.name, current.orderfood, current.bill)
	fmt.Println()
	fmt.Println("Confirm Your Order (y/n) : ")
	fmt.Scanln(&confirmation)
	if confirmation == "y" {
		balance := userlist.Balance(id)
		if balance > current.bill {
			userlist.Recharge(id, -current.bill)
			fmt.Println("Your Order is confirmed !")
		} else {
			fmt.Println("Please Reacharge your account ...!")
		}
	} else {
		fmt.Println("Your order is cancelled, thank you! Come Agian!")
	}

}
