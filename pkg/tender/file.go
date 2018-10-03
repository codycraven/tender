package tender

import "net/http"

// File is a single file serving Tender.
type File struct{}

// DeployTender for File struct attaches a single route to serve a single file.
func (f *File) DeployTender(path, route string) error {
	http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path)
	})
	return nil
}
