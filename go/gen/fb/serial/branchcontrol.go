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

type BranchControl struct {
	_tab flatbuffers.Table
}

func InitBranchControlRoot(o *BranchControl, buf []byte, offset flatbuffers.UOffsetT) error {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	return o.Init(buf, n+offset)
}

func TryGetRootAsBranchControl(buf []byte, offset flatbuffers.UOffsetT) (*BranchControl, error) {
	x := &BranchControl{}
	return x, InitBranchControlRoot(x, buf, offset)
}

func TryGetSizePrefixedRootAsBranchControl(buf []byte, offset flatbuffers.UOffsetT) (*BranchControl, error) {
	x := &BranchControl{}
	return x, InitBranchControlRoot(x, buf, offset+flatbuffers.SizeUint32)
}

func (rcv *BranchControl) Init(buf []byte, i flatbuffers.UOffsetT) error {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
	if BranchControlNumFields < rcv.Table().NumFields() {
		return flatbuffers.ErrTableHasUnknownFields
	}
	return nil
}

func (rcv *BranchControl) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *BranchControl) TryAccessTbl(obj *BranchControlAccess) (*BranchControlAccess, error) {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(BranchControlAccess)
		}
		obj.Init(rcv._tab.Bytes, x)
		if BranchControlAccessNumFields < obj.Table().NumFields() {
			return nil, flatbuffers.ErrTableHasUnknownFields
		}
		return obj, nil
	}
	return nil, nil
}

func (rcv *BranchControl) TryNamespaceTbl(obj *BranchControlNamespace) (*BranchControlNamespace, error) {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(BranchControlNamespace)
		}
		obj.Init(rcv._tab.Bytes, x)
		if BranchControlNamespaceNumFields < obj.Table().NumFields() {
			return nil, flatbuffers.ErrTableHasUnknownFields
		}
		return obj, nil
	}
	return nil, nil
}

const BranchControlNumFields = 2

func BranchControlStart(builder *flatbuffers.Builder) {
	builder.StartObject(BranchControlNumFields)
}
func BranchControlAddAccessTbl(builder *flatbuffers.Builder, accessTbl flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(accessTbl), 0)
}
func BranchControlAddNamespaceTbl(builder *flatbuffers.Builder, namespaceTbl flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(namespaceTbl), 0)
}
func BranchControlEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

type BranchControlAccess struct {
	_tab flatbuffers.Table
}

func InitBranchControlAccessRoot(o *BranchControlAccess, buf []byte, offset flatbuffers.UOffsetT) error {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	return o.Init(buf, n+offset)
}

func TryGetRootAsBranchControlAccess(buf []byte, offset flatbuffers.UOffsetT) (*BranchControlAccess, error) {
	x := &BranchControlAccess{}
	return x, InitBranchControlAccessRoot(x, buf, offset)
}

func TryGetSizePrefixedRootAsBranchControlAccess(buf []byte, offset flatbuffers.UOffsetT) (*BranchControlAccess, error) {
	x := &BranchControlAccess{}
	return x, InitBranchControlAccessRoot(x, buf, offset+flatbuffers.SizeUint32)
}

func (rcv *BranchControlAccess) Init(buf []byte, i flatbuffers.UOffsetT) error {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
	if BranchControlAccessNumFields < rcv.Table().NumFields() {
		return flatbuffers.ErrTableHasUnknownFields
	}
	return nil
}

func (rcv *BranchControlAccess) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *BranchControlAccess) TryBinlog(obj *BranchControlBinlog) (*BranchControlBinlog, error) {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(BranchControlBinlog)
		}
		obj.Init(rcv._tab.Bytes, x)
		if BranchControlBinlogNumFields < obj.Table().NumFields() {
			return nil, flatbuffers.ErrTableHasUnknownFields
		}
		return obj, nil
	}
	return nil, nil
}

func (rcv *BranchControlAccess) TryDatabases(obj *BranchControlMatchExpression, j int) (bool, error) {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		if BranchControlMatchExpressionNumFields < obj.Table().NumFields() {
			return false, flatbuffers.ErrTableHasUnknownFields
		}
		return true, nil
	}
	return false, nil
}

func (rcv *BranchControlAccess) DatabasesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *BranchControlAccess) TryBranches(obj *BranchControlMatchExpression, j int) (bool, error) {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		if BranchControlMatchExpressionNumFields < obj.Table().NumFields() {
			return false, flatbuffers.ErrTableHasUnknownFields
		}
		return true, nil
	}
	return false, nil
}

