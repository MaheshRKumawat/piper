package piper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/IBM/code-engine-go-sdk/codeenginev2"
	"github.com/IBM/go-sdk-core/v5/core"
)

func generateJob(filename string, projectID string, ceClient *codeenginev2.CodeEngineV2) {
	j := PiedPiper{}
	j.GetConf(filename)

	createBuildOptions := ceClient.NewCreateBuildOptions(
		projectID,
		j.Job.Name,
		j.RepositoryName,
		j.RegistryAccessSecret,
		j.Job.SourceCode,
		"dockerfile",
	)

	build, _, err := ceClient.CreateBuild(createBuildOptions)

	if err != nil {
		panic(err)
	}
	b, _ := json.MarshalIndent(build, "", "  ")
	log.Println(string(b))
}

func auth(region string) (ceClient *codeenginev2.CodeEngineV2) {
	authenticator, err := core.NewIamAuthenticatorBuilder().
		SetApiKey(os.Getenv("CE_API_KEY")).
		Build()

	if err != nil {
		panic(err)
	}

	ceClient, _ = codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
		Authenticator: authenticator,
		URL:           "https://api." + region + ".codeengine.cloud.ibm.com/v2",
	})

	return ceClient
}

func getBluemixConf() (projectID string, region string) {
	home, err := os.UserHomeDir()
	content, err := ioutil.ReadFile(filepath.Join(home, "/.bluemix/plugins/code-engine/config.json"))
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	var payload map[string]interface{}
	err = json.Unmarshal(content, &payload)
	projectID = fmt.Sprint(payload["projectID"])
	region = fmt.Sprint(payload["region"])
	if err != nil {
		log.Fatal("Error during unmarshal: ", err)
	}
	return projectID, region
}