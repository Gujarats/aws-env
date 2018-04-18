package main

import (
	"bytes"
	"errors"
	"flag"
	"os"
	"strings"
	"syscall"
	"unicode"

	"github.com/Gujarats/logger"
)

const (
	AccessKey = "aws_access_key_id"
	SecretKey = "aws_secret_access_key"

	// Environment Variable
	AccessKeyEnv = "AWS_ACCESS_KEY_ID"
	SecretKeyEnv = "AWS_SECRET_ACCESS_KEY"
	AwsTokenEnv  = "AWS_SESSION_TOKEN"
)

var profile string

func init() {
	flag.StringVar(&profile, "profile", "default", "specify your aws profile (called using default)")
	flag.Parse()
}

func main() {
	config := getConfig()

	dataCredentials, err := OpenFile(config.AwsConfigPath)
	if err != nil {
		logger.Debug("Error :: ", err)
		os.Exit(1)
	}

	// define which profile to use from config or flag
	var useProfile string
	// override profile if using flag
	if profile != "default" {
		useProfile = profile
	} else {
		useProfile = config.Profile
	}

	awsCred := getCredentials(dataCredentials, useProfile)
	err = awsCred.exportCredentials()
	if err != nil {
		logger.Debug("Error :: ", err)
		os.Exit(1)
	}
}

type AwsCredentials struct {
	AccessKey string
	SecretKey string
	Token     string
}

func getCredentials(data []byte, profile string) *AwsCredentials {
	awsCredentials := &AwsCredentials{}
	profileIndex := bytes.Index(data, []byte("["+profile+"]"))
	if profileIndex == -1 {
		return awsCredentials
	}

	// check if found profile is valid profile
	//firstP := bytes.Index(data[profileIndex-1], []byte("["))
	//secondP := bytes.Index(data[profileIndex:], []byte("]"))

	// get the access key
	accessKeyIndex := bytes.Index(data[profileIndex:], []byte(`=`))
	enter := bytes.Index(data[profileIndex+accessKeyIndex:], []byte("\n"))

	// +1 avoid `=` added
	accesKey := data[profileIndex+accessKeyIndex+1 : profileIndex+accessKeyIndex+enter]
	profileIndex = profileIndex + accessKeyIndex + enter
	awsCredentials.AccessKey = removeSpace(string(accesKey))

	// get the secret key
	secretKeyIndex := bytes.Index(data[profileIndex:], []byte(`=`))
	enter = bytes.Index(data[profileIndex+secretKeyIndex:], []byte("\n"))
	if enter == -1 {
		enter = len(data[profileIndex+secretKeyIndex:])
	}

	// +1 avoid `=` added
	secretKey := data[profileIndex+secretKeyIndex+1 : profileIndex+secretKeyIndex+enter]
	awsCredentials.SecretKey = removeSpace(string(secretKey))

	logger.Debug("credential exported to environment variable :: ", *awsCredentials)

	return awsCredentials
}

// Exporting aws credentials to env variable
func (a *AwsCredentials) exportCredentials() error {
	if a == nil {
		return errors.New("Please check if your profile is exist in Aws credentials")
	}

	err := os.Setenv(AccessKeyEnv, a.AccessKey)
	if err != nil {
		return err
	}

	err = os.Setenv(SecretKeyEnv, a.SecretKey)
	if err != nil {
		return err
	}

	syscall.Exec(os.Getenv("SHELL"), []string{os.Getenv("SHELL")}, syscall.Environ())

	return nil
}

// Remove white space from string the fastest way
// See here https://stackoverflow.com/questions/32081808/strip-all-whitespace-from-a-string-in-golang
func removeSpace(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}
