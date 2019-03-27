type CreateNewsRequest struct {
	Title string
	Body  string
}

func (a CreateNewsRequest) Validate() (errs url.Values) {
	if a.Title == "" {
		errs.Add("title", "The title field is required!")
	}
	if a.Body == "" {
		errs.Add("body", "The body field is required!")
	}
	if len(a.Body) < 50 || len(a.Body) > 500 {
		errs.Add("body", "The title field must be between 50-500 chars!")
	}
	return errs
}
