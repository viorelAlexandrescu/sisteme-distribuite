package main

import "fmt"

func printMessage(message string){
  fmt.Println(message)
}

func main() {
  x:="Hello World! From Golang!"
  printMessage("My message is:" + x)
}
