package servicego

import (
	"log"
	"os"
	"path/filepath"
)

var defaultServicesRoot = []string{"/", "opt", "servicego"}

func ServicesRoot() string { return filepath.Join(defaultServicesRoot...) }

func ServiceName() string { return filepath.Base(MustExecutable()) }
func ServiceDir() string  { return filepath.Join(ServicesRoot(), ServiceName()) }
func ServicePath() string { return filepath.Join(ServiceDir(), ServiceName()) }

func MustExecutable() string {
	executable, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	return executable
}
