package listing

type Repository interface {
	GetAllArticles() []Article
}

type Service interface {
	GetAllArticles() []Article
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetAllArticles() []Article {
	return s.r.GetAllArticles()
}
