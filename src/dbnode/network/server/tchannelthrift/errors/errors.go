// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package errors

import (
	"fmt"

	"github.com/m3db/m3/src/dbnode/generated/thrift/rpc"
)

func newError(errType rpc.ErrorType, errCode rpc.ErrorCode, err error) *rpc.Error {
	rpcErr := rpc.NewError()
	rpcErr.Type = errType
	rpcErr.Code = errCode
	rpcErr.Message = fmt.Sprintf("%v", err)
	return rpcErr
}

// IsInternalError returns whether the error is an internal error
func IsInternalError(err *rpc.Error) bool {
	return err != nil && err.Type == rpc.ErrorType_INTERNAL_ERROR
}

// IsBadRequestError returns whether the error is a bad request error
func IsBadRequestError(err *rpc.Error) bool {
	return err != nil && err.Type == rpc.ErrorType_BAD_REQUEST
}

// IsResourceExhaustedError returns whether the error is a resource exhausted error.
func IsResourceExhaustedError(err *rpc.Error) bool {
	return err != nil && err.Code == rpc.ErrorCode_RESOURCE_EXHAUSTED
}

// NewInternalError creates a new internal error
func NewInternalError(err error) *rpc.Error {
	return newError(rpc.ErrorType_INTERNAL_ERROR, rpc.ErrorCode_NONE, err)
}

// NewBadRequestError creates a new bad request error
func NewBadRequestError(err error) *rpc.Error {
	return newError(rpc.ErrorType_BAD_REQUEST, rpc.ErrorCode_NONE, err)
}

// NewResourceExhaustedError creates a new resource exhausted error.
func NewResourceExhaustedError(err error) *rpc.Error {
	// NB: To maintain better backwards compatibility, using BAD_REQUEST error type coupled with
	// RESOURCE_EXHAUSTED error code
	return newError(rpc.ErrorType_BAD_REQUEST, rpc.ErrorCode_RESOURCE_EXHAUSTED, err)
}

// NewWriteBatchRawError creates a new write batch error
func NewWriteBatchRawError(index int, err error) *rpc.WriteBatchRawError {
	batchErr := rpc.NewWriteBatchRawError()
	batchErr.Index = int64(index)
	batchErr.Err = NewInternalError(err)
	return batchErr
}

// NewBadRequestWriteBatchRawError creates a new bad request write batch error
func NewBadRequestWriteBatchRawError(index int, err error) *rpc.WriteBatchRawError {
	batchErr := rpc.NewWriteBatchRawError()
	batchErr.Index = int64(index)
	batchErr.Err = NewBadRequestError(err)
	return batchErr
}
