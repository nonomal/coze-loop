// Code generated by Kitex v0.13.1. DO NOT EDIT.

package fileservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	file "github.com/coze-dev/coze-loop/backend/kitex_gen/coze/loop/foundation/file"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"UploadLoopFileInner": kitex.NewMethodInfo(
		uploadLoopFileInnerHandler,
		newFileServiceUploadLoopFileInnerArgs,
		newFileServiceUploadLoopFileInnerResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"SignUploadFile": kitex.NewMethodInfo(
		signUploadFileHandler,
		newFileServiceSignUploadFileArgs,
		newFileServiceSignUploadFileResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"SignDownloadFile": kitex.NewMethodInfo(
		signDownloadFileHandler,
		newFileServiceSignDownloadFileArgs,
		newFileServiceSignDownloadFileResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	fileServiceServiceInfo = NewServiceInfo()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return fileServiceServiceInfo
}

// NewServiceInfo creates a new ServiceInfo
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo()
}

func newServiceInfo() *kitex.ServiceInfo {
	serviceName := "FileService"
	handlerType := (*file.FileService)(nil)
	extra := map[string]interface{}{
		"PackageName": "file",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         serviceMethods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.13.1",
		Extra:           extra,
	}
	return svcInfo
}

func uploadLoopFileInnerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*file.FileServiceUploadLoopFileInnerArgs)
	realResult := result.(*file.FileServiceUploadLoopFileInnerResult)
	success, err := handler.(file.FileService).UploadLoopFileInner(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newFileServiceUploadLoopFileInnerArgs() interface{} {
	return file.NewFileServiceUploadLoopFileInnerArgs()
}

func newFileServiceUploadLoopFileInnerResult() interface{} {
	return file.NewFileServiceUploadLoopFileInnerResult()
}

func signUploadFileHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*file.FileServiceSignUploadFileArgs)
	realResult := result.(*file.FileServiceSignUploadFileResult)
	success, err := handler.(file.FileService).SignUploadFile(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newFileServiceSignUploadFileArgs() interface{} {
	return file.NewFileServiceSignUploadFileArgs()
}

func newFileServiceSignUploadFileResult() interface{} {
	return file.NewFileServiceSignUploadFileResult()
}

func signDownloadFileHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*file.FileServiceSignDownloadFileArgs)
	realResult := result.(*file.FileServiceSignDownloadFileResult)
	success, err := handler.(file.FileService).SignDownloadFile(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newFileServiceSignDownloadFileArgs() interface{} {
	return file.NewFileServiceSignDownloadFileArgs()
}

func newFileServiceSignDownloadFileResult() interface{} {
	return file.NewFileServiceSignDownloadFileResult()
}

type kClient struct {
	c  client.Client
	sc client.Streaming
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c:  c,
		sc: c.(client.Streaming),
	}
}

func (p *kClient) UploadLoopFileInner(ctx context.Context, req *file.UploadLoopFileInnerRequest) (r *file.UploadLoopFileInnerResponse, err error) {
	var _args file.FileServiceUploadLoopFileInnerArgs
	_args.Req = req
	var _result file.FileServiceUploadLoopFileInnerResult
	if err = p.c.Call(ctx, "UploadLoopFileInner", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SignUploadFile(ctx context.Context, req *file.SignUploadFileRequest) (r *file.SignUploadFileResponse, err error) {
	var _args file.FileServiceSignUploadFileArgs
	_args.Req = req
	var _result file.FileServiceSignUploadFileResult
	if err = p.c.Call(ctx, "SignUploadFile", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SignDownloadFile(ctx context.Context, req *file.SignDownloadFileRequest) (r *file.SignDownloadFileResponse, err error) {
	var _args file.FileServiceSignDownloadFileArgs
	_args.Req = req
	var _result file.FileServiceSignDownloadFileResult
	if err = p.c.Call(ctx, "SignDownloadFile", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
