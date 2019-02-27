// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package lexmodelbuildingserviceiface provides an interface to enable mocking the Amazon Lex Model Building Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package lexmodelbuildingserviceiface

import (
	github.com/Beeketing/aws-sdk-go/aws"
	github.com/Beeketing/aws-sdk-go/aws/request"
	github.com/Beeketing/aws-sdk-go/service/lexmodelbuildingservice"
)

// LexModelBuildingServiceAPI provides an interface to enable mocking the
// lexmodelbuildingservice.LexModelBuildingService service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Lex Model Building Service.
//    func myFunc(svc lexmodelbuildingserviceiface.LexModelBuildingServiceAPI) bool {
//        // Make svc.CreateBotVersion request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := lexmodelbuildingservice.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockLexModelBuildingServiceClient struct {
//        lexmodelbuildingserviceiface.LexModelBuildingServiceAPI
//    }
//    func (m *mockLexModelBuildingServiceClient) CreateBotVersion(input *lexmodelbuildingservice.CreateBotVersionInput) (*lexmodelbuildingservice.CreateBotVersionOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockLexModelBuildingServiceClient{}
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
type LexModelBuildingServiceAPI interface {
	CreateBotVersion(*lexmodelbuildingservice.CreateBotVersionInput) (*lexmodelbuildingservice.CreateBotVersionOutput, error)
	CreateBotVersionWithContext(aws.Context, *lexmodelbuildingservice.CreateBotVersionInput, ...request.Option) (*lexmodelbuildingservice.CreateBotVersionOutput, error)
	CreateBotVersionRequest(*lexmodelbuildingservice.CreateBotVersionInput) (*request.Request, *lexmodelbuildingservice.CreateBotVersionOutput)

	CreateIntentVersion(*lexmodelbuildingservice.CreateIntentVersionInput) (*lexmodelbuildingservice.CreateIntentVersionOutput, error)
	CreateIntentVersionWithContext(aws.Context, *lexmodelbuildingservice.CreateIntentVersionInput, ...request.Option) (*lexmodelbuildingservice.CreateIntentVersionOutput, error)
	CreateIntentVersionRequest(*lexmodelbuildingservice.CreateIntentVersionInput) (*request.Request, *lexmodelbuildingservice.CreateIntentVersionOutput)

	CreateSlotTypeVersion(*lexmodelbuildingservice.CreateSlotTypeVersionInput) (*lexmodelbuildingservice.CreateSlotTypeVersionOutput, error)
	CreateSlotTypeVersionWithContext(aws.Context, *lexmodelbuildingservice.CreateSlotTypeVersionInput, ...request.Option) (*lexmodelbuildingservice.CreateSlotTypeVersionOutput, error)
	CreateSlotTypeVersionRequest(*lexmodelbuildingservice.CreateSlotTypeVersionInput) (*request.Request, *lexmodelbuildingservice.CreateSlotTypeVersionOutput)

	DeleteBot(*lexmodelbuildingservice.DeleteBotInput) (*lexmodelbuildingservice.DeleteBotOutput, error)
	DeleteBotWithContext(aws.Context, *lexmodelbuildingservice.DeleteBotInput, ...request.Option) (*lexmodelbuildingservice.DeleteBotOutput, error)
	DeleteBotRequest(*lexmodelbuildingservice.DeleteBotInput) (*request.Request, *lexmodelbuildingservice.DeleteBotOutput)

	DeleteBotAlias(*lexmodelbuildingservice.DeleteBotAliasInput) (*lexmodelbuildingservice.DeleteBotAliasOutput, error)
	DeleteBotAliasWithContext(aws.Context, *lexmodelbuildingservice.DeleteBotAliasInput, ...request.Option) (*lexmodelbuildingservice.DeleteBotAliasOutput, error)
	DeleteBotAliasRequest(*lexmodelbuildingservice.DeleteBotAliasInput) (*request.Request, *lexmodelbuildingservice.DeleteBotAliasOutput)

	DeleteBotChannelAssociation(*lexmodelbuildingservice.DeleteBotChannelAssociationInput) (*lexmodelbuildingservice.DeleteBotChannelAssociationOutput, error)
	DeleteBotChannelAssociationWithContext(aws.Context, *lexmodelbuildingservice.DeleteBotChannelAssociationInput, ...request.Option) (*lexmodelbuildingservice.DeleteBotChannelAssociationOutput, error)
	DeleteBotChannelAssociationRequest(*lexmodelbuildingservice.DeleteBotChannelAssociationInput) (*request.Request, *lexmodelbuildingservice.DeleteBotChannelAssociationOutput)

	DeleteBotVersion(*lexmodelbuildingservice.DeleteBotVersionInput) (*lexmodelbuildingservice.DeleteBotVersionOutput, error)
	DeleteBotVersionWithContext(aws.Context, *lexmodelbuildingservice.DeleteBotVersionInput, ...request.Option) (*lexmodelbuildingservice.DeleteBotVersionOutput, error)
	DeleteBotVersionRequest(*lexmodelbuildingservice.DeleteBotVersionInput) (*request.Request, *lexmodelbuildingservice.DeleteBotVersionOutput)

	DeleteIntent(*lexmodelbuildingservice.DeleteIntentInput) (*lexmodelbuildingservice.DeleteIntentOutput, error)
	DeleteIntentWithContext(aws.Context, *lexmodelbuildingservice.DeleteIntentInput, ...request.Option) (*lexmodelbuildingservice.DeleteIntentOutput, error)
	DeleteIntentRequest(*lexmodelbuildingservice.DeleteIntentInput) (*request.Request, *lexmodelbuildingservice.DeleteIntentOutput)

	DeleteIntentVersion(*lexmodelbuildingservice.DeleteIntentVersionInput) (*lexmodelbuildingservice.DeleteIntentVersionOutput, error)
	DeleteIntentVersionWithContext(aws.Context, *lexmodelbuildingservice.DeleteIntentVersionInput, ...request.Option) (*lexmodelbuildingservice.DeleteIntentVersionOutput, error)
	DeleteIntentVersionRequest(*lexmodelbuildingservice.DeleteIntentVersionInput) (*request.Request, *lexmodelbuildingservice.DeleteIntentVersionOutput)

	DeleteSlotType(*lexmodelbuildingservice.DeleteSlotTypeInput) (*lexmodelbuildingservice.DeleteSlotTypeOutput, error)
	DeleteSlotTypeWithContext(aws.Context, *lexmodelbuildingservice.DeleteSlotTypeInput, ...request.Option) (*lexmodelbuildingservice.DeleteSlotTypeOutput, error)
	DeleteSlotTypeRequest(*lexmodelbuildingservice.DeleteSlotTypeInput) (*request.Request, *lexmodelbuildingservice.DeleteSlotTypeOutput)

	DeleteSlotTypeVersion(*lexmodelbuildingservice.DeleteSlotTypeVersionInput) (*lexmodelbuildingservice.DeleteSlotTypeVersionOutput, error)
	DeleteSlotTypeVersionWithContext(aws.Context, *lexmodelbuildingservice.DeleteSlotTypeVersionInput, ...request.Option) (*lexmodelbuildingservice.DeleteSlotTypeVersionOutput, error)
	DeleteSlotTypeVersionRequest(*lexmodelbuildingservice.DeleteSlotTypeVersionInput) (*request.Request, *lexmodelbuildingservice.DeleteSlotTypeVersionOutput)

	DeleteUtterances(*lexmodelbuildingservice.DeleteUtterancesInput) (*lexmodelbuildingservice.DeleteUtterancesOutput, error)
	DeleteUtterancesWithContext(aws.Context, *lexmodelbuildingservice.DeleteUtterancesInput, ...request.Option) (*lexmodelbuildingservice.DeleteUtterancesOutput, error)
	DeleteUtterancesRequest(*lexmodelbuildingservice.DeleteUtterancesInput) (*request.Request, *lexmodelbuildingservice.DeleteUtterancesOutput)

	GetBot(*lexmodelbuildingservice.GetBotInput) (*lexmodelbuildingservice.GetBotOutput, error)
	GetBotWithContext(aws.Context, *lexmodelbuildingservice.GetBotInput, ...request.Option) (*lexmodelbuildingservice.GetBotOutput, error)
	GetBotRequest(*lexmodelbuildingservice.GetBotInput) (*request.Request, *lexmodelbuildingservice.GetBotOutput)

	GetBotAlias(*lexmodelbuildingservice.GetBotAliasInput) (*lexmodelbuildingservice.GetBotAliasOutput, error)
	GetBotAliasWithContext(aws.Context, *lexmodelbuildingservice.GetBotAliasInput, ...request.Option) (*lexmodelbuildingservice.GetBotAliasOutput, error)
	GetBotAliasRequest(*lexmodelbuildingservice.GetBotAliasInput) (*request.Request, *lexmodelbuildingservice.GetBotAliasOutput)

	GetBotAliases(*lexmodelbuildingservice.GetBotAliasesInput) (*lexmodelbuildingservice.GetBotAliasesOutput, error)
	GetBotAliasesWithContext(aws.Context, *lexmodelbuildingservice.GetBotAliasesInput, ...request.Option) (*lexmodelbuildingservice.GetBotAliasesOutput, error)
	GetBotAliasesRequest(*lexmodelbuildingservice.GetBotAliasesInput) (*request.Request, *lexmodelbuildingservice.GetBotAliasesOutput)

	GetBotAliasesPages(*lexmodelbuildingservice.GetBotAliasesInput, func(*lexmodelbuildingservice.GetBotAliasesOutput, bool) bool) error
	GetBotAliasesPagesWithContext(aws.Context, *lexmodelbuildingservice.GetBotAliasesInput, func(*lexmodelbuildingservice.GetBotAliasesOutput, bool) bool, ...request.Option) error

	GetBotChannelAssociation(*lexmodelbuildingservice.GetBotChannelAssociationInput) (*lexmodelbuildingservice.GetBotChannelAssociationOutput, error)
	GetBotChannelAssociationWithContext(aws.Context, *lexmodelbuildingservice.GetBotChannelAssociationInput, ...request.Option) (*lexmodelbuildingservice.GetBotChannelAssociationOutput, error)
	GetBotChannelAssociationRequest(*lexmodelbuildingservice.GetBotChannelAssociationInput) (*request.Request, *lexmodelbuildingservice.GetBotChannelAssociationOutput)

	GetBotChannelAssociations(*lexmodelbuildingservice.GetBotChannelAssociationsInput) (*lexmodelbuildingservice.GetBotChannelAssociationsOutput, error)
	GetBotChannelAssociationsWithContext(aws.Context, *lexmodelbuildingservice.GetBotChannelAssociationsInput, ...request.Option) (*lexmodelbuildingservice.GetBotChannelAssociationsOutput, error)
	GetBotChannelAssociationsRequest(*lexmodelbuildingservice.GetBotChannelAssociationsInput) (*request.Request, *lexmodelbuildingservice.GetBotChannelAssociationsOutput)

	GetBotChannelAssociationsPages(*lexmodelbuildingservice.GetBotChannelAssociationsInput, func(*lexmodelbuildingservice.GetBotChannelAssociationsOutput, bool) bool) error
	GetBotChannelAssociationsPagesWithContext(aws.Context, *lexmodelbuildingservice.GetBotChannelAssociationsInput, func(*lexmodelbuildingservice.GetBotChannelAssociationsOutput, bool) bool, ...request.Option) error

	GetBotVersions(*lexmodelbuildingservice.GetBotVersionsInput) (*lexmodelbuildingservice.GetBotVersionsOutput, error)
	GetBotVersionsWithContext(aws.Context, *lexmodelbuildingservice.GetBotVersionsInput, ...request.Option) (*lexmodelbuildingservice.GetBotVersionsOutput, error)
	GetBotVersionsRequest(*lexmodelbuildingservice.GetBotVersionsInput) (*request.Request, *lexmodelbuildingservice.GetBotVersionsOutput)

	GetBotVersionsPages(*lexmodelbuildingservice.GetBotVersionsInput, func(*lexmodelbuildingservice.GetBotVersionsOutput, bool) bool) error
	GetBotVersionsPagesWithContext(aws.Context, *lexmodelbuildingservice.GetBotVersionsInput, func(*lexmodelbuildingservice.GetBotVersionsOutput, bool) bool, ...request.Option) error

	GetBots(*lexmodelbuildingservice.GetBotsInput) (*lexmodelbuildingservice.GetBotsOutput, error)
	GetBotsWithContext(aws.Context, *lexmodelbuildingservice.GetBotsInput, ...request.Option) (*lexmodelbuildingservice.GetBotsOutput, error)
	GetBotsRequest(*lexmodelbuildingservice.GetBotsInput) (*request.Request, *lexmodelbuildingservice.GetBotsOutput)

	GetBotsPages(*lexmodelbuildingservice.GetBotsInput, func(*lexmodelbuildingservice.GetBotsOutput, bool) bool) error
	GetBotsPagesWithContext(aws.Context, *lexmodelbuildingservice.GetBotsInput, func(*lexmodelbuildingservice.GetBotsOutput, bool) bool, ...request.Option) error

	GetBuiltinIntent(*lexmodelbuildingservice.GetBuiltinIntentInput) (*lexmodelbuildingservice.GetBuiltinIntentOutput, error)
	GetBuiltinIntentWithContext(aws.Context, *lexmodelbuildingservice.GetBuiltinIntentInput, ...request.Option) (*lexmodelbuildingservice.GetBuiltinIntentOutput, error)
	GetBuiltinIntentRequest(*lexmodelbuildingservice.GetBuiltinIntentInput) (*request.Request, *lexmodelbuildingservice.GetBuiltinIntentOutput)

	GetBuiltinIntents(*lexmodelbuildingservice.GetBuiltinIntentsInput) (*lexmodelbuildingservice.GetBuiltinIntentsOutput, error)
	GetBuiltinIntentsWithContext(aws.Context, *lexmodelbuildingservice.GetBuiltinIntentsInput, ...request.Option) (*lexmodelbuildingservice.GetBuiltinIntentsOutput, error)
	GetBuiltinIntentsRequest(*lexmodelbuildingservice.GetBuiltinIntentsInput) (*request.Request, *lexmodelbuildingservice.GetBuiltinIntentsOutput)

	GetBuiltinIntentsPages(*lexmodelbuildingservice.GetBuiltinIntentsInput, func(*lexmodelbuildingservice.GetBuiltinIntentsOutput, bool) bool) error
	GetBuiltinIntentsPagesWithContext(aws.Context, *lexmodelbuildingservice.GetBuiltinIntentsInput, func(*lexmodelbuildingservice.GetBuiltinIntentsOutput, bool) bool, ...request.Option) error

	GetBuiltinSlotTypes(*lexmodelbuildingservice.GetBuiltinSlotTypesInput) (*lexmodelbuildingservice.GetBuiltinSlotTypesOutput, error)
	GetBuiltinSlotTypesWithContext(aws.Context, *lexmodelbuildingservice.GetBuiltinSlotTypesInput, ...request.Option) (*lexmodelbuildingservice.GetBuiltinSlotTypesOutput, error)
	GetBuiltinSlotTypesRequest(*lexmodelbuildingservice.GetBuiltinSlotTypesInput) (*request.Request, *lexmodelbuildingservice.GetBuiltinSlotTypesOutput)

	GetBuiltinSlotTypesPages(*lexmodelbuildingservice.GetBuiltinSlotTypesInput, func(*lexmodelbuildingservice.GetBuiltinSlotTypesOutput, bool) bool) error
	GetBuiltinSlotTypesPagesWithContext(aws.Context, *lexmodelbuildingservice.GetBuiltinSlotTypesInput, func(*lexmodelbuildingservice.GetBuiltinSlotTypesOutput, bool) bool, ...request.Option) error

	GetExport(*lexmodelbuildingservice.GetExportInput) (*lexmodelbuildingservice.GetExportOutput, error)
	GetExportWithContext(aws.Context, *lexmodelbuildingservice.GetExportInput, ...request.Option) (*lexmodelbuildingservice.GetExportOutput, error)
	GetExportRequest(*lexmodelbuildingservice.GetExportInput) (*request.Request, *lexmodelbuildingservice.GetExportOutput)

	GetImport(*lexmodelbuildingservice.GetImportInput) (*lexmodelbuildingservice.GetImportOutput, error)
	GetImportWithContext(aws.Context, *lexmodelbuildingservice.GetImportInput, ...request.Option) (*lexmodelbuildingservice.GetImportOutput, error)
	GetImportRequest(*lexmodelbuildingservice.GetImportInput) (*request.Request, *lexmodelbuildingservice.GetImportOutput)

	GetIntent(*lexmodelbuildingservice.GetIntentInput) (*lexmodelbuildingservice.GetIntentOutput, error)
	GetIntentWithContext(aws.Context, *lexmodelbuildingservice.GetIntentInput, ...request.Option) (*lexmodelbuildingservice.GetIntentOutput, error)
	GetIntentRequest(*lexmodelbuildingservice.GetIntentInput) (*request.Request, *lexmodelbuildingservice.GetIntentOutput)

	GetIntentVersions(*lexmodelbuildingservice.GetIntentVersionsInput) (*lexmodelbuildingservice.GetIntentVersionsOutput, error)
	GetIntentVersionsWithContext(aws.Context, *lexmodelbuildingservice.GetIntentVersionsInput, ...request.Option) (*lexmodelbuildingservice.GetIntentVersionsOutput, error)
	GetIntentVersionsRequest(*lexmodelbuildingservice.GetIntentVersionsInput) (*request.Request, *lexmodelbuildingservice.GetIntentVersionsOutput)

	GetIntentVersionsPages(*lexmodelbuildingservice.GetIntentVersionsInput, func(*lexmodelbuildingservice.GetIntentVersionsOutput, bool) bool) error
	GetIntentVersionsPagesWithContext(aws.Context, *lexmodelbuildingservice.GetIntentVersionsInput, func(*lexmodelbuildingservice.GetIntentVersionsOutput, bool) bool, ...request.Option) error

	GetIntents(*lexmodelbuildingservice.GetIntentsInput) (*lexmodelbuildingservice.GetIntentsOutput, error)
	GetIntentsWithContext(aws.Context, *lexmodelbuildingservice.GetIntentsInput, ...request.Option) (*lexmodelbuildingservice.GetIntentsOutput, error)
	GetIntentsRequest(*lexmodelbuildingservice.GetIntentsInput) (*request.Request, *lexmodelbuildingservice.GetIntentsOutput)

	GetIntentsPages(*lexmodelbuildingservice.GetIntentsInput, func(*lexmodelbuildingservice.GetIntentsOutput, bool) bool) error
	GetIntentsPagesWithContext(aws.Context, *lexmodelbuildingservice.GetIntentsInput, func(*lexmodelbuildingservice.GetIntentsOutput, bool) bool, ...request.Option) error

	GetSlotType(*lexmodelbuildingservice.GetSlotTypeInput) (*lexmodelbuildingservice.GetSlotTypeOutput, error)
	GetSlotTypeWithContext(aws.Context, *lexmodelbuildingservice.GetSlotTypeInput, ...request.Option) (*lexmodelbuildingservice.GetSlotTypeOutput, error)
	GetSlotTypeRequest(*lexmodelbuildingservice.GetSlotTypeInput) (*request.Request, *lexmodelbuildingservice.GetSlotTypeOutput)

	GetSlotTypeVersions(*lexmodelbuildingservice.GetSlotTypeVersionsInput) (*lexmodelbuildingservice.GetSlotTypeVersionsOutput, error)
	GetSlotTypeVersionsWithContext(aws.Context, *lexmodelbuildingservice.GetSlotTypeVersionsInput, ...request.Option) (*lexmodelbuildingservice.GetSlotTypeVersionsOutput, error)
	GetSlotTypeVersionsRequest(*lexmodelbuildingservice.GetSlotTypeVersionsInput) (*request.Request, *lexmodelbuildingservice.GetSlotTypeVersionsOutput)

	GetSlotTypeVersionsPages(*lexmodelbuildingservice.GetSlotTypeVersionsInput, func(*lexmodelbuildingservice.GetSlotTypeVersionsOutput, bool) bool) error
	GetSlotTypeVersionsPagesWithContext(aws.Context, *lexmodelbuildingservice.GetSlotTypeVersionsInput, func(*lexmodelbuildingservice.GetSlotTypeVersionsOutput, bool) bool, ...request.Option) error

	GetSlotTypes(*lexmodelbuildingservice.GetSlotTypesInput) (*lexmodelbuildingservice.GetSlotTypesOutput, error)
	GetSlotTypesWithContext(aws.Context, *lexmodelbuildingservice.GetSlotTypesInput, ...request.Option) (*lexmodelbuildingservice.GetSlotTypesOutput, error)
	GetSlotTypesRequest(*lexmodelbuildingservice.GetSlotTypesInput) (*request.Request, *lexmodelbuildingservice.GetSlotTypesOutput)

	GetSlotTypesPages(*lexmodelbuildingservice.GetSlotTypesInput, func(*lexmodelbuildingservice.GetSlotTypesOutput, bool) bool) error
	GetSlotTypesPagesWithContext(aws.Context, *lexmodelbuildingservice.GetSlotTypesInput, func(*lexmodelbuildingservice.GetSlotTypesOutput, bool) bool, ...request.Option) error

	GetUtterancesView(*lexmodelbuildingservice.GetUtterancesViewInput) (*lexmodelbuildingservice.GetUtterancesViewOutput, error)
	GetUtterancesViewWithContext(aws.Context, *lexmodelbuildingservice.GetUtterancesViewInput, ...request.Option) (*lexmodelbuildingservice.GetUtterancesViewOutput, error)
	GetUtterancesViewRequest(*lexmodelbuildingservice.GetUtterancesViewInput) (*request.Request, *lexmodelbuildingservice.GetUtterancesViewOutput)

	PutBot(*lexmodelbuildingservice.PutBotInput) (*lexmodelbuildingservice.PutBotOutput, error)
	PutBotWithContext(aws.Context, *lexmodelbuildingservice.PutBotInput, ...request.Option) (*lexmodelbuildingservice.PutBotOutput, error)
	PutBotRequest(*lexmodelbuildingservice.PutBotInput) (*request.Request, *lexmodelbuildingservice.PutBotOutput)

	PutBotAlias(*lexmodelbuildingservice.PutBotAliasInput) (*lexmodelbuildingservice.PutBotAliasOutput, error)
	PutBotAliasWithContext(aws.Context, *lexmodelbuildingservice.PutBotAliasInput, ...request.Option) (*lexmodelbuildingservice.PutBotAliasOutput, error)
	PutBotAliasRequest(*lexmodelbuildingservice.PutBotAliasInput) (*request.Request, *lexmodelbuildingservice.PutBotAliasOutput)

	PutIntent(*lexmodelbuildingservice.PutIntentInput) (*lexmodelbuildingservice.PutIntentOutput, error)
	PutIntentWithContext(aws.Context, *lexmodelbuildingservice.PutIntentInput, ...request.Option) (*lexmodelbuildingservice.PutIntentOutput, error)
	PutIntentRequest(*lexmodelbuildingservice.PutIntentInput) (*request.Request, *lexmodelbuildingservice.PutIntentOutput)

	PutSlotType(*lexmodelbuildingservice.PutSlotTypeInput) (*lexmodelbuildingservice.PutSlotTypeOutput, error)
	PutSlotTypeWithContext(aws.Context, *lexmodelbuildingservice.PutSlotTypeInput, ...request.Option) (*lexmodelbuildingservice.PutSlotTypeOutput, error)
	PutSlotTypeRequest(*lexmodelbuildingservice.PutSlotTypeInput) (*request.Request, *lexmodelbuildingservice.PutSlotTypeOutput)

	StartImport(*lexmodelbuildingservice.StartImportInput) (*lexmodelbuildingservice.StartImportOutput, error)
	StartImportWithContext(aws.Context, *lexmodelbuildingservice.StartImportInput, ...request.Option) (*lexmodelbuildingservice.StartImportOutput, error)
	StartImportRequest(*lexmodelbuildingservice.StartImportInput) (*request.Request, *lexmodelbuildingservice.StartImportOutput)
}

var _ LexModelBuildingServiceAPI = (*lexmodelbuildingservice.LexModelBuildingService)(nil)
