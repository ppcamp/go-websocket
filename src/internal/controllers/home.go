package controllers

import (
	"log"
	"net/http"
	"path/filepath"
	"src/internal/config"

	"github.com/sirupsen/logrus"
)

func Home(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)

	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	p, err := filepath.Abs(filepath.Join(".", config.App.PublicFolder))
	if err != nil {
		logrus.Fatal(err)
	}

	http.ServeFile(w, r, p)
}
