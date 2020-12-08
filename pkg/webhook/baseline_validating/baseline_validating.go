/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package baseline_validating

import (
	"fmt"

	v1 "k8s.io/api/admission/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
	"webhook/pkg/webhook/baseline_validating/podenv"
	"webhook/pkg/webhook/utils"
)

// baseline pod基线
func BaseLineValidating(ar v1.AdmissionReview) *v1.AdmissionResponse {

	klog.V(2).Info("handpay pod baseline")
	podResource := metav1.GroupVersionResource{Group: "", Version: "v1", Resource: "pods"}
	if ar.Request.Resource != podResource {
		err := fmt.Errorf("expect resource to be %s", podResource)
		klog.Error(err)
		return utils.ToV1AdmissionResponse(err)
	}

	raw := ar.Request.Object.Raw
	pod := corev1.Pod{}
	deserializer := utils.Codecs.UniversalDeserializer()
	if _, _, err := deserializer.Decode(raw, nil, &pod); err != nil {
		klog.Error(err)
		return utils.ToV1AdmissionResponse(err)
	}
	reviewResponse := v1.AdmissionResponse{}
	reviewResponse.Allowed = true

	if !podenv.IsExistEnv("TZ", pod) {
		reviewResponse.Allowed = false
		reviewResponse.Result = &metav1.Status{
			Message: "验证消息GG",
			Reason:  metav1.StatusReasonForbidden,
			Code:    500,
		}
	} else {
		reviewResponse.Result = &metav1.Status{
			Message: "GoodJob",
			Reason:  metav1.StatusReasonForbidden,
			Code:    0,
		}
	}

	return &reviewResponse
}
