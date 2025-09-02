package blogrenderer_test

import (
	"bytes"
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

	t.Run("it converts a single post to html", func(t *testing.T) {
		b := bytes.Buffer{}
		err := blogrenderer.Render(&b, aPost)

		if err != nil {
			t.Fatalf("shoudlnt got an error here")
		}

		approvals.VerifyString(t, b.String())
	})

}
