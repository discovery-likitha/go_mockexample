package main

import (
	"testing"
	"fmt"
	//"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/stretchr/testify/assert"
)

type MockedS3ListObj struct {
	s3iface.S3API
	Resp s3.ListObjectsV2Output
}
func (m MockedS3ListObj) ListObjectsV2(in *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	// Only need to return mocked response output
	return &m.Resp, nil
}
func TestS3ListOjectsMock(t *testing.T) {
	mockSvc := &MockedS3ListObj{}
	testcient := S3Client{ Client :mockSvc}
	result,err := testcient.List(bucket, prefix, delimiter, startAfter)
	fmt.Println(err)
	fmt.Println(result)
	assert.NoError(t, err)
}