func (rcv *BranchControlAccess) BranchesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *BranchControlAccess) TryUsers(obj *BranchControlMatchExpression, j int) (bool, error) {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		if BranchControlMatchExpressionNumFields < obj.Table().NumFields() {
			return false, flatbuffers.ErrTableHasUnknownFields
		}
		return true, nil
	}
	return false, nil
}

func (rcv *BranchControlAccess) UsersLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *BranchControlAccess) TryHosts(obj *BranchControlMatchExpression, j int) (bool, error) {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		if BranchControlMatchExpressionNumFields < obj.Table().NumFields() {
			return false, flatbuffers.ErrTableHasUnknownFields
		}
		return true, nil
	}
	return false, nil
}

func (rcv *BranchControlAccess) HostsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *BranchControlAccess) TryValues(obj *BranchControlAccessValue, j int) (bool, error) {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		if BranchControlAccessValueNumFields < obj.Table().NumFields() {
			return false, flatbuffers.ErrTableHasUnknownFields
		}
		return true, nil
	}
	return false, nil
}

func (rcv *BranchControlAccess) ValuesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

const BranchControlAccessNumFields = 6

func BranchControlAccessStart(builder *flatbuffers.Builder) {
	builder.StartObject(BranchControlAccessNumFields)
}
func BranchControlAccessAddBinlog(builder *flatbuffers.Builder, binlog flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(binlog), 0)
}
func BranchControlAccessAddDatabases(builder *flatbuffers.Builder, databases flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(databases), 0)
}
func BranchControlAccessStartDatabasesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func BranchControlAccessAddBranches(builder *flatbuffers.Builder, branches flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(branches), 0)
}
func BranchControlAccessStartBranchesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func BranchControlAccessAddUsers(builder *flatbuffers.Builder, users flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(users), 0)
}
func BranchControlAccessStartUsersVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func BranchControlAccessAddHosts(builder *flatbuffers.Builder, hosts flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(hosts), 0)
}
func BranchControlAccessStartHostsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func BranchControlAccessAddValues(builder *flatbuffers.Builder, values flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(values), 0)
}
func BranchControlAccessStartValuesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func BranchControlAccessEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

type BranchControlAccessValue struct {
	_tab flatbuffers.Table
}

func InitBranchControlAccessValueRoot(o *BranchControlAccessValue, buf []byte, offset flatbuffers.UOffsetT) error {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	return o.Init(buf, n+offset)
}

func TryGetRootAsBranchControlAccessValue(buf []byte, offset flatbuffers.UOffsetT) (*BranchControlAccessValue, error) {
	x := &BranchControlAccessValue{}
	return x, InitBranchControlAccessValueRoot(x, buf, offset)
}

func TryGetSizePrefixedRootAsBranchControlAccessValue(buf []byte, offset flatbuffers.UOffsetT) (*BranchControlAccessValue, error) {
	x := &BranchControlAccessValue{}
	return x, InitBranchControlAccessValueRoot(x, buf, offset+flatbuffers.SizeUint32)
}

func (rcv *BranchControlAccessValue) Init(buf []byte, i flatbuffers.UOffsetT) error {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
	if BranchControlAccessValueNumFields < rcv.Table().NumFields() {
		return flatbuffers.ErrTableHasUnknownFields
	}
	return nil
}

func (rcv *BranchControlAccessValue) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *BranchControlAccessValue) Database() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *BranchControlAccessValue) Branch() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *BranchControlAccessValue) User() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *BranchControlAccessValue) Host() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *BranchControlAccessValue) Permissions() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *BranchControlAccessValue) MutatePermissions(n uint64) bool {
	return rcv._tab.MutateUint64Slot(12, n)
}

const BranchControlAccessValueNumFields = 5

func BranchControlAccessValueStart(builder *flatbuffers.Builder) {
	builder.StartObject(BranchControlAccessValueNumFields)
}
func BranchControlAccessValueAddDatabase(builder *flatbuffers.Builder, database flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(database), 0)
}
func BranchControlAccessValueAddBranch(builder *flatbuffers.Builder, branch flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(branch), 0)
}
func BranchControlAccessValueAddUser(builder *flatbuffers.Builder, user flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(user), 0)
}
func BranchControlAccessValueAddHost(builder *flatbuffers.Builder, host flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(host), 0)
}
func BranchControlAccessValueAddPermissions(builder *flatbuffers.Builder, permissions uint64) {
	builder.PrependUint64Slot(4, permissions, 0)
}
func BranchControlAccessValueEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

