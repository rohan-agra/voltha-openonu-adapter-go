/*
 * Copyright 2018-present Open Networking Foundation

 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at

 * http://www.apache.org/licenses/LICENSE-2.0

 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

//Package adaptercoreonu provides the utility for onu devices, flows and statistics
package adaptercoreonu

import (
	gp "github.com/google/gopacket"
	"github.com/opencord/omci-lib-go"
)

type MessageType uint8

const (
	TestMsg MessageType = iota
	OMCI
)

func (m MessageType) String() string {
	names := [...]string{
		"TestMsg",
		"OMCI",
	}
	return names[m]
}

type Message struct {
	Type MessageType
	Data interface{}
}

type TestMessageType uint8

const (
	noOperation TestMessageType = iota
	LoadMibTemplateOk
	LoadMibTemplateFailed
	TimeOutOccurred
	AbortMessageProcessing
)

//TODO: place holder to have a second interface variant - to be replaced by real variant later on
type TestMessage struct {
	TestMessageVal TestMessageType
}

type OmciMessage struct {
	//OnuSN   *openolt.SerialNumber
	//OnuID   uint32
	OmciMsg    *omci.OMCI
	OmciPacket *gp.Packet
}
