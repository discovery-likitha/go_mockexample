package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/aws/aws-sdk-go/aws"
	"fmt"
)

const (
	bucket           string = "vqtf-qa-eu-west-1-datalake-bucket"
	prefix           string = "input/"
	delimiter        string = "/"
	startAfter       string = "input/"
)

type S3Client interface {
	List(bucket string, prefix string, delimiter string, startAfter string) (*s3.ListObjectsV2Output, error)
}

type S3ClientObj struct {
	Client s3iface.S3API
}

func (instance *S3ClientObj) List(bucket string, prefix string, delimiter string, startAfter string) (*s3.ListObjectsV2Output, error) {
	var bucketInputObject = &s3.ListObjectsV2Input{Bucket: aws.String(bucket), Delimiter: aws.String(delimiter), Prefix: aws.String(prefix), StartAfter: aws.String(startAfter),}
	response, err := instance.Client.ListObjectsV2(bucketInputObject)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	s3obj := S3ClientObj{Client: s3.New(sess)}
	//fmt.Println(s3obj)
	res, _:= s3obj.List(bucket, prefix, delimiter, startAfter)
	fmt.Println(res)
}