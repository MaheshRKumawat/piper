package piper

import (
	"log"
	"os"

	cos "github.com/MaheshRKumawat/COS_Connection"
)

func main() {
	var p PiedPiper
	p.GetConf("job.piedpiper.yml")

	c := cos.COS_Instance{
		ApiKey:            os.Getenv("API_KEY"),
		ServiceInstanceID: os.Getenv("RESOURCE_INSTANCE_ID"),
		AuthEndpoint:      os.Getenv("AUTH_ENDPOINT"),
		ServiceEndpoint:   os.Getenv("SERVICE_ENDPOINT"),
		BucketName:        os.Getenv("BUCKET_NAME"),
	}

	_, ob_keys, client, _ := cos.Connect(c)

	if os.Args[1] == "input" {
		for _, ip := range p.Inputs {
			if !cos.Check_keys(ob_keys, ip) {
				log.Printf("%v key not present in Cloud Object Storage bucket, hence exiting", ip)
				os.Setenv("ALL_KEYS_PRESENT", "false")
				os.Exit(0)
			}
		}
		log.Println("All input keys present")
		os.Setenv("ALL_KEYS_PRESENT", "true")

		for _, ip := range p.Inputs {
			cos.Read_file_from_cos(c, ip, client)
		}
	} else if os.Args[1] == "output" {
		for _, op := range p.Outputs {
			cos.Write_file_to_cos(c, op, client)
		}
		log.Println("All output object keys pushed")
	} else {
		log.Println("Specify Arguments in bash script while running go file")
		os.Exit(1)
	}
}
