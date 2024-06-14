package services_tests

import (
	"TopDoctorsBackendChallenge/internal/models"
	"context"
	"github.com/golang/mock/gomock"
	"reflect"
)

type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryRecorder
}

type MockRepositoryRecorder struct {
	mock *MockRepository
}

func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{
		ctrl: ctrl,
	}
	mock.recorder = &MockRepositoryRecorder{mock: mock}
	return mock
}

func (m *MockRepository) EXPECT() *MockRepositoryRecorder {
	return m.recorder
}

func (m *MockRepository) CreateDiagnosis(ctx context.Context, diagnosis models.Diagnosis) (models.Diagnosis, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDiagnosis", ctx, diagnosis)
	ret0, _ := ret[0].(models.Diagnosis)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockRepository) SearchDiagnoses(ctx context.Context, filters models.SearchDiagnosesFilters) ([]models.Diagnosis, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchDiagnoses", ctx, filters)
	ret0, _ := ret[0].([]models.Diagnosis)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockRepository) DoesPatientExist(ctx context.Context, patientID string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoesPatientExist", ctx, patientID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRepositoryRecorder) CreateDiagnosis(ctx, diagnosis interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDiagnosis", reflect.TypeOf((*MockRepository)(nil).CreateDiagnosis), ctx, diagnosis)
}

func (mr *MockRepositoryRecorder) SearchDiagnoses(ctx, filters interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchDiagnoses", reflect.TypeOf((*MockRepository)(nil).SearchDiagnoses), ctx, filters)
}

func (mr *MockRepositoryRecorder) DoesPatientExist(ctx, patientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoesPatientExist", reflect.TypeOf((*MockRepository)(nil).DoesPatientExist), ctx, patientID)
}
