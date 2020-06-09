package mocks

import (
	s3 "github.com/aws/aws-sdk-go/service/s3"
	mock "github.com/stretchr/testify/mock"
)

type S3Client struct {
	mock.Mock
}

// ListObjectsV2 provides a mock function with given fields: _a0
func (_m *S3Client) List(bucket string, prefix string, delimiter string, startAfter string) (*s3.ListObjectsV2Output, error) {
	ret := _m.Called(bucket, prefix, delimiter, startAfter)
	var r0 *s3.ListObjectsV2Output
	if rf, ok := ret.Get(0).(func(string,string,string,string) *s3.ListObjectsV2Output); ok {
		r0 = rf(bucket, prefix, delimiter, startAfter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.ListObjectsV2Output)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string,string,string,string) error); ok {
		r1 = rf(bucket, prefix, delimiter, startAfter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
