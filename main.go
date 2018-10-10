package main

import (
  "fmt"
  "os"
  "./calc"
  "errors"
  "strconv"
  // "github.com/c-bata/go-prompt"
  "github.com/manifoldco/promptui"
)


func main() {
  // init calculator
  if os.Args[1] == "calc" {
    fmt.Println("result:", calc.Calc(os.Args[2], os.Args[3:]))
  }

  if os.Args[1] == "opencalc" {
  }

  if os.Args[1] == "prompt" {
    prompter()
  }


}

func prompter() {
    validate := func(i string) error {
    _, err := strconv.ParseFloat(i, 64)
    if err != nil {
      return errors.New("Invalid Number")
    }
    return nil
  }

  prompt := promptui.Prompt{
    Label: "Number",
    Validate: validate,
  }

  result, err := prompt.Run()

  if err != nil {
    fmt.Printf("Promt failed %v\n", err)
    return
  }

  fmt.Printf("You chose %q\n ", result)
  fmt.Printf("You chose \n", result) // what the hell does %q do??
}
