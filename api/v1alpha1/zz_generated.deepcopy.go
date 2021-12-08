// +build !ignore_autogenerated

/*

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
	"github.com/brancz/locutus/feedback"
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIRBAC) DeepCopyInto(out *APIRBAC) {
	*out = *in
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]RBACRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RoleBindings != nil {
		in, out := &in.RoleBindings, &out.RoleBindings
		*out = make([]RBACRoleBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIRBAC.
func (in *APIRBAC) DeepCopy() *APIRBAC {
	if in == nil {
		return nil
	}
	out := new(APIRBAC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APISpec) DeepCopyInto(out *APISpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	out.TLS = in.TLS
	in.RBAC.DeepCopyInto(&out.RBAC)
	if in.Tenants != nil {
		in, out := &in.Tenants, &out.Tenants
		*out = make([]APITenant, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APISpec.
func (in *APISpec) DeepCopy() *APISpec {
	if in == nil {
		return nil
	}
	out := new(APISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APITenant) DeepCopyInto(out *APITenant) {
	*out = *in
	if in.OIDC != nil {
		in, out := &in.OIDC, &out.OIDC
		*out = new(TenantOIDC)
		**out = **in
	}
	if in.MTLS != nil {
		in, out := &in.MTLS, &out.MTLS
		*out = new(TenantMTLS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APITenant.
func (in *APITenant) DeepCopy() *APITenant {
	if in == nil {
		return nil
	}
	out := new(APITenant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompactSpec) DeepCopyInto(out *CompactSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	in.VolumeClaimTemplate.DeepCopyInto(&out.VolumeClaimTemplate)
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompactSpec.
func (in *CompactSpec) DeepCopy() *CompactSpec {
	if in == nil {
		return nil
	}
	out := new(CompactSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hashring) DeepCopyInto(out *Hashring) {
	*out = *in
	if in.Tenants != nil {
		in, out := &in.Tenants, &out.Tenants
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hashring.
func (in *Hashring) DeepCopy() *Hashring {
	if in == nil {
		return nil
	}
	out := new(Hashring)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LokiObjectStorageConfigSpec) DeepCopyInto(out *LokiObjectStorageConfigSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LokiObjectStorageConfigSpec.
func (in *LokiObjectStorageConfigSpec) DeepCopy() *LokiObjectStorageConfigSpec {
	if in == nil {
		return nil
	}
	out := new(LokiObjectStorageConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LokiSpec) DeepCopyInto(out *LokiSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = make(map[string]int32, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.VolumeClaimTemplate.DeepCopyInto(&out.VolumeClaimTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LokiSpec.
func (in *LokiSpec) DeepCopy() *LokiSpec {
	if in == nil {
		return nil
	}
	out := new(LokiSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStorageConfig) DeepCopyInto(out *ObjectStorageConfig) {
	*out = *in
	if in.Thanos != nil {
		in, out := &in.Thanos, &out.Thanos
		*out = new(ThanosObjectStorageConfigSpec)
		**out = **in
	}
	if in.Loki != nil {
		in, out := &in.Loki, &out.Loki
		*out = new(LokiObjectStorageConfigSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStorageConfig.
func (in *ObjectStorageConfig) DeepCopy() *ObjectStorageConfig {
	if in == nil {
		return nil
	}
	out := new(ObjectStorageConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Observatorium) DeepCopyInto(out *Observatorium) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Observatorium.
func (in *Observatorium) DeepCopy() *Observatorium {
	if in == nil {
		return nil
	}
	out := new(Observatorium)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Observatorium) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservatoriumList) DeepCopyInto(out *ObservatoriumList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Observatorium, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservatoriumList.
func (in *ObservatoriumList) DeepCopy() *ObservatoriumList {
	if in == nil {
		return nil
	}
	out := new(ObservatoriumList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObservatoriumList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservatoriumSpec) DeepCopyInto(out *ObservatoriumSpec) {
	*out = *in
	in.ObjectStorageConfig.DeepCopyInto(&out.ObjectStorageConfig)
	if in.Hashrings != nil {
		in, out := &in.Hashrings, &out.Hashrings
		*out = make([]*Hashring, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Hashring)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	in.Compact.DeepCopyInto(&out.Compact)
	in.ThanosReceiveController.DeepCopyInto(&out.ThanosReceiveController)
	in.Receivers.DeepCopyInto(&out.Receivers)
	in.QueryFrontend.DeepCopyInto(&out.QueryFrontend)
	in.Store.DeepCopyInto(&out.Store)
	in.Rule.DeepCopyInto(&out.Rule)
	in.API.DeepCopyInto(&out.API)
	in.Query.DeepCopyInto(&out.Query)
	if in.Loki != nil {
		in, out := &in.Loki, &out.Loki
		*out = new(LokiSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservatoriumSpec.
func (in *ObservatoriumSpec) DeepCopy() *ObservatoriumSpec {
	if in == nil {
		return nil
	}
	out := new(ObservatoriumSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservatoriumStatus) DeepCopyInto(out *ObservatoriumStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*feedback.StatusCondition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(feedback.StatusCondition)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservatoriumStatus.
func (in *ObservatoriumStatus) DeepCopy() *ObservatoriumStatus {
	if in == nil {
		return nil
	}
	out := new(ObservatoriumStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryFrontendSpec) DeepCopyInto(out *QueryFrontendSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryFrontendSpec.
func (in *QueryFrontendSpec) DeepCopy() *QueryFrontendSpec {
	if in == nil {
		return nil
	}
	out := new(QueryFrontendSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuerySpec) DeepCopyInto(out *QuerySpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuerySpec.
func (in *QuerySpec) DeepCopy() *QuerySpec {
	if in == nil {
		return nil
	}
	out := new(QuerySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RBACRole) DeepCopyInto(out *RBACRole) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tenants != nil {
		in, out := &in.Tenants, &out.Tenants
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]Permission, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RBACRole.
func (in *RBACRole) DeepCopy() *RBACRole {
	if in == nil {
		return nil
	}
	out := new(RBACRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RBACRoleBinding) DeepCopyInto(out *RBACRoleBinding) {
	*out = *in
	if in.Subjects != nil {
		in, out := &in.Subjects, &out.Subjects
		*out = make([]Subject, len(*in))
		copy(*out, *in)
	}
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RBACRoleBinding.
func (in *RBACRoleBinding) DeepCopy() *RBACRoleBinding {
	if in == nil {
		return nil
	}
	out := new(RBACRoleBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReceiversSpec) DeepCopyInto(out *ReceiversSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	in.VolumeClaimTemplate.DeepCopyInto(&out.VolumeClaimTemplate)
	if in.ReplicationFactor != nil {
		in, out := &in.ReplicationFactor, &out.ReplicationFactor
		*out = new(int32)
		**out = **in
	}
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReceiversSpec.
func (in *ReceiversSpec) DeepCopy() *ReceiversSpec {
	if in == nil {
		return nil
	}
	out := new(ReceiversSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleConfig) DeepCopyInto(out *RuleConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleConfig.
func (in *RuleConfig) DeepCopy() *RuleConfig {
	if in == nil {
		return nil
	}
	out := new(RuleConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleSpec) DeepCopyInto(out *RuleSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	in.VolumeClaimTemplate.DeepCopyInto(&out.VolumeClaimTemplate)
	if in.RulesConfig != nil {
		in, out := &in.RulesConfig, &out.RulesConfig
		*out = make([]RuleConfig, len(*in))
		copy(*out, *in)
	}
	if in.AlertmanagerURLs != nil {
		in, out := &in.AlertmanagerURLs, &out.AlertmanagerURLs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Resources.DeepCopyInto(&out.Resources)
	in.ReloaderResources.DeepCopyInto(&out.ReloaderResources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleSpec.
func (in *RuleSpec) DeepCopy() *RuleSpec {
	if in == nil {
		return nil
	}
	out := new(RuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoreCacheSpec) DeepCopyInto(out *StoreCacheSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.MemoryLimitMB != nil {
		in, out := &in.MemoryLimitMB, &out.MemoryLimitMB
		*out = new(int32)
		**out = **in
	}
	in.Resources.DeepCopyInto(&out.Resources)
	in.ExporterResources.DeepCopyInto(&out.ExporterResources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoreCacheSpec.
func (in *StoreCacheSpec) DeepCopy() *StoreCacheSpec {
	if in == nil {
		return nil
	}
	out := new(StoreCacheSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoreSpec) DeepCopyInto(out *StoreSpec) {
	*out = *in
	in.VolumeClaimTemplate.DeepCopyInto(&out.VolumeClaimTemplate)
	if in.Shards != nil {
		in, out := &in.Shards, &out.Shards
		*out = new(int32)
		**out = **in
	}
	in.Cache.DeepCopyInto(&out.Cache)
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoreSpec.
func (in *StoreSpec) DeepCopy() *StoreSpec {
	if in == nil {
		return nil
	}
	out := new(StoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subject) DeepCopyInto(out *Subject) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subject.
func (in *Subject) DeepCopy() *Subject {
	if in == nil {
		return nil
	}
	out := new(Subject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLS) DeepCopyInto(out *TLS) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLS.
func (in *TLS) DeepCopy() *TLS {
	if in == nil {
		return nil
	}
	out := new(TLS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantMTLS) DeepCopyInto(out *TenantMTLS) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantMTLS.
func (in *TenantMTLS) DeepCopy() *TenantMTLS {
	if in == nil {
		return nil
	}
	out := new(TenantMTLS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantOIDC) DeepCopyInto(out *TenantOIDC) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantOIDC.
func (in *TenantOIDC) DeepCopy() *TenantOIDC {
	if in == nil {
		return nil
	}
	out := new(TenantOIDC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosObjectStorageConfigSpec) DeepCopyInto(out *ThanosObjectStorageConfigSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosObjectStorageConfigSpec.
func (in *ThanosObjectStorageConfigSpec) DeepCopy() *ThanosObjectStorageConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ThanosObjectStorageConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosReceiveControllerSpec) DeepCopyInto(out *ThanosReceiveControllerSpec) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosReceiveControllerSpec.
func (in *ThanosReceiveControllerSpec) DeepCopy() *ThanosReceiveControllerSpec {
	if in == nil {
		return nil
	}
	out := new(ThanosReceiveControllerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeClaimTemplate) DeepCopyInto(out *VolumeClaimTemplate) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeClaimTemplate.
func (in *VolumeClaimTemplate) DeepCopy() *VolumeClaimTemplate {
	if in == nil {
		return nil
	}
	out := new(VolumeClaimTemplate)
	in.DeepCopyInto(out)
	return out
}
