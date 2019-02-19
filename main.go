package main

import (
	"io/ioutil"
	"fmt"
	test "./go_out"
	"github.com/golang/protobuf/proto"
)

func main() {
	in, err := ioutil.ReadFile("./log.txt")
	if err != nil {
		fmt.Printf("The open file error is : %v.\n", err)
		return
	}
	person := test.Person{}
	if err := proto.Unmarshal(in, &person); err != nil {
		fmt.Printf("The unmarshal error is : %v.\n", err)
		return
	}
	fmt.Printf("The person name is : %v.\n", *person.Name)
	fmt.Printf("The person id is : %v.\n", *person.Id)
	fmt.Printf("The person email is : %v.\n", *person.Email)
	phones := person.GetPhones()
	for i, v := range phones {
		fmt.Printf("The %v 's phone is : %v.\n", i, *v.Number);
		fmt.Printf("The %v 's phone type is : %v.\n", i, *v.Type)
	}
}
