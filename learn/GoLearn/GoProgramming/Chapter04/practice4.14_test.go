package Chapter03

import (
	"html/template"
	"net/http"
	"testing"
)

func start0414() {

	const templ = `<p>A: {{.A}} </p>  <p>B: {{.B}}</p>`
	t := template.Must(template.New("escape").Parse(templ))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data struct {
			A string
			B template.HTML
		}
		data.A = "<b>Hello!</b>"
		data.B = "<b>Hello!</b>"

		t.Execute(w, data)
	})
	http.ListenAndServe(":8888", nil)
}

// go test -v -run Test0414 practice4.14_test.go
func Test0414(t *testing.T) {
	start0414()
}
