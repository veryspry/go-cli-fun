package main

import (
  "fmt"
  "os"
  "./calc"
  "errors"
  "strconv"
  // "github.com/c-bata/go-prompt"
  p "github.com/manifoldco/promptui"
)


func main() {
  // init calculator
  if os.Args[1] == "calc" {
    fmt.Println("result:", calc.Calc(os.Args[2], os.Args[3:]))
  }

  if os.Args[1] == "opencalc" {
  }
  prompter()


  // fmt.Println("Please select table.")
	// t := prompt.Input("> ", completer)
	// fmt.Println("You selected " + t)
}

func prompter() {
  validate := func(i string) error {
    _, err := strconv.ParseFloat(i, 64)
    if err != nil {
      return errors.New("Invalid Number")
    }
    return nil
  }

  prompt := p.Prompt{
    Label: "Number",
    Validate: validate,
  }

  result := prompt.Run()

  if err != nil {
    fmt.Printf("Promt failed %v\n", err)
    err
  }

  fmt.Printf("You chose: ", result)
}


// func completer(d prompt.Document) []prompt.Suggest {
// 	s := []prompt.Suggest {
// 		{Text: "users", Description: "Store the username and age"},
// 		{Text: "articles", Description: "Store the article text posted by user"},
// 		{Text: "comments", Description: "Store the text commented to articles"},
// 	}
// 	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
// }
