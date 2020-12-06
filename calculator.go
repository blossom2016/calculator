package main
import(
    "fmt"    
    )
    
func main(){
  var operator string
  var num1,num2 int 
fmt.Println("For simple calculation")
fmt.Println("Enter Num1 ")
fmt.Scanln(&num1)
fmt.Println("Enter Num2 ")
fmt.Scanln(&num2)
fmt.Println("Pick an Operation + - * /")
fmt.Scanln(&operator)
switch operator{
  case "+":
  x:=add(num1,num2)
  fmt.Println(x)
  break
  case "-":
  x:=sub(num1,num2)
  fmt.Println(x)
  break
  case "*":
  x:=mult(num1,num2)
  fmt.Println(x)
  break
  case "/":
  x:=divv(num1,num2)
  fmt.Println(x)
  break

default:
fmt.Println("No operation selected" )
}


}
func add(x int, y int) int{
sum:= x+ y
return sum
}

func sub(x int, y int) int{
subt:= x- y
return subt
}
func divv(x int, y int) int{
divd:= x/y
return divd
}
func mult(x int, y int) int{
mul:= x* y
return mul
}