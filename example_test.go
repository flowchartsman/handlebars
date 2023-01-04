package handlebars_test

import (
	"fmt"

	"github.com/steeringwaves/go-handlebars"
)

func Example() {
	h := handlebars.New()
	source := "<h1>{{title}}</h1><p>{{body.content}}</p>"

	ctx := map[string]interface{}{
		"title": "foo",
		"body":  map[string]string{"content": "bar"},
	}

	// parse template
	tpl := h.MustParse(source)

	// evaluate template with context
	output := tpl.MustExec(ctx)

	// alternatively, for one shots:
	// output :=  MustRender(source, ctx)

	fmt.Print(output)
	// Output: <h1>foo</h1><p>bar</p>
}

func Example_struct() {
	h := handlebars.New()
	source := `<div class="post">
  <h1>By {{fullName author}}</h1>
  <div class="body">{{body}}</div>

  <h1>Comments</h1>

  {{#each comments}}
  <h2>By {{fullName author}}</h2>
  <div class="body">{{content}}</div>
  {{/each}}
</div>`

	type Person struct {
		FirstName string
		LastName  string
	}

	type Comment struct {
		Author Person
		Body   string `handlebars:"content"`
	}

	type Post struct {
		Author   Person
		Body     string
		Comments []Comment
	}

	ctx := Post{
		Person{"Jean", "Valjean"},
		"Life is difficult",
		[]Comment{
			{
				Person{"Marcel", "Beliveau"},
				"LOL!",
			},
		},
	}

	h.RegisterHelper("fullName", func(person Person) string {
		return person.FirstName + " " + person.LastName
	})

	output := h.MustRender(source, ctx)

	fmt.Print(output)
	// Output: <div class="post">
	//   <h1>By Jean Valjean</h1>
	//   <div class="body">Life is difficult</div>
	//
	//   <h1>Comments</h1>
	//
	//   <h2>By Marcel Beliveau</h2>
	//   <div class="body">LOL!</div>
	// </div>
}

func ExampleRender() {
	h := handlebars.New()
	tpl := "<h1>{{title}}</h1><p>{{body.content}}</p>"

	ctx := map[string]interface{}{
		"title": "foo",
		"body":  map[string]string{"content": "bar"},
	}

	// render template with context
	output, err := h.Render(tpl, ctx)
	if err != nil {
		panic(err)
	}

	fmt.Print(output)
	// Output: <h1>foo</h1><p>bar</p>
}

func ExampleMustRender() {
	h := handlebars.New()
	tpl := "<h1>{{title}}</h1><p>{{body.content}}</p>"

	ctx := map[string]interface{}{
		"title": "foo",
		"body":  map[string]string{"content": "bar"},
	}

	// render template with context
	output := h.MustRender(tpl, ctx)

	fmt.Print(output)
	// Output: <h1>foo</h1><p>bar</p>
}