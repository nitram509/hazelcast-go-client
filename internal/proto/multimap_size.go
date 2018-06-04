// Copyright (c) 2008-2018, Hazelcast, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package proto

import ()

func multimapSizeCalculateSize(name string) int {
	// Calculates the request payload size
	dataSize := 0
	dataSize += stringCalculateSize(name)
	return dataSize
}

// MultiMapSizeEncodeRequest creates and encodes a client message
// with the given parameters.
// It returns the encoded client message.
func MultiMapSizeEncodeRequest(name string) *ClientMessage {
	// Encode request into clientMessage
	clientMessage := NewClientMessage(nil, multimapSizeCalculateSize(name))
	clientMessage.SetMessageType(multimapSize)
	clientMessage.IsRetryable = true
	clientMessage.AppendString(name)
	clientMessage.UpdateFrameLength()
	return clientMessage
}

// MultiMapSizeDecodeResponse decodes the given client message.
// It returns a function which returns the response parameters.
func MultiMapSizeDecodeResponse(clientMessage *ClientMessage) func() (response int32) {
	// Decode response from client message
	return func() (response int32) {
		response = clientMessage.ReadInt32()
		return
	}
}
