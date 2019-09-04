package liveprobe

import (
	"log"
	"os"
)

// MakeAlive creates a temporary file for liveliness check.
func MakeAlive(name string) {
	_, err := os.Create(name)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("server was marked as alive")
}

// MakeDead deletes the temp file indicating that the service is dead
func MakeDead(name string) {
	os.Remove(name)
	log.Println("service was marked as dead")
}

// CheckLiveliness checks if the temp file still exists for liveliness check.
func CheckLiveliness(name string) bool {
	if _, err := os.Stat(name); err == nil {
		return true
	}

	return false
}
