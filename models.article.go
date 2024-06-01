package main

type Article struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
}

var articleList = []Article{
	Article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	Article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

func getAllArticles() []Article {
	return articleList
}