package main

import "fmt"

// func chagnenum(num *int){
// 	*num = 1
// 	fmt.Println("memory ofnum in changenum", &num)
// }
// func main(){

// 	num := 10
// 	fmt.Println("memory of num in main", &num)
// 	chagnenum(&num)
// 	fmt.Println("num in main", num)
// 	fmt.Println("memory of num in main", &num)
// }

// func main() {
// 	var num1 int = 10
// 	var num2 int = 20
// 	fmt.Println("before swap", num1, num2)
// 	swap(&num1, &num2)
// 	fmt.Println("after swap", num1, num2)
// }
// func swap(num1 *int, num2 *int) { //using pointer
// 	temp := *num1
// 	*num1 = *num2
// 	*num2 = temp
            
// }
// func main(){
// 	num1 := 10
// 	num2 := 20
// 	fmt.Println("before swap", num1, num2)
// 	swap(&num1, &num2)
// 	fmt.Println("after swap", num1, num2)
// 	fmt.Println()
// }
// func swap(num1 *int, num2 *int) {
// 	temp := *num1
// 	*num1 = *num2
// 	*num2 = temp
// }


// updating pattern in the system 
// type User struct{
// 	Name string
// 	Email string
// }

// func updateuser(u *User){
// 	u.Name = "abcd"
// 	u.Email = "updated@example.com"
	
// }
// func main(){
// 	u := User{Name: "Original", Email: "old@example.com"}
// 	updateuser(&u)
// 	fmt.Println(u)
// }

type egg struct {
	customer string;
    amount int;
}
func customerch(bablu *egg){
	bablu.amount = 23;
	bablu.customer = "alex"
	fmt.Println("this is bablu shop",*bablu)
}
func main(){
  santu :=egg{customer : "BELA" , amount: 3}
  customerch(&santu)
  fmt.Println("now the customer are switched",santu)
}