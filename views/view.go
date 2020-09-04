package views

import "html/template"

//NewView func creates our templates
func NewView(files ...string) *View {
	files = append(files, "views/layouts/footer.gohtml")
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
	}
}

//View type for NewView func
type View struct {
	Template *template.Template
}
