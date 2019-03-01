// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// +build go1.10,integration

package emr_test

import (
	"context"
	"testing"
	"time"

	"github.com/Beeketing/aws-sdk-go/aws"
	"github.com/Beeketing/aws-sdk-go/aws/awserr"
	"github.com/Beeketing/aws-sdk-go/aws/request"
	"github.com/Beeketing/aws-sdk-go/awstesting/integration"
	"github.com/Beeketing/aws-sdk-go/service/emr"
)

var _ aws.Config
var _ awserr.Error
var _ request.Request

func TestInteg_00_ListClusters(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	sess := integration.SessionWithDefaultRegion("us-west-2")
	svc := emr.New(sess)
	params := &emr.ListClustersInput{}
	_, err := svc.ListClustersWithContext(ctx, params)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}
func TestInteg_01_DescribeCluster(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	sess := integration.SessionWithDefaultRegion("us-west-2")
	svc := emr.New(sess)
	params := &emr.DescribeClusterInput{
		ClusterId: aws.String("fake_cluster"),
	}
	_, err := svc.DescribeClusterWithContext(ctx, params)
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
