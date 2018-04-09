// +build !ignore_autogenerated

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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiserver "k8s.io/apiserver/pkg/apis/apiserver"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_AdmissionConfiguration_To_apiserver_AdmissionConfiguration,
		Convert_apiserver_AdmissionConfiguration_To_v1alpha1_AdmissionConfiguration,
		Convert_v1alpha1_AdmissionPluginConfiguration_To_apiserver_AdmissionPluginConfiguration,
		Convert_apiserver_AdmissionPluginConfiguration_To_v1alpha1_AdmissionPluginConfiguration,
	)
}

func autoConvert_v1alpha1_AdmissionConfiguration_To_apiserver_AdmissionConfiguration(in *AdmissionConfiguration, out *apiserver.AdmissionConfiguration, s conversion.Scope) error {
	out.Plugins = *(*[]apiserver.AdmissionPluginConfiguration)(unsafe.Pointer(&in.Plugins))
	return nil
}

// Convert_v1alpha1_AdmissionConfiguration_To_apiserver_AdmissionConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_AdmissionConfiguration_To_apiserver_AdmissionConfiguration(in *AdmissionConfiguration, out *apiserver.AdmissionConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_AdmissionConfiguration_To_apiserver_AdmissionConfiguration(in, out, s)
}

func autoConvert_apiserver_AdmissionConfiguration_To_v1alpha1_AdmissionConfiguration(in *apiserver.AdmissionConfiguration, out *AdmissionConfiguration, s conversion.Scope) error {
	out.Plugins = *(*[]AdmissionPluginConfiguration)(unsafe.Pointer(&in.Plugins))
	return nil
}

// Convert_apiserver_AdmissionConfiguration_To_v1alpha1_AdmissionConfiguration is an autogenerated conversion function.
func Convert_apiserver_AdmissionConfiguration_To_v1alpha1_AdmissionConfiguration(in *apiserver.AdmissionConfiguration, out *AdmissionConfiguration, s conversion.Scope) error {
	return autoConvert_apiserver_AdmissionConfiguration_To_v1alpha1_AdmissionConfiguration(in, out, s)
}

func autoConvert_v1alpha1_AdmissionPluginConfiguration_To_apiserver_AdmissionPluginConfiguration(in *AdmissionPluginConfiguration, out *apiserver.AdmissionPluginConfiguration, s conversion.Scope) error {
	out.Name = in.Name
	out.Path = in.Path
	out.Configuration = (*runtime.Unknown)(unsafe.Pointer(in.Configuration))
	return nil
}

// Convert_v1alpha1_AdmissionPluginConfiguration_To_apiserver_AdmissionPluginConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_AdmissionPluginConfiguration_To_apiserver_AdmissionPluginConfiguration(in *AdmissionPluginConfiguration, out *apiserver.AdmissionPluginConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_AdmissionPluginConfiguration_To_apiserver_AdmissionPluginConfiguration(in, out, s)
}

func autoConvert_apiserver_AdmissionPluginConfiguration_To_v1alpha1_AdmissionPluginConfiguration(in *apiserver.AdmissionPluginConfiguration, out *AdmissionPluginConfiguration, s conversion.Scope) error {
	out.Name = in.Name
	out.Path = in.Path
	out.Configuration = (*runtime.Unknown)(unsafe.Pointer(in.Configuration))
	return nil
}

// Convert_apiserver_AdmissionPluginConfiguration_To_v1alpha1_AdmissionPluginConfiguration is an autogenerated conversion function.
func Convert_apiserver_AdmissionPluginConfiguration_To_v1alpha1_AdmissionPluginConfiguration(in *apiserver.AdmissionPluginConfiguration, out *AdmissionPluginConfiguration, s conversion.Scope) error {
	return autoConvert_apiserver_AdmissionPluginConfiguration_To_v1alpha1_AdmissionPluginConfiguration(in, out, s)
}
