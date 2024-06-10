package ownadventure

import (
	"net/http"
	"text/template"
)

const layoutTemp = "layout.html"

type BookServer struct {
	book Book
	t    *template.Template
	p    func(r *http.Request) string
}
type HandlerOption func(h *BookServer)

func WithTmpl(tmpl *template.Template) HandlerOption {
	return func(h *BookServer) {
		h.t = tmpl
	}
}

func WithPathFunc(fn func(r *http.Request) string) HandlerOption {
	return func(h *BookServer) {
		h.p = fn
	}

}

func NewBookServer(book Book, opts ...HandlerOption) (BookServer, error) {
	tmpl := template.Must(template.ParseFiles(layoutTemp))
	result := BookServer{book: book, t: tmpl, p: fn}

	for _, v := range opts {
		v(&result)
	}

	return result, nil
}

func (b BookServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if story, ok := b.book[b.p(r)]; ok {
		if err := b.t.Execute(w, story); err != nil {
			http.Error(w, "Server Error", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Not Found", http.StatusNotFound)
}

func fn(r *http.Request) string {
	url := r.URL.Path
	if url == "" || url == "/" {
		url = "/intro"
	}

	url = url[1:]
	return url
}
