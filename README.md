# latex-get
Quick and dirty tool to start LaTeX projects from a selection of templates.

## Installation
Just clone this repository to your local machine. Executable binaries, suffixed with system and compile date, are provided within the `bin/` directory.

## Usage 
*	List available templates:
	`go run *.go list`
*	Show description for a template (science_thesis):
	`go run *.go describe science_thesis`
*	Copy an instance of the selected template (science_thesis) into a destination path:
	`go run *.go create science_thesis /path/to/destination/`

## Roadmap
I might add the possibility to set a config file in the future. That way, most placeholder content within the templates can be filled with real data.