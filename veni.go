package veni

import(
	"net/http"
	//"os"
	//"fmt"
    //"log"
    //"bufio"
)

/*
type HTMLDir struct {
    d http.Dir
}

func (d HTMLDir) Open(name string) (http.File, error) {
    fmt.Println(name)
    // Try name as supplied
    file, err := d.d.Open(name)
    if os.IsNotExist(err) {
        // Not found, try with .html
        editFileName := name+".html"
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
        fmt.Println("New Index Filename: "+newFileName)
        if file, err = d.d.Open(newFileName); err != nil {
            fmt.Println("Kicking default http.FileServer...")
            return nil, err
        }
    }
    //we found a file
    fmt.Println("File Found! -> Processing HTML")
    //return processHTML(file), err
    return file, err
}

func (d HTMLDir) processHTML(target http.File) http.File{
    fmt.Println("Attempt to process HTML")
    scanner := bufio.NewScanner(target)

    const maxCapacity int = 100000  // your required line length
    buf := make([]byte, maxCapacity)
    scanner.Buffer(buf, maxCapacity)

    for scanner.Scan() {
        textFound := scanner.Text()
        fmt.Println(textFound)
        //d.generatedFile = append(d.generatedFile, textFound)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	return target
}
*/

type veniContext struct {
    d http.Dir
    generatedFile []string
    err error
}

func (d veniContext) generateHTMLHandler(w http.ResponseWriter, r *http.Request){
    
}

func Load(fs http.FileSystem) http.Handler{
    context := veniContext{fs, nil, nil}
    if context.err!=true {
        return http.FileServer(fs)
    }
    return context.generateHTMLHandler
}
