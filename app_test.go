package main

import (
	"log"
	"testing"

	"github.com/Gujarats/logger"
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

		// /n in the end of the line
		{
			data: []byte(`[testing-user]
aws_access_key_id = AKKKJJJY62LYY25PLLLL
aws_secret_access_key = qqUTe7gMPjjIWNSaLMMM+7ZFlILmJKUJ142gZ+ll

[s3-test-read-only]
aws_access_key_id = AKKKJOFRRA6JEHAFTYYY
aws_secret_access_key = mFF8JknsquQnCQuyb1j6IRjycu3RoFpGJdZiKnnn
`),
			profile: "s3-test-read-only",
			awsCredentials: AwsCredentials{
				AccessKey: "AKKKJOFRRA6JEHAFTYYY",
				SecretKey: "mFF8JknsquQnCQuyb1j6IRjycu3RoFpGJdZiKnnn",
			},
		},

		// without /n in the end of the line
		{
			data: []byte(`[testing-user]
aws_access_key_id = AKKKJJJY62LYY25PLLLL
aws_secret_access_key = qqUTe7gMPjjIWNSaLMMM+7ZFlILmJKUJ142gZ+ll

[s3-test-read-only]
aws_access_key_id = AKKKJOFRRA6JEHAFTYYY
aws_secret_access_key = mFF8JknsquQnCQuyb1j6IRjycu3RoFpGJdZiKnnn`),
			profile: "s3-test-read-only",
			awsCredentials: AwsCredentials{
				AccessKey: "AKKKJOFRRA6JEHAFTYYY",
				SecretKey: "mFF8JknsquQnCQuyb1j6IRjycu3RoFpGJdZiKnnn",
			},
		},

		// find the undefined profile but exist in the substring
		{
			data: []byte(`[testing-user]
aws_access_key_id = AKKKJJJY62LYY25PLLLL
aws_secret_access_key = qqUTe7gMPjjIWNSaLMMM+7ZFlILmJKUJ142gZ+ll

[s3-test-read-only]
aws_access_key_id = AKKKJOFRRA6JEHAFTYYY
aws_secret_access_key = mFF8JknsquQnCQuyb1j6IRjycu3RoFpGJdZiKnnn`),
			profile:        "testing",
			awsCredentials: AwsCredentials{},
		},
	}

	for index, testObject := range testObjects {
		result := getCredentials(testObject.data, testObject.profile)
		logger.Debug("Test Index :: ", index)

		if result.AccessKey != testObject.awsCredentials.AccessKey {
			log.Fatalf("expected %+v, result = %+v\n", testObject.awsCredentials, result)
		}

		if result.SecretKey != testObject.awsCredentials.SecretKey {
			log.Fatalf("expected %+v, result = %+v\n", testObject.awsCredentials, result)
		}
	}
}
