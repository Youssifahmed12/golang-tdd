package blogrenderer

import (
	"embed"
	"html/template"
	"io"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

type PostRenderer struct {
	templ *template.Template
}

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ}, nil
}

func (pr *PostRenderer) Render(w io.Writer, p Post) error {

	if err := pr.templ.ExecuteTemplate(w, "blog.gohtml", p); err != nil {
		return err
	}

	return nil
}
