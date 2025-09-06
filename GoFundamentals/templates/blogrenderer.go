package blogrenderer

import (
	"bytes"
	"embed"
	"html/template"
	"io"
	"strings"

	"github.com/yuin/goldmark"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        template.HTML
}

type PostRenderer struct {
	templ *template.Template
}

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ}, nil
}

func (pr *PostRenderer) Render(w io.Writer, p Post) error {

	htmlBody, err := markDownToHtml(string(p.Body))
	if err != nil {
		return err
	}

	p.Body = htmlBody

	return pr.templ.ExecuteTemplate(w, "blog.gohtml", p)
}

func (pr *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return pr.templ.ExecuteTemplate(w, "index.gohtml", posts)
}

func markDownToHtml(input string) (template.HTML, error) {
	b := bytes.Buffer{}
	if err := goldmark.Convert([]byte(input), &b); err != nil {
		return "", err
	}
	return template.HTML(b.String()), nil
}
