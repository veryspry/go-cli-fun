package flagfun

import (
  flag "github.com/ogier/pflag"
  "fmt"
  "strings"
)

// This is essentially a user model
type User struct {
  firstName, lastName string
  age int
}

var user string
var matt = new(User)

func checkFlag() {
  // if user does not supply flags, print usage
  // we can clean this up later by putting this into its own function
   if flag.NFlag() == 0 {
      fmt.Printf("Usage: %s [options]\n", os.Args[0])
      fmt.Println("Options:")
      flag.PrintDefaults()
      os.Exit(1)
   }
  users := strings.Split(user, ",")
  fmt.Printf("Searching user(s): %s\n", users)
}

func setUser(u User, fn string, ln string, a int) {
  fmt.Println("before:", u)
  u.firstName = fn
  u.lastName = ln
  u.age = a
  fmt.Println("after:", u)
}

func init() {
  // point our flag to the user var
  flag.StringVarP(&user, "user", "u", "", "Search Users")

  flag.Parse()
  checkFlag()
  setUser(*matt, "matt", "ehlinger", 25)
}
