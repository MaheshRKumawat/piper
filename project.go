package piper

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func GenerateJobFile() {
	p := PiedPiper{}
	project := Project_PiedPiper{}
	project.GetConf("project.piedpiper.yml")
	p.DefaultConf()

	folders := project.Folders

	for _, folder := range folders {
		p.Job.Name = folder
		p.Job.SourceCode = project.GithubURL
		p.Job.Folder = folder
		p.RegistryAccessSecret = project.DockerhubAccess
		p.RepositoryName = project.CodeEngineProject + "-image"
		p.Namespace = project.Namespace
		p.Tag = folder
		yamlData, _ := yaml.Marshal(&p)
		fileName := "job.piedpiper.yaml"
		file := filepath.Join(folder, filepath.Base(fileName))
		err := ioutil.WriteFile(file, yamlData, 0644)
		if err != nil {
			panic("Unable to write data into the file")
		}
	}
}

func build(){
}

func run() {
	if os.Args[1] == "generate" {
		GenerateJobFile()
	} else if os.Args[1] == "build" {
		build()
	}
}
