//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2020 The KEDA Authors

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	auditv1 "k8s.io/apiserver/pkg/apis/audit/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuditConfig) DeepCopyInto(out *AuditConfig) {
	*out = *in
	in.Policy.DeepCopyInto(&out.Policy)
	out.AuditLifetime = in.AuditLifetime
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuditConfig.
func (in *AuditConfig) DeepCopy() *AuditConfig {
	if in == nil {
		return nil
	}
	out := new(AuditConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuditLifetime) DeepCopyInto(out *AuditLifetime) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuditLifetime.
func (in *AuditLifetime) DeepCopy() *AuditLifetime {
	if in == nil {
		return nil
	}
	out := new(AuditLifetime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuditPolicy) DeepCopyInto(out *AuditPolicy) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]auditv1.PolicyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OmitStages != nil {
		in, out := &in.OmitStages, &out.OmitStages
		*out = make([]auditv1.Stage, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuditPolicy.
func (in *AuditPolicy) DeepCopy() *AuditPolicy {
	if in == nil {
		return nil
	}
	out := new(AuditPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericDeploymentSpec) DeepCopyInto(out *GenericDeploymentSpec) {
	*out = *in
	if in.DeploymentAnnotations != nil {
		in, out := &in.DeploymentAnnotations, &out.DeploymentAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DeploymentLabels != nil {
		in, out := &in.DeploymentLabels, &out.DeploymentLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodAnnotations != nil {
		in, out := &in.PodAnnotations, &out.PodAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodLabels != nil {
		in, out := &in.PodLabels, &out.PodLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericDeploymentSpec.
func (in *GenericDeploymentSpec) DeepCopy() *GenericDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(GenericDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KedaAdmissionWebhooksSpec) DeepCopyInto(out *KedaAdmissionWebhooksSpec) {
	*out = *in
	in.GenericDeploymentSpec.DeepCopyInto(&out.GenericDeploymentSpec)
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KedaAdmissionWebhooksSpec.
func (in *KedaAdmissionWebhooksSpec) DeepCopy() *KedaAdmissionWebhooksSpec {
	if in == nil {
		return nil
	}
	out := new(KedaAdmissionWebhooksSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KedaController) DeepCopyInto(out *KedaController) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KedaController.
func (in *KedaController) DeepCopy() *KedaController {
	if in == nil {
		return nil
	}
	out := new(KedaController)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KedaController) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KedaControllerList) DeepCopyInto(out *KedaControllerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KedaController, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KedaControllerList.
func (in *KedaControllerList) DeepCopy() *KedaControllerList {
	if in == nil {
		return nil
	}
	out := new(KedaControllerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KedaControllerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KedaControllerSpec) DeepCopyInto(out *KedaControllerSpec) {
	*out = *in
	in.Operator.DeepCopyInto(&out.Operator)
	in.MetricsServer.DeepCopyInto(&out.MetricsServer)
	in.AdmissionWebhooks.DeepCopyInto(&out.AdmissionWebhooks)
	in.ServiceAccount.DeepCopyInto(&out.ServiceAccount)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KedaControllerSpec.
func (in *KedaControllerSpec) DeepCopy() *KedaControllerSpec {
	if in == nil {
		return nil
	}
	out := new(KedaControllerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KedaControllerStatus) DeepCopyInto(out *KedaControllerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KedaControllerStatus.
func (in *KedaControllerStatus) DeepCopy() *KedaControllerStatus {
	if in == nil {
		return nil
	}
	out := new(KedaControllerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KedaMetricsServerSpec) DeepCopyInto(out *KedaMetricsServerSpec) {
	*out = *in
	in.GenericDeploymentSpec.DeepCopyInto(&out.GenericDeploymentSpec)
	in.AuditConfig.DeepCopyInto(&out.AuditConfig)
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KedaMetricsServerSpec.
func (in *KedaMetricsServerSpec) DeepCopy() *KedaMetricsServerSpec {
	if in == nil {
		return nil
	}
	out := new(KedaMetricsServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KedaOperatorSpec) DeepCopyInto(out *KedaOperatorSpec) {
	*out = *in
	in.GenericDeploymentSpec.DeepCopyInto(&out.GenericDeploymentSpec)
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KedaOperatorSpec.
func (in *KedaOperatorSpec) DeepCopy() *KedaOperatorSpec {
	if in == nil {
		return nil
	}
	out := new(KedaOperatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KedaServiceAccountSpec) DeepCopyInto(out *KedaServiceAccountSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KedaServiceAccountSpec.
func (in *KedaServiceAccountSpec) DeepCopy() *KedaServiceAccountSpec {
	if in == nil {
		return nil
	}
	out := new(KedaServiceAccountSpec)
	in.DeepCopyInto(out)
	return out
}
