package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

// Struct Method
func (authI authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf(
		"Authorization: Basic %s:%s",
		authI.username,
		authI.password,
	)
}

func authenticate(authInfo authenticationInfo){
	fmt.Println(authInfo.getBasicAuth())
}

func main(){
	curr_user := authenticationInfo{
		"gfssp",
		"1234pwd",
	}

	fmt.Println(curr_user.getBasicAuth())
}