package piper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/IBM/code-engine-go-sdk/codeenginev2"
	"github.com/IBM/go-sdk-core/core"
)

func generateJob(filename string, projectID string, ceClient *codeenginev2.NewCodeEngineV2) {
	j := PiedPiper{}
	j.GetConf(filename)

	var codeEngineService *codeenginev2.CodeEngineV2

	// projectID, region := getBluemixConf()
	// ceClient := auth(region)

	createBuildOptions := ceClient.NewCreateBuildOptions(
		projectID,
		j.Job.Name,
		j.RepositoryName,
		j.RegistryAccessSecret,
		j.Job.SourceCode,
		"dockerfile",
	)

	build, _, err := codeEngineService.CreateBuild(createBuildOptions)

	if err != nil {
		panic(err)
	}
	b, _ := json.MarshalIndent(build, "", "  ")
	fmt.Println(string(b))
}

func auth(region string) (ceClient *codeenginev2.NewCodeEngineV2) {
	// Initialize the IAM authenticator using an API key
	authenticator := &core.IamAuthenticator{
		ApiKey:       os.Getenv("CE_API_KEY"),
		ClientId:     "bx",
		ClientSecret: "bx",
		URL:          "https://iam.cloud.ibm.com",
	}

	ceClient, _ = codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
		Authenticator: authenticator,
		URL:           "https://api." + region + ".codeengine.cloud.ibm.com/v2",
	})

	return ceClient
}

func getBluemixConf() (projectID string, region string) {
	home, err := os.UserHomeDir()
	content, err := ioutil.ReadFile(home + "/.bluemix/plugins/code-engine/config.json")
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
