/*
 *   Copyright 2016 Manik Surtani
 *
 *   Licensed under the Apache License, Version 2.0 (the "License");
 *   you may not use this file except in compliance with the License.
 *   You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *   Unless required by applicable law or agreed to in writing, software
 *   distributed under the License is distributed on an "AS IS" BASIS,
 *   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *   See the License for the specific language governing permissions and
 *   limitations under the License.
 */

package grpc
import (
	"net"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"golang.org/x/net/context"
	"github.com/maniksurtani/quotaservice/logging"
	"github.com/maniksurtani/quotaservice"
	qspb "github.com/maniksurtani/quotaservice/protos"
	"github.com/maniksurtani/quotaservice/lifecycle"
	"github.com/maniksurtani/quotaservice/configs"

)

// gRPC-backed implementation of an RPC endpoint
type GrpcEndpoint struct {
	cfgs          *configs.Configs
	grpcServer    *grpc.Server
	currentStatus lifecycle.Status
	qs            quotaservice.QuotaService
}

func (this *GrpcEndpoint) Init(cfgs *configs.Configs, qs quotaservice.QuotaService) {
	this.cfgs = cfgs
	this.qs = qs
}

func (this *GrpcEndpoint) Start() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", this.cfgs.Port))
	if err != nil {
		logging.Fatalf("Cannot start server on port %v. Error %v", this.cfgs.Port, err)
		panic(fmt.Sprintf("Cannot start server on port %v. Error %v", this.cfgs.Port, err))
	}

	grpclog.SetLogger(logging.GetLogger())
	this.grpcServer = grpc.NewServer()
	// Each service should be registered
	qspb.RegisterQuotaServiceServer(this.grpcServer, this)
	go this.grpcServer.Serve(lis)
	this.currentStatus = lifecycle.Started
	logging.Printf("Starting server on port %v", this.cfgs.Port)
	logging.Printf("Server status: %v", this.currentStatus)

}

func (this *GrpcEndpoint) Stop() {
	this.currentStatus = lifecycle.Stopped
}

func (this *GrpcEndpoint) Allow(ctx context.Context, req *qspb.AllowRequest) (*qspb.AllowResponse, error) {
	rsp := new(qspb.AllowResponse)
	// TODO(manik) validate inputs
	granted, err := this.qs.Allow(req.BucketName, int(req.TokensRequested), convert(req.EmptyBucketPolicy))

	if err != nil {
		if qsErr, ok := err.(quotaservice.QuotaServiceError); ok {
			switch qsErr.Reason {
			case quotaservice.ER_NO_SUCH_BUCKET:
				rsp.Status = qspb.AllowResponse_REJECTED
			case quotaservice.ER_REJECTED:
				rsp.Status = qspb.AllowResponse_REJECTED
			case quotaservice.ER_TIMED_OUT_WAITING:
				rsp.Status = qspb.AllowResponse_TIMED_OUT
			}
		} else {
			return nil, err
		}
	} else {
		rsp.Status = qspb.AllowResponse_OK
		rsp.TokensGranted = int32(granted)
	}
	return rsp, nil
}

func convert(o qspb.AllowRequest_EmptyBucketPolicyOverride) quotaservice.EmptyBucketPolicyOverride {
	switch o {
	case qspb.AllowRequest_REJECT:
		return quotaservice.EBP_REJECT
	case qspb.AllowRequest_WAIT:
		return quotaservice.EBP_WAIT
	case qspb.AllowRequest_SERVER_DEFAULTS:
		return quotaservice.EBP_SERVER_DEFAULTS
	default:
		panic(fmt.Sprintf("Unknown enum value %+v", o))
	}
}