type BranchControlNamespace struct {
	_tab flatbuffers.Table
}

func InitBranchControlNamespaceRoot(o *BranchControlNamespace, buf []byte, offset flatbuffers.UOffsetT) error {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	return o.Init(buf, n+offset)
}

func TryGetRootAsBranchControlNamespace(buf []byte, offset flatbuffers.UOffsetT) (*BranchControlNamespace, error) {
	x := &BranchControlNamespace{}
	return x, InitBranchControlNamespaceRoot(x, buf, offset)
}

func TryGetSizePrefixedRootAsBranchControlNamespace(buf []byte, offset flatbuffers.UOffsetT) (*BranchControlNamespace, error) {
	x := &BranchControlNamespace{}
	return x, InitBranchControlNamespaceRoot(x, buf, offset+flatbuffers.SizeUint32)
}

func (rcv *BranchControlNamespace) Init(buf []byte, i flatbuffers.UOffsetT) error {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
	if BranchControlNamespaceNumFields < rcv.Table().NumFields() {
		return flatbuffers.ErrTableHasUnknownFields
	}
	return nil
}

func (rcv *BranchControlNamespace) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *BranchControlNamespace) TryBinlog(obj *BranchControlBinlog) (*BranchControlBinlog, error) {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(BranchControlBinlog)
		}
		obj.Init(rcv._tab.Bytes, x)
		if BranchControlBinlogNumFields < obj.Table().NumFields() {
			return nil, flatbuffers.ErrTableHasUnknownFields
		}
		return obj, nil
	}
	return nil, nil
}

func (rcv *BranchControlNamespace) TryDatabases(obj *BranchControlMatchExpression, j int) (bool, error) {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		if BranchControlMatchExpressionNumFields < obj.Table().NumFields() {
			return false, flatbuffers.ErrTableHasUnknownFields
		}
		return true, nil
	}
	return false, nil
}

func (rcv *BranchControlNamespace) DatabasesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *BranchControlNamespace) TryBranches(obj *BranchControlMatchExpression, j int) (bool, error) {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		if BranchControlMatchExpressionNumFields < obj.Table().NumFields() {
			return false, flatbuffers.ErrTableHasUnknownFields
		}
		return true, nil
	}
	return false, nil
}

func (rcv *BranchControlNamespace) BranchesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *BranchControlNamespace) TryUsers(obj *BranchControlMatchExpression, j int) (bool, error) {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		if BranchControlMatchExpressionNumFields < obj.Table().NumFields() {
			return false, flatbuffers.ErrTableHasUnknownFields
		}
		return true, nil
	}
	return false, nil
}

func (rcv *BranchControlNamespace) UsersLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *BranchControlNamespace) TryHosts(obj *BranchControlMatchExpression, j int) (bool, error) {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		if BranchControlMatchExpressionNumFields < obj.Table().NumFields() {
			return false, flatbuffers.ErrTableHasUnknownFields
		}
		return true, nil
	}
	return false, nil
}

func (rcv *BranchControlNamespace) HostsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *BranchControlNamespace) TryValues(obj *BranchControlNamespaceValue, j int) (bool, error) {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		if BranchControlNamespaceValueNumFields < obj.Table().NumFields() {
			return false, flatbuffers.ErrTableHasUnknownFields
		}
		return true, nil
	}
	return false, nil
}

func (rcv *BranchControlNamespace) ValuesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

const BranchControlNamespaceNumFields = 6

func BranchControlNamespaceStart(builder *flatbuffers.Builder) {
	builder.StartObject(BranchControlNamespaceNumFields)
}
func BranchControlNamespaceAddBinlog(builder *flatbuffers.Builder, binlog flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(binlog), 0)
}
func BranchControlNamespaceAddDatabases(builder *flatbuffers.Builder, databases flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(databases), 0)
}
func BranchControlNamespaceStartDatabasesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func BranchControlNamespaceAddBranches(builder *flatbuffers.Builder, branches flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(branches), 0)
}
func BranchControlNamespaceStartBranchesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func BranchControlNamespaceAddUsers(builder *flatbuffers.Builder, users flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(users), 0)
}
func BranchControlNamespaceStartUsersVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func BranchControlNamespaceAddHosts(builder *flatbuffers.Builder, hosts flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(hosts), 0)
}
func BranchControlNamespaceStartHostsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func BranchControlNamespaceAddValues(builder *flatbuffers.Builder, values flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(values), 0)
}
func BranchControlNamespaceStartValuesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func BranchControlNamespaceEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

