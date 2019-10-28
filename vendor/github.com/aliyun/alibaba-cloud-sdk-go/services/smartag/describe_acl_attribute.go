package smartag

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeACLAttribute invokes the smartag.DescribeACLAttribute API synchronously
// api document: https://help.aliyun.com/api/smartag/describeaclattribute.html
func (client *Client) DescribeACLAttribute(request *DescribeACLAttributeRequest) (response *DescribeACLAttributeResponse, err error) {
	response = CreateDescribeACLAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeACLAttributeWithChan invokes the smartag.DescribeACLAttribute API asynchronously
// api document: https://help.aliyun.com/api/smartag/describeaclattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeACLAttributeWithChan(request *DescribeACLAttributeRequest) (<-chan *DescribeACLAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeACLAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeACLAttribute(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeACLAttributeWithCallback invokes the smartag.DescribeACLAttribute API asynchronously
// api document: https://help.aliyun.com/api/smartag/describeaclattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeACLAttributeWithCallback(request *DescribeACLAttributeRequest, callback func(response *DescribeACLAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeACLAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeACLAttribute(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeACLAttributeRequest is the request struct for api DescribeACLAttribute
type DescribeACLAttributeRequest struct {
	*requests.RpcRequest
	AclId                string           `position:"Query" name:"AclId"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	Direction            string           `position:"Query" name:"Direction"`
	Order                string           `position:"Query" name:"Order"`
}

// DescribeACLAttributeResponse is the response struct for api DescribeACLAttribute
type DescribeACLAttributeResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	PageNumber int    `json:"PageNumber" xml:"PageNumber"`
	PageSize   int    `json:"PageSize" xml:"PageSize"`
	Acrs       Acrs   `json:"Acrs" xml:"Acrs"`
}

// CreateDescribeACLAttributeRequest creates a request to invoke DescribeACLAttribute API
func CreateDescribeACLAttributeRequest() (request *DescribeACLAttributeRequest) {
	request = &DescribeACLAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DescribeACLAttribute", "smartag", "openAPI")
	return
}

// CreateDescribeACLAttributeResponse creates a response to parse from DescribeACLAttribute response
func CreateDescribeACLAttributeResponse() (response *DescribeACLAttributeResponse) {
	response = &DescribeACLAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}