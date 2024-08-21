package veni_legacy

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func (d *veniContext) Open(name string) (http.File, error) {
	fmt.Println(name)
	// Try name as supplied
	file, err := d.d.Open(name)
	if os.IsNotExist(err) {
		// Not found, try with .html
		editFileName := name + ".html"
		if file, err = d.d.Open(editFileName); err != nil {
			fmt.Println("Kicking default http.FileServer...")
			return nil, err
		}
	}
	//check if this is a directory, if so attempt to load index.html
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Cant check fileInfo default http.FileServer...")
		return nil, err
	}

	// IsDir is short for fileInfo.Mode().IsDir()
	if fileInfo.IsDir() {
		//attempt with index.html
		newFileName := name + "index.html"
		fmt.Println("New Index Filename: " + newFileName)
		if file, err = d.d.Open(newFileName); err != nil {
			fmt.Println("Kicking default http.FileServer...")
			return nil, err
		}
	}
	//we found a file
	fmt.Println("File Found!")
	return file, err
}

func isVENI(str string) bool {
	return strings.Contains(str, "<veni-")
}

func isStartTag(str string) bool {
	return !strings.Contains(str, "/")
}

func isFinalTag(str string) bool {
	return strings.Contains(str, "<html/>") || strings.Contains(str, "</html>")
}

func (d *veniContext) assignVeniComponent(targetLine string) {
	if targetLine != "" {
		initialTargets := strings.Fields(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(targetLine, "/", " "), "<", ""), ">", ""))
		d.veniTargets = append(d.veniTargets, initialTargets[0])
	}
}

func (d *veniContext) generateVeniComponent() error {
	if d.veniTargets != nil {
		//create an error buffer for failure case
		errorBuffer := []string{"veni", "vidi", "vici"}
		copy(errorBuffer, d.generatedFile)
		//this means there are in fact targets present
		d.generatedFile = append(d.generatedFile, "<script>")
		for i, v := range d.veniTargets {
			file, e := d.Open(v)
			if e == nil {
				fmt.Println(v)
				fmt.Println(i)
				//load the generated script into the section here
				d.generatedFile = append(d.generatedFile, "customElements.define(")
				d.generatedFile = append(d.generatedFile, "'"+v+"',")
				d.generatedFile = append(d.generatedFile, "class extends HTMLElement {constructor() {super();template=document.createElement('template');template.innerHTML='")
				//generated file output
				e = d.scanFileBody(file)
				if e != nil {
					d.generatedFile = errorBuffer
					return e
				}
				d.generatedFile = append(d.generatedFile, "';let templateContent = template.content;const shadowRoot = this.attachShadow({ mode: 'open' });shadowRoot.appendChild(templateContent.cloneNode(true));}},);\n")
			} else {
				fmt.Println(e)
			}
		}
		d.generatedFile = append(d.generatedFile, "</script>")
	}
	return nil
}

func (d *veniContext) scanFileBody(target http.File) error {
	var err error
	defer target.Close()
	fmt.Println("Attempt to process HTML")
	scanner := bufio.NewScanner(target)

	const maxCapacity int = 100000 // your required line length
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	inBody := false

	for scanner.Scan() {
		textFound := scanner.Text()

		if inBody {
			d.generatedFile = append(d.generatedFile, textFound)
		}
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return err
}

func (d *veniContext) processHTML(target http.File) error {
	var err error
	defer target.Close()
	fmt.Println("Attempt to process HTML")
	scanner := bufio.NewScanner(target)

	const maxCapacity int = 100000 // your required line length
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	for scanner.Scan() {
		textFound := scanner.Text()
		fmt.Println(textFound)
		if isStartTag(textFound) {
			if isVENI(textFound) {
				d.assignVeniComponent(textFound)
			}
		}
		if isFinalTag(textFound) {
			err = d.generateVeniComponent()
		}
		d.generatedFile = append(d.generatedFile, textFound)
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return err
}

type veniContext struct {
	d             http.Dir
	generatedFile []string
	err           error
	veniTargets   []string
}

func (d *veniContext) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d.generatedFile = []string{}
	fmt.Printf("Req: %s %s\n", r.Host, r.URL.Path)
	name := r.URL.Path
	file, err := d.Open(name)
	if err != nil {
		fmt.Println("Issue opening file, please review")
		log.Fatal(err)
	}
	err = d.processHTML(file)
	if err != nil {
		fmt.Println("Issue reading file, please review")
		log.Fatal(err)
	}
	//close my file
	fmt.Println("File processed!")
	//write line by line to resp
	charCount := 0
	d.generatedFile = []string{}
	w.Header().Set("Content-Type", "text/html")
	for i, v := range d.generatedFile {
		value := []byte(v + "\n")
		tmpCount, erWrite := w.Write(value)
		charCount += tmpCount
		if erWrite != nil {
			fmt.Printf("Issue writing please review at line: %n\n", i)
			log.Fatal(erWrite)
		} else {
			fmt.Printf("Wrote line: %n\n", i)
		}
	}
}

func Load(fs http.FileSystem) http.Handler {
	context := &veniContext{fs.(http.Dir), nil, nil, nil}
	if context.err != nil {
		return http.FileServer(fs)
	}
	return context
}
