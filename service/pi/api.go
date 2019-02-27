// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pi

import (
	"fmt"
	"time"

	"github.com/Beeketing/aws-sdk-go/aws"
	"github.com/Beeketing/aws-sdk-go/aws/awsutil"
	"github.com/Beeketing/aws-sdk-go/aws/request"
)

const opDescribeDimensionKeys = "DescribeDimensionKeys"

// DescribeDimensionKeysRequest generates a "aws/request.Request" representing the
// client's request for the DescribeDimensionKeys operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See DescribeDimensionKeys for more information on using the DescribeDimensionKeys
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the DescribeDimensionKeysRequest method.
//    req, resp := client.DescribeDimensionKeysRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/pi-2018-02-27/DescribeDimensionKeys
func (c *PI) DescribeDimensionKeysRequest(input *DescribeDimensionKeysInput) (req *request.Request, output *DescribeDimensionKeysOutput) {
	op := &request.Operation{
		Name:       opDescribeDimensionKeys,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDimensionKeysInput{}
	}

	output = &DescribeDimensionKeysOutput{}
	req = c.newRequest(op, input, output)
	return
}

// DescribeDimensionKeys API operation for AWS Performance Insights.
//
// For a specific time period, retrieve the top N dimension keys for a metric.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS Performance Insights's
// API operation DescribeDimensionKeys for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeInvalidArgumentException "InvalidArgumentException"
//   One of the arguments provided is invalid for this request.
//
//   * ErrCodeInternalServiceError "InternalServiceError"
//   The request failed due to an unknown error.
//
//   * ErrCodeNotAuthorizedException "NotAuthorizedException"
//   The user is not authorized to perform this request.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/pi-2018-02-27/DescribeDimensionKeys
func (c *PI) DescribeDimensionKeys(input *DescribeDimensionKeysInput) (*DescribeDimensionKeysOutput, error) {
	req, out := c.DescribeDimensionKeysRequest(input)
	return out, req.Send()
}

// DescribeDimensionKeysWithContext is the same as DescribeDimensionKeys with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDimensionKeys for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PI) DescribeDimensionKeysWithContext(ctx aws.Context, input *DescribeDimensionKeysInput, opts ...request.Option) (*DescribeDimensionKeysOutput, error) {
	req, out := c.DescribeDimensionKeysRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetResourceMetrics = "GetResourceMetrics"

// GetResourceMetricsRequest generates a "aws/request.Request" representing the
// client's request for the GetResourceMetrics operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetResourceMetrics for more information on using the GetResourceMetrics
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the GetResourceMetricsRequest method.
//    req, resp := client.GetResourceMetricsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/pi-2018-02-27/GetResourceMetrics
func (c *PI) GetResourceMetricsRequest(input *GetResourceMetricsInput) (req *request.Request, output *GetResourceMetricsOutput) {
	op := &request.Operation{
		Name:       opGetResourceMetrics,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetResourceMetricsInput{}
	}

	output = &GetResourceMetricsOutput{}
	req = c.newRequest(op, input, output)
	return
}

// GetResourceMetrics API operation for AWS Performance Insights.
//
// Retrieve Performance Insights metrics for a set of data sources, over a time
// period. You can provide specific dimension groups and dimensions, and provide
// aggregation and filtering criteria for each group.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS Performance Insights's
// API operation GetResourceMetrics for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeInvalidArgumentException "InvalidArgumentException"
//   One of the arguments provided is invalid for this request.
//
//   * ErrCodeInternalServiceError "InternalServiceError"
//   The request failed due to an unknown error.
//
//   * ErrCodeNotAuthorizedException "NotAuthorizedException"
//   The user is not authorized to perform this request.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/pi-2018-02-27/GetResourceMetrics
func (c *PI) GetResourceMetrics(input *GetResourceMetricsInput) (*GetResourceMetricsOutput, error) {
	req, out := c.GetResourceMetricsRequest(input)
	return out, req.Send()
}

// GetResourceMetricsWithContext is the same as GetResourceMetrics with the addition of
// the ability to pass a context and additional request options.
//
// See GetResourceMetrics for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PI) GetResourceMetricsWithContext(ctx aws.Context, input *GetResourceMetricsInput, opts ...request.Option) (*GetResourceMetricsOutput, error) {
	req, out := c.GetResourceMetricsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// A timestamp, and a single numerical value, which together represent a measurement
