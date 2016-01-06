//Package yamlToJson provides a simple interface to generate swagger html structure from an yaml file.
package main

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/NexwayGroup/nx-lib/tools"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/ghodss/yaml"
)

var (
	pathFolder     string
	force          bool
	recursive      bool
	output         string
	indexPath      string
	templateString string
	active         bool
	selectString   string
)

// main set up the commands line and the flags
func main() {
	app := cli.NewApp()
	app.Name = "go-init"
	app.Usage = "Initialize a go project"
	app.Version = "1.1.0"
	// global flags
	//Log Level and Config Path
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:   "force, f",
			EnvVar: "FORCE",
			Usage:  "force the creation of the folder",
		},
		cli.BoolFlag{
			Name:   "recursive, r",
			EnvVar: "RECURSIVE",
			Usage:  "Make a recursive search for the yaml files",
		},
		cli.StringFlag{
			Name:   "index, i",
			EnvVar: "INDEX",
			Usage:  "Path to the index.html file",
		},
	}
	app.Action = func(c *cli.Context) {
		if len(c.Args()) == 0 {
			return
		}
		if len(c.Args()) == 1 {
			log.Print("Please specify the ouput folder")
			return
		}
		if len(c.Args()) == 2 {
			output = c.Args()[1]
			pathFolder = c.Args()[0]

		} else {
			log.Print("too many argument")
			return
		}
		Convert(c)
	}
	app.Before = func(ctx *cli.Context) error {
		force = ctx.Bool("force")
		recursive = ctx.Bool("recursive")
		indexPath = ctx.String("index")
		return nil
	}
	app.Run(os.Args)
}

func Convert(ctx *cli.Context) {
	selectString += `<FORM>` + "\n" + `<SELECT class="form-control" >` + "\n" + `<OPTION>` + "\n"
	os.RemoveAll(output)
	findFile(pathFolder, 2)
	os.MkdirAll(output, 0755)
	CreateIndex()
	CreateTemplate()
}

func CreateTemplate() {
	templateFile, err := os.Create(path.Join(output, "template.html"))
	if err != nil {
		log.Fatal("error while creating the template file ", err)
	}
	defer templateFile.Close()
	selectString += `</SELECT>` + "\n" + `</FORM>` + "\n"
	templateFile.WriteString(selectString)
	templateFile.WriteString(templateString)
}

// copyFile copy a file from a destination to a source
func copyFile(destination, source string) error {
	file, err := os.Open(source)
	defer file.Close()
	if err != nil {
		return err
	}
	destinationFile, err := os.Create(destination)
	defer destinationFile.Close()
	if err != nil {
		log.Print(err)
		return err
	}
	if _, err := io.Copy(destinationFile, file); err != nil {
		destinationFile.Close()
		return err
	}
	return nil
}

func CreateIndex() {
	if indexPath != "" {
		copyFile(path.Join(output, "index.html"), path.Join(indexPath, "index.html"))
		copyFile(path.Join(output, "nav.css"), path.Join(indexPath, "nav.css"))
		copyFile(path.Join(output, "main.css"), path.Join(indexPath, "main.css"))
	}
}

func findFile(pathFile string, size int) {
	if strings.Contains(pathFile, ".yaml") {
		outputFolderName := strings.Replace(pathFile, ".yaml", "", -1)
		ConvertFile(pathFile, path.Join(output, outputFolderName), outputFolderName)
		return
	}
	files, _ := ioutil.ReadDir(pathFile)
	for _, f := range files {
		var parent string
		if f.IsDir() && recursive {
			if size == 2 {
				selectString += "<OPTION>" + f.Name() + "\n"
				templateString += `<div class="hide" id="` + f.Name() + `">`
			} else {
				parent = "collapse" + f.Name()
				templateString += `<button class="btn btn-secondary btn-block" data-toggle="collapse" data-target="#` + parent + `" aria-expanded="false" aria-controls="` + parent + `"><i class="glyphicon fa fa-android pull-left"></i>` + f.Name() + `<span class="glyphicon glyphicon-menu-down  pull-right" aria-hidden="true" ></span></button>` + "\n"
				templateString += `<div class="collapse" id="` + parent + `">`
			}
			size++
			findFile(path.Join(pathFile, f.Name()), size)
			size--
			templateString += `</div>` + "\n"
		}
		if strings.Contains(f.Name(), ".yaml") {
			outputFolderName := strings.Replace(f.Name(), ".yaml", "", -1)
			ConvertFile(path.Join(pathFile, f.Name()), path.Join(output, pathFile, outputFolderName), outputFolderName)
		}
	}
}

func ConvertFile(filePath, outputFolder, fileName string) {
	source, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	y, err := yaml.YAMLToJSON(source)
	if err != nil {
		log.Fatal("An error occur while decoding your yaml")
	}
	nameFile := strings.Replace(filePath, ".yaml", ".json", 1)
	exist, err := tools.Exists(nameFile)
	if err != nil {
		log.Fatal("an error occure while trying to find the file: ", nameFile)
		return
	}
	if exist && !force {
		log.Print("The file :", nameFile, "already exist use -f to force the creation of the Json")
		return
	}
	file, err := os.Create(nameFile)
	defer file.Close()
	if err != nil {
		log.Error("error while creating ", nameFile, " err:", err.Error())
		return
	}
	_, err = file.Write(y)
	if err != nil {
		log.Debug("Error while writting file")
		return
	}
	exec.Command("bootprint", "swagger", filePath, outputFolder).Output()
	os.Remove(nameFile)
	indexName := path.Join(outputFolder, "index.html")
	newName := path.Join(outputFolder, fileName+".html")
	htmlPath := path.Join(strings.Replace(filePath, ".yaml", "", 1), fileName+".html")
	os.Rename(indexName, newName)
	cssName := path.Join(outputFolder, "main.css")
	os.Remove(cssName)
	cssMapName := path.Join(outputFolder, "main.css.map")
	os.Remove(cssMapName)
	if !active {
		templateString += `<li class="active">` + "\n" + `<a onclick="Load('` + htmlPath + `')">` + fileName + `</a>` + "\n" + `</li>` + "\n"
		active = true
		return
	}
	templateString += `<li>` + "\n" + `<a onclick="Load('` + htmlPath + `')">` + fileName + `</a>` + "\n" + `</li>` + "\n"
}
