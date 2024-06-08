// Copyright 2022-2023 Dolthub, Inc.
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

// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package serial

import (
	flatbuffers "github.com/dolthub/flatbuffers/v23/go"
)

type StoreRoot struct {
	_tab flatbuffers.Table
}

func InitStoreRootRoot(o *StoreRoot, buf []byte, offset flatbuffers.UOffsetT) error {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	return o.Init(buf, n+offset)
}

func TryGetRootAsStoreRoot(buf []byte, offset flatbuffers.UOffsetT) (*StoreRoot, error) {
	x := &StoreRoot{}
	return x, InitStoreRootRoot(x, buf, offset)
}

func TryGetSizePrefixedRootAsStoreRoot(buf []byte, offset flatbuffers.UOffsetT) (*StoreRoot, error) {
	x := &StoreRoot{}
	return x, InitStoreRootRoot(x, buf, offset+flatbuffers.SizeUint32)
}

func (rcv *StoreRoot) Init(buf []byte, i flatbuffers.UOffsetT) error {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
	if StoreRootNumFields < rcv.Table().NumFields() {
		return flatbuffers.ErrTableHasUnknownFields
	}
	return nil
}

func (rcv *StoreRoot) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *StoreRoot) AddressMap(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *StoreRoot) AddressMapLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *StoreRoot) AddressMapBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *StoreRoot) MutateAddressMap(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

const StoreRootNumFields = 1

func StoreRootStart(builder *flatbuffers.Builder) {
	builder.StartObject(StoreRootNumFields)
}
func StoreRootAddAddressMap(builder *flatbuffers.Builder, addressMap flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(addressMap), 0)
}
func StoreRootStartAddressMapVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func StoreRootEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}