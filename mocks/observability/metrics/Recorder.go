// Code generated by mockery v1.0.0
package metrics

import metrics "github.com/slok/kubewebhook/pkg/observability/metrics"
import mock "github.com/stretchr/testify/mock"
import time "time"
import v1beta1 "k8s.io/api/admission/v1beta1"

// Recorder is an autogenerated mock type for the Recorder type
type Recorder struct {
	mock.Mock
}

// IncAdmissionReview provides a mock function with given fields: webhook, namespace, resource, operation, kind
func (_m *Recorder) IncAdmissionReview(webhook string, namespace string, resource string, operation v1beta1.Operation, kind metrics.ReviewKind) {
	_m.Called(webhook, namespace, resource, operation, kind)
}

// IncAdmissionReviewError provides a mock function with given fields: webhook, namespace, resource, operation, kind
func (_m *Recorder) IncAdmissionReviewError(webhook string, namespace string, resource string, operation v1beta1.Operation, kind metrics.ReviewKind) {
	_m.Called(webhook, namespace, resource, operation, kind)
}

// ObserveAdmissionReviewDuration provides a mock function with given fields: webhook, namespace, resource, operation, kind, startTime
func (_m *Recorder) ObserveAdmissionReviewDuration(webhook string, namespace string, resource string, operation v1beta1.Operation, kind metrics.ReviewKind, startTime time.Time) {
	_m.Called(webhook, namespace, resource, operation, kind, startTime)
}