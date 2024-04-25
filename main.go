package main

import (
	"fmt"
	"time"
)

type Item struct {
	Name string
	Date time.Time
	Tags string
	Link string
}

func addItem(name, tags, link string) *Item {
	item := &Item{
		Name: name,
		Date: time.Now(),
		Tags: tags,
		Link: link,
	}
	return item
}

func main() {

	item1 := addItem("GB", "education", "https://gb.ru")
	fmt.Println("Name:", item1.Name)
	fmt.Println("Date:", item1.Date.Format("2006-01-02 15:04:05"))
	fmt.Println("Tags:", item1.Tags)
	fmt.Println("Link:", item1.Link)
}
