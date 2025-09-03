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
			Title: "hello world",
			Body: `# First recipe!
Welcome to my **amazing blog**. I am going to write about my family recipes, and make sure I write a long, irrelevant and boring story about my family before you get to the actual instructions.`,

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

	t.Run("it renders an index of posts", func(t *testing.T) {
		b := bytes.Buffer{}
		posts := []blogrenderer.Post{{Title: "Hello World"}, {Title: "Hello World2"}}

		if err := pr.RenderIndex(&b, posts); err != nil {
			t.Fatal(err)
		}

		want := `<ol><li><a href="/post/hello-world">Hello World</a></li><li><a href="/post/hello-world-2">Hello World 2</a></li></ol>`
		got := b.String()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
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
