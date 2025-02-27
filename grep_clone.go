// ---------------------------------------------------------
// EXERCISE: Grep Clone
//
//  Create a grep clone. grep is a command-line utility for
//  searching plain-text data for lines that match a specific
//  pattern.
//
//  1. Feed the shakespeare.txt to your program.
//
//  2. Accept a command-line argument for the pattern
//
//  3. Only print the lines that contains that pattern
//
//  4. If no pattern is provided, print all the lines
//
//
// EXPECTED OUTPUT
//
//  go run main.go come < shakespeare.txt
//
//    come night come romeo come thou day in night
//    come gentle night come loving black-browed night
//
// ---------------------------------------------------------
package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  )


func main() {
  var pattern int
  input:=bufio.NewScanner(os.Stdin)
  if args:=os.Args[1:];len(args)==1{
    pattern =args[0]
    }else{
    pattern=args[1]
    }
  for input.Scan(){
    if strings.Contains(input.Text(),pattern){
      fmt.Println(input.Text())
      }
    }
}
