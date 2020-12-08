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

package podenv

import (
	corev1 "k8s.io/api/core/v1"
)

// key   为环境变量名称
// value 为环境变量值
func AddEnv(key string, value string, pod corev1.Pod) corev1.Pod {
	for k, _ := range pod.Spec.Containers {
		hasValue := isExistEnv(key, pod.Spec.Containers[k].Env)
		switch {
		case hasValue == true:
			// 存在标签 不进行操作
		case hasValue == false:
			podenv := pod.Spec.Containers[k].Env
			pod.Spec.Containers[k].Env = append(podenv, corev1.EnvVar{
				Name:  key,
				Value: value,
			})

		default:
			// 白给
		}
	}

	return pod
}

func isExistEnv(key string, env []corev1.EnvVar) (hasValue bool) {
	for _, v := range env {
		if v.Name == key {
			return true
		}
	}
	return false
}
