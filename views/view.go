package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	//LayoutDir is to used as path
	LayoutDir string = "views/layouts/"
	//TemplateExt is the extension for LayoutDir
	TemplateExt string = ".gohtml"
)

//NewView func creates our templates
func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}

//View type for NewView func
type View struct {
	Template *template.Template
	Layout   string
}

//Render is used for rendering templates
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}
