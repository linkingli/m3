// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/dbnode/persist/types.go

// Copyright (c) 2020 Uber Technologies, Inc.
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

// Package persist is a generated GoMock package.
package persist

import (
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/pborman/uuid"
)

// MockManager is a mock of Manager interface
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// StartFlushPersist mocks base method
func (m *MockManager) StartFlushPersist() (FlushPreparer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartFlushPersist")
	ret0, _ := ret[0].(FlushPreparer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartFlushPersist indicates an expected call of StartFlushPersist
func (mr *MockManagerMockRecorder) StartFlushPersist() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartFlushPersist", reflect.TypeOf((*MockManager)(nil).StartFlushPersist))
}

// StartSnapshotPersist mocks base method
func (m *MockManager) StartSnapshotPersist(snapshotID uuid.UUID) (SnapshotPreparer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartSnapshotPersist", snapshotID)
	ret0, _ := ret[0].(SnapshotPreparer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartSnapshotPersist indicates an expected call of StartSnapshotPersist
func (mr *MockManagerMockRecorder) StartSnapshotPersist(snapshotID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartSnapshotPersist", reflect.TypeOf((*MockManager)(nil).StartSnapshotPersist), snapshotID)
}

// StartIndexPersist mocks base method
func (m *MockManager) StartIndexPersist() (IndexFlush, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartIndexPersist")
	ret0, _ := ret[0].(IndexFlush)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartIndexPersist indicates an expected call of StartIndexPersist
func (mr *MockManagerMockRecorder) StartIndexPersist() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartIndexPersist", reflect.TypeOf((*MockManager)(nil).StartIndexPersist))
}

// Close mocks base method
func (m *MockManager) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockManagerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockManager)(nil).Close))
}

// MockPreparer is a mock of Preparer interface
type MockPreparer struct {
	ctrl     *gomock.Controller
	recorder *MockPreparerMockRecorder
}

// MockPreparerMockRecorder is the mock recorder for MockPreparer
type MockPreparerMockRecorder struct {
	mock *MockPreparer
}

// NewMockPreparer creates a new mock instance
func NewMockPreparer(ctrl *gomock.Controller) *MockPreparer {
	mock := &MockPreparer{ctrl: ctrl}
	mock.recorder = &MockPreparerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPreparer) EXPECT() *MockPreparerMockRecorder {
	return m.recorder
}

