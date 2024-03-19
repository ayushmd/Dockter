package templ

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderAndExecute(filename string, w http.ResponseWriter, data interface{}) {
	tmpl := template.Must(template.ParseFiles(fmt.Sprintf("master\\templ\\templates\\%s", filename)))
	tmpl.Execute(w, data)
}

func RenderAndExecuteTempl(filename string, w http.ResponseWriter, name string, data interface{}) {
	tmpl := template.Must(template.ParseFiles(fmt.Sprintf("master\\templ\\templates\\%s", filename)))
	tmpl.ExecuteTemplate(w, "", data)
}
