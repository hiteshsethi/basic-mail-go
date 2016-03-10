#basic-mail-go
# Introduction
Basic Mail Sending Through SMTP Using Golang
# Usage
```
package main

import (
	"basicmail"
	"fmt"
)

func main(){
	err := basicmail.SendEmail("to1@email.com,to2@email.com","from@email.com","Testing Mail Go","Some Body","127.0.0.1",25)
	if err != nil {
		fmt.Println("Some Error")
	}else {
		fmt.Println("Success")
	}
}
```
### Developed by Hitesh sethi

