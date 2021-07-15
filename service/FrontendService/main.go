package FrontendService

import (
	"embed"
	"encoding/json"
	"os"
	"strings"
	"time"
)

//go:embed dist
var dist embed.FS

type item struct {
	Title            string     `json:"title"`
	ProtectionStatus bool       `json:"protectionStatus"`
	Items            []fileItem `json:"items"`
	Link             string     `json:"link"`
}

type fileItem struct {
	Name  string    `json:"name"`
	Mtime time.Time `json:"mtime"`
}

func getJSON(services map[string]bool, files map[string][]string) ([]byte, error) {
	result := []item{}
	for service, protectionStatus := range services {
		currentFiles := []fileItem{}
		for _, file := range files[service] {
			fi, err := os.Stat(file)
			if err != nil {
				panic(err)
			}
			currentFiles = append(currentFiles, fileItem{
				Name:  fi.Name(),
				Mtime: fi.ModTime(),
			})
		}
		current := item{
			Title:            service,
			ProtectionStatus: protectionStatus,
			Items:            currentFiles,
			Link:             "/" + service + "/",
		}
		result = append(result, current)
	}
	return json.Marshal(result)
}

func GenerateHtml(services map[string]bool, files map[string][]string) (string, string) {
	tpl, err := dist.ReadFile("dist/index_tpl.html")
	if err != nil {
		panic(err)
	}
	js, err := dist.ReadFile("dist/app.js")
	if err != nil {
		panic(err)
	}
	json, err := getJSON(services, files)
	if err != nil {
		panic(err)
	}

	newContent := strings.Replace(string(tpl), "#ITEMS#", string(json), -1)

	return newContent, string(js)
}
