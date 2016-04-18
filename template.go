package migo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"text/template"
)

// Migo is base type
type Migo struct {
	BasePath    string
	TemplateMap map[string][]string
}

// New returns new instance of Migo
func New(basePath string) Migo {
	t := Migo{BasePath: basePath}

	t.TemplateMap = createTemplateMap(&t)

	return t
}

func createTemplateMap(t *Migo) map[string][]string {
	m := make(map[string][]string)

	//loop through base path and create key {path/filename} for each template
	if templates, err := ioutil.ReadDir(t.BasePath); len(templates) > 0 && err == nil {
		for _, template := range templates {
			dirKey := ""

			if template.IsDir() {
				dirKey = template.Name()

				// each key should have values of templatefile and layout (wether sub layout or shared layout)
				if fileNames, _ := t.getTemplatesFromDir(dirKey); len(fileNames) > 0 {
					for _, file := range fileNames {
						n := strings.Split(file[0], "/")
						nameWExtension := n[len(n)-1]                                                // last item in slice is filename
						key := fmt.Sprintf("%v/%v", dirKey, nameWExtension[0:len(nameWExtension)-5]) // cut off extension name
						m[key] = file                                                                // nice key values i.e. "account/index" = "templates/account/index.tmpl"
					}
				}
			} else {
				continue
			}
		}
	}

	return m
}

func (t *Migo) getTemplatesFromDir(dir string) ([][]string, error) {
	var fileNames [][]string

	if dir == "shared" {
		return make([][]string, 0), nil
	}

	if files, err := ioutil.ReadDir(t.BasePath + "/" + dir); len(files) > 0 && err == nil {
		for _, file := range files {
			var fileMap []string
			if !file.IsDir() {
				fileName := file.Name()

				fileMap = append(fileMap, fmt.Sprintf("%v/%v/%v", t.BasePath, dir, fileName))

				// if this template dir doesn't have a shared section with layout.tmpl use globalTemplate (basepath/shared/layout.tmpl)
				if !t.hasSubLayout(dir) {
					if _, err := os.Stat(fmt.Sprintf("%v/shared/layout.tmpl", t.BasePath)); err == nil {
						fileMap = append(fileMap, fmt.Sprintf("%v/shared/layout.tmpl", t.BasePath))
					}
				} else {
					fileMap = append(fileMap, fmt.Sprintf("%v/%v/shared/layout.tmpl", t.BasePath, dir))
				}

				fileNames = append(fileNames, fileMap)
			}
		}

		return fileNames, nil
	}

	err := fmt.Errorf("directory %v not found in templates folder", dir)
	return fileNames, err
}

func (t *Migo) hasSubLayout(dir string) bool {
	if _, err := os.Stat(fmt.Sprintf("%v/%v/shared/layout.tmpl", t.BasePath, dir)); err == nil {
		return true
	}

	return false
}

// Render accepts takes templates in map for key given and writes to http ResponseWriter
func (t *Migo) Render(rw http.ResponseWriter, key string, data interface{}) {
	tmpl, _ := template.ParseFiles(t.TemplateMap[key]...)
	tmpl.Execute(rw, data)
}
