package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Options interface {
	addArticle(article Article) Account
}

var data = map[string]Account{}

func main() {
	u := uuid.New()
	myAccount := Account{
		uuid:     u.String(),
		email:    "guilherme@email.com",
		username: "john_doe",
		password: "pwd#1234",
		articles: nil,
	}

	updateDB(&myAccount)

	myArticle := Article{
		title:      "First Title",
		content:    "First content",
		publicDate: time.Now().Format("2006-01-02"),
	}

	addArticle(myArticle, myAccount)
	if len(data[myAccount.uuid].articles) > 0 {
		fmt.Println("Article published successfully!", data[myAccount.uuid])
	} else {
		fmt.Println("Failed to publish article")
	}
}

type Article struct {
	title      string
	content    string
	publicDate string
}

type Account struct {
	uuid     string
	email    string
	username string
	password string
	articles []Article
}

func addArticle(myArticle Article, myAccount Account) {
	account, ok := data[myAccount.uuid]
	if !ok {
		return
	}
	account.articles = append(account.articles, myArticle)
	data[myAccount.uuid] = account
}

func updateDB(a *Account) {
	data[a.uuid] = *a
}
