# latex-get
Quick and dirty tool to start LaTeX projects from a selection of templates.

## Installation
Just clone this repository to your local machine. If you have a working Go environment, you can now `go run *.go`.
Since this is a fairly early stage of development, no binary is provided with the source yet.

## Usage 
-	List available templates
	`go run *.go list`
-	Show description for a template (science_thesis)
	`go run *.go describe science_thesis`
-	Copy an instance of the selected template (science_thesis) into a destination path
	`go run *.go create science_thesis /path/to/destination/`
