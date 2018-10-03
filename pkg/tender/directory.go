package tender

import (
	"net/http"
	"os"
)

// Directory is a directory serving Tender.
type Directory struct{}

// DeployTender for Directory to serve directory.
func (d *Directory) DeployTender(path, route string) error {
	fs := http.FileServer(http.Dir(path))
	h := http.StripPrefix(route, fs)
	http.Handle(route, h)
	return nil
}

// DirectoryNoListing is a directory serving Tender that does not list the
// contents of the directory upon 404.
type DirectoryNoListing struct {
	Fs http.FileSystem
}

// DeployTender for DirectoryNoListing to serve directory.
func (d *DirectoryNoListing) DeployTender(path, route string) error {
	d.Fs = http.Dir(path)
	h := http.StripPrefix(route, http.FileServer(d))
	http.Handle(route, h)
	return nil
}

// Open is the http.FileSystem method wrapping http.FileSystem.Open that
// prevents directory listing.
func (d DirectoryNoListing) Open(name string) (http.File, error) {
	f, err := d.Fs.Open(name)
	if err != nil {
		return nil, err
	}
	stat, err := f.Stat()
	if stat.IsDir() {
		return nil, os.ErrNotExist
	}
	return f, nil
}