type BranchControlNamespaceValue struct {
	_tab flatbuffers.Table
}

func InitBranchControlNamespaceValueRoot(o *BranchControlNamespaceValue, buf []byte, offset flatbuffers.UOffsetT) error {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	return o.Init(buf, n+offset)
}

func TryGetRootAsBranchControlNamespaceValue(buf []byte, offset flatbuffers.UOffsetT) (*BranchControlNamespaceValue, error) {
	x := &BranchControlNamespaceValue{}
	return x, InitBranchControlNamespaceValueRoot(x, buf, offset)
}

func TryGetSizePrefixedRootAsBranchControlNamespaceValue(buf []byte, offset flatbuffers.UOffsetT) (*BranchControlNamespaceValue, error) {
	x := &BranchControlNamespaceValue{}
	return x, InitBranchControlNamespaceValueRoot(x, buf, offset+flatbuffers.SizeUint32)
}

func (rcv *BranchControlNamespaceValue) Init(buf []byte, i flatbuffers.UOffsetT) error {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
	if BranchControlNamespaceValueNumFields < rcv.Table().NumFields() {
		return flatbuffers.ErrTableHasUnknownFields
	}
	return nil
}

func (rcv *BranchControlNamespaceValue) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *BranchControlNamespaceValue) Database() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *BranchControlNamespaceValue) Branch() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *BranchControlNamespaceValue) User() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *BranchControlNamespaceValue) Host() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

const BranchControlNamespaceValueNumFields = 4

func BranchControlNamespaceValueStart(builder *flatbuffers.Builder) {
	builder.StartObject(BranchControlNamespaceValueNumFields)
}
func BranchControlNamespaceValueAddDatabase(builder *flatbuffers.Builder, database flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(database), 0)
}
func BranchControlNamespaceValueAddBranch(builder *flatbuffers.Builder, branch flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(branch), 0)
}
func BranchControlNamespaceValueAddUser(builder *flatbuffers.Builder, user flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(user), 0)
}
func BranchControlNamespaceValueAddHost(builder *flatbuffers.Builder, host flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(host), 0)
}
func BranchControlNamespaceValueEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

type BranchControlBinlog struct {
	_tab flatbuffers.Table
}

func InitBranchControlBinlogRoot(o *BranchControlBinlog, buf []byte, offset flatbuffers.UOffsetT) error {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	return o.Init(buf, n+offset)
}

func TryGetRootAsBranchControlBinlog(buf []byte, offset flatbuffers.UOffsetT) (*BranchControlBinlog, error) {
	x := &BranchControlBinlog{}
	return x, InitBranchControlBinlogRoot(x, buf, offset)
}

func TryGetSizePrefixedRootAsBranchControlBinlog(buf []byte, offset flatbuffers.UOffsetT) (*BranchControlBinlog, error) {
	x := &BranchControlBinlog{}
	return x, InitBranchControlBinlogRoot(x, buf, offset+flatbuffers.SizeUint32)
}

func (rcv *BranchControlBinlog) Init(buf []byte, i flatbuffers.UOffsetT) error {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
	if BranchControlBinlogNumFields < rcv.Table().NumFields() {
		return flatbuffers.ErrTableHasUnknownFields
	}
	return nil
}

func (rcv *BranchControlBinlog) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *BranchControlBinlog) TryRows(obj *BranchControlBinlogRow, j int) (bool, error) {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		if BranchControlBinlogRowNumFields < obj.Table().NumFields() {
			return false, flatbuffers.ErrTableHasUnknownFields
		}
		return true, nil
	}
	return false, nil
}

func (rcv *BranchControlBinlog) RowsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

const BranchControlBinlogNumFields = 1

func BranchControlBinlogStart(builder *flatbuffers.Builder) {
	builder.StartObject(BranchControlBinlogNumFields)
}
func BranchControlBinlogAddRows(builder *flatbuffers.Builder, rows flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(rows), 0)
}
func BranchControlBinlogStartRowsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func BranchControlBinlogEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

type BranchControlBinlogRow struct {
	_tab flatbuffers.Table
}

func InitBranchControlBinlogRowRoot(o *BranchControlBinlogRow, buf []byte, offset flatbuffers.UOffsetT) error {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	return o.Init(buf, n+offset)
}

func TryGetRootAsBranchControlBinlogRow(buf []byte, offset flatbuffers.UOffsetT) (*BranchControlBinlogRow, error) {
	x := &BranchControlBinlogRow{}
	return x, InitBranchControlBinlogRowRoot(x, buf, offset)
}