// PrepareData mocks base method
func (m *MockPreparer) PrepareData(opts DataPrepareOptions) (PreparedDataPersist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareData", opts)
	ret0, _ := ret[0].(PreparedDataPersist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareData indicates an expected call of PrepareData
func (mr *MockPreparerMockRecorder) PrepareData(opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareData", reflect.TypeOf((*MockPreparer)(nil).PrepareData), opts)
}

// MockFlushPreparer is a mock of FlushPreparer interface
type MockFlushPreparer struct {
	ctrl     *gomock.Controller
	recorder *MockFlushPreparerMockRecorder
}

// MockFlushPreparerMockRecorder is the mock recorder for MockFlushPreparer
type MockFlushPreparerMockRecorder struct {
	mock *MockFlushPreparer
}

// NewMockFlushPreparer creates a new mock instance
func NewMockFlushPreparer(ctrl *gomock.Controller) *MockFlushPreparer {
	mock := &MockFlushPreparer{ctrl: ctrl}
	mock.recorder = &MockFlushPreparerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFlushPreparer) EXPECT() *MockFlushPreparerMockRecorder {
	return m.recorder
}

// PrepareData mocks base method
func (m *MockFlushPreparer) PrepareData(opts DataPrepareOptions) (PreparedDataPersist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareData", opts)
	ret0, _ := ret[0].(PreparedDataPersist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareData indicates an expected call of PrepareData
func (mr *MockFlushPreparerMockRecorder) PrepareData(opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareData", reflect.TypeOf((*MockFlushPreparer)(nil).PrepareData), opts)
}

// DoneFlush mocks base method
func (m *MockFlushPreparer) DoneFlush() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoneFlush")
	ret0, _ := ret[0].(error)
	return ret0
}

// DoneFlush indicates an expected call of DoneFlush
func (mr *MockFlushPreparerMockRecorder) DoneFlush() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoneFlush", reflect.TypeOf((*MockFlushPreparer)(nil).DoneFlush))
}

// MockSnapshotPreparer is a mock of SnapshotPreparer interface
type MockSnapshotPreparer struct {
	ctrl     *gomock.Controller
	recorder *MockSnapshotPreparerMockRecorder
}

// MockSnapshotPreparerMockRecorder is the mock recorder for MockSnapshotPreparer
type MockSnapshotPreparerMockRecorder struct {
	mock *MockSnapshotPreparer
}

// NewMockSnapshotPreparer creates a new mock instance
func NewMockSnapshotPreparer(ctrl *gomock.Controller) *MockSnapshotPreparer {
	mock := &MockSnapshotPreparer{ctrl: ctrl}
	mock.recorder = &MockSnapshotPreparerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSnapshotPreparer) EXPECT() *MockSnapshotPreparerMockRecorder {
	return m.recorder
}

// PrepareData mocks base method
func (m *MockSnapshotPreparer) PrepareData(opts DataPrepareOptions) (PreparedDataPersist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareData", opts)
	ret0, _ := ret[0].(PreparedDataPersist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareData indicates an expected call of PrepareData
func (mr *MockSnapshotPreparerMockRecorder) PrepareData(opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareData", reflect.TypeOf((*MockSnapshotPreparer)(nil).PrepareData), opts)
}

// PrepareIndexSnapshot mocks base method
func (m *MockSnapshotPreparer) PrepareIndexSnapshot(opts IndexPrepareOptions) (PreparedIndexSnapshotPersist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareIndexSnapshot", opts)
	ret0, _ := ret[0].(PreparedIndexSnapshotPersist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareIndexSnapshot indicates an expected call of PrepareIndexSnapshot
func (mr *MockSnapshotPreparerMockRecorder) PrepareIndexSnapshot(opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareIndexSnapshot", reflect.TypeOf((*MockSnapshotPreparer)(nil).PrepareIndexSnapshot), opts)
}

// DoneSnapshot mocks base method
func (m *MockSnapshotPreparer) DoneSnapshot(snapshotUUID uuid.UUID, commitLogIdentifier CommitLogFile) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoneSnapshot", snapshotUUID, commitLogIdentifier)
	ret0, _ := ret[0].(error)
	return ret0
}

// DoneSnapshot indicates an expected call of DoneSnapshot
func (mr *MockSnapshotPreparerMockRecorder) DoneSnapshot(snapshotUUID, commitLogIdentifier interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoneSnapshot", reflect.TypeOf((*MockSnapshotPreparer)(nil).DoneSnapshot), snapshotUUID, commitLogIdentifier)
}

// MockIndexFlush is a mock of IndexFlush interface
type MockIndexFlush struct {
	ctrl     *gomock.Controller
	recorder *MockIndexFlushMockRecorder
}

// MockIndexFlushMockRecorder is the mock recorder for MockIndexFlush
type MockIndexFlushMockRecorder struct {
	mock *MockIndexFlush
}

// NewMockIndexFlush creates a new mock instance
func NewMockIndexFlush(ctrl *gomock.Controller) *MockIndexFlush {
	mock := &MockIndexFlush{ctrl: ctrl}
	mock.recorder = &MockIndexFlushMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIndexFlush) EXPECT() *MockIndexFlushMockRecorder {
	return m.recorder
}

// PrepareIndexFlush mocks base method
func (m *MockIndexFlush) PrepareIndexFlush(opts IndexPrepareOptions) (PreparedIndexFlushPersist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareIndexFlush", opts)
	ret0, _ := ret[0].(PreparedIndexFlushPersist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareIndexFlush indicates an expected call of PrepareIndexFlush
func (mr *MockIndexFlushMockRecorder) PrepareIndexFlush(opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareIndexFlush", reflect.TypeOf((*MockIndexFlush)(nil).PrepareIndexFlush), opts)
}

// DoneIndex mocks base method
func (m *MockIndexFlush) DoneIndex() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoneIndex")
	ret0, _ := ret[0].(error)
	return ret0
}

// DoneIndex indicates an expected call of DoneIndex
func (mr *MockIndexFlushMockRecorder) DoneIndex() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoneIndex", reflect.TypeOf((*MockIndexFlush)(nil).DoneIndex))
}

// MockOnFlushSeries is a mock of OnFlushSeries interface
type MockOnFlushSeries struct {
	ctrl     *gomock.Controller
	recorder *MockOnFlushSeriesMockRecorder
}

// MockOnFlushSeriesMockRecorder is the mock recorder for MockOnFlushSeries
type MockOnFlushSeriesMockRecorder struct {
	mock *MockOnFlushSeries
}

// NewMockOnFlushSeries creates a new mock instance
func NewMockOnFlushSeries(ctrl *gomock.Controller) *MockOnFlushSeries {
	mock := &MockOnFlushSeries{ctrl: ctrl}
	mock.recorder = &MockOnFlushSeriesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOnFlushSeries) EXPECT() *MockOnFlushSeriesMockRecorder {
	return m.recorder
}

// OnFlushNewSeries mocks base method
func (m *MockOnFlushSeries) OnFlushNewSeries(arg0 OnFlushNewSeriesEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnFlushNewSeries", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnFlushNewSeries indicates an expected call of OnFlushNewSeries
func (mr *MockOnFlushSeriesMockRecorder) OnFlushNewSeries(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnFlushNewSeries", reflect.TypeOf((*MockOnFlushSeries)(nil).OnFlushNewSeries), arg0)
}
