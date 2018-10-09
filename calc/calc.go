package calc

import (
  "fmt"
  "strconv"
)

func Calc(o string, n []string) float64 {
  var nums []float64
  // convert slice of strings to float64
  for i := 0; i < len(n); i++ {
    t, err := strconv.ParseFloat(n[i], 64)
    if err != nil {
      panic("String conversion failed!")
    }
    nums = append(nums, t)
  }
  if o == "add" {
    return Sum(nums)
  }
  if o == "subt" {
    return Diff(nums)
  }
  if o == "mult" {
    return Product(nums)
  }
  if o == "div" {
    return Quotient(nums)
  }
  err := o + " is not a valid calculator operation!"
  panic(err)
}

func Sum(nums []float64) float64 {
  fmt.Println("nums", nums)
  var res float64 = 0
  for i := 0; i < len(nums); i++ {
    res += nums[i]
  }
  return res
}

func Diff(nums []float64) float64 {
  var res float64 = 0
  for i := 0; i < len(nums); i++ {
    res -= nums[i]
  }
  return res
}

func Product(nums []float64) float64 {
  // error handling for no nums
  if len(nums) == 0 {
    return 0
  }
  var res float64 = nums[0]
  for i := 1; i < len(nums); i++ {
    res *= nums[i]
  }
  return res
}

func Quotient(nums []float64) float64 {
  if len(nums) == 0 {
    return 0
  }
  var res float64 = nums[0]
  for i := 1; i < len(nums); i++ {
    res /= nums[i]
  }
  return res
}
