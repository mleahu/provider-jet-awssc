//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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

package v1alpha2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OutputsObservation) DeepCopyInto(out *OutputsObservation) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutputsObservation.
func (in *OutputsObservation) DeepCopy() *OutputsObservation {
	if in == nil {
		return nil
	}
	out := new(OutputsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OutputsParameters) DeepCopyInto(out *OutputsParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutputsParameters.
func (in *OutputsParameters) DeepCopy() *OutputsParameters {
	if in == nil {
		return nil
	}
	out := new(OutputsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisionedProduct) DeepCopyInto(out *ProvisionedProduct) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisionedProduct.
func (in *ProvisionedProduct) DeepCopy() *ProvisionedProduct {
	if in == nil {
		return nil
	}
	out := new(ProvisionedProduct)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProvisionedProduct) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisionedProductList) DeepCopyInto(out *ProvisionedProductList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProvisionedProduct, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisionedProductList.
func (in *ProvisionedProductList) DeepCopy() *ProvisionedProductList {
	if in == nil {
		return nil
	}
	out := new(ProvisionedProductList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProvisionedProductList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisionedProductObservation) DeepCopyInto(out *ProvisionedProductObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.CloudwatchDashboardNames != nil {
		in, out := &in.CloudwatchDashboardNames, &out.CloudwatchDashboardNames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.CreatedTime != nil {
		in, out := &in.CreatedTime, &out.CreatedTime
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.LastProvisioningRecordID != nil {
		in, out := &in.LastProvisioningRecordID, &out.LastProvisioningRecordID
		*out = new(string)
		**out = **in
	}
	if in.LastRecordID != nil {
		in, out := &in.LastRecordID, &out.LastRecordID
		*out = new(string)
		**out = **in
	}
	if in.LastSuccessfulProvisioningRecordID != nil {
		in, out := &in.LastSuccessfulProvisioningRecordID, &out.LastSuccessfulProvisioningRecordID
		*out = new(string)
		**out = **in
	}
	if in.LaunchRoleArn != nil {
		in, out := &in.LaunchRoleArn, &out.LaunchRoleArn
		*out = new(string)
		**out = **in
	}
	if in.Outputs != nil {
		in, out := &in.Outputs, &out.Outputs
		*out = make([]OutputsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.StatusMessage != nil {
		in, out := &in.StatusMessage, &out.StatusMessage
		*out = new(string)
		**out = **in
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisionedProductObservation.
func (in *ProvisionedProductObservation) DeepCopy() *ProvisionedProductObservation {
	if in == nil {
		return nil
	}
	out := new(ProvisionedProductObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisionedProductParameters) DeepCopyInto(out *ProvisionedProductParameters) {
	*out = *in
	if in.AcceptLanguage != nil {
		in, out := &in.AcceptLanguage, &out.AcceptLanguage
		*out = new(string)
		**out = **in
	}
	if in.IgnoreErrors != nil {
		in, out := &in.IgnoreErrors, &out.IgnoreErrors
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NotificationArns != nil {
		in, out := &in.NotificationArns, &out.NotificationArns
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PathID != nil {
		in, out := &in.PathID, &out.PathID
		*out = new(string)
		**out = **in
	}
	if in.PathName != nil {
		in, out := &in.PathName, &out.PathName
		*out = new(string)
		**out = **in
	}
	if in.ProductID != nil {
		in, out := &in.ProductID, &out.ProductID
		*out = new(string)
		**out = **in
	}
	if in.ProductName != nil {
		in, out := &in.ProductName, &out.ProductName
		*out = new(string)
		**out = **in
	}
	if in.ProvisioningArtifactID != nil {
		in, out := &in.ProvisioningArtifactID, &out.ProvisioningArtifactID
		*out = new(string)
		**out = **in
	}
	if in.ProvisioningArtifactName != nil {
		in, out := &in.ProvisioningArtifactName, &out.ProvisioningArtifactName
		*out = new(string)
		**out = **in
	}
	if in.ProvisioningParameters != nil {
		in, out := &in.ProvisioningParameters, &out.ProvisioningParameters
		*out = make([]ProvisioningParametersParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.RetainPhysicalResources != nil {
		in, out := &in.RetainPhysicalResources, &out.RetainPhysicalResources
		*out = new(bool)
		**out = **in
	}
	if in.StackSetProvisioningPreferences != nil {
		in, out := &in.StackSetProvisioningPreferences, &out.StackSetProvisioningPreferences
		*out = make([]StackSetProvisioningPreferencesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisionedProductParameters.
func (in *ProvisionedProductParameters) DeepCopy() *ProvisionedProductParameters {
	if in == nil {
		return nil
	}
	out := new(ProvisionedProductParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisionedProductSpec) DeepCopyInto(out *ProvisionedProductSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisionedProductSpec.
func (in *ProvisionedProductSpec) DeepCopy() *ProvisionedProductSpec {
	if in == nil {
		return nil
	}
	out := new(ProvisionedProductSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisionedProductStatus) DeepCopyInto(out *ProvisionedProductStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisionedProductStatus.
func (in *ProvisionedProductStatus) DeepCopy() *ProvisionedProductStatus {
	if in == nil {
		return nil
	}
	out := new(ProvisionedProductStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisioningArtifact) DeepCopyInto(out *ProvisioningArtifact) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisioningArtifact.
func (in *ProvisioningArtifact) DeepCopy() *ProvisioningArtifact {
	if in == nil {
		return nil
	}
	out := new(ProvisioningArtifact)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProvisioningArtifact) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisioningArtifactList) DeepCopyInto(out *ProvisioningArtifactList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProvisioningArtifact, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisioningArtifactList.
func (in *ProvisioningArtifactList) DeepCopy() *ProvisioningArtifactList {
	if in == nil {
		return nil
	}
	out := new(ProvisioningArtifactList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProvisioningArtifactList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisioningArtifactObservation) DeepCopyInto(out *ProvisioningArtifactObservation) {
	*out = *in
	if in.CreatedTime != nil {
		in, out := &in.CreatedTime, &out.CreatedTime
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisioningArtifactObservation.
func (in *ProvisioningArtifactObservation) DeepCopy() *ProvisioningArtifactObservation {
	if in == nil {
		return nil
	}
	out := new(ProvisioningArtifactObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisioningArtifactParameters) DeepCopyInto(out *ProvisioningArtifactParameters) {
	*out = *in
	if in.AcceptLanguage != nil {
		in, out := &in.AcceptLanguage, &out.AcceptLanguage
		*out = new(string)
		**out = **in
	}
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisableTemplateValidation != nil {
		in, out := &in.DisableTemplateValidation, &out.DisableTemplateValidation
		*out = new(bool)
		**out = **in
	}
	if in.Guidance != nil {
		in, out := &in.Guidance, &out.Guidance
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ProductID != nil {
		in, out := &in.ProductID, &out.ProductID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.TemplatePhysicalID != nil {
		in, out := &in.TemplatePhysicalID, &out.TemplatePhysicalID
		*out = new(string)
		**out = **in
	}
	if in.TemplateURL != nil {
		in, out := &in.TemplateURL, &out.TemplateURL
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisioningArtifactParameters.
func (in *ProvisioningArtifactParameters) DeepCopy() *ProvisioningArtifactParameters {
	if in == nil {
		return nil
	}
	out := new(ProvisioningArtifactParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisioningArtifactSpec) DeepCopyInto(out *ProvisioningArtifactSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisioningArtifactSpec.
func (in *ProvisioningArtifactSpec) DeepCopy() *ProvisioningArtifactSpec {
	if in == nil {
		return nil
	}
	out := new(ProvisioningArtifactSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisioningArtifactStatus) DeepCopyInto(out *ProvisioningArtifactStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisioningArtifactStatus.
func (in *ProvisioningArtifactStatus) DeepCopy() *ProvisioningArtifactStatus {
	if in == nil {
		return nil
	}
	out := new(ProvisioningArtifactStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisioningParametersObservation) DeepCopyInto(out *ProvisioningParametersObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisioningParametersObservation.
func (in *ProvisioningParametersObservation) DeepCopy() *ProvisioningParametersObservation {
	if in == nil {
		return nil
	}
	out := new(ProvisioningParametersObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisioningParametersParameters) DeepCopyInto(out *ProvisioningParametersParameters) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.UsePreviousValue != nil {
		in, out := &in.UsePreviousValue, &out.UsePreviousValue
		*out = new(bool)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisioningParametersParameters.
func (in *ProvisioningParametersParameters) DeepCopy() *ProvisioningParametersParameters {
	if in == nil {
		return nil
	}
	out := new(ProvisioningParametersParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackSetProvisioningPreferencesObservation) DeepCopyInto(out *StackSetProvisioningPreferencesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackSetProvisioningPreferencesObservation.
func (in *StackSetProvisioningPreferencesObservation) DeepCopy() *StackSetProvisioningPreferencesObservation {
	if in == nil {
		return nil
	}
	out := new(StackSetProvisioningPreferencesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackSetProvisioningPreferencesParameters) DeepCopyInto(out *StackSetProvisioningPreferencesParameters) {
	*out = *in
	if in.Accounts != nil {
		in, out := &in.Accounts, &out.Accounts
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.FailureToleranceCount != nil {
		in, out := &in.FailureToleranceCount, &out.FailureToleranceCount
		*out = new(float64)
		**out = **in
	}
	if in.FailureTolerancePercentage != nil {
		in, out := &in.FailureTolerancePercentage, &out.FailureTolerancePercentage
		*out = new(float64)
		**out = **in
	}
	if in.MaxConcurrencyCount != nil {
		in, out := &in.MaxConcurrencyCount, &out.MaxConcurrencyCount
		*out = new(float64)
		**out = **in
	}
	if in.MaxConcurrencyPercentage != nil {
		in, out := &in.MaxConcurrencyPercentage, &out.MaxConcurrencyPercentage
		*out = new(float64)
		**out = **in
	}
	if in.Regions != nil {
		in, out := &in.Regions, &out.Regions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackSetProvisioningPreferencesParameters.
func (in *StackSetProvisioningPreferencesParameters) DeepCopy() *StackSetProvisioningPreferencesParameters {
	if in == nil {
		return nil
	}
	out := new(StackSetProvisioningPreferencesParameters)
	in.DeepCopyInto(out)
	return out
}
