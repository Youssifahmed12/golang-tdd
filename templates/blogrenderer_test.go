package blogrenderer_test

import (
	"bytes"
	"io"
	"testing"

	blogrenderer "github.com/Youssifahmed12/golang-tdd/templates"
	approvals "github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	pr, err := blogrenderer.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}
	t.Run("it converts a single post to html", func(t *testing.T) {
		b := bytes.Buffer{}

		if err := pr.Render(&b, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, b.String())
	})

}
func BenchmarkRender(b *testing.B) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := blogrenderer.NewPostRenderer()

	if err != nil {
		b.Fatal(err)
	}

	for b.Loop() {
		postRenderer.Render(io.Discard, aPost)
	}
}
