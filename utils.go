package piper

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type InstanceResource struct {
	Vcpu             float64 `yaml:"vcpu"`
	Memory           int     `yaml:"memory"`
	EphimeralStorage float64 `yaml:"ephimeral_storage"`
}

type Runtime struct {
	InstanceResource InstanceResource `yaml:"instance_resource"`
	Mode             string           `yaml:"mode"`
	Retries          int              `yaml:"retries"`
	Timeout          int              `yaml:"timeout"`
}

type Job struct {
	Name       string `yaml:"name"`
	SourceCode string `yaml:"source_code"`
	Branch     string `yaml:"branch"`
	Folder     string `yaml:"folder"`
}

type PiedPiper struct {
	Version              string   `yaml:"version"`
	Job                  Job      `yaml:"job"`
	Inputs               []string `yaml:"inputs"`
	Outputs              []string `yaml:"outputs"`
	RegistryServer       string   `yaml:"registry_server"`
	RegistryAccessSecret string   `yaml:"registry_access_secret"`
	Namespace            string   `yaml:"namespace"`
	RepositoryName       string   `yaml:"repository_name"`
	Tag                  string   `yaml:"tag"`
	Runtime              Runtime  `yaml:"runtime"`
}

type Project_PiedPiper struct {
	Version           string   `yaml:"version"`
	CodeEngineProject string   `yaml:"code_engine_project"`
	GithubURL         string   `yaml:"github_url"`
	Folders           []string `yaml:"folders"`
	Secrets           string   `yaml:"secrets"`
	Configmaps        string   `yaml:"configmaps"`
	Namespace         string   `yaml:"namespace"`
	DockerhubAccess   string   `yaml:"dockerhub_access"`
	CosBucketName     string   `yaml:"cos_bucket_name"`
}

func (p *PiedPiper) GetConf(filename string) *PiedPiper {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, p)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return p
}

func (p *PiedPiper) DefaultConf() *PiedPiper {
	p.Version = "v1.0"
	p.Job = *p.Job.DefaultConf()
	p.Inputs = append(p.Inputs, "")
	p.Outputs = append(p.Outputs, "")
	p.RegistryServer = "https://index.docker.io/v1/"
	p.RegistryAccessSecret = ""
	p.Namespace = ""
	p.RepositoryName = ""
	p.Tag = ""
	p.Runtime = *p.Runtime.DefaultConf()
	return p
}

func (i *InstanceResource) DefaultConf() *InstanceResource {
	i.Vcpu = 1
	i.Memory = 4
	i.EphimeralStorage = 0.4
	return i
}

func (r *Runtime) DefaultConf() *Runtime {
	r.InstanceResource = *r.InstanceResource.DefaultConf()
	r.Mode = "Task"
	r.Retries = 3
	r.Timeout = 7200
	return r
}

func (j *Job) DefaultConf() *Job {
	j.Name = ""
	j.SourceCode = ""
	j.Branch = "main"
	j.Folder = ""
	return j
}

func (project *Project_PiedPiper) GetConf(filename string) *Project_PiedPiper {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, project)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return project
}
