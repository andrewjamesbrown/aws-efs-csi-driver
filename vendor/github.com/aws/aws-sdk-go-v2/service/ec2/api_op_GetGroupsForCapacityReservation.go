// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the resource groups to which a Capacity Reservation has been added.
func (c *Client) GetGroupsForCapacityReservation(ctx context.Context, params *GetGroupsForCapacityReservationInput, optFns ...func(*Options)) (*GetGroupsForCapacityReservationOutput, error) {
	if params == nil {
		params = &GetGroupsForCapacityReservationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetGroupsForCapacityReservation", params, optFns, c.addOperationGetGroupsForCapacityReservationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetGroupsForCapacityReservationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetGroupsForCapacityReservationInput struct {

	// The ID of the Capacity Reservation. If you specify a Capacity Reservation that
	// is shared with you, the operation returns only Capacity Reservation groups that
	// you own.
	//
	// This member is required.
	CapacityReservationId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see [Pagination].
	//
	// [Pagination]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination
	MaxResults *int32

	// The token to use to retrieve the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetGroupsForCapacityReservationOutput struct {

	// Information about the resource groups to which the Capacity Reservation has
	// been added.
	CapacityReservationGroups []types.CapacityReservationGroup

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetGroupsForCapacityReservationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpGetGroupsForCapacityReservation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpGetGroupsForCapacityReservation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetGroupsForCapacityReservation"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpGetGroupsForCapacityReservationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetGroupsForCapacityReservation(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

// GetGroupsForCapacityReservationPaginatorOptions is the paginator options for
// GetGroupsForCapacityReservation
type GetGroupsForCapacityReservationPaginatorOptions struct {
	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see [Pagination].
	//
	// [Pagination]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetGroupsForCapacityReservationPaginator is a paginator for
// GetGroupsForCapacityReservation
type GetGroupsForCapacityReservationPaginator struct {
	options   GetGroupsForCapacityReservationPaginatorOptions
	client    GetGroupsForCapacityReservationAPIClient
	params    *GetGroupsForCapacityReservationInput
	nextToken *string
	firstPage bool
}

// NewGetGroupsForCapacityReservationPaginator returns a new
// GetGroupsForCapacityReservationPaginator
func NewGetGroupsForCapacityReservationPaginator(client GetGroupsForCapacityReservationAPIClient, params *GetGroupsForCapacityReservationInput, optFns ...func(*GetGroupsForCapacityReservationPaginatorOptions)) *GetGroupsForCapacityReservationPaginator {
	if params == nil {
		params = &GetGroupsForCapacityReservationInput{}
	}

	options := GetGroupsForCapacityReservationPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetGroupsForCapacityReservationPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetGroupsForCapacityReservationPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetGroupsForCapacityReservation page.
func (p *GetGroupsForCapacityReservationPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetGroupsForCapacityReservationOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.GetGroupsForCapacityReservation(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// GetGroupsForCapacityReservationAPIClient is a client that implements the
// GetGroupsForCapacityReservation operation.
type GetGroupsForCapacityReservationAPIClient interface {
	GetGroupsForCapacityReservation(context.Context, *GetGroupsForCapacityReservationInput, ...func(*Options)) (*GetGroupsForCapacityReservationOutput, error)
}

var _ GetGroupsForCapacityReservationAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opGetGroupsForCapacityReservation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetGroupsForCapacityReservation",
	}
}