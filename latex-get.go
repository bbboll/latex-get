package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// command types
const (
	CommandCreate   = "create"
	CommandList     = "list"
	CommandDescribe = "describe"
)

// filename choices
const (
	FilenameDescription = ".templ_description"
)

func main() {
	// provide usage information
	flag.Usage = func() {
		help()
		flag.PrintDefaults()
	}
	flag.Parse()

	// grab user input
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Please provide a command to be executed!")
		help()
		return
	}

	// execute command
	command := args[0]
	switch command {
	case CommandCreate:
		if len(args) < 3 {
			fmt.Println("Please provide template name and output directory!")
			return
		}
		create(args[1], args[2])
	case CommandList:
		list()
		return
	case CommandDescribe:
		if len(args) < 2 {
			fmt.Println("Please provide template name!")
			return
		}
		describe(args[1])
		return
	default:
		fmt.Printf("Command not found: %s\n", command)
		help()
		return
	}
}

func help() {
	fmt.Println("This tool provides access to a selection of LaTeX templates.\n\n" +
		"Usage: $ latex-get [-flags] [command] <template name> <out directory>\n\n" +
		"Available commands:\n" +
		"	create:		paste template into out directory\n" +
		"	list:		list available templates\n" +
		"	describe: 	show description for template")
}

// create a new latex file structure from
// the passed template name
func create(templ, outDir string) {
	// find template directory
	tmplDir, dErr := getTemplateDir()
	if dErr != nil {
		fmt.Println(dErr.Error())
	}

	// find template
	tmplPath, tErr := getTemplate(tmplDir, templ)
	if tErr != nil {
		fmt.Printf("Template '%s' not found.\n", templ)
	}

	// copy to out dir if possible
	if dirExists(outDir) {
		// copy template to out dir
		copyDirContents(tmplPath, outDir)
	} else {
		fmt.Printf("Output directory '%s' does not exists.", outDir)
	}
}

// list available templates
func list() {
	// find template directory
	tmplDir, err := getTemplateDir()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// list contents
	contents, err := ioutil.ReadDir(tmplDir)
	if err != nil {
		return
	}

	// print tempalte names
	println()
	for _, template := range contents {
		fmt.Printf("%s\n%s\n\n", 
			template.Name(), 
			getTemplateDescription(filepath.Join(tmplDir, template.Name())))
	}
}

// describe a passed template
func describe(templ string) {
	// find template directory
	tmplDir, err := getTemplateDir()
	if err != nil {
		fmt.Println(err.Error())
	}

	descr := getTemplateDescription(filepath.Join(tmplDir, templ))
	fmt.Printf("Template %s:\n%s", templ, descr)
}

func getTemplateDir() (string, error) {
	candidates := []string{"./templates", "../templates"}
	for _, dir := range candidates {
		if dirExists(dir) {
			return dir, nil
		}
	}
	return "", errors.New("Template directory not found.")
}

// path to template directory, error if it doesn't exists
func getTemplate(tmplDir, name string) (string, error) {
	dir := filepath.Join(tmplDir, name)
	if dirExists(dir) {
		return dir, nil
	}
	return "", errors.New("Template not found")
}

func getTemplateDescription(tmplPath string) string {
	descr, err :=  ioutil.ReadFile(filepath.Join(tmplPath, FilenameDescription))
	if err != nil {
		return "No description available."
	}
	return string(descr)
}
