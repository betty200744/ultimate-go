/*
 *
 * Copyright 2017 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package mock_helloworld_test

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
	"gopkg.in/go-playground/assert.v1"
	"testing"
	"ultimate-go/awesome-go/grpc-go-mock/helloworld"
	"ultimate-go/awesome-go/grpc-go-mock/mock_helloworld"
)

/*type s struct {
	grpctest.Tester
}

func Test(t *testing.T) {
	grpctest.RunSubTests(t, s{})
}*/

// rpcMsg implements the gomock.Matcher interface
type rpcMsg struct {
	msg proto.Message
}

func (r *rpcMsg) Matches(msg interface{}) bool {
	m, ok := msg.(proto.Message)
	if !ok {
		return false
	}
	return proto.Equal(m, r.msg)
}

func (r *rpcMsg) String() string {
	return fmt.Sprintf("is %s", r.msg)
}

func TestMockGreeterServer_SayHello(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockGreeterClient := mock_helloworld.NewMockGreeterClient(ctrl)
	// 单纯mock调用
	mockGreeterClient.EXPECT().SayHello(gomock.Any(), gomock.Any())
	r, _ := mockGreeterClient.SayHello(context.Background(), &helloworld.HelloRequest{Name: "betty"})
	assert.Equal(t, r, nil)
}

func TestMockGreeterServer_SayHello_Expect(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockGreeterClient := mock_helloworld.NewMockGreeterClient(ctrl)
	ctx := context.Background()
	req := &helloworld.HelloRequest{Name: "unit_test"}
	// ctx, req注入
	mockGreeterClient.EXPECT().SayHello(ctx, req)
	r, _ := mockGreeterClient.SayHello(context.Background(), &helloworld.HelloRequest{Name: "unit_test"})
	assert.Equal(t, r, nil)
}

func TestMockGreeterServer_SayHello_Do(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockGreeterClient := mock_helloworld.NewMockGreeterClient(ctrl)
	// Do, 注入动作，但无返回
	mockGreeterClient.EXPECT().SayHello(gomock.Any(), gomock.Any()).Do(func(arg1, arg2 interface{}) { fmt.Println("SayHello arg1", arg1) })
	r, _ := mockGreeterClient.SayHello(context.Background(), &helloworld.HelloRequest{Name: "betty"})
	assert.Equal(t, r, nil)
}

func TestMockGreeterServer_SayHello_Return(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockGreeterClient := mock_helloworld.NewMockGreeterClient(ctrl)
	ctx := context.Background()
	req := &helloworld.HelloRequest{Name: "unit_test"}
	// Return, 注入返回
	mockGreeterClient.EXPECT().SayHello(ctx, req).Return(&helloworld.HelloReply{Message: "Mocked Interface"}, nil)
	r, _ := mockGreeterClient.SayHello(context.Background(), &helloworld.HelloRequest{Name: "unit_test"})
	assert.Equal(t, r.Message, "Mocked Interface")
}

func TestMockGreeterServer_SayHello_DoAndReturn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockGreeterClient := mock_helloworld.NewMockGreeterClient(ctrl)
	// DoAndReturn, 注入动作及返回
	mockGreeterClient.EXPECT().SayHello(gomock.Any(), gomock.Any()).DoAndReturn(func(a interface{}, b interface{}) (*helloworld.HelloReply, error) {
		fmt.Println("SayHello arg1", a)
		fmt.Println("SayHello arg2", b)
		return &helloworld.HelloReply{Message: "Mocked Interface"}, nil
	})
	r, _ := mockGreeterClient.SayHello(context.Background(), &helloworld.HelloRequest{Name: "betty"})
	assert.Equal(t, r.Message, "Mocked Interface")
}
