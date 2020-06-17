package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	oauthTokenHTML, readFileError := ioutil.ReadFile("../../../../assets/html/oauth_token.html")
	if readFileError != nil {
		panic(readFileError)
	}
	constantFile := fmt.Sprintf("package assets\n\nconst OAuthTokenPage=`\n%s\n`", string(oauthTokenHTML))
	writeFileError := ioutil.WriteFile("./assets/oauth_token.go", []byte(constantFile), os.ModePerm)
	if writeFileError != nil {
		panic(writeFileError)
	}

	oauthTokenErrorHTML, readFileError := ioutil.ReadFile("../../../../assets/html/oauth_token_error.html")
	if readFileError != nil {
		panic(readFileError)
	}
	constantFile = fmt.Sprintf("package assets\n\nconst OAuthTokenErrorPage=`\n%s\n`", string(oauthTokenErrorHTML))
	writeFileError = ioutil.WriteFile("./assets/oauth_token_error.go", []byte(constantFile), os.ModePerm)
	if writeFileError != nil {
		panic(writeFileError)
	}
}
