package main

import "fmt"

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

//for embeded structs(inheritance of structs)
// type messageToSend struct {
// 	message   string
// 	sender    user
// 	recipient user
//   user
// }

// type user struct {
// 	name   string
// 	number int
// }

//Struct methods in go
type authenticationInfo struct {
	username string
	password string
}

func (authI authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", authI.username, authI.password)
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name == "" {
		return false
	}
	if mToSend.recipient.name == "" {
		return false
	}
	if mToSend.sender.number == 0 {
		return false
	}
	if mToSend.recipient.number == 0 {
		return false
	}
	return true
}

func main() {
	mts := messageToSend{}
	mts.sender.name = "Aniket"
	mts.recipient.name = "Pranit"
	mts.sender.number = 0
	mts.recipient.number = 9423061943
	flag := canSendMessage(mts)
	fmt.Println(flag)

	au := authenticationInfo{}
	au.username = "Aniket"
	au.password = "Aniket@1234"

	authorizarion := au.getBasicAuth()
	fmt.Printf(authorizarion)

}
