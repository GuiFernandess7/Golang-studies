package main

import (
	"fmt"
	"strings"
	"time"

	uuid "github.com/nu7hatch/gouuid"
)

type Options interface {
	add_article(article Article) Account
}

func main() {
    articles := []Account{}
    u, err := uuid.NewV4()
    if err != nil {
        fmt.Println("Failed to generate UUID:", err)
        return
    }

    myAccount := Account{
        uuid: 	   u.String(),
        email:    "guilherme@email.com",
        username: "john_doe",
        password: "pwd#1234",
		articles: nil,
    }

	update_db(&myAccount, &articles)

	my_article := Article{
		title: "First Title",
		content: "First content",
		publi_date: strings.Split(time.Now().String(), " ")[0],
	}

    myAccount.add_article(my_article)
    if len(articles) > 0 {
        if articles[len(articles)-1].email == myAccount.email {
            fmt.Println("Article published successfully!", articles)
        } else {
            fmt.Println("Failed to publish article")
        }
    } else {
        fmt.Println("Failed to read articles list")
    }
}

type Article struct {
    title     string
    content   string
    publi_date string
}

type Account struct {
	uuid     string
    email    string
    username string
    password string
	articles []Article
}

type PremiumAccount struct {
	uuid 	   string
    email      string
    username   string
    password   string
    premium_id string
	articles []Article
}

func (a *Account) add_article(article Article) *Account {
    a.articles = append(a.articles, article)
    return a
}

func (a PremiumAccount) add_article(article Article) PremiumAccount {
	a.articles = append(a.articles, article)
	return a
}

func update_db(a *Account, db *[]Account) {
	*db = append(*db, *a)
}