// at a particular point in time.
type DataPoint struct {
	_ struct{} `type:"structure"`

	// The time, in epoch format, associated with a particular Value.
	//
	// Timestamp is a required field
	Timestamp *time.Time `type:"timestamp" required:"true"`

	// The actual value associated with a particular Timestamp.
	//
	// Value is a required field
	Value *float64 `type:"double" required:"true"`
}

// String returns the string representation
func (s DataPoint) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DataPoint) GoString() string {
	return s.String()
}

// SetTimestamp sets the Timestamp field's value.
func (s *DataPoint) SetTimestamp(v time.Time) *DataPoint {
	s.Timestamp = &v
	return s
}

// SetValue sets the Value field's value.
func (s *DataPoint) SetValue(v float64) *DataPoint {
	s.Value = &v
	return s
}

type DescribeDimensionKeysInput struct {
	_ struct{} `type:"structure"`

	// The date and time specifying the end of the requested time series data. The
	// value specified is exclusive - data points less than (but not equal to) EndTime
	// will be returned.
	//
	// The value for EndTime must be later than the value for StartTime.
	//
	// EndTime is a required field
	EndTime *time.Time `type:"timestamp" required:"true"`

	// One or more filters to apply in the request. Restrictions:
	//
	//    * Any number of filters by the same dimension, as specified in the GroupBy
	//    or Partition parameters.
	//
	//    * A single filter for any other dimension in this dimension group.
	Filter map[string]*string `type:"map"`

	// A specification for how to aggregate the data points from a query result.
	// You must specify a valid dimension group. Performance Insights will return
	// all of the dimensions within that group, unless you provide the names of
	// specific dimensions within that group. You can also request that Performance
	// Insights return a limited number of values for a dimension.
	//
	// GroupBy is a required field
	GroupBy *DimensionGroup `type:"structure" required:"true"`

	// An immutable, AWS Region-unique identifier for a data source. Performance
	// Insights gathers metrics from this data source.
	//
	// To use an Amazon RDS instance as a data source, you specify its DbiResourceId
	// value - for example: db-FAIHNTYBKTGAUSUZQYPDS2GW4A
	//
	// Identifier is a required field
	Identifier *string `type:"string" required:"true"`

	// The maximum number of items to return in the response. If more items exist
	// than the specified MaxRecords value, a pagination token is included in the
	// response so that the remaining results can be retrieved.
	MaxResults *int64 `type:"integer"`

	// The name of a Performance Insights metric to be measured.
	//
	// Valid values for Metric are:
	//
	//    * db.load.avg - a scaled representation of the number of active sessions
	//    for the database engine.
	//
	//    * db.sampledload.avg - the raw number of active sessions for the database
	//    engine.
	//
	// Metric is a required field
	Metric *string `type:"string" required:"true"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the token, up to
	// the value specified by MaxRecords.
	NextToken *string `type:"string"`

	// For each dimension specified in GroupBy, specify a secondary dimension to
	// further subdivide the partition keys in the response.
	PartitionBy *DimensionGroup `type:"structure"`

	// The granularity, in seconds, of the data points returned from Performance
	// Insights. A period can be as short as one second, or as long as one day (86400
	// seconds). Valid values are:
	//
	//    * 1 (one second)
	//
	//    * 60 (one minute)
	//
	//    * 300 (five minutes)
	//
	//    * 3600 (one hour)
	//
	//    * 86400 (twenty-four hours)
	//
	// If you don't specify PeriodInSeconds, then Performance Insights will choose
	// a value for you, with a goal of returning roughly 100-200 data points in
	// the response.
	PeriodInSeconds *int64 `type:"integer"`

	// The AWS service for which Performance Insights will return metrics. The only
	// valid value for ServiceType is: RDS
	//
	// ServiceType is a required field
	ServiceType *string `type:"string" required:"true" enum:"ServiceType"`

	// The date and time specifying the beginning of the requested time series data.
	// You can't specify a StartTime that's earlier than 7 days ago. The value specified
	// is inclusive - data points equal to or greater than StartTime will be returned.
	//
	// The value for StartTime must be earlier than the value for EndTime.
	//
	// StartTime is a required field
	StartTime *time.Time `type:"timestamp" required:"true"`
}

// String returns the string representation
func (s DescribeDimensionKeysInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDimensionKeysInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDimensionKeysInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeDimensionKeysInput"}
	if s.EndTime == nil {
		invalidParams.Add(request.NewErrParamRequired("EndTime"))
	}
	if s.GroupBy == nil {
		invalidParams.Add(request.NewErrParamRequired("GroupBy"))
	}
	if s.Identifier == nil {
		invalidParams.Add(request.NewErrParamRequired("Identifier"))
	}
	if s.Metric == nil {
		invalidParams.Add(request.NewErrParamRequired("Metric"))
	}
	if s.ServiceType == nil {
		invalidParams.Add(request.NewErrParamRequired("ServiceType"))
	}
	if s.StartTime == nil {
		invalidParams.Add(request.NewErrParamRequired("StartTime"))
	}
	if s.GroupBy != nil {
		if err := s.GroupBy.Validate(); err != nil {
			invalidParams.AddNested("GroupBy", err.(request.ErrInvalidParams))
		}
	}
	if s.PartitionBy != nil {
		if err := s.PartitionBy.Validate(); err != nil {
			invalidParams.AddNested("PartitionBy", err.(request.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEndTime sets the EndTime field's value.
func (s *DescribeDimensionKeysInput) SetEndTime(v time.Time) *DescribeDimensionKeysInput {
	s.EndTime = &v
	return s
}

// SetFilter sets the Filter field's value.
func (s *DescribeDimensionKeysInput) SetFilter(v map[string]*string) *DescribeDimensionKeysInput {
	s.Filter = v
	return s
}

// SetGroupBy sets the GroupBy field's value.
func (s *DescribeDimensionKeysInput) SetGroupBy(v *DimensionGroup) *DescribeDimensionKeysInput {
	s.GroupBy = v
	return s
}

// SetIdentifier sets the Identifier field's value.
func (s *DescribeDimensionKeysInput) SetIdentifier(v string) *DescribeDimensionKeysInput {
	s.Identifier = &v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *DescribeDimensionKeysInput) SetMaxResults(v int64) *DescribeDimensionKeysInput {
	s.MaxResults = &v
	return s
}

// SetMetric sets the Metric field's value.
func (s *DescribeDimensionKeysInput) SetMetric(v string) *DescribeDimensionKeysInput {
	s.Metric = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeDimensionKeysInput) SetNextToken(v string) *DescribeDimensionKeysInput {
	s.NextToken = &v
	return s
}

// SetPartitionBy sets the PartitionBy field's value.
func (s *DescribeDimensionKeysInput) SetPartitionBy(v *DimensionGroup) *DescribeDimensionKeysInput {
	s.PartitionBy = v
	return s
}

// SetPeriodInSeconds sets the PeriodInSeconds field's value.
func (s *DescribeDimensionKeysInput) SetPeriodInSeconds(v int64) *DescribeDimensionKeysInput {
	s.PeriodInSeconds = &v
	return s
}

// SetServiceType sets the ServiceType field's value.
func (s *DescribeDimensionKeysInput) SetServiceType(v string) *DescribeDimensionKeysInput {
	s.ServiceType = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *DescribeDimensionKeysInput) SetStartTime(v time.Time) *DescribeDimensionKeysInput {
	s.StartTime = &v
	return s
}

type DescribeDimensionKeysOutput struct {
	_ struct{} `type:"structure"`

	// The end time for the returned dimension keys, after alignment to a granular
	// boundary (as specified by PeriodInSeconds). AlignedEndTime will be greater
	// than or equal to the value of the user-specified Endtime.
	AlignedEndTime *time.Time `type:"timestamp"`

	// The start time for the returned dimension keys, after alignment to a granular
	// boundary (as specified by PeriodInSeconds). AlignedStartTime will be less
	// than or equal to the value of the user-specified StartTime.
	AlignedStartTime *time.Time `type:"timestamp"`

	// The dimension keys that were requested.
	Keys []*DimensionKeyDescription `type:"list"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the token, up to
	// the value specified by MaxRecords.
	NextToken *string `type:"string"`

	// If PartitionBy was present in the request, PartitionKeys contains the breakdown
	// of dimension keys by the specified partitions.
	PartitionKeys []*ResponsePartitionKey `type:"list"`
}

