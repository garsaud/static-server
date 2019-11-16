package main

import (
	"net/http"
)

func main() {
	var fs = http.Dir("www");
	var indexPage = "index.html"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var name = r.URL.Path;

		// Inspect requested path
		f, err := fs.Open(name)
		if err != nil {
			// Fallback to index.html if resource wasn’t found
			f, _ = fs.Open(indexPage)
		}
		defer f.Close()
		d, _ := f.Stat()

		// Use content of /index.html if requested path is a directory
		if d.IsDir() {
			name = name + "/" + indexPage
			f, err = fs.Open(name)
			if err != nil {
				// Fallback to index.html if resource wasn’t found
				f, _ = fs.Open(indexPage)
			}
			defer f.Close()
			d, _ = f.Stat()
		}

		http.ServeContent(w, r, d.Name(), d.ModTime(), f)
	})

	_ = http.ListenAndServe(":80", nil)
}
