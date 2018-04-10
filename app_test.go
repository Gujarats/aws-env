package main

import (
	"log"
	"reflect"
	"testing"
)

func TestExportCredentials(t *testing.T) {
	testObjects := []struct {
		data           []byte
		profile        string
		awsCredentials AwsCredentials
	}{
		{
			data: []byte(`[testing-user]
aws_access_key_id = AKKKJJJY62LYY25PLLLL
aws_secret_access_key = qqUTe7gMPjjIWNSaLMMM+7ZFlILmJKUJ142gZ+ll

[s3-test-read-only]
aws_access_key_id = AKKKJOFRRA6JEHAFTYYY
aws_secret_access_key = mFF8JknsquQnCQuyb1j6IRjycu3RoFpGJdZiKnnn`),
			profile: "testing-user",
			awsCredentials: AwsCredentials{
				AccessKey: "AKKKJJJY62LYY25PLLLL",
				SecretKey: "qqUTe7gMPjjIWNSaLMMM+7ZFlILmJKUJ142gZ+ll",
			},
		},
	}

	for _, testObject := range testObjects {
		result := getCredentials(testObject.data, testObject.profile)
		if reflect.DeepEqual(result, testObject.awsCredentials) {
			log.Fatalf("expected %+v, result = %+v\n", testObject.awsCredentials, result)
		}
	}
}