// String returns the string representation
func (s DescribeDimensionKeysOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDimensionKeysOutput) GoString() string {
	return s.String()
}

// SetAlignedEndTime sets the AlignedEndTime field's value.
func (s *DescribeDimensionKeysOutput) SetAlignedEndTime(v time.Time) *DescribeDimensionKeysOutput {
	s.AlignedEndTime = &v
	return s
}

// SetAlignedStartTime sets the AlignedStartTime field's value.
func (s *DescribeDimensionKeysOutput) SetAlignedStartTime(v time.Time) *DescribeDimensionKeysOutput {
	s.AlignedStartTime = &v
	return s
}

// SetKeys sets the Keys field's value.
func (s *DescribeDimensionKeysOutput) SetKeys(v []*DimensionKeyDescription) *DescribeDimensionKeysOutput {
	s.Keys = v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeDimensionKeysOutput) SetNextToken(v string) *DescribeDimensionKeysOutput {
	s.NextToken = &v
	return s
}

// SetPartitionKeys sets the PartitionKeys field's value.
func (s *DescribeDimensionKeysOutput) SetPartitionKeys(v []*ResponsePartitionKey) *DescribeDimensionKeysOutput {
	s.PartitionKeys = v
	return s
}

// A logical grouping of Performance Insights metrics for a related subject
// area. For example, the db.sql dimension group consists of the following dimensions:
// db.sql.id, db.sql.db_id, db.sql.statement, and db.sql.tokenized_id.
type DimensionGroup struct {
	_ struct{} `type:"structure"`

	// A list of specific dimensions from a dimension group. If this parameter is
	// not present, then it signifies that all of the dimensions in the group were
	// requested, or are present in the response.
	//
	// Valid values for elements in the Dimensions array are:
	//
	//    * db.user.id
	//
	//    * db.user.name
	//
	//    * db.host.id
	//
	//    * db.host.name
	//
	//    * db.sql.id
	//
	//    * db.sql.db_id
	//
	//    * db.sql.statement
	//
	//    * db.sql.tokenized_id
	//
	//    * db.sql_tokenized.id
	//
	//    * db.sql_tokenized.db_id
	//
	//    * db.sql_tokenized.statement
	//
	//    * db.wait_event.name
	//
	//    * db.wait_event.type
	//
	//    * db.wait_event_type.name
	Dimensions []*string `min:"1" type:"list"`

	// The name of the dimension group. Valid values are:
	//
	//    * db.user
	//
	//    * db.host
	//
	//    * db.sql
	//
	//    * db.sql_tokenized
	//
	//    * db.wait_event
	//
	//    * db.wait_event_type
	//
	// Group is a required field
	Group *string `type:"string" required:"true"`

	// The maximum number of items to fetch for this dimension group.
	Limit *int64 `min:"1" type:"integer"`
}

