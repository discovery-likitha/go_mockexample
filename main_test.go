package main

import (
	//"fmt"
	"testing"
	//"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	//"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/discovery-likitha/go_mockexample/mocks"
	"github.com/stretchr/testify/assert"
)

// type S3Test struct {
// 	s3Client S3Client
// }

func TestS3ListOjectsMock(t *testing.T) {
	var (
		validbucket     = "vqtf-qa-eu-west-1-datalake-bucket"
		validprefix     = "input/"
		validdelimiter  = "/"
		validstartAfter = "input/"
	)
	//s3test = &S3Test{}
	mock := &mocks.S3Client{}
	//s3test.s3Client = S3ClientObj{Client: mock}
	mockedInputObject := &s3.ListObjectsV2Input{
		Bucket:     &validbucket,
		Delimiter:  &validdelimiter,
		Prefix:     &validprefix,
		StartAfter: &validstartAfter}
	mock.On("ListObjectsV2", mockedInputObject).Return(&s3.ListObjectsV2Output{}, nil)
	_, err := mock.List(&validbucket, &validprefix, &validdelimiter, &validstartAfter)
	assert.NoError(t, err)
}
