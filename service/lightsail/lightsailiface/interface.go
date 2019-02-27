// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package lightsailiface provides an interface to enable mocking the Amazon Lightsail service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package lightsailiface

import (
	"github.com/Beeketing/aws-sdk-go/aws"
	"github.com/Beeketing/aws-sdk-go/aws/request"
	"github.com/Beeketing/aws-sdk-go/service/lightsail"
)

// LightsailAPI provides an interface to enable mocking the
// lightsail.Lightsail service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Lightsail.
//    func myFunc(svc lightsailiface.LightsailAPI) bool {
//        // Make svc.AllocateStaticIp request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := lightsail.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockLightsailClient struct {
//        lightsailiface.LightsailAPI
//    }
//    func (m *mockLightsailClient) AllocateStaticIp(input *lightsail.AllocateStaticIpInput) (*lightsail.AllocateStaticIpOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockLightsailClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type LightsailAPI interface {
	AllocateStaticIp(*lightsail.AllocateStaticIpInput) (*lightsail.AllocateStaticIpOutput, error)
	AllocateStaticIpWithContext(aws.Context, *lightsail.AllocateStaticIpInput, ...request.Option) (*lightsail.AllocateStaticIpOutput, error)
	AllocateStaticIpRequest(*lightsail.AllocateStaticIpInput) (*request.Request, *lightsail.AllocateStaticIpOutput)

	AttachDisk(*lightsail.AttachDiskInput) (*lightsail.AttachDiskOutput, error)
	AttachDiskWithContext(aws.Context, *lightsail.AttachDiskInput, ...request.Option) (*lightsail.AttachDiskOutput, error)
	AttachDiskRequest(*lightsail.AttachDiskInput) (*request.Request, *lightsail.AttachDiskOutput)

	AttachInstancesToLoadBalancer(*lightsail.AttachInstancesToLoadBalancerInput) (*lightsail.AttachInstancesToLoadBalancerOutput, error)
	AttachInstancesToLoadBalancerWithContext(aws.Context, *lightsail.AttachInstancesToLoadBalancerInput, ...request.Option) (*lightsail.AttachInstancesToLoadBalancerOutput, error)
	AttachInstancesToLoadBalancerRequest(*lightsail.AttachInstancesToLoadBalancerInput) (*request.Request, *lightsail.AttachInstancesToLoadBalancerOutput)

	AttachLoadBalancerTlsCertificate(*lightsail.AttachLoadBalancerTlsCertificateInput) (*lightsail.AttachLoadBalancerTlsCertificateOutput, error)
	AttachLoadBalancerTlsCertificateWithContext(aws.Context, *lightsail.AttachLoadBalancerTlsCertificateInput, ...request.Option) (*lightsail.AttachLoadBalancerTlsCertificateOutput, error)
	AttachLoadBalancerTlsCertificateRequest(*lightsail.AttachLoadBalancerTlsCertificateInput) (*request.Request, *lightsail.AttachLoadBalancerTlsCertificateOutput)

	AttachStaticIp(*lightsail.AttachStaticIpInput) (*lightsail.AttachStaticIpOutput, error)
	AttachStaticIpWithContext(aws.Context, *lightsail.AttachStaticIpInput, ...request.Option) (*lightsail.AttachStaticIpOutput, error)
	AttachStaticIpRequest(*lightsail.AttachStaticIpInput) (*request.Request, *lightsail.AttachStaticIpOutput)

	CloseInstancePublicPorts(*lightsail.CloseInstancePublicPortsInput) (*lightsail.CloseInstancePublicPortsOutput, error)
	CloseInstancePublicPortsWithContext(aws.Context, *lightsail.CloseInstancePublicPortsInput, ...request.Option) (*lightsail.CloseInstancePublicPortsOutput, error)
	CloseInstancePublicPortsRequest(*lightsail.CloseInstancePublicPortsInput) (*request.Request, *lightsail.CloseInstancePublicPortsOutput)

	CopySnapshot(*lightsail.CopySnapshotInput) (*lightsail.CopySnapshotOutput, error)
	CopySnapshotWithContext(aws.Context, *lightsail.CopySnapshotInput, ...request.Option) (*lightsail.CopySnapshotOutput, error)
	CopySnapshotRequest(*lightsail.CopySnapshotInput) (*request.Request, *lightsail.CopySnapshotOutput)

	CreateCloudFormationStack(*lightsail.CreateCloudFormationStackInput) (*lightsail.CreateCloudFormationStackOutput, error)
	CreateCloudFormationStackWithContext(aws.Context, *lightsail.CreateCloudFormationStackInput, ...request.Option) (*lightsail.CreateCloudFormationStackOutput, error)
	CreateCloudFormationStackRequest(*lightsail.CreateCloudFormationStackInput) (*request.Request, *lightsail.CreateCloudFormationStackOutput)

	CreateDisk(*lightsail.CreateDiskInput) (*lightsail.CreateDiskOutput, error)
	CreateDiskWithContext(aws.Context, *lightsail.CreateDiskInput, ...request.Option) (*lightsail.CreateDiskOutput, error)
	CreateDiskRequest(*lightsail.CreateDiskInput) (*request.Request, *lightsail.CreateDiskOutput)

	CreateDiskFromSnapshot(*lightsail.CreateDiskFromSnapshotInput) (*lightsail.CreateDiskFromSnapshotOutput, error)
	CreateDiskFromSnapshotWithContext(aws.Context, *lightsail.CreateDiskFromSnapshotInput, ...request.Option) (*lightsail.CreateDiskFromSnapshotOutput, error)
	CreateDiskFromSnapshotRequest(*lightsail.CreateDiskFromSnapshotInput) (*request.Request, *lightsail.CreateDiskFromSnapshotOutput)

	CreateDiskSnapshot(*lightsail.CreateDiskSnapshotInput) (*lightsail.CreateDiskSnapshotOutput, error)
	CreateDiskSnapshotWithContext(aws.Context, *lightsail.CreateDiskSnapshotInput, ...request.Option) (*lightsail.CreateDiskSnapshotOutput, error)
	CreateDiskSnapshotRequest(*lightsail.CreateDiskSnapshotInput) (*request.Request, *lightsail.CreateDiskSnapshotOutput)

	CreateDomain(*lightsail.CreateDomainInput) (*lightsail.CreateDomainOutput, error)
	CreateDomainWithContext(aws.Context, *lightsail.CreateDomainInput, ...request.Option) (*lightsail.CreateDomainOutput, error)
	CreateDomainRequest(*lightsail.CreateDomainInput) (*request.Request, *lightsail.CreateDomainOutput)

	CreateDomainEntry(*lightsail.CreateDomainEntryInput) (*lightsail.CreateDomainEntryOutput, error)
	CreateDomainEntryWithContext(aws.Context, *lightsail.CreateDomainEntryInput, ...request.Option) (*lightsail.CreateDomainEntryOutput, error)
	CreateDomainEntryRequest(*lightsail.CreateDomainEntryInput) (*request.Request, *lightsail.CreateDomainEntryOutput)

	CreateInstanceSnapshot(*lightsail.CreateInstanceSnapshotInput) (*lightsail.CreateInstanceSnapshotOutput, error)
	CreateInstanceSnapshotWithContext(aws.Context, *lightsail.CreateInstanceSnapshotInput, ...request.Option) (*lightsail.CreateInstanceSnapshotOutput, error)
	CreateInstanceSnapshotRequest(*lightsail.CreateInstanceSnapshotInput) (*request.Request, *lightsail.CreateInstanceSnapshotOutput)

	CreateInstances(*lightsail.CreateInstancesInput) (*lightsail.CreateInstancesOutput, error)
	CreateInstancesWithContext(aws.Context, *lightsail.CreateInstancesInput, ...request.Option) (*lightsail.CreateInstancesOutput, error)
	CreateInstancesRequest(*lightsail.CreateInstancesInput) (*request.Request, *lightsail.CreateInstancesOutput)

	CreateInstancesFromSnapshot(*lightsail.CreateInstancesFromSnapshotInput) (*lightsail.CreateInstancesFromSnapshotOutput, error)
	CreateInstancesFromSnapshotWithContext(aws.Context, *lightsail.CreateInstancesFromSnapshotInput, ...request.Option) (*lightsail.CreateInstancesFromSnapshotOutput, error)
	CreateInstancesFromSnapshotRequest(*lightsail.CreateInstancesFromSnapshotInput) (*request.Request, *lightsail.CreateInstancesFromSnapshotOutput)

	CreateKeyPair(*lightsail.CreateKeyPairInput) (*lightsail.CreateKeyPairOutput, error)
	CreateKeyPairWithContext(aws.Context, *lightsail.CreateKeyPairInput, ...request.Option) (*lightsail.CreateKeyPairOutput, error)
	CreateKeyPairRequest(*lightsail.CreateKeyPairInput) (*request.Request, *lightsail.CreateKeyPairOutput)

	CreateLoadBalancer(*lightsail.CreateLoadBalancerInput) (*lightsail.CreateLoadBalancerOutput, error)
	CreateLoadBalancerWithContext(aws.Context, *lightsail.CreateLoadBalancerInput, ...request.Option) (*lightsail.CreateLoadBalancerOutput, error)
	CreateLoadBalancerRequest(*lightsail.CreateLoadBalancerInput) (*request.Request, *lightsail.CreateLoadBalancerOutput)

	CreateLoadBalancerTlsCertificate(*lightsail.CreateLoadBalancerTlsCertificateInput) (*lightsail.CreateLoadBalancerTlsCertificateOutput, error)
	CreateLoadBalancerTlsCertificateWithContext(aws.Context, *lightsail.CreateLoadBalancerTlsCertificateInput, ...request.Option) (*lightsail.CreateLoadBalancerTlsCertificateOutput, error)
	CreateLoadBalancerTlsCertificateRequest(*lightsail.CreateLoadBalancerTlsCertificateInput) (*request.Request, *lightsail.CreateLoadBalancerTlsCertificateOutput)

	CreateRelationalDatabase(*lightsail.CreateRelationalDatabaseInput) (*lightsail.CreateRelationalDatabaseOutput, error)
	CreateRelationalDatabaseWithContext(aws.Context, *lightsail.CreateRelationalDatabaseInput, ...request.Option) (*lightsail.CreateRelationalDatabaseOutput, error)
	CreateRelationalDatabaseRequest(*lightsail.CreateRelationalDatabaseInput) (*request.Request, *lightsail.CreateRelationalDatabaseOutput)

	CreateRelationalDatabaseFromSnapshot(*lightsail.CreateRelationalDatabaseFromSnapshotInput) (*lightsail.CreateRelationalDatabaseFromSnapshotOutput, error)
	CreateRelationalDatabaseFromSnapshotWithContext(aws.Context, *lightsail.CreateRelationalDatabaseFromSnapshotInput, ...request.Option) (*lightsail.CreateRelationalDatabaseFromSnapshotOutput, error)
	CreateRelationalDatabaseFromSnapshotRequest(*lightsail.CreateRelationalDatabaseFromSnapshotInput) (*request.Request, *lightsail.CreateRelationalDatabaseFromSnapshotOutput)

	CreateRelationalDatabaseSnapshot(*lightsail.CreateRelationalDatabaseSnapshotInput) (*lightsail.CreateRelationalDatabaseSnapshotOutput, error)
	CreateRelationalDatabaseSnapshotWithContext(aws.Context, *lightsail.CreateRelationalDatabaseSnapshotInput, ...request.Option) (*lightsail.CreateRelationalDatabaseSnapshotOutput, error)
	CreateRelationalDatabaseSnapshotRequest(*lightsail.CreateRelationalDatabaseSnapshotInput) (*request.Request, *lightsail.CreateRelationalDatabaseSnapshotOutput)

	DeleteDisk(*lightsail.DeleteDiskInput) (*lightsail.DeleteDiskOutput, error)
	DeleteDiskWithContext(aws.Context, *lightsail.DeleteDiskInput, ...request.Option) (*lightsail.DeleteDiskOutput, error)
	DeleteDiskRequest(*lightsail.DeleteDiskInput) (*request.Request, *lightsail.DeleteDiskOutput)

	DeleteDiskSnapshot(*lightsail.DeleteDiskSnapshotInput) (*lightsail.DeleteDiskSnapshotOutput, error)
	DeleteDiskSnapshotWithContext(aws.Context, *lightsail.DeleteDiskSnapshotInput, ...request.Option) (*lightsail.DeleteDiskSnapshotOutput, error)
	DeleteDiskSnapshotRequest(*lightsail.DeleteDiskSnapshotInput) (*request.Request, *lightsail.DeleteDiskSnapshotOutput)

	DeleteDomain(*lightsail.DeleteDomainInput) (*lightsail.DeleteDomainOutput, error)
	DeleteDomainWithContext(aws.Context, *lightsail.DeleteDomainInput, ...request.Option) (*lightsail.DeleteDomainOutput, error)
	DeleteDomainRequest(*lightsail.DeleteDomainInput) (*request.Request, *lightsail.DeleteDomainOutput)

	DeleteDomainEntry(*lightsail.DeleteDomainEntryInput) (*lightsail.DeleteDomainEntryOutput, error)
	DeleteDomainEntryWithContext(aws.Context, *lightsail.DeleteDomainEntryInput, ...request.Option) (*lightsail.DeleteDomainEntryOutput, error)
	DeleteDomainEntryRequest(*lightsail.DeleteDomainEntryInput) (*request.Request, *lightsail.DeleteDomainEntryOutput)

	DeleteInstance(*lightsail.DeleteInstanceInput) (*lightsail.DeleteInstanceOutput, error)
	DeleteInstanceWithContext(aws.Context, *lightsail.DeleteInstanceInput, ...request.Option) (*lightsail.DeleteInstanceOutput, error)
	DeleteInstanceRequest(*lightsail.DeleteInstanceInput) (*request.Request, *lightsail.DeleteInstanceOutput)

	DeleteInstanceSnapshot(*lightsail.DeleteInstanceSnapshotInput) (*lightsail.DeleteInstanceSnapshotOutput, error)
	DeleteInstanceSnapshotWithContext(aws.Context, *lightsail.DeleteInstanceSnapshotInput, ...request.Option) (*lightsail.DeleteInstanceSnapshotOutput, error)
	DeleteInstanceSnapshotRequest(*lightsail.DeleteInstanceSnapshotInput) (*request.Request, *lightsail.DeleteInstanceSnapshotOutput)

	DeleteKeyPair(*lightsail.DeleteKeyPairInput) (*lightsail.DeleteKeyPairOutput, error)
	DeleteKeyPairWithContext(aws.Context, *lightsail.DeleteKeyPairInput, ...request.Option) (*lightsail.DeleteKeyPairOutput, error)
	DeleteKeyPairRequest(*lightsail.DeleteKeyPairInput) (*request.Request, *lightsail.DeleteKeyPairOutput)

	DeleteLoadBalancer(*lightsail.DeleteLoadBalancerInput) (*lightsail.DeleteLoadBalancerOutput, error)
	DeleteLoadBalancerWithContext(aws.Context, *lightsail.DeleteLoadBalancerInput, ...request.Option) (*lightsail.DeleteLoadBalancerOutput, error)
	DeleteLoadBalancerRequest(*lightsail.DeleteLoadBalancerInput) (*request.Request, *lightsail.DeleteLoadBalancerOutput)

	DeleteLoadBalancerTlsCertificate(*lightsail.DeleteLoadBalancerTlsCertificateInput) (*lightsail.DeleteLoadBalancerTlsCertificateOutput, error)
	DeleteLoadBalancerTlsCertificateWithContext(aws.Context, *lightsail.DeleteLoadBalancerTlsCertificateInput, ...request.Option) (*lightsail.DeleteLoadBalancerTlsCertificateOutput, error)
	DeleteLoadBalancerTlsCertificateRequest(*lightsail.DeleteLoadBalancerTlsCertificateInput) (*request.Request, *lightsail.DeleteLoadBalancerTlsCertificateOutput)

	DeleteRelationalDatabase(*lightsail.DeleteRelationalDatabaseInput) (*lightsail.DeleteRelationalDatabaseOutput, error)
	DeleteRelationalDatabaseWithContext(aws.Context, *lightsail.DeleteRelationalDatabaseInput, ...request.Option) (*lightsail.DeleteRelationalDatabaseOutput, error)
	DeleteRelationalDatabaseRequest(*lightsail.DeleteRelationalDatabaseInput) (*request.Request, *lightsail.DeleteRelationalDatabaseOutput)

	DeleteRelationalDatabaseSnapshot(*lightsail.DeleteRelationalDatabaseSnapshotInput) (*lightsail.DeleteRelationalDatabaseSnapshotOutput, error)
	DeleteRelationalDatabaseSnapshotWithContext(aws.Context, *lightsail.DeleteRelationalDatabaseSnapshotInput, ...request.Option) (*lightsail.DeleteRelationalDatabaseSnapshotOutput, error)
	DeleteRelationalDatabaseSnapshotRequest(*lightsail.DeleteRelationalDatabaseSnapshotInput) (*request.Request, *lightsail.DeleteRelationalDatabaseSnapshotOutput)

	DetachDisk(*lightsail.DetachDiskInput) (*lightsail.DetachDiskOutput, error)
	DetachDiskWithContext(aws.Context, *lightsail.DetachDiskInput, ...request.Option) (*lightsail.DetachDiskOutput, error)
	DetachDiskRequest(*lightsail.DetachDiskInput) (*request.Request, *lightsail.DetachDiskOutput)

	DetachInstancesFromLoadBalancer(*lightsail.DetachInstancesFromLoadBalancerInput) (*lightsail.DetachInstancesFromLoadBalancerOutput, error)
	DetachInstancesFromLoadBalancerWithContext(aws.Context, *lightsail.DetachInstancesFromLoadBalancerInput, ...request.Option) (*lightsail.DetachInstancesFromLoadBalancerOutput, error)
	DetachInstancesFromLoadBalancerRequest(*lightsail.DetachInstancesFromLoadBalancerInput) (*request.Request, *lightsail.DetachInstancesFromLoadBalancerOutput)

	DetachStaticIp(*lightsail.DetachStaticIpInput) (*lightsail.DetachStaticIpOutput, error)
	DetachStaticIpWithContext(aws.Context, *lightsail.DetachStaticIpInput, ...request.Option) (*lightsail.DetachStaticIpOutput, error)
	DetachStaticIpRequest(*lightsail.DetachStaticIpInput) (*request.Request, *lightsail.DetachStaticIpOutput)

	DownloadDefaultKeyPair(*lightsail.DownloadDefaultKeyPairInput) (*lightsail.DownloadDefaultKeyPairOutput, error)
	DownloadDefaultKeyPairWithContext(aws.Context, *lightsail.DownloadDefaultKeyPairInput, ...request.Option) (*lightsail.DownloadDefaultKeyPairOutput, error)
	DownloadDefaultKeyPairRequest(*lightsail.DownloadDefaultKeyPairInput) (*request.Request, *lightsail.DownloadDefaultKeyPairOutput)

	ExportSnapshot(*lightsail.ExportSnapshotInput) (*lightsail.ExportSnapshotOutput, error)
	ExportSnapshotWithContext(aws.Context, *lightsail.ExportSnapshotInput, ...request.Option) (*lightsail.ExportSnapshotOutput, error)
	ExportSnapshotRequest(*lightsail.ExportSnapshotInput) (*request.Request, *lightsail.ExportSnapshotOutput)

	GetActiveNames(*lightsail.GetActiveNamesInput) (*lightsail.GetActiveNamesOutput, error)
	GetActiveNamesWithContext(aws.Context, *lightsail.GetActiveNamesInput, ...request.Option) (*lightsail.GetActiveNamesOutput, error)
	GetActiveNamesRequest(*lightsail.GetActiveNamesInput) (*request.Request, *lightsail.GetActiveNamesOutput)

	GetBlueprints(*lightsail.GetBlueprintsInput) (*lightsail.GetBlueprintsOutput, error)
	GetBlueprintsWithContext(aws.Context, *lightsail.GetBlueprintsInput, ...request.Option) (*lightsail.GetBlueprintsOutput, error)
	GetBlueprintsRequest(*lightsail.GetBlueprintsInput) (*request.Request, *lightsail.GetBlueprintsOutput)

	GetBundles(*lightsail.GetBundlesInput) (*lightsail.GetBundlesOutput, error)
	GetBundlesWithContext(aws.Context, *lightsail.GetBundlesInput, ...request.Option) (*lightsail.GetBundlesOutput, error)
	GetBundlesRequest(*lightsail.GetBundlesInput) (*request.Request, *lightsail.GetBundlesOutput)

	GetCloudFormationStackRecords(*lightsail.GetCloudFormationStackRecordsInput) (*lightsail.GetCloudFormationStackRecordsOutput, error)
	GetCloudFormationStackRecordsWithContext(aws.Context, *lightsail.GetCloudFormationStackRecordsInput, ...request.Option) (*lightsail.GetCloudFormationStackRecordsOutput, error)
	GetCloudFormationStackRecordsRequest(*lightsail.GetCloudFormationStackRecordsInput) (*request.Request, *lightsail.GetCloudFormationStackRecordsOutput)

	GetDisk(*lightsail.GetDiskInput) (*lightsail.GetDiskOutput, error)
	GetDiskWithContext(aws.Context, *lightsail.GetDiskInput, ...request.Option) (*lightsail.GetDiskOutput, error)
	GetDiskRequest(*lightsail.GetDiskInput) (*request.Request, *lightsail.GetDiskOutput)

	GetDiskSnapshot(*lightsail.GetDiskSnapshotInput) (*lightsail.GetDiskSnapshotOutput, error)
	GetDiskSnapshotWithContext(aws.Context, *lightsail.GetDiskSnapshotInput, ...request.Option) (*lightsail.GetDiskSnapshotOutput, error)
	GetDiskSnapshotRequest(*lightsail.GetDiskSnapshotInput) (*request.Request, *lightsail.GetDiskSnapshotOutput)

	GetDiskSnapshots(*lightsail.GetDiskSnapshotsInput) (*lightsail.GetDiskSnapshotsOutput, error)
	GetDiskSnapshotsWithContext(aws.Context, *lightsail.GetDiskSnapshotsInput, ...request.Option) (*lightsail.GetDiskSnapshotsOutput, error)
	GetDiskSnapshotsRequest(*lightsail.GetDiskSnapshotsInput) (*request.Request, *lightsail.GetDiskSnapshotsOutput)

	GetDisks(*lightsail.GetDisksInput) (*lightsail.GetDisksOutput, error)
	GetDisksWithContext(aws.Context, *lightsail.GetDisksInput, ...request.Option) (*lightsail.GetDisksOutput, error)
	GetDisksRequest(*lightsail.GetDisksInput) (*request.Request, *lightsail.GetDisksOutput)

	GetDomain(*lightsail.GetDomainInput) (*lightsail.GetDomainOutput, error)
	GetDomainWithContext(aws.Context, *lightsail.GetDomainInput, ...request.Option) (*lightsail.GetDomainOutput, error)
	GetDomainRequest(*lightsail.GetDomainInput) (*request.Request, *lightsail.GetDomainOutput)

	GetDomains(*lightsail.GetDomainsInput) (*lightsail.GetDomainsOutput, error)
	GetDomainsWithContext(aws.Context, *lightsail.GetDomainsInput, ...request.Option) (*lightsail.GetDomainsOutput, error)
	GetDomainsRequest(*lightsail.GetDomainsInput) (*request.Request, *lightsail.GetDomainsOutput)

	GetExportSnapshotRecords(*lightsail.GetExportSnapshotRecordsInput) (*lightsail.GetExportSnapshotRecordsOutput, error)
	GetExportSnapshotRecordsWithContext(aws.Context, *lightsail.GetExportSnapshotRecordsInput, ...request.Option) (*lightsail.GetExportSnapshotRecordsOutput, error)
	GetExportSnapshotRecordsRequest(*lightsail.GetExportSnapshotRecordsInput) (*request.Request, *lightsail.GetExportSnapshotRecordsOutput)

	GetInstance(*lightsail.GetInstanceInput) (*lightsail.GetInstanceOutput, error)
	GetInstanceWithContext(aws.Context, *lightsail.GetInstanceInput, ...request.Option) (*lightsail.GetInstanceOutput, error)
	GetInstanceRequest(*lightsail.GetInstanceInput) (*request.Request, *lightsail.GetInstanceOutput)

	GetInstanceAccessDetails(*lightsail.GetInstanceAccessDetailsInput) (*lightsail.GetInstanceAccessDetailsOutput, error)
	GetInstanceAccessDetailsWithContext(aws.Context, *lightsail.GetInstanceAccessDetailsInput, ...request.Option) (*lightsail.GetInstanceAccessDetailsOutput, error)
	GetInstanceAccessDetailsRequest(*lightsail.GetInstanceAccessDetailsInput) (*request.Request, *lightsail.GetInstanceAccessDetailsOutput)

	GetInstanceMetricData(*lightsail.GetInstanceMetricDataInput) (*lightsail.GetInstanceMetricDataOutput, error)
	GetInstanceMetricDataWithContext(aws.Context, *lightsail.GetInstanceMetricDataInput, ...request.Option) (*lightsail.GetInstanceMetricDataOutput, error)
	GetInstanceMetricDataRequest(*lightsail.GetInstanceMetricDataInput) (*request.Request, *lightsail.GetInstanceMetricDataOutput)

	GetInstancePortStates(*lightsail.GetInstancePortStatesInput) (*lightsail.GetInstancePortStatesOutput, error)
	GetInstancePortStatesWithContext(aws.Context, *lightsail.GetInstancePortStatesInput, ...request.Option) (*lightsail.GetInstancePortStatesOutput, error)
	GetInstancePortStatesRequest(*lightsail.GetInstancePortStatesInput) (*request.Request, *lightsail.GetInstancePortStatesOutput)

	GetInstanceSnapshot(*lightsail.GetInstanceSnapshotInput) (*lightsail.GetInstanceSnapshotOutput, error)
	GetInstanceSnapshotWithContext(aws.Context, *lightsail.GetInstanceSnapshotInput, ...request.Option) (*lightsail.GetInstanceSnapshotOutput, error)
	GetInstanceSnapshotRequest(*lightsail.GetInstanceSnapshotInput) (*request.Request, *lightsail.GetInstanceSnapshotOutput)

	GetInstanceSnapshots(*lightsail.GetInstanceSnapshotsInput) (*lightsail.GetInstanceSnapshotsOutput, error)
	GetInstanceSnapshotsWithContext(aws.Context, *lightsail.GetInstanceSnapshotsInput, ...request.Option) (*lightsail.GetInstanceSnapshotsOutput, error)
	GetInstanceSnapshotsRequest(*lightsail.GetInstanceSnapshotsInput) (*request.Request, *lightsail.GetInstanceSnapshotsOutput)

	GetInstanceState(*lightsail.GetInstanceStateInput) (*lightsail.GetInstanceStateOutput, error)
	GetInstanceStateWithContext(aws.Context, *lightsail.GetInstanceStateInput, ...request.Option) (*lightsail.GetInstanceStateOutput, error)
	GetInstanceStateRequest(*lightsail.GetInstanceStateInput) (*request.Request, *lightsail.GetInstanceStateOutput)

	GetInstances(*lightsail.GetInstancesInput) (*lightsail.GetInstancesOutput, error)
	GetInstancesWithContext(aws.Context, *lightsail.GetInstancesInput, ...request.Option) (*lightsail.GetInstancesOutput, error)
	GetInstancesRequest(*lightsail.GetInstancesInput) (*request.Request, *lightsail.GetInstancesOutput)

	GetKeyPair(*lightsail.GetKeyPairInput) (*lightsail.GetKeyPairOutput, error)
	GetKeyPairWithContext(aws.Context, *lightsail.GetKeyPairInput, ...request.Option) (*lightsail.GetKeyPairOutput, error)
	GetKeyPairRequest(*lightsail.GetKeyPairInput) (*request.Request, *lightsail.GetKeyPairOutput)

	GetKeyPairs(*lightsail.GetKeyPairsInput) (*lightsail.GetKeyPairsOutput, error)
	GetKeyPairsWithContext(aws.Context, *lightsail.GetKeyPairsInput, ...request.Option) (*lightsail.GetKeyPairsOutput, error)
	GetKeyPairsRequest(*lightsail.GetKeyPairsInput) (*request.Request, *lightsail.GetKeyPairsOutput)

	GetLoadBalancer(*lightsail.GetLoadBalancerInput) (*lightsail.GetLoadBalancerOutput, error)
	GetLoadBalancerWithContext(aws.Context, *lightsail.GetLoadBalancerInput, ...request.Option) (*lightsail.GetLoadBalancerOutput, error)
	GetLoadBalancerRequest(*lightsail.GetLoadBalancerInput) (*request.Request, *lightsail.GetLoadBalancerOutput)

	GetLoadBalancerMetricData(*lightsail.GetLoadBalancerMetricDataInput) (*lightsail.GetLoadBalancerMetricDataOutput, error)
	GetLoadBalancerMetricDataWithContext(aws.Context, *lightsail.GetLoadBalancerMetricDataInput, ...request.Option) (*lightsail.GetLoadBalancerMetricDataOutput, error)
	GetLoadBalancerMetricDataRequest(*lightsail.GetLoadBalancerMetricDataInput) (*request.Request, *lightsail.GetLoadBalancerMetricDataOutput)

	GetLoadBalancerTlsCertificates(*lightsail.GetLoadBalancerTlsCertificatesInput) (*lightsail.GetLoadBalancerTlsCertificatesOutput, error)
	GetLoadBalancerTlsCertificatesWithContext(aws.Context, *lightsail.GetLoadBalancerTlsCertificatesInput, ...request.Option) (*lightsail.GetLoadBalancerTlsCertificatesOutput, error)
	GetLoadBalancerTlsCertificatesRequest(*lightsail.GetLoadBalancerTlsCertificatesInput) (*request.Request, *lightsail.GetLoadBalancerTlsCertificatesOutput)

	GetLoadBalancers(*lightsail.GetLoadBalancersInput) (*lightsail.GetLoadBalancersOutput, error)
	GetLoadBalancersWithContext(aws.Context, *lightsail.GetLoadBalancersInput, ...request.Option) (*lightsail.GetLoadBalancersOutput, error)
	GetLoadBalancersRequest(*lightsail.GetLoadBalancersInput) (*request.Request, *lightsail.GetLoadBalancersOutput)

	GetOperation(*lightsail.GetOperationInput) (*lightsail.GetOperationOutput, error)
	GetOperationWithContext(aws.Context, *lightsail.GetOperationInput, ...request.Option) (*lightsail.GetOperationOutput, error)
	GetOperationRequest(*lightsail.GetOperationInput) (*request.Request, *lightsail.GetOperationOutput)

	GetOperations(*lightsail.GetOperationsInput) (*lightsail.GetOperationsOutput, error)
	GetOperationsWithContext(aws.Context, *lightsail.GetOperationsInput, ...request.Option) (*lightsail.GetOperationsOutput, error)
	GetOperationsRequest(*lightsail.GetOperationsInput) (*request.Request, *lightsail.GetOperationsOutput)

	GetOperationsForResource(*lightsail.GetOperationsForResourceInput) (*lightsail.GetOperationsForResourceOutput, error)
	GetOperationsForResourceWithContext(aws.Context, *lightsail.GetOperationsForResourceInput, ...request.Option) (*lightsail.GetOperationsForResourceOutput, error)
	GetOperationsForResourceRequest(*lightsail.GetOperationsForResourceInput) (*request.Request, *lightsail.GetOperationsForResourceOutput)

	GetRegions(*lightsail.GetRegionsInput) (*lightsail.GetRegionsOutput, error)
	GetRegionsWithContext(aws.Context, *lightsail.GetRegionsInput, ...request.Option) (*lightsail.GetRegionsOutput, error)
	GetRegionsRequest(*lightsail.GetRegionsInput) (*request.Request, *lightsail.GetRegionsOutput)

	GetRelationalDatabase(*lightsail.GetRelationalDatabaseInput) (*lightsail.GetRelationalDatabaseOutput, error)
	GetRelationalDatabaseWithContext(aws.Context, *lightsail.GetRelationalDatabaseInput, ...request.Option) (*lightsail.GetRelationalDatabaseOutput, error)
	GetRelationalDatabaseRequest(*lightsail.GetRelationalDatabaseInput) (*request.Request, *lightsail.GetRelationalDatabaseOutput)

	GetRelationalDatabaseBlueprints(*lightsail.GetRelationalDatabaseBlueprintsInput) (*lightsail.GetRelationalDatabaseBlueprintsOutput, error)
	GetRelationalDatabaseBlueprintsWithContext(aws.Context, *lightsail.GetRelationalDatabaseBlueprintsInput, ...request.Option) (*lightsail.GetRelationalDatabaseBlueprintsOutput, error)
	GetRelationalDatabaseBlueprintsRequest(*lightsail.GetRelationalDatabaseBlueprintsInput) (*request.Request, *lightsail.GetRelationalDatabaseBlueprintsOutput)

	GetRelationalDatabaseBundles(*lightsail.GetRelationalDatabaseBundlesInput) (*lightsail.GetRelationalDatabaseBundlesOutput, error)
	GetRelationalDatabaseBundlesWithContext(aws.Context, *lightsail.GetRelationalDatabaseBundlesInput, ...request.Option) (*lightsail.GetRelationalDatabaseBundlesOutput, error)
	GetRelationalDatabaseBundlesRequest(*lightsail.GetRelationalDatabaseBundlesInput) (*request.Request, *lightsail.GetRelationalDatabaseBundlesOutput)

	GetRelationalDatabaseEvents(*lightsail.GetRelationalDatabaseEventsInput) (*lightsail.GetRelationalDatabaseEventsOutput, error)
	GetRelationalDatabaseEventsWithContext(aws.Context, *lightsail.GetRelationalDatabaseEventsInput, ...request.Option) (*lightsail.GetRelationalDatabaseEventsOutput, error)
	GetRelationalDatabaseEventsRequest(*lightsail.GetRelationalDatabaseEventsInput) (*request.Request, *lightsail.GetRelationalDatabaseEventsOutput)

	GetRelationalDatabaseLogEvents(*lightsail.GetRelationalDatabaseLogEventsInput) (*lightsail.GetRelationalDatabaseLogEventsOutput, error)
	GetRelationalDatabaseLogEventsWithContext(aws.Context, *lightsail.GetRelationalDatabaseLogEventsInput, ...request.Option) (*lightsail.GetRelationalDatabaseLogEventsOutput, error)
	GetRelationalDatabaseLogEventsRequest(*lightsail.GetRelationalDatabaseLogEventsInput) (*request.Request, *lightsail.GetRelationalDatabaseLogEventsOutput)

	GetRelationalDatabaseLogStreams(*lightsail.GetRelationalDatabaseLogStreamsInput) (*lightsail.GetRelationalDatabaseLogStreamsOutput, error)
	GetRelationalDatabaseLogStreamsWithContext(aws.Context, *lightsail.GetRelationalDatabaseLogStreamsInput, ...request.Option) (*lightsail.GetRelationalDatabaseLogStreamsOutput, error)
	GetRelationalDatabaseLogStreamsRequest(*lightsail.GetRelationalDatabaseLogStreamsInput) (*request.Request, *lightsail.GetRelationalDatabaseLogStreamsOutput)

	GetRelationalDatabaseMasterUserPassword(*lightsail.GetRelationalDatabaseMasterUserPasswordInput) (*lightsail.GetRelationalDatabaseMasterUserPasswordOutput, error)
	GetRelationalDatabaseMasterUserPasswordWithContext(aws.Context, *lightsail.GetRelationalDatabaseMasterUserPasswordInput, ...request.Option) (*lightsail.GetRelationalDatabaseMasterUserPasswordOutput, error)
	GetRelationalDatabaseMasterUserPasswordRequest(*lightsail.GetRelationalDatabaseMasterUserPasswordInput) (*request.Request, *lightsail.GetRelationalDatabaseMasterUserPasswordOutput)

	GetRelationalDatabaseMetricData(*lightsail.GetRelationalDatabaseMetricDataInput) (*lightsail.GetRelationalDatabaseMetricDataOutput, error)
	GetRelationalDatabaseMetricDataWithContext(aws.Context, *lightsail.GetRelationalDatabaseMetricDataInput, ...request.Option) (*lightsail.GetRelationalDatabaseMetricDataOutput, error)
	GetRelationalDatabaseMetricDataRequest(*lightsail.GetRelationalDatabaseMetricDataInput) (*request.Request, *lightsail.GetRelationalDatabaseMetricDataOutput)

	GetRelationalDatabaseParameters(*lightsail.GetRelationalDatabaseParametersInput) (*lightsail.GetRelationalDatabaseParametersOutput, error)
	GetRelationalDatabaseParametersWithContext(aws.Context, *lightsail.GetRelationalDatabaseParametersInput, ...request.Option) (*lightsail.GetRelationalDatabaseParametersOutput, error)
	GetRelationalDatabaseParametersRequest(*lightsail.GetRelationalDatabaseParametersInput) (*request.Request, *lightsail.GetRelationalDatabaseParametersOutput)

	GetRelationalDatabaseSnapshot(*lightsail.GetRelationalDatabaseSnapshotInput) (*lightsail.GetRelationalDatabaseSnapshotOutput, error)
	GetRelationalDatabaseSnapshotWithContext(aws.Context, *lightsail.GetRelationalDatabaseSnapshotInput, ...request.Option) (*lightsail.GetRelationalDatabaseSnapshotOutput, error)
	GetRelationalDatabaseSnapshotRequest(*lightsail.GetRelationalDatabaseSnapshotInput) (*request.Request, *lightsail.GetRelationalDatabaseSnapshotOutput)

	GetRelationalDatabaseSnapshots(*lightsail.GetRelationalDatabaseSnapshotsInput) (*lightsail.GetRelationalDatabaseSnapshotsOutput, error)
	GetRelationalDatabaseSnapshotsWithContext(aws.Context, *lightsail.GetRelationalDatabaseSnapshotsInput, ...request.Option) (*lightsail.GetRelationalDatabaseSnapshotsOutput, error)
	GetRelationalDatabaseSnapshotsRequest(*lightsail.GetRelationalDatabaseSnapshotsInput) (*request.Request, *lightsail.GetRelationalDatabaseSnapshotsOutput)

	GetRelationalDatabases(*lightsail.GetRelationalDatabasesInput) (*lightsail.GetRelationalDatabasesOutput, error)
	GetRelationalDatabasesWithContext(aws.Context, *lightsail.GetRelationalDatabasesInput, ...request.Option) (*lightsail.GetRelationalDatabasesOutput, error)
	GetRelationalDatabasesRequest(*lightsail.GetRelationalDatabasesInput) (*request.Request, *lightsail.GetRelationalDatabasesOutput)

	GetStaticIp(*lightsail.GetStaticIpInput) (*lightsail.GetStaticIpOutput, error)
	GetStaticIpWithContext(aws.Context, *lightsail.GetStaticIpInput, ...request.Option) (*lightsail.GetStaticIpOutput, error)
	GetStaticIpRequest(*lightsail.GetStaticIpInput) (*request.Request, *lightsail.GetStaticIpOutput)

	GetStaticIps(*lightsail.GetStaticIpsInput) (*lightsail.GetStaticIpsOutput, error)
	GetStaticIpsWithContext(aws.Context, *lightsail.GetStaticIpsInput, ...request.Option) (*lightsail.GetStaticIpsOutput, error)
	GetStaticIpsRequest(*lightsail.GetStaticIpsInput) (*request.Request, *lightsail.GetStaticIpsOutput)

	ImportKeyPair(*lightsail.ImportKeyPairInput) (*lightsail.ImportKeyPairOutput, error)
	ImportKeyPairWithContext(aws.Context, *lightsail.ImportKeyPairInput, ...request.Option) (*lightsail.ImportKeyPairOutput, error)
	ImportKeyPairRequest(*lightsail.ImportKeyPairInput) (*request.Request, *lightsail.ImportKeyPairOutput)

	IsVpcPeered(*lightsail.IsVpcPeeredInput) (*lightsail.IsVpcPeeredOutput, error)
	IsVpcPeeredWithContext(aws.Context, *lightsail.IsVpcPeeredInput, ...request.Option) (*lightsail.IsVpcPeeredOutput, error)
	IsVpcPeeredRequest(*lightsail.IsVpcPeeredInput) (*request.Request, *lightsail.IsVpcPeeredOutput)

	OpenInstancePublicPorts(*lightsail.OpenInstancePublicPortsInput) (*lightsail.OpenInstancePublicPortsOutput, error)
	OpenInstancePublicPortsWithContext(aws.Context, *lightsail.OpenInstancePublicPortsInput, ...request.Option) (*lightsail.OpenInstancePublicPortsOutput, error)
	OpenInstancePublicPortsRequest(*lightsail.OpenInstancePublicPortsInput) (*request.Request, *lightsail.OpenInstancePublicPortsOutput)

	PeerVpc(*lightsail.PeerVpcInput) (*lightsail.PeerVpcOutput, error)
	PeerVpcWithContext(aws.Context, *lightsail.PeerVpcInput, ...request.Option) (*lightsail.PeerVpcOutput, error)
	PeerVpcRequest(*lightsail.PeerVpcInput) (*request.Request, *lightsail.PeerVpcOutput)

	PutInstancePublicPorts(*lightsail.PutInstancePublicPortsInput) (*lightsail.PutInstancePublicPortsOutput, error)
	PutInstancePublicPortsWithContext(aws.Context, *lightsail.PutInstancePublicPortsInput, ...request.Option) (*lightsail.PutInstancePublicPortsOutput, error)
	PutInstancePublicPortsRequest(*lightsail.PutInstancePublicPortsInput) (*request.Request, *lightsail.PutInstancePublicPortsOutput)

	RebootInstance(*lightsail.RebootInstanceInput) (*lightsail.RebootInstanceOutput, error)
	RebootInstanceWithContext(aws.Context, *lightsail.RebootInstanceInput, ...request.Option) (*lightsail.RebootInstanceOutput, error)
	RebootInstanceRequest(*lightsail.RebootInstanceInput) (*request.Request, *lightsail.RebootInstanceOutput)

	RebootRelationalDatabase(*lightsail.RebootRelationalDatabaseInput) (*lightsail.RebootRelationalDatabaseOutput, error)
	RebootRelationalDatabaseWithContext(aws.Context, *lightsail.RebootRelationalDatabaseInput, ...request.Option) (*lightsail.RebootRelationalDatabaseOutput, error)
	RebootRelationalDatabaseRequest(*lightsail.RebootRelationalDatabaseInput) (*request.Request, *lightsail.RebootRelationalDatabaseOutput)

	ReleaseStaticIp(*lightsail.ReleaseStaticIpInput) (*lightsail.ReleaseStaticIpOutput, error)
	ReleaseStaticIpWithContext(aws.Context, *lightsail.ReleaseStaticIpInput, ...request.Option) (*lightsail.ReleaseStaticIpOutput, error)
	ReleaseStaticIpRequest(*lightsail.ReleaseStaticIpInput) (*request.Request, *lightsail.ReleaseStaticIpOutput)

	StartInstance(*lightsail.StartInstanceInput) (*lightsail.StartInstanceOutput, error)
	StartInstanceWithContext(aws.Context, *lightsail.StartInstanceInput, ...request.Option) (*lightsail.StartInstanceOutput, error)
	StartInstanceRequest(*lightsail.StartInstanceInput) (*request.Request, *lightsail.StartInstanceOutput)

	StartRelationalDatabase(*lightsail.StartRelationalDatabaseInput) (*lightsail.StartRelationalDatabaseOutput, error)
	StartRelationalDatabaseWithContext(aws.Context, *lightsail.StartRelationalDatabaseInput, ...request.Option) (*lightsail.StartRelationalDatabaseOutput, error)
	StartRelationalDatabaseRequest(*lightsail.StartRelationalDatabaseInput) (*request.Request, *lightsail.StartRelationalDatabaseOutput)

	StopInstance(*lightsail.StopInstanceInput) (*lightsail.StopInstanceOutput, error)
	StopInstanceWithContext(aws.Context, *lightsail.StopInstanceInput, ...request.Option) (*lightsail.StopInstanceOutput, error)
	StopInstanceRequest(*lightsail.StopInstanceInput) (*request.Request, *lightsail.StopInstanceOutput)

	StopRelationalDatabase(*lightsail.StopRelationalDatabaseInput) (*lightsail.StopRelationalDatabaseOutput, error)
	StopRelationalDatabaseWithContext(aws.Context, *lightsail.StopRelationalDatabaseInput, ...request.Option) (*lightsail.StopRelationalDatabaseOutput, error)
	StopRelationalDatabaseRequest(*lightsail.StopRelationalDatabaseInput) (*request.Request, *lightsail.StopRelationalDatabaseOutput)

	TagResource(*lightsail.TagResourceInput) (*lightsail.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *lightsail.TagResourceInput, ...request.Option) (*lightsail.TagResourceOutput, error)
	TagResourceRequest(*lightsail.TagResourceInput) (*request.Request, *lightsail.TagResourceOutput)

	UnpeerVpc(*lightsail.UnpeerVpcInput) (*lightsail.UnpeerVpcOutput, error)
	UnpeerVpcWithContext(aws.Context, *lightsail.UnpeerVpcInput, ...request.Option) (*lightsail.UnpeerVpcOutput, error)
	UnpeerVpcRequest(*lightsail.UnpeerVpcInput) (*request.Request, *lightsail.UnpeerVpcOutput)

	UntagResource(*lightsail.UntagResourceInput) (*lightsail.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *lightsail.UntagResourceInput, ...request.Option) (*lightsail.UntagResourceOutput, error)
	UntagResourceRequest(*lightsail.UntagResourceInput) (*request.Request, *lightsail.UntagResourceOutput)

	UpdateDomainEntry(*lightsail.UpdateDomainEntryInput) (*lightsail.UpdateDomainEntryOutput, error)
	UpdateDomainEntryWithContext(aws.Context, *lightsail.UpdateDomainEntryInput, ...request.Option) (*lightsail.UpdateDomainEntryOutput, error)
	UpdateDomainEntryRequest(*lightsail.UpdateDomainEntryInput) (*request.Request, *lightsail.UpdateDomainEntryOutput)

	UpdateLoadBalancerAttribute(*lightsail.UpdateLoadBalancerAttributeInput) (*lightsail.UpdateLoadBalancerAttributeOutput, error)
	UpdateLoadBalancerAttributeWithContext(aws.Context, *lightsail.UpdateLoadBalancerAttributeInput, ...request.Option) (*lightsail.UpdateLoadBalancerAttributeOutput, error)
	UpdateLoadBalancerAttributeRequest(*lightsail.UpdateLoadBalancerAttributeInput) (*request.Request, *lightsail.UpdateLoadBalancerAttributeOutput)

	UpdateRelationalDatabase(*lightsail.UpdateRelationalDatabaseInput) (*lightsail.UpdateRelationalDatabaseOutput, error)
	UpdateRelationalDatabaseWithContext(aws.Context, *lightsail.UpdateRelationalDatabaseInput, ...request.Option) (*lightsail.UpdateRelationalDatabaseOutput, error)
	UpdateRelationalDatabaseRequest(*lightsail.UpdateRelationalDatabaseInput) (*request.Request, *lightsail.UpdateRelationalDatabaseOutput)

	UpdateRelationalDatabaseParameters(*lightsail.UpdateRelationalDatabaseParametersInput) (*lightsail.UpdateRelationalDatabaseParametersOutput, error)
	UpdateRelationalDatabaseParametersWithContext(aws.Context, *lightsail.UpdateRelationalDatabaseParametersInput, ...request.Option) (*lightsail.UpdateRelationalDatabaseParametersOutput, error)
	UpdateRelationalDatabaseParametersRequest(*lightsail.UpdateRelationalDatabaseParametersInput) (*request.Request, *lightsail.UpdateRelationalDatabaseParametersOutput)
}

var _ LightsailAPI = (*lightsail.Lightsail)(nil)
