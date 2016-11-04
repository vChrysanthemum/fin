package tinyiron

import (
	"net/http"
	"text/template"
)

type View struct {
	filenames []string
	tmpl      *template.Template
}

func (p *Server) Render(path string, w http.ResponseWriter, r *http.Request, ViewData map[string]interface{}) {
	w.Header().Add("Content-Type", "text/html;charset=utf-8")
	w.Header().Add("Server", "tinyiron")

	var tmpl *template.Template

	if p.isTMPLAutoRefresh {
		tmpl = template.Must(template.New(path).Delims("<?", "?>").ParseFiles(p.views[path].filenames...))
	} else {
		if nil == p.views[path].tmpl {
			p.views[path].tmpl = template.Must(template.New(path).Delims("<?", "?>").ParseFiles(p.views[path].filenames...))
		}
		tmpl = p.views[path].tmpl
	}

	tmpl.ExecuteTemplate(w, "frame", ViewData)
}

func (p *Server) AssignView(path string, _viewFilenames []string) {
	viewFilenames := []string{}
	for _, v := range _viewFilenames {
		if "" == v {
			break
		}
		viewFilenames = append(viewFilenames, p.ViewDir+"/"+v)
	}
	p.views[path] = &View{viewFilenames, nil}
}
