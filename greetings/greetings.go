package greetings

import (
	"fmt"
	"errors"
)

// This is a function
// we have variable followed by type
// at the end we have return type
func Hello(name string) (string,error){
  // If no name is provided, we return an error with message
  if name == ""{
	return "", errors.New("empty name")
  }
  // message variable is dynamically typeset
  // := operator is used for declaring and initializing
  message := fmt.Sprintf("Hi, %v. Welcome!",name)
  // so normally this would be the case for variable declaration
  var tmp string
  tmp = "GOはなんですか。"
  // interestingly code will not compile if there are unused variables
  fmt.Println(tmp)
  // the nil specifies the error if it did not occur
  return message,nil
}