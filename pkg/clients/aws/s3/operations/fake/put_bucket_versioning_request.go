// Code generated by mockery v1.0.0. DO NOT EDIT.

package fake

import mock "github.com/stretchr/testify/mock"

import s3 "github.com/aws/aws-sdk-go-v2/service/s3"

// PutBucketVersioningRequest is an autogenerated mock type for the PutBucketVersioningRequest type
type PutBucketVersioningRequest struct {
	mock.Mock
}

// Send provides a mock function with given fields:
func (_m *PutBucketVersioningRequest) Send() (*s3.PutBucketVersioningOutput, error) {
	ret := _m.Called()

	var r0 *s3.PutBucketVersioningOutput
	if rf, ok := ret.Get(0).(func() *s3.PutBucketVersioningOutput); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketVersioningOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
