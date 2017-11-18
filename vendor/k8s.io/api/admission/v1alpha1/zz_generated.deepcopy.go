// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdmissionRequest) DeepCopyInto(out *AdmissionRequest) {
	*out = *in
	out.Kind = in.Kind
	out.Resource = in.Resource
	in.UserInfo.DeepCopyInto(&out.UserInfo)
	in.Object.DeepCopyInto(&out.Object)
	in.OldObject.DeepCopyInto(&out.OldObject)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdmissionRequest.
func (in *AdmissionRequest) DeepCopy() *AdmissionRequest {
	if in == nil {
		return nil
	}
	out := new(AdmissionRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdmissionResponse) DeepCopyInto(out *AdmissionResponse) {
	*out = *in
	if in.Result != nil {
		in, out := &in.Result, &out.Result
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Status)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Patch != nil {
		in, out := &in.Patch, &out.Patch
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.PatchType != nil {
		in, out := &in.PatchType, &out.PatchType
		if *in == nil {
			*out = nil
		} else {
			*out = new(PatchType)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdmissionResponse.
func (in *AdmissionResponse) DeepCopy() *AdmissionResponse {
	if in == nil {
		return nil
	}
	out := new(AdmissionResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdmissionReview) DeepCopyInto(out *AdmissionReview) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Request != nil {
		in, out := &in.Request, &out.Request
		if *in == nil {
			*out = nil
		} else {
			*out = new(AdmissionRequest)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Response != nil {
		in, out := &in.Response, &out.Response
		if *in == nil {
			*out = nil
		} else {
			*out = new(AdmissionResponse)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdmissionReview.
func (in *AdmissionReview) DeepCopy() *AdmissionReview {
	if in == nil {
		return nil
	}
	out := new(AdmissionReview)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AdmissionReview) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}
