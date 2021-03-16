//assess1.go

package main

import "fmt"

type User struct {
	name   string
	mobile int
	place  string
}

type Order struct {
	name             string
	quantity         int
	username         string
	place            string
	out_for_delivery bool
}

type Rep struct {
	mobile   int
	place    string
	assigned bool
}

var products [6]string = [6]string{"tv", "mobile", "watch", "fridge", "cooler", "fan"}

var slc []User
var orders []Order
var reps []Rep

var currentUser string
var admin string = "Rahul"

func preAdd() {
	o1 := Order{name: "mobile", username: "Abhi", quantity: 2, place: "Vizag"}
	o2 := Order{name: "watches", username: "Abhi", quantity: 5, place: "Vizag"}
	o3 := Order{name: "fan", username: "Mukesh", quantity: 3, place: "Kadapa"}
	o4 := Order{name: "cooler", username: "Mukesh", quantity: 2, place: "Kadapa"}
	o5 := Order{name: "tv", username: "Silva", quantity: 10, place: "Hyderabad"}
	o6 := Order{name: "fridge", username: "Silva", quantity: 1, place: "Hyderabad"}
	o7 := Order{name: "watches", username: "Silva", quantity: 30, place: "Hyderabad"}
	orders = append(orders, o1, o3, o5, o2, o6, o4, o7)

	u1 := User{name: "Abhi", place: "Vizag", mobile: 9874561231}
	u2 := User{name: "Silva", place: "Hyderabad", mobile: 9874561232}
	u3 := User{name: "Mukesh", place: "Kadapa", mobile: 9874561233}
	slc = append(slc, u1, u2, u3)

	r1 := Rep{place: "Vizag", mobile: 9874561234}
	r2 := Rep{place: "Hyderabad", mobile: 9874561235}
	r3 := Rep{place: "Kadapa", mobile: 9874561236}
	r4 := Rep{place: "Hyderabad", mobile: 9874561237}
	reps = append(reps, r1, r2, r3, r4)

}

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
		fmt.Println("Invalid Number")
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

func add_rep_function() {
	for !add_rep() {

	}
}

func add_rep() bool {
	fmt.Println("Please enter delivery representative details")
	var r Rep
	fmt.Println("Enter Mobile Number : ")
	fmt.Scanln(&r.mobile)

	fmt.Println("Enter Place : ")
	fmt.Scanln(&r.place)

	return AddReps(r)
}

func AddReps(rep Rep) bool {
	if rep.mobile < 6999999999 || rep.mobile > 9999999999 {
		fmt.Println("Invalid Number")
		return false
	}
	for _, v := range reps {
		if rep.mobile == v.mobile {
			fmt.Println("Mobile number already exists")
			return false
		}
	}

	reps = append(reps, rep)
	return true
}

func AssignRep() {
	var place string
	var name string
	fmt.Println("Enter the name of the user to be assigned:")
	fmt.Scanln(&name)
	fmt.Println("Enter the location:")
	fmt.Scanln(&place)

	var flag bool = true
	var order_notfound bool = true

	var flag2 bool = true
	var rep_notfound bool = true

	for k, v := range orders {
		if v.username == name {
			order_notfound = false
			// fmt.Println(admin)
			// fmt.Println(currentUser)
			if currentUser == admin {
				flag = false
				for a, b := range reps {
					if b.place == place {
						rep_notfound = false
						if !b.assigned || (flag == false) {
							if currentUser == admin {
								flag2 = false
								(&orders[k]).out_for_delivery = true
								(&reps[a]).assigned = true
								break
							}
						}

					}
				}

			}
		}
	}
	if flag == true {
		fmt.Println("You don't have permission to do that")
	}
	if order_notfound == true {
		fmt.Println("Order not found")
	}

	if flag2 == true {
		fmt.Println("You don't have permission to do that 2")
	}
	if rep_notfound == true {
		fmt.Println("Rep not found")
	}

}

func GetThem() {
	fmt.Println(slc[:])
	fmt.Println(reps[:])
}
func main() {
	preAdd()
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
		fmt.Println("Enter 8 to Add Reps")
		fmt.Println("Enter 9 to Assign Reps to Orders")
		fmt.Println("Enter 10 to Print both Users and Reps")
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
		} else if op == 8 {
			add_rep_function()
		} else if op == 9 {
			AssignRep()
		} else if op == 10 {
			GetThem()
		} else {
			break
		}
	}
}


// func AssignRep() {
// 	var place string
// 	var name string
// 	fmt.Println("Enter the name of the user to be assigned:")
// 	fmt.Scanln(&name)
// 	fmt.Println("Enter the location:")
// 	fmt.Scanln(&place)

// 	var flag bool = true
// 	var order_notfound bool = true

// 	for k, v := range orders {
// 		if v.username == name {
// 			order_notfound = false
// 			// fmt.Println(admin)
// 			// fmt.Println(currentUser)
// 			if currentUser == admin {
// 				flag = false

// 				(&orders[k]).out_for_delivery = true
// 			}
// 		}
// 	}
// 	if flag == true {
// 		fmt.Println("You don't have permission to do that")
// 	}
// 	if order_notfound {
// 		fmt.Println("Order not found")
// 	}

// 	var flag2 bool = true
// 	var rep_notfound bool = true

// 	for a, b := range reps {
// 		if b.place == place {
// 			rep_notfound = false
// 			if !b.assigned {
// 				if currentUser == admin {
// 					flag2 = false
// 					(&reps[a]).assigned = true
// 					//break
// 				}
// 			}

// 		}
// 	}
// 	if flag2 == true {
// 		fmt.Println("You don't have permission to do that 2")
// 	}
// 	if rep_notfound {
// 		fmt.Println("Rep not found")
// 	}

// 	// for i := 0; i < len(reps); i++ {
// 	// 	if location == reps[i].place && reps[i].free == false {
// 	// 		fmt.Println("Delivery representative assigned succesfully!")
// 	// 		reps[i].free = true
// 	// 		users[match].out_for_delivery = true
// 	// 		break
// 	// 	} else {
// 	// 		fmt.Println("Delivery Rep Not Found!")
// 	// 	}
// 	// }

// }

