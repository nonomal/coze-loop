// Code generated by Kitex v0.13.1. DO NOT EDIT.

package evaluatorservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	evaluator "github.com/coze-dev/coze-loop/backend/kitex_gen/coze/loop/evaluation/evaluator"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	ListEvaluators(ctx context.Context, request *evaluator.ListEvaluatorsRequest, callOptions ...callopt.Option) (r *evaluator.ListEvaluatorsResponse, err error)
	BatchGetEvaluators(ctx context.Context, request *evaluator.BatchGetEvaluatorsRequest, callOptions ...callopt.Option) (r *evaluator.BatchGetEvaluatorsResponse, err error)
	GetEvaluator(ctx context.Context, request *evaluator.GetEvaluatorRequest, callOptions ...callopt.Option) (r *evaluator.GetEvaluatorResponse, err error)
	CreateEvaluator(ctx context.Context, request *evaluator.CreateEvaluatorRequest, callOptions ...callopt.Option) (r *evaluator.CreateEvaluatorResponse, err error)
	UpdateEvaluator(ctx context.Context, request *evaluator.UpdateEvaluatorRequest, callOptions ...callopt.Option) (r *evaluator.UpdateEvaluatorResponse, err error)
	UpdateEvaluatorDraft(ctx context.Context, request *evaluator.UpdateEvaluatorDraftRequest, callOptions ...callopt.Option) (r *evaluator.UpdateEvaluatorDraftResponse, err error)
	DeleteEvaluator(ctx context.Context, request *evaluator.DeleteEvaluatorRequest, callOptions ...callopt.Option) (r *evaluator.DeleteEvaluatorResponse, err error)
	CheckEvaluatorName(ctx context.Context, request *evaluator.CheckEvaluatorNameRequest, callOptions ...callopt.Option) (r *evaluator.CheckEvaluatorNameResponse, err error)
	ListEvaluatorVersions(ctx context.Context, request *evaluator.ListEvaluatorVersionsRequest, callOptions ...callopt.Option) (r *evaluator.ListEvaluatorVersionsResponse, err error)
	GetEvaluatorVersion(ctx context.Context, request *evaluator.GetEvaluatorVersionRequest, callOptions ...callopt.Option) (r *evaluator.GetEvaluatorVersionResponse, err error)
	BatchGetEvaluatorVersions(ctx context.Context, request *evaluator.BatchGetEvaluatorVersionsRequest, callOptions ...callopt.Option) (r *evaluator.BatchGetEvaluatorVersionsResponse, err error)
	SubmitEvaluatorVersion(ctx context.Context, request *evaluator.SubmitEvaluatorVersionRequest, callOptions ...callopt.Option) (r *evaluator.SubmitEvaluatorVersionResponse, err error)
	ListTemplates(ctx context.Context, request *evaluator.ListTemplatesRequest, callOptions ...callopt.Option) (r *evaluator.ListTemplatesResponse, err error)
	GetTemplateInfo(ctx context.Context, request *evaluator.GetTemplateInfoRequest, callOptions ...callopt.Option) (r *evaluator.GetTemplateInfoResponse, err error)
	GetDefaultPromptEvaluatorTools(ctx context.Context, req *evaluator.GetDefaultPromptEvaluatorToolsRequest, callOptions ...callopt.Option) (r *evaluator.GetDefaultPromptEvaluatorToolsResponse, err error)
	RunEvaluator(ctx context.Context, req *evaluator.RunEvaluatorRequest, callOptions ...callopt.Option) (r *evaluator.RunEvaluatorResponse, err error)
	DebugEvaluator(ctx context.Context, req *evaluator.DebugEvaluatorRequest, callOptions ...callopt.Option) (r *evaluator.DebugEvaluatorResponse, err error)
	UpdateEvaluatorRecord(ctx context.Context, req *evaluator.UpdateEvaluatorRecordRequest, callOptions ...callopt.Option) (r *evaluator.UpdateEvaluatorRecordResponse, err error)
	GetEvaluatorRecord(ctx context.Context, req *evaluator.GetEvaluatorRecordRequest, callOptions ...callopt.Option) (r *evaluator.GetEvaluatorRecordResponse, err error)
	BatchGetEvaluatorRecords(ctx context.Context, req *evaluator.BatchGetEvaluatorRecordsRequest, callOptions ...callopt.Option) (r *evaluator.BatchGetEvaluatorRecordsResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kEvaluatorServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kEvaluatorServiceClient struct {
	*kClient
}

func (p *kEvaluatorServiceClient) ListEvaluators(ctx context.Context, request *evaluator.ListEvaluatorsRequest, callOptions ...callopt.Option) (r *evaluator.ListEvaluatorsResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListEvaluators(ctx, request)
}

func (p *kEvaluatorServiceClient) BatchGetEvaluators(ctx context.Context, request *evaluator.BatchGetEvaluatorsRequest, callOptions ...callopt.Option) (r *evaluator.BatchGetEvaluatorsResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.BatchGetEvaluators(ctx, request)
}

func (p *kEvaluatorServiceClient) GetEvaluator(ctx context.Context, request *evaluator.GetEvaluatorRequest, callOptions ...callopt.Option) (r *evaluator.GetEvaluatorResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetEvaluator(ctx, request)
}

func (p *kEvaluatorServiceClient) CreateEvaluator(ctx context.Context, request *evaluator.CreateEvaluatorRequest, callOptions ...callopt.Option) (r *evaluator.CreateEvaluatorResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateEvaluator(ctx, request)
}

func (p *kEvaluatorServiceClient) UpdateEvaluator(ctx context.Context, request *evaluator.UpdateEvaluatorRequest, callOptions ...callopt.Option) (r *evaluator.UpdateEvaluatorResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateEvaluator(ctx, request)
}

func (p *kEvaluatorServiceClient) UpdateEvaluatorDraft(ctx context.Context, request *evaluator.UpdateEvaluatorDraftRequest, callOptions ...callopt.Option) (r *evaluator.UpdateEvaluatorDraftResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateEvaluatorDraft(ctx, request)
}

func (p *kEvaluatorServiceClient) DeleteEvaluator(ctx context.Context, request *evaluator.DeleteEvaluatorRequest, callOptions ...callopt.Option) (r *evaluator.DeleteEvaluatorResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteEvaluator(ctx, request)
}

func (p *kEvaluatorServiceClient) CheckEvaluatorName(ctx context.Context, request *evaluator.CheckEvaluatorNameRequest, callOptions ...callopt.Option) (r *evaluator.CheckEvaluatorNameResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CheckEvaluatorName(ctx, request)
}

func (p *kEvaluatorServiceClient) ListEvaluatorVersions(ctx context.Context, request *evaluator.ListEvaluatorVersionsRequest, callOptions ...callopt.Option) (r *evaluator.ListEvaluatorVersionsResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListEvaluatorVersions(ctx, request)
}

func (p *kEvaluatorServiceClient) GetEvaluatorVersion(ctx context.Context, request *evaluator.GetEvaluatorVersionRequest, callOptions ...callopt.Option) (r *evaluator.GetEvaluatorVersionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetEvaluatorVersion(ctx, request)
}

func (p *kEvaluatorServiceClient) BatchGetEvaluatorVersions(ctx context.Context, request *evaluator.BatchGetEvaluatorVersionsRequest, callOptions ...callopt.Option) (r *evaluator.BatchGetEvaluatorVersionsResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.BatchGetEvaluatorVersions(ctx, request)
}

func (p *kEvaluatorServiceClient) SubmitEvaluatorVersion(ctx context.Context, request *evaluator.SubmitEvaluatorVersionRequest, callOptions ...callopt.Option) (r *evaluator.SubmitEvaluatorVersionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SubmitEvaluatorVersion(ctx, request)
}

func (p *kEvaluatorServiceClient) ListTemplates(ctx context.Context, request *evaluator.ListTemplatesRequest, callOptions ...callopt.Option) (r *evaluator.ListTemplatesResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListTemplates(ctx, request)
}

func (p *kEvaluatorServiceClient) GetTemplateInfo(ctx context.Context, request *evaluator.GetTemplateInfoRequest, callOptions ...callopt.Option) (r *evaluator.GetTemplateInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetTemplateInfo(ctx, request)
}

func (p *kEvaluatorServiceClient) GetDefaultPromptEvaluatorTools(ctx context.Context, req *evaluator.GetDefaultPromptEvaluatorToolsRequest, callOptions ...callopt.Option) (r *evaluator.GetDefaultPromptEvaluatorToolsResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetDefaultPromptEvaluatorTools(ctx, req)
}

func (p *kEvaluatorServiceClient) RunEvaluator(ctx context.Context, req *evaluator.RunEvaluatorRequest, callOptions ...callopt.Option) (r *evaluator.RunEvaluatorResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RunEvaluator(ctx, req)
}

func (p *kEvaluatorServiceClient) DebugEvaluator(ctx context.Context, req *evaluator.DebugEvaluatorRequest, callOptions ...callopt.Option) (r *evaluator.DebugEvaluatorResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DebugEvaluator(ctx, req)
}

func (p *kEvaluatorServiceClient) UpdateEvaluatorRecord(ctx context.Context, req *evaluator.UpdateEvaluatorRecordRequest, callOptions ...callopt.Option) (r *evaluator.UpdateEvaluatorRecordResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateEvaluatorRecord(ctx, req)
}

func (p *kEvaluatorServiceClient) GetEvaluatorRecord(ctx context.Context, req *evaluator.GetEvaluatorRecordRequest, callOptions ...callopt.Option) (r *evaluator.GetEvaluatorRecordResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetEvaluatorRecord(ctx, req)
}

func (p *kEvaluatorServiceClient) BatchGetEvaluatorRecords(ctx context.Context, req *evaluator.BatchGetEvaluatorRecordsRequest, callOptions ...callopt.Option) (r *evaluator.BatchGetEvaluatorRecordsResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.BatchGetEvaluatorRecords(ctx, req)
}