func TryGetSizePrefixedRootAsBranchControlBinlogRow(buf []byte, offset flatbuffers.UOffsetT) (*BranchControlBinlogRow, error) {
	x := &BranchControlBinlogRow{}
	return x, InitBranchControlBinlogRowRoot(x, buf, offset+flatbuffers.SizeUint32)
}

func (rcv *BranchControlBinlogRow) Init(buf []byte, i flatbuffers.UOffsetT) error {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
	if BranchControlBinlogRowNumFields < rcv.Table().NumFields() {
		return flatbuffers.ErrTableHasUnknownFields
	}
	return nil
}

func (rcv *BranchControlBinlogRow) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *BranchControlBinlogRow) IsInsert() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *BranchControlBinlogRow) MutateIsInsert(n bool) bool {
	return rcv._tab.MutateBoolSlot(4, n)
}

func (rcv *BranchControlBinlogRow) Database() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *BranchControlBinlogRow) Branch() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *BranchControlBinlogRow) User() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *BranchControlBinlogRow) Host() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *BranchControlBinlogRow) Permissions() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *BranchControlBinlogRow) MutatePermissions(n uint64) bool {
	return rcv._tab.MutateUint64Slot(14, n)
}

const BranchControlBinlogRowNumFields = 6

func BranchControlBinlogRowStart(builder *flatbuffers.Builder) {
	builder.StartObject(BranchControlBinlogRowNumFields)
}
func BranchControlBinlogRowAddIsInsert(builder *flatbuffers.Builder, isInsert bool) {
	builder.PrependBoolSlot(0, isInsert, false)
}
func BranchControlBinlogRowAddDatabase(builder *flatbuffers.Builder, database flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(database), 0)
}
func BranchControlBinlogRowAddBranch(builder *flatbuffers.Builder, branch flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(branch), 0)
}
func BranchControlBinlogRowAddUser(builder *flatbuffers.Builder, user flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(user), 0)
}
func BranchControlBinlogRowAddHost(builder *flatbuffers.Builder, host flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(host), 0)
}
func BranchControlBinlogRowAddPermissions(builder *flatbuffers.Builder, permissions uint64) {
	builder.PrependUint64Slot(5, permissions, 0)
}
func BranchControlBinlogRowEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

type BranchControlMatchExpression struct {
	_tab flatbuffers.Table
}

func InitBranchControlMatchExpressionRoot(o *BranchControlMatchExpression, buf []byte, offset flatbuffers.UOffsetT) error {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	return o.Init(buf, n+offset)
}

func TryGetRootAsBranchControlMatchExpression(buf []byte, offset flatbuffers.UOffsetT) (*BranchControlMatchExpression, error) {
	x := &BranchControlMatchExpression{}
	return x, InitBranchControlMatchExpressionRoot(x, buf, offset)
}

func TryGetSizePrefixedRootAsBranchControlMatchExpression(buf []byte, offset flatbuffers.UOffsetT) (*BranchControlMatchExpression, error) {
	x := &BranchControlMatchExpression{}
	return x, InitBranchControlMatchExpressionRoot(x, buf, offset+flatbuffers.SizeUint32)
}

func (rcv *BranchControlMatchExpression) Init(buf []byte, i flatbuffers.UOffsetT) error {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
	if BranchControlMatchExpressionNumFields < rcv.Table().NumFields() {
		return flatbuffers.ErrTableHasUnknownFields
	}
	return nil
}

func (rcv *BranchControlMatchExpression) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *BranchControlMatchExpression) Index() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *BranchControlMatchExpression) MutateIndex(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *BranchControlMatchExpression) SortOrders(j int) int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *BranchControlMatchExpression) SortOrdersLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *BranchControlMatchExpression) MutateSortOrders(j int, n int32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

const BranchControlMatchExpressionNumFields = 2

func BranchControlMatchExpressionStart(builder *flatbuffers.Builder) {
	builder.StartObject(BranchControlMatchExpressionNumFields)
}
func BranchControlMatchExpressionAddIndex(builder *flatbuffers.Builder, index uint32) {
	builder.PrependUint32Slot(0, index, 0)
}
func BranchControlMatchExpressionAddSortOrders(builder *flatbuffers.Builder, sortOrders flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(sortOrders), 0)
}
func BranchControlMatchExpressionStartSortOrdersVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func BranchControlMatchExpressionEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}