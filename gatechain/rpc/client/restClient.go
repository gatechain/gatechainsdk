// Copyright (C) 2019 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/gatechain/gatechainsdk/gatechain/node/appinterface"
	"github.com/gatechain/gatechainsdk/gatechain/rpc/spec/common"
	"github.com/gatechain/gatechainsdk/gatechain/rpc/spec/v1"
)

const (
	authHeader           = "X-Gatemint-API-Token"
	healthCheckEndpoint  = "/health"
	apiVersionPathPrefix = "/v1"
)

// unversionedPaths ais a set of paths that should not be prefixed by the API version
var unversionedPaths = map[string]bool{
	"/versions": true,
	"/health":   true,
}

// rawRequestPaths is a set of paths where the body should not be urlencoded
var rawRequestPaths = map[string]bool{
	"/transactions":         true,
	"/broadcastTx":          true,
	"/tx":                   true,
	"/getParticipationKey":  true,
	"/genParticipationKey":  true,
	"/getConAccount":        true,
	"/listLocalConAccounts": true,
	"/txsearch":             true,
}

// RestClient manages the REST interface for a calling user.
type RestClient struct {
	serverURL url.URL
	apiToken  string
}

// MakeRestClient is the factory for constructing a RestClient for a given endpoint
func MakeRestClient(url url.URL, apiToken string) RestClient {
	return RestClient{
		serverURL: url,
		apiToken:  apiToken,
	}
}

// extractError checks if the response signifies an error (for now, StatusCode != 200).
// If so, it returns the error.
// Otherwise, it returns nil.
func extractError(resp *http.Response) error {
	if resp.StatusCode == 200 {
		return nil
	}

	errorBuf, _ := ioutil.ReadAll(resp.Body) // ignore returned error
	return fmt.Errorf("HTTP %v: %s", resp.Status, errorBuf)
}

// submitForm is a helper used for submitting (ex.) GETs and POSTs to the server
func (client RestClient) submitForm(response interface{}, path string, request interface{}, requestMethod string, encodeJSON bool) error {
	var err error
	queryURL := client.serverURL
	queryURL.Path = path

	// Handle version prefix
	if !unversionedPaths[path] {
		queryURL.Path = strings.Join([]string{apiVersionPathPrefix, path}, "")
	}

	var req *http.Request
	var body io.Reader

	if request != nil {
		if rawRequestPaths[path] {
			reqBytes, ok := request.([]byte)
			if !ok {
				return fmt.Errorf("couldn't decode raw request as bytes")
			}
			body = bytes.NewBuffer(reqBytes)
		} else {
			v, err := query.Values(request)
			if err != nil {
				return err
			}

			queryURL.RawQuery = v.Encode()
			if encodeJSON {
				jsonValue, _ := json.Marshal(request)
				body = bytes.NewBuffer(jsonValue)
			}
		}
	}

	req, err = http.NewRequest(requestMethod, queryURL.String(), body)
	if err != nil {
		return err
	}

	// If we add another endpoint that does not require auth, we should add a
	// requiresAuth argument to submitForm rather than checking here
	if path != healthCheckEndpoint {
		req.Header.Set(authHeader, client.apiToken)
	}

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	err = extractError(resp)
	if err != nil {
		return err
	}

	dec := json.NewDecoder(resp.Body)
	return dec.Decode(&response)
}

// get performs a GET request to the specific path against the server
func (client RestClient) get(response interface{}, path string, request interface{}) error {
	return client.submitForm(response, path, request, "GET", false /* encodeJSON */)
}

// post sends a POST request to the given path with the given request object.
// No query parameters will be sent if request is nil.
// response must be a pointer to an object as post writes the response there.
func (client RestClient) post(response interface{}, path string, request interface{}) error {
	return client.submitForm(response, path, request, "POST", true /* encodeJSON */)
}

// Status retrieves the StatusResponse from the running node
// the StatusResponse includes data like the consensus version and current round
// Not supported
func (client RestClient) Status() (response v1.NodeStatus, err error) {
	err = client.get(&response, "/status", nil)
	return
}

// Versions retrieves the VersionResponse from the running node
// the VersionResponse includes data like version number and genesis ID
func (client RestClient) Versions() (response common.Version, err error) {
	err = client.get(&response, "/versions", nil)
	return
}

// Block gets the block info for the given round
func (client RestClient) Block(round uint64) (response v1.Block, err error) {
	err = client.get(&response, fmt.Sprintf("/block/%d", round), nil)
	return
}

func (client RestClient) Query(path string, data []byte, opts appinterface.QueryOptions) (response v1.Response, err error) {
	requestQuery := new(appinterface.RequestQuery)
	requestQuery.Path = path
	requestQuery.Data = data
	requestQuery.Height = opts.Height
	requestQuery.Prove = opts.Prove
	err = client.post(&response, fmt.Sprintf("/app-query"), requestQuery)
	return
}

func (client RestClient) BroadcastTx(bytes []byte) (response v1.ResultBroadcastTx, err error) {
	err = client.post(&response, "/broadcastTx", bytes)
	return
}

func (client RestClient) Tx(bytes []byte) (response appinterface.ResponseTx, err error) {
	err = client.post(&response, "/tx", bytes)
	return
}

// GenParticipationKeysTo
func (client RestClient) GenParticipationKey(bytes []byte) (response v1.ParticipationKeyResponse, err error) {
	err = client.get(&response, "/genParticipationKey", bytes)
	return
}

// GetParticipationKeys
func (client RestClient) GetParticipationKey(bytes []byte) (response v1.ParticipationKeyResponse, err error) {
	err = client.get(&response, "/getParticipationKey", bytes)
	return
}

// GetConAccount
func (client RestClient) GetConAccount(bytes []byte) (response v1.ResultConAccount, err error) {
	err = client.get(&response, "/getConAccount", bytes)
	return
}

// ListLocalConAccounts
func (client RestClient) ListLocalConAccounts() (response []appinterface.ParticipationData, err error) {
	err = client.get(&response, "/listLocalConAccounts", nil)
	return
}
