package X

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:title`
	Content string `json:content`
}

var X int
X = 3

var ArticleList []Article
ArticleList = []Article{
	Article{ID: 1, Title: "Article 1", Content: "Body of Article 1"},
	Article{ID: 1, Title: "Article 1", Content: "Body of Article 1"},
}

func getAllArticles() []Article {
	return ArticleList
}