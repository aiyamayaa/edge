// Copyright 2019 Intel Corporation, Inc. All rights reserved
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ngcaf

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Linger please
var (
	_ context.Context
)

// TrafficInfluenceSubscriptionGetAllAPIService type
type TrafficInfluenceSubscriptionGetAllAPIService service

func (a *TrafficInfluenceSubscriptionGetAllAPIService) handleGetAllResponse(
	ts *[]TrafficInfluSub, r *http.Response, body []byte) error {

	if r.StatusCode == 200 {
		err := json.Unmarshal(body, ts)
		if err != nil {
			log.Errf("Error decoding response body %s, ", err.Error())
		}
		return err
	}

	return handleGetErrorResp(r, body)
}

/*
SubscriptionsGetAll read all of the active
subscriptions for the AF
read all of the active subscriptions for the AF
 * @param ctx context.Context - for authentication, logging, cancellation,
 * deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param afID Identifier of the AF

@return []TrafficInfluSub
*/
func (a *TrafficInfluenceSubscriptionGetAllAPIService) SubscriptionsGetAll(
	ctx context.Context, afID string) ([]TrafficInfluSub,
	*http.Response, error) {
	var (
		method  = strings.ToUpper("Get")
		getBody interface{}
		ret     []TrafficInfluSub
	)

	// create path and map variables
	path := a.client.cfg.NEFBasePath + "/{afId}/subscriptions"
	path = strings.Replace(path, "{"+"afId"+"}",
		fmt.Sprintf("%v", afID), -1)

	headerParams := make(map[string]string)
	// to determine the Content-Type header
	contentTypes := []string{"application/json"}
	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headerParams["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{"application/json"}
	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headerParams["Accept"] = headerAccept
	}
	r, err := a.client.prepareRequest(ctx, path, method,
		getBody, headerParams)

	if err != nil {
		return ret, nil, err
	}

	resp, err := a.client.callAPI(r)
	if err != nil || resp == nil {
		return ret, resp, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Errf("response body was not closed properly")
		}
	}()

	if err != nil {
		log.Errf("http response body could not be read")
		return ret, resp, err
	}

	if err = a.handleGetAllResponse(&ret, resp,
		respBody); err != nil {
		return ret, resp, err
	}

	return ret, resp, nil
}
