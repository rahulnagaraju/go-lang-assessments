//assess1.go

package main

import "fmt"

type User struct {
	name   string
	mobile int
	place  string
}

type Order struct {
	name     string
	quantity int
	username string
}

var products [6]string = [6]string{"tv", "mobile", "watch", "fridge", "cooler", "fan"}

var slc []User
var orders []Order

var currentUser string

func GetProducts() {
	fmt.Println("These are the products we have,select among them")
	for i := range products {
		fmt.Println(products[i])
	}
}

func PrintOrders() {
	fmt.Println("These are the orders taken:")
	fmt.Println(orders[:])
}

func PrintMyOrders() {
	fmt.Println("These are the orders taken:")
	for k, v := range orders {
		if currentUser == v.username {
			fmt.Printf("Product:%s,Qty:%d,Index:%d\n", v.name, v.quantity, k)

		}

	}
}

func product_inputs() (string, int) {
	fmt.Println("Enter the product name ")
	var temp1 string
	fmt.Scanln(&temp1)
	fmt.Println("Enter the quantity ")
	var temp2 int
	fmt.Scanln(&temp2)
	return temp1, temp2
}

func CheckProducts(productName string) bool {
	var flag bool = true
	for i := 0; (i < 6) && (flag); i++ {
		if productName == products[i] {
			flag = false
		}
	}
	// fmt.Println("I am here")
	//fmt.Println(flag)
	return flag
}

func add_user_function() {
	for !add_user() {

	}
}

func add_user() bool {
	currentUser = ""
	fmt.Println("Please enter your credentials")
	var u User
	fmt.Println("Enter your Name : ")
	fmt.Scanln(&u.name)

	fmt.Println("Enter your Mobile Number : ")
	fmt.Scanln(&u.mobile)

	fmt.Println("Enter your Place : ")
	fmt.Scanln(&u.place)

	return AddUsers(u)
}

func AddUsers(user User) bool {
	if user.mobile < 6999999999 || user.mobile > 9999999999 {
		fmt.Printf("Invalid Number")
		return false
	}
	for _, v := range slc {
		if user.mobile == v.mobile {
			fmt.Println("Mobile number already exists")
			return false
		}
	}

	slc = append(slc, user)
	currentUser = user.name
	return true
}

func place_order() {
	temp1, temp2 := product_inputs()

	var flag bool = true

	for CheckProducts(temp1) {
		// fmt.Println("Before")
		// fmt.Println(temp1)
		// fmt.Println("After")
		fmt.Println("We don't have that product..Please enter again:")
		flag = false
		break

		// place_order()

	}
	if flag {
		PlaceOrder(temp1, temp2)
	}

}

func PlaceOrder(product string, quantity int) {
	var order1 Order
	order1.name = product
	order1.quantity = quantity
	order1.username = currentUser
	orders = append(orders, order1)
}

func change_order() {
	temp1, temp2 := product_inputs()

	var flag bool = true

	for CheckProducts(temp1) {
		fmt.Println("We don't have that product..Please enter again:")
		flag = false
		break
	}
	if flag {
		ChangeOrder(temp1, temp2)
	}
}

func ChangeOrder(productName string, quantity int) {
	var flag bool = true
	for k, v := range orders {
		if v.name == productName {
			if currentUser == v.username {
				flag = false
				(&orders[k]).quantity = quantity
			}
		}
	}
	if flag {
		fmt.Println("You don't have permission to do that")
	}
}

func cancel_order() {
	fmt.Println("Enter the product name ")
	var temp1 string
	fmt.Scanln(&temp1)

	var flag bool = true

	for CheckProducts(temp1) {
		fmt.Println("We don't have that product..Please enter again:")
		flag = false
		break
	}
	if flag {
		CancelOrder(temp1)
	}
}

func CancelOrder(productName string) {
	var flag bool = true
	for k, v := range orders {
		if v.name == productName {
			if currentUser == v.username {
				flag = false
				orders[k] = orders[len(orders)-1] // Copy last element to index k.
				//orders[len(orders)-1] = NULL   // Erase last element Rahul here (write zero value).
				orders = orders[:len(orders)-1] // Truncate slice.
			}

		}
	}
	if flag {
		fmt.Println("You don't have permission to do that")
	}
}

func main() {

	add_user_function()
	var op int

	for true {
		fmt.Printf("\n")
		fmt.Println("Enter 1 to Get Products")
		fmt.Println("Enter 2 to Place Orders")
		fmt.Println("Enter 3 to Modify Products")
		fmt.Println("Enter 4 to Cancel Products")
		fmt.Println("Enter 5 to Print Products")
		fmt.Println("Enter 6 to Print Products ordered from me")
		fmt.Println("Enter 7 to Login to another account")
		fmt.Scanln(&op)
		if op == 1 {
			GetProducts()
		} else if op == 2 {
			place_order()
		} else if op == 3 {
			change_order()
		} else if op == 4 {
			cancel_order()
		} else if op == 5 {
			PrintOrders()
		} else if op == 6 {
			PrintMyOrders()
		} else if op == 7 {
			add_user_function()
		} else {
			break
		}
	}
}
