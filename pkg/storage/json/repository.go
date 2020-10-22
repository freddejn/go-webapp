package json

import (
	"encoding/json"
	"fmt"

	"github.com/freddejn/go-webapp/pkg/domain/listing"
)

const (
	dir               = "/data/"
	CollectionArticle = "article"
)

type Storage struct {
	db *jsonDriver
}

func NewStorage() (*Storage, error) {
	var err error
	fmt.Println("stored object")

	if err != nil {
		return nil, err
	}
	s := new(Storage)
	return s, nil
}

func (s *Storage) GetAllArticles() []listing.Article {
	fmt.Println("get all articles")
	list := []listing.Article{}
	records, err := s.db.ReadAll(CollectionArticle)
	if err != nil {
		return list
	}
	for i := 0; i < len(records); i++ {
		var a Article
		var article listing.Article
		if err := json.Unmarshal([]byte(records[i]), &a); err != nil {
			return list
		}
		article.Title = a.Title
		article.Content = a.Content
		list = append(list, article)

	}
	return list
}

type jsonDriver struct {
	dir string
}

func (j *jsonDriver) ReadAll(collection string) ([]string, error) {
	return []string{"{\"Title\":\"Title1\", \"Content\":\"content1\"}", "{\"Title\":\"Peter\", \"Content\":\"Jones\"}"}, nil
}
