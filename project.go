package piper

import (
	"io/ioutil"
	"log"
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
		fileName := "job.piedpiper.yml"
		file := filepath.Join(folder, filepath.Base(fileName))
		err := ioutil.WriteFile(file, yamlData, 0644)
		if err != nil {
			panic("Unable to write job.piedpiper.yml into the folder")
		}

		dockerfile_content := []byte(
			`FROM golang:1.20
		RUN apt-get update && \
			apt-get upgrade -y && \
			apt-get install -y git
		RUN git clone https://github.com/MaheshRKumawat/piper
		COPY ./bash.sh .
		RUN chmod +x bash.sh
		CMD ["./bash.sh"]`)

		dockerfile := filepath.Join(folder, filepath.Base("Dockerfile"))
		err = ioutil.WriteFile(dockerfile, dockerfile_content, 0644)
		if err != nil {
			panic("Unable to write dockerfile into the folder")
		}

		bash_content := []byte(
		`#!/bin/bash

		cd 
		go build -o /cos ./piper/cos.go

		/cos input	
		
		if [[ $ALL_KEYS_PRESENT == "true"]]; then
			# Your commands to execute the file
		else
			echo "Exit from Bash"
			return
		fi

		/cos output`)

		bash := filepath.Join(folder, filepath.Base("bash.sh"))
		err = ioutil.WriteFile(bash, bash_content, 0644)
		if err != nil {
			panic("Unable to write bash file into the folder")
		}
	}
}

func build() {
	project := Project_PiedPiper{}
	project.GetConf("project.piedpiper.yml")
	folders := project.Folders

	projectID, region := getBluemixConf()
	ceClient := auth(region)

	for _, folder := range folders {
		fileName := "job.piedpiper.yml"
		file := filepath.Join(folder, filepath.Base(fileName))
		generateJob(file, projectID, ceClient)
	}

	log.Println("Build Job Successful")
}

func run() {
	if os.Args[1] == "generate" {
		GenerateJobFile()
	} else if os.Args[1] == "build" {
		build()
	}
}
