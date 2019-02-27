// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// +build go1.10,integration

package cloudtrail_test

import (
	"context"
	"testing"
	"time"

	"github.com/Beeketing/aws-sdk-go/aws"
	"github.com/Beeketing/aws-sdk-go/aws/awserr"
	"github.com/Beeketing/aws-sdk-go/aws/request"
	"github.com/Beeketing/aws-sdk-go/awstesting/integration"
	"github.com/Beeketing/aws-sdk-go/service/cloudtrail"
)

var _ aws.Config
var _ awserr.Error
var _ request.Request

func TestInteg_00_DescribeTrails(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	sess := integration.SessionWithDefaultRegion("us-west-2")
	svc := cloudtrail.New(sess)
	params := &cloudtrail.DescribeTrailsInput{}
	_, err := svc.DescribeTrailsWithContext(ctx, params)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}
func TestInteg_01_DeleteTrail(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	sess := integration.SessionWithDefaultRegion("us-west-2")
	svc := cloudtrail.New(sess)
	params := &cloudtrail.DeleteTrailInput{
		Name: aws.String("faketrail"),
	}
	_, err := svc.DeleteTrailWithContext(ctx, params)
	if err == nil {
		t.Fatalf("expect request to fail")
	}
	aerr, ok := err.(awserr.RequestFailure)
	if !ok {
		t.Fatalf("expect awserr, was %T", err)
	}
	if v := aerr.Code(); v == request.ErrCodeSerialization {
		t.Errorf("expect API error code got serialization failure")
	}
}
