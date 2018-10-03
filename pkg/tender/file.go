package tender

import "net/http"

type File struct{}

// DeployTender for File struct attaches a single .
func (f *File) DeployTender(path, route string) error {
	http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path)
	})
	return nil
}
