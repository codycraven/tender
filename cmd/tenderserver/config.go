package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/codycraven/tender/pkg/tender"
)

type config struct {
	Tenders []Tender
	Port    int
}

// UnmarshalJSON interface method to
func (c *config) UnmarshalJSON(b []byte) error {
	var raw struct {
		Tenders []map[string]string `json:"tenders"`
		Port    int                 `json:"port"`
	}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	c.Port = raw.Port

	// To support third party tenders we could generate this...
	for _, v := range raw.Tenders {
		var t Tender
		switch v["type"] {
		case "directory":
			t = &tender.Directory{}
			t.DeployTender(v["path"], v["route"])
		case "directory no listing":
			t = &tender.DirectoryNoListing{}
			t.DeployTender(v["path"], v["route"])
		case "file":
			t = &tender.File{}
			t.DeployTender(v["path"], v["route"])
		default:
			return fmt.Errorf("tender does not have supported type %v", v)
		}
		log.Println("Deployed", v["type"], "tender for", v["route"])
		c.Tenders = append(c.Tenders, t)
	}
	return nil
}
