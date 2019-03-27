package article

type createArticle struct {
  Title string `json:"title" validate:"require,min=3,max=30"`
}

func Create(req *http.Request) error {
  c := createArticle{}
  // ... unmarshal the request body into the request struct

  if err := v := validator.New().Validate(c); err != nil {
    return err
  }

  // Map from input to your domain entity and use it
  article := Article {
    Title: c.Title
  }
}
