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

// Code generated by terrajet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha2 "github.com/crossplane-contrib/provider-jet-aws/apis/autoscaling/v1alpha2"
	v1alpha2ec2 "github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha2"
	v1alpha2ecr "github.com/crossplane-contrib/provider-jet-aws/apis/ecr/v1alpha2"
	v1alpha2ecrpublic "github.com/crossplane-contrib/provider-jet-aws/apis/ecrpublic/v1alpha2"
	v1alpha2ecs "github.com/crossplane-contrib/provider-jet-aws/apis/ecs/v1alpha2"
	v1alpha2eks "github.com/crossplane-contrib/provider-jet-aws/apis/eks/v1alpha2"
	v1alpha2elasticache "github.com/crossplane-contrib/provider-jet-aws/apis/elasticache/v1alpha2"
	v1alpha2elasticsearch "github.com/crossplane-contrib/provider-jet-aws/apis/elasticsearch/v1alpha2"
	v1alpha2elbv2 "github.com/crossplane-contrib/provider-jet-aws/apis/elbv2/v1alpha2"
	v1alpha2globalaccelerator "github.com/crossplane-contrib/provider-jet-aws/apis/globalaccelerator/v1alpha2"
	v1alpha2iam "github.com/crossplane-contrib/provider-jet-aws/apis/iam/v1alpha2"
	v1alpha2kms "github.com/crossplane-contrib/provider-jet-aws/apis/kms/v1alpha2"
	v1alpha2mq "github.com/crossplane-contrib/provider-jet-aws/apis/mq/v1alpha2"
	v1alpha2neptune "github.com/crossplane-contrib/provider-jet-aws/apis/neptune/v1alpha2"
	v1alpha2rds "github.com/crossplane-contrib/provider-jet-aws/apis/rds/v1alpha2"
	v1alpha2route53 "github.com/crossplane-contrib/provider-jet-aws/apis/route53/v1alpha2"
	v1alpha1 "github.com/crossplane-contrib/provider-jet-aws/apis/route53resolver/v1alpha1"
	v1alpha2s3 "github.com/crossplane-contrib/provider-jet-aws/apis/s3/v1alpha2"
	v1alpha1apis "github.com/crossplane-contrib/provider-jet-aws/apis/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha2.SchemeBuilder.AddToScheme,
		v1alpha2ec2.SchemeBuilder.AddToScheme,
		v1alpha2ecr.SchemeBuilder.AddToScheme,
		v1alpha2ecrpublic.SchemeBuilder.AddToScheme,
		v1alpha2ecs.SchemeBuilder.AddToScheme,
		v1alpha2eks.SchemeBuilder.AddToScheme,
		v1alpha2elasticache.SchemeBuilder.AddToScheme,
		v1alpha2elasticsearch.SchemeBuilder.AddToScheme,
		v1alpha2elbv2.SchemeBuilder.AddToScheme,
		v1alpha2globalaccelerator.SchemeBuilder.AddToScheme,
		v1alpha2iam.SchemeBuilder.AddToScheme,
		v1alpha2kms.SchemeBuilder.AddToScheme,
		v1alpha2mq.SchemeBuilder.AddToScheme,
		v1alpha2neptune.SchemeBuilder.AddToScheme,
		v1alpha2rds.SchemeBuilder.AddToScheme,
		v1alpha2route53.SchemeBuilder.AddToScheme,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha2s3.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