// String returns the string representation
func (s DimensionGroup) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DimensionGroup) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DimensionGroup) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DimensionGroup"}
	if s.Dimensions != nil && len(s.Dimensions) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Dimensions", 1))
	}
	if s.Group == nil {
		invalidParams.Add(request.NewErrParamRequired("Group"))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(request.NewErrParamMinValue("Limit", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDimensions sets the Dimensions field's value.
func (s *DimensionGroup) SetDimensions(v []*string) *DimensionGroup {
	s.Dimensions = v
	return s
}

// SetGroup sets the Group field's value.
func (s *DimensionGroup) SetGroup(v string) *DimensionGroup {
	s.Group = &v
	return s
}

// SetLimit sets the Limit field's value.
func (s *DimensionGroup) SetLimit(v int64) *DimensionGroup {
	s.Limit = &v
	return s
}

// An array of descriptions and aggregated values for each dimension within
// a dimension group.
type DimensionKeyDescription struct {
	_ struct{} `type:"structure"`

	// A map of name-value pairs for the dimensions in the group.
	Dimensions map[string]*string `type:"map"`

	// If PartitionBy was specified, PartitionKeys contains the dimensions that
	// were.
	Partitions []*float64 `type:"list"`

	// The aggregated metric value for the dimension(s), over the requested time
	// range.
	Total *float64 `type:"double"`
}

// String returns the string representation
func (s DimensionKeyDescription) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DimensionKeyDescription) GoString() string {
	return s.String()
}

// SetDimensions sets the Dimensions field's value.
func (s *DimensionKeyDescription) SetDimensions(v map[string]*string) *DimensionKeyDescription {
	s.Dimensions = v
	return s
}

// SetPartitions sets the Partitions field's value.
func (s *DimensionKeyDescription) SetPartitions(v []*float64) *DimensionKeyDescription {
	s.Partitions = v
	return s
}

// SetTotal sets the Total field's value.
func (s *DimensionKeyDescription) SetTotal(v float64) *DimensionKeyDescription {
	s.Total = &v
	return s
}

type GetResourceMetricsInput struct {
	_ struct{} `type:"structure"`

	// The date and time specifiying the end of the requested time series data.
	// The value specified is exclusive - data points less than (but not equal to)
	// EndTime will be returned.
	//
	// The value for EndTime must be later than the value for StartTime.
	//
	// EndTime is a required field
	EndTime *time.Time `type:"timestamp" required:"true"`

	// An immutable, AWS Region-unique identifier for a data source. Performance
	// Insights gathers metrics from this data source.
	//
	// To use an Amazon RDS instance as a data source, you specify its DbiResourceId
	// value - for example: db-FAIHNTYBKTGAUSUZQYPDS2GW4A
	//
	// Identifier is a required field
	Identifier *string `type:"string" required:"true"`

	// The maximum number of items to return in the response. If more items exist
	// than the specified MaxRecords value, a pagination token is included in the
	// response so that the remaining results can be retrieved.
	MaxResults *int64 `type:"integer"`

	// An array of one or more queries to perform. Each query must specify a Performance
	// Insights metric, and can optionally specify aggregation and filtering criteria.
	//
	// MetricQueries is a required field
	MetricQueries []*MetricQuery `min:"1" type:"list" required:"true"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the token, up to
	// the value specified by MaxRecords.
	NextToken *string `type:"string"`

	// The granularity, in seconds, of the data points returned from Performance
	// Insights. A period can be as short as one second, or as long as one day (86400
	// seconds). Valid values are:
	//
	//    * 1 (one second)
	//
	//    * 60 (one minute)
	//
	//    * 300 (five minutes)
	//
	//    * 3600 (one hour)
	//
	//    * 86400 (twenty-four hours)
	//
	// If you don't specify PeriodInSeconds, then Performance Insights will choose
	// a value for you, with a goal of returning roughly 100-200 data points in
	// the response.
	PeriodInSeconds *int64 `type:"integer"`

	// The AWS service for which Performance Insights will return metrics. The only
	// valid value for ServiceType is: RDS
	//
	// ServiceType is a required field
	ServiceType *string `type:"string" required:"true" enum:"ServiceType"`

	// The date and time specifying the beginning of the requested time series data.
	// You can't specify a StartTime that's earlier than 7 days ago. The value specified
	// is inclusive - data points equal to or greater than StartTime will be returned.
	//
	// The value for StartTime must be earlier than the value for EndTime.
	//
	// StartTime is a required field
	StartTime *time.Time `type:"timestamp" required:"true"`
}

// String returns the string representation
func (s GetResourceMetricsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetResourceMetricsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetResourceMetricsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetResourceMetricsInput"}
	if s.EndTime == nil {
		invalidParams.Add(request.NewErrParamRequired("EndTime"))
	}
	if s.Identifier == nil {
		invalidParams.Add(request.NewErrParamRequired("Identifier"))
	}
	if s.MetricQueries == nil {
		invalidParams.Add(request.NewErrParamRequired("MetricQueries"))
	}
	if s.MetricQueries != nil && len(s.MetricQueries) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("MetricQueries", 1))
	}
	if s.ServiceType == nil {
		invalidParams.Add(request.NewErrParamRequired("ServiceType"))
	}
	if s.StartTime == nil {
		invalidParams.Add(request.NewErrParamRequired("StartTime"))
	}
	if s.MetricQueries != nil {
		for i, v := range s.MetricQueries {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "MetricQueries", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEndTime sets the EndTime field's value.
func (s *GetResourceMetricsInput) SetEndTime(v time.Time) *GetResourceMetricsInput {
	s.EndTime = &v
	return s
}

// SetIdentifier sets the Identifier field's value.
func (s *GetResourceMetricsInput) SetIdentifier(v string) *GetResourceMetricsInput {
	s.Identifier = &v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *GetResourceMetricsInput) SetMaxResults(v int64) *GetResourceMetricsInput {
	s.MaxResults = &v
	return s
}

// SetMetricQueries sets the MetricQueries field's value.
func (s *GetResourceMetricsInput) SetMetricQueries(v []*MetricQuery) *GetResourceMetricsInput {
	s.MetricQueries = v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *GetResourceMetricsInput) SetNextToken(v string) *GetResourceMetricsInput {
	s.NextToken = &v
	return s
}

// SetPeriodInSeconds sets the PeriodInSeconds field's value.
func (s *GetResourceMetricsInput) SetPeriodInSeconds(v int64) *GetResourceMetricsInput {
	s.PeriodInSeconds = &v
	return s
}

// SetServiceType sets the ServiceType field's value.
func (s *GetResourceMetricsInput) SetServiceType(v string) *GetResourceMetricsInput {
	s.ServiceType = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *GetResourceMetricsInput) SetStartTime(v time.Time) *GetResourceMetricsInput {
	s.StartTime = &v
	return s
}

type GetResourceMetricsOutput struct {
	_ struct{} `type:"structure"`

	// The end time for the returned metrics, after alignment to a granular boundary
	// (as specified by PeriodInSeconds). AlignedEndTime will be greater than or
	// equal to the value of the user-specified Endtime.
	AlignedEndTime *time.Time `type:"timestamp"`

	// The start time for the returned metrics, after alignment to a granular boundary
	// (as specified by PeriodInSeconds). AlignedStartTime will be less than or
	// equal to the value of the user-specified StartTime.
	AlignedStartTime *time.Time `type:"timestamp"`

	// An immutable, AWS Region-unique identifier for a data source. Performance
	// Insights gathers metrics from this data source.
	//
	// To use an Amazon RDS instance as a data source, you specify its DbiResourceId
	// value - for example: db-FAIHNTYBKTGAUSUZQYPDS2GW4A
	Identifier *string `type:"string"`

	// An array of metric results,, where each array element contains all of the
	// data points for a particular dimension.
	MetricList []*MetricKeyDataPoints `type:"list"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the token, up to
	// the value specified by MaxRecords.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetResourceMetricsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetResourceMetricsOutput) GoString() string {
	return s.String()
}

// SetAlignedEndTime sets the AlignedEndTime field's value.
func (s *GetResourceMetricsOutput) SetAlignedEndTime(v time.Time) *GetResourceMetricsOutput {
	s.AlignedEndTime = &v
	return s
}

// SetAlignedStartTime sets the AlignedStartTime field's value.
func (s *GetResourceMetricsOutput) SetAlignedStartTime(v time.Time) *GetResourceMetricsOutput {
	s.AlignedStartTime = &v
	return s
}

// SetIdentifier sets the Identifier field's value.
func (s *GetResourceMetricsOutput) SetIdentifier(v string) *GetResourceMetricsOutput {
	s.Identifier = &v
	return s
}

// SetMetricList sets the MetricList field's value.
func (s *GetResourceMetricsOutput) SetMetricList(v []*MetricKeyDataPoints) *GetResourceMetricsOutput {
	s.MetricList = v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *GetResourceMetricsOutput) SetNextToken(v string) *GetResourceMetricsOutput {
	s.NextToken = &v
	return s
}

// A time-ordered series of data points, correpsonding to a dimension of a Performance
// Insights metric.
type MetricKeyDataPoints struct {
	_ struct{} `type:"structure"`

	// An array of timestamp-value pairs, representing measurements over a period
	// of time.
	DataPoints []*DataPoint `type:"list"`

	// The dimension(s) to which the data points apply.
	Key *ResponseResourceMetricKey `type:"structure"`
}

// String returns the string representation
func (s MetricKeyDataPoints) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s MetricKeyDataPoints) GoString() string {
	return s.String()
}

// SetDataPoints sets the DataPoints field's value.
func (s *MetricKeyDataPoints) SetDataPoints(v []*DataPoint) *MetricKeyDataPoints {
	s.DataPoints = v
	return s
}

// SetKey sets the Key field's value.
func (s *MetricKeyDataPoints) SetKey(v *ResponseResourceMetricKey) *MetricKeyDataPoints {
	s.Key = v
	return s
}

// A single query to be processed. You must provide the metric to query. If
// no other parameters are specified, Performance Insights returns all of the
// data points for that metric. You can optionally request that the data points
// be aggregated by dimension group ( GroupBy), and return only those data points
// that match your criteria (Filter).
type MetricQuery struct {
	_ struct{} `type:"structure"`

	// One or more filters to apply in the request. Restrictions:
	//
	//    * Any number of filters by the same dimension, as specified in the GroupBy
	//    parameter.
	//
	//    * A single filter for any other dimension in this dimension group.
	Filter map[string]*string `type:"map"`

	// A specification for how to aggregate the data points from a query result.
	// You must specify a valid dimension group. Performance Insights will return
	// all of the dimensions within that group, unless you provide the names of
	// specific dimensions within that group. You can also request that Performance
	// Insights return a limited number of values for a dimension.
	GroupBy *DimensionGroup `type:"structure"`

	// The name of a Performance Insights metric to be measured.
	//
	// Valid values for Metric are:
	//
	//    * db.load.avg - a scaled representation of the number of active sessions
	//    for the database engine.
	//
	//    * db.sampledload.avg - the raw number of active sessions for the database
	//    engine.
	//
	// Metric is a required field
	Metric *string `type:"string" required:"true"`
}

// String returns the string representation
func (s MetricQuery) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s MetricQuery) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *MetricQuery) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "MetricQuery"}
	if s.Metric == nil {
		invalidParams.Add(request.NewErrParamRequired("Metric"))
	}
	if s.GroupBy != nil {
		if err := s.GroupBy.Validate(); err != nil {
			invalidParams.AddNested("GroupBy", err.(request.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetFilter sets the Filter field's value.
func (s *MetricQuery) SetFilter(v map[string]*string) *MetricQuery {
	s.Filter = v
	return s
}

// SetGroupBy sets the GroupBy field's value.
func (s *MetricQuery) SetGroupBy(v *DimensionGroup) *MetricQuery {
	s.GroupBy = v
	return s
}

// SetMetric sets the Metric field's value.
func (s *MetricQuery) SetMetric(v string) *MetricQuery {
	s.Metric = &v
	return s
}

// If PartitionBy was specified in a DescribeDimensionKeys request, the dimensions
// are returned in an array. Each element in the array specifies one dimension.
type ResponsePartitionKey struct {
	_ struct{} `type:"structure"`

	// A dimension map that contains the dimension(s) for this partition.
	//
	// Dimensions is a required field
	Dimensions map[string]*string `type:"map" required:"true"`
}

// String returns the string representation
func (s ResponsePartitionKey) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ResponsePartitionKey) GoString() string {
	return s.String()
}

// SetDimensions sets the Dimensions field's value.
func (s *ResponsePartitionKey) SetDimensions(v map[string]*string) *ResponsePartitionKey {
	s.Dimensions = v
	return s
}

// An object describing a Performance Insights metric and one or more dimensions
// for that metric.
type ResponseResourceMetricKey struct {
	_ struct{} `type:"structure"`

	// The valid dimensions for the metric.
	Dimensions map[string]*string `type:"map"`

	// The name of a Performance Insights metric to be measured.
	//
	// Valid values for Metric are:
	//
	//    * db.load.avg - a scaled representation of the number of active sessions
	//    for the database engine.
	//
	//    * db.sampledload.avg - the raw number of active sessions for the database
	//    engine.
	//
	// Metric is a required field
	Metric *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ResponseResourceMetricKey) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ResponseResourceMetricKey) GoString() string {
	return s.String()
}

// SetDimensions sets the Dimensions field's value.
func (s *ResponseResourceMetricKey) SetDimensions(v map[string]*string) *ResponseResourceMetricKey {
	s.Dimensions = v
	return s
}

// SetMetric sets the Metric field's value.
func (s *ResponseResourceMetricKey) SetMetric(v string) *ResponseResourceMetricKey {
	s.Metric = &v
	return s
}

const (
	// ServiceTypeRds is a ServiceType enum value
	ServiceTypeRds = "RDS"
)
