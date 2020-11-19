package main

import (
	"awesomeProject4/adventure"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	jsonFile, err := os.Open("gopher.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal(err)
	}

	chapters, err := adventure.ChaptersToJson(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	tpl := template.Must(template.New("").Parse(storyTmpl))
	h := adventure.NewHandler(chapters,
		adventure.WithTemplate(tpl),
		adventure.WithPathFunc(pathFn),
	)
	// Create a ServeMux to route our requests
	mux := http.NewServeMux()

	mux.Handle("/chapters/", h)

	mux.Handle("/", adventure.NewHandler(chapters))

	fmt.Printf("Starting the server on port: %d\n", 8080)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 8080), mux))

}

func pathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	return path[len("/story/"):]
}

var storyTmpl = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <section class="page">
      <h1>{{.Title}}</h1>
      {{range .Story}}
        <p>{{.}}</p>
      {{end}}
      <ul>
      {{range .Options}}
        <li><a href="/story/{{.Arc}}">{{.Text}}</a></li>
      {{end}}
      </ul>
    </section>
    <style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FCF6FC;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #797;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: underline;
        color: #555;
      }
      a:active,
      a:hover {
        color: #222;
      }
      p {
        text-indent: 1em;
      }
    </style>
  </body>
</html>`
