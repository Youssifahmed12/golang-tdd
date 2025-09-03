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
	if err := pr.templ.ExecuteTemplate(w, "blog.gohtml", p); err != nil {
		return err
	}

	return nil
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	indexTemplate := `<ol>{{range .}}<li><a href="/post/{{sanitiseTitle .Title}}">{{.Title}}</a></li>{{end}}</ol>`

	templ, err := template.New("index").Funcs(template.FuncMap{
		"sanitiseTitle": func(title string) string {
			return strings.ToLower(strings.Replace(title, " ", "-", -1))
		},
	}).Parse(indexTemplate)
	if err != nil {
		return err
	}

	if err := templ.Execute(w, posts); err != nil {
		return err
	}

	return nil
}

func markDownToHtml(input string) (template.HTML, error) {
	b := bytes.Buffer{}
	if err := goldmark.Convert([]byte(input), &b); err != nil {
		return "", err
	}
	return template.HTML(b.String()), nil
}
