// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package acmpcaiface provides an interface to enable mocking the AWS Certificate Manager Private Certificate Authority service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package acmpcaiface

import (
	github.com/Beeketing/aws-sdk-go/aws"
	github.com/Beeketing/aws-sdk-go/aws/request"
	github.com/Beeketing/aws-sdk-go/service/acmpca"
)

// ACMPCAAPI provides an interface to enable mocking the
// acmpca.ACMPCA service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Certificate Manager Private Certificate Authority.
//    func myFunc(svc acmpcaiface.ACMPCAAPI) bool {
//        // Make svc.CreateCertificateAuthority request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := acmpca.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockACMPCAClient struct {
//        acmpcaiface.ACMPCAAPI
//    }
//    func (m *mockACMPCAClient) CreateCertificateAuthority(input *acmpca.CreateCertificateAuthorityInput) (*acmpca.CreateCertificateAuthorityOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockACMPCAClient{}
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
type ACMPCAAPI interface {
	CreateCertificateAuthority(*acmpca.CreateCertificateAuthorityInput) (*acmpca.CreateCertificateAuthorityOutput, error)
	CreateCertificateAuthorityWithContext(aws.Context, *acmpca.CreateCertificateAuthorityInput, ...request.Option) (*acmpca.CreateCertificateAuthorityOutput, error)
	CreateCertificateAuthorityRequest(*acmpca.CreateCertificateAuthorityInput) (*request.Request, *acmpca.CreateCertificateAuthorityOutput)

	CreateCertificateAuthorityAuditReport(*acmpca.CreateCertificateAuthorityAuditReportInput) (*acmpca.CreateCertificateAuthorityAuditReportOutput, error)
	CreateCertificateAuthorityAuditReportWithContext(aws.Context, *acmpca.CreateCertificateAuthorityAuditReportInput, ...request.Option) (*acmpca.CreateCertificateAuthorityAuditReportOutput, error)
	CreateCertificateAuthorityAuditReportRequest(*acmpca.CreateCertificateAuthorityAuditReportInput) (*request.Request, *acmpca.CreateCertificateAuthorityAuditReportOutput)

	DeleteCertificateAuthority(*acmpca.DeleteCertificateAuthorityInput) (*acmpca.DeleteCertificateAuthorityOutput, error)
	DeleteCertificateAuthorityWithContext(aws.Context, *acmpca.DeleteCertificateAuthorityInput, ...request.Option) (*acmpca.DeleteCertificateAuthorityOutput, error)
	DeleteCertificateAuthorityRequest(*acmpca.DeleteCertificateAuthorityInput) (*request.Request, *acmpca.DeleteCertificateAuthorityOutput)

	DescribeCertificateAuthority(*acmpca.DescribeCertificateAuthorityInput) (*acmpca.DescribeCertificateAuthorityOutput, error)
	DescribeCertificateAuthorityWithContext(aws.Context, *acmpca.DescribeCertificateAuthorityInput, ...request.Option) (*acmpca.DescribeCertificateAuthorityOutput, error)
	DescribeCertificateAuthorityRequest(*acmpca.DescribeCertificateAuthorityInput) (*request.Request, *acmpca.DescribeCertificateAuthorityOutput)

	DescribeCertificateAuthorityAuditReport(*acmpca.DescribeCertificateAuthorityAuditReportInput) (*acmpca.DescribeCertificateAuthorityAuditReportOutput, error)
	DescribeCertificateAuthorityAuditReportWithContext(aws.Context, *acmpca.DescribeCertificateAuthorityAuditReportInput, ...request.Option) (*acmpca.DescribeCertificateAuthorityAuditReportOutput, error)
	DescribeCertificateAuthorityAuditReportRequest(*acmpca.DescribeCertificateAuthorityAuditReportInput) (*request.Request, *acmpca.DescribeCertificateAuthorityAuditReportOutput)

	GetCertificate(*acmpca.GetCertificateInput) (*acmpca.GetCertificateOutput, error)
	GetCertificateWithContext(aws.Context, *acmpca.GetCertificateInput, ...request.Option) (*acmpca.GetCertificateOutput, error)
	GetCertificateRequest(*acmpca.GetCertificateInput) (*request.Request, *acmpca.GetCertificateOutput)

	GetCertificateAuthorityCertificate(*acmpca.GetCertificateAuthorityCertificateInput) (*acmpca.GetCertificateAuthorityCertificateOutput, error)
	GetCertificateAuthorityCertificateWithContext(aws.Context, *acmpca.GetCertificateAuthorityCertificateInput, ...request.Option) (*acmpca.GetCertificateAuthorityCertificateOutput, error)
	GetCertificateAuthorityCertificateRequest(*acmpca.GetCertificateAuthorityCertificateInput) (*request.Request, *acmpca.GetCertificateAuthorityCertificateOutput)

	GetCertificateAuthorityCsr(*acmpca.GetCertificateAuthorityCsrInput) (*acmpca.GetCertificateAuthorityCsrOutput, error)
	GetCertificateAuthorityCsrWithContext(aws.Context, *acmpca.GetCertificateAuthorityCsrInput, ...request.Option) (*acmpca.GetCertificateAuthorityCsrOutput, error)
	GetCertificateAuthorityCsrRequest(*acmpca.GetCertificateAuthorityCsrInput) (*request.Request, *acmpca.GetCertificateAuthorityCsrOutput)

	ImportCertificateAuthorityCertificate(*acmpca.ImportCertificateAuthorityCertificateInput) (*acmpca.ImportCertificateAuthorityCertificateOutput, error)
	ImportCertificateAuthorityCertificateWithContext(aws.Context, *acmpca.ImportCertificateAuthorityCertificateInput, ...request.Option) (*acmpca.ImportCertificateAuthorityCertificateOutput, error)
	ImportCertificateAuthorityCertificateRequest(*acmpca.ImportCertificateAuthorityCertificateInput) (*request.Request, *acmpca.ImportCertificateAuthorityCertificateOutput)

	IssueCertificate(*acmpca.IssueCertificateInput) (*acmpca.IssueCertificateOutput, error)
	IssueCertificateWithContext(aws.Context, *acmpca.IssueCertificateInput, ...request.Option) (*acmpca.IssueCertificateOutput, error)
	IssueCertificateRequest(*acmpca.IssueCertificateInput) (*request.Request, *acmpca.IssueCertificateOutput)

	ListCertificateAuthorities(*acmpca.ListCertificateAuthoritiesInput) (*acmpca.ListCertificateAuthoritiesOutput, error)
	ListCertificateAuthoritiesWithContext(aws.Context, *acmpca.ListCertificateAuthoritiesInput, ...request.Option) (*acmpca.ListCertificateAuthoritiesOutput, error)
	ListCertificateAuthoritiesRequest(*acmpca.ListCertificateAuthoritiesInput) (*request.Request, *acmpca.ListCertificateAuthoritiesOutput)

	ListTags(*acmpca.ListTagsInput) (*acmpca.ListTagsOutput, error)
	ListTagsWithContext(aws.Context, *acmpca.ListTagsInput, ...request.Option) (*acmpca.ListTagsOutput, error)
	ListTagsRequest(*acmpca.ListTagsInput) (*request.Request, *acmpca.ListTagsOutput)

	RestoreCertificateAuthority(*acmpca.RestoreCertificateAuthorityInput) (*acmpca.RestoreCertificateAuthorityOutput, error)
	RestoreCertificateAuthorityWithContext(aws.Context, *acmpca.RestoreCertificateAuthorityInput, ...request.Option) (*acmpca.RestoreCertificateAuthorityOutput, error)
	RestoreCertificateAuthorityRequest(*acmpca.RestoreCertificateAuthorityInput) (*request.Request, *acmpca.RestoreCertificateAuthorityOutput)

	RevokeCertificate(*acmpca.RevokeCertificateInput) (*acmpca.RevokeCertificateOutput, error)
	RevokeCertificateWithContext(aws.Context, *acmpca.RevokeCertificateInput, ...request.Option) (*acmpca.RevokeCertificateOutput, error)
	RevokeCertificateRequest(*acmpca.RevokeCertificateInput) (*request.Request, *acmpca.RevokeCertificateOutput)

	TagCertificateAuthority(*acmpca.TagCertificateAuthorityInput) (*acmpca.TagCertificateAuthorityOutput, error)
	TagCertificateAuthorityWithContext(aws.Context, *acmpca.TagCertificateAuthorityInput, ...request.Option) (*acmpca.TagCertificateAuthorityOutput, error)
	TagCertificateAuthorityRequest(*acmpca.TagCertificateAuthorityInput) (*request.Request, *acmpca.TagCertificateAuthorityOutput)

	UntagCertificateAuthority(*acmpca.UntagCertificateAuthorityInput) (*acmpca.UntagCertificateAuthorityOutput, error)
	UntagCertificateAuthorityWithContext(aws.Context, *acmpca.UntagCertificateAuthorityInput, ...request.Option) (*acmpca.UntagCertificateAuthorityOutput, error)
	UntagCertificateAuthorityRequest(*acmpca.UntagCertificateAuthorityInput) (*request.Request, *acmpca.UntagCertificateAuthorityOutput)

	UpdateCertificateAuthority(*acmpca.UpdateCertificateAuthorityInput) (*acmpca.UpdateCertificateAuthorityOutput, error)
	UpdateCertificateAuthorityWithContext(aws.Context, *acmpca.UpdateCertificateAuthorityInput, ...request.Option) (*acmpca.UpdateCertificateAuthorityOutput, error)
	UpdateCertificateAuthorityRequest(*acmpca.UpdateCertificateAuthorityInput) (*request.Request, *acmpca.UpdateCertificateAuthorityOutput)
}

var _ ACMPCAAPI = (*acmpca.ACMPCA)(nil)
