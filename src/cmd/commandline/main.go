package main

import "fmt"

func main() {
	fmt.Print("Enter AoC day you want to execute (1-12): ")

  var day string
  fmt.Scanln(&day)

  switch day {
  case "1":
    printDayHeader("1")
    fmt.Printf("TBD")
    return
  default:
    fmt.Printf("Invalid Day: %s", day)
    return
  }
}

func printDayHeader(day string) {
	fmt.Printf("Day %s: \n", day)
}

func printProblemTitleAndAnswer(title string, data int) {
	fmt.Printf("\t%s: %d\n", title, data)
}
