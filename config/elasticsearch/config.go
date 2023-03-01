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

package elasticsearch

import (
	"errors"
	"strings"

	"github.com/crossplane/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-jet-aws/config/common"
)

// Configure adds configurations for elasticsearch group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_elasticsearch_domain", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.SetIdentifierArgumentFn = func(base map[string]interface{}, externalName string) {
			base["domain_name"] = externalName
		}
		r.ExternalName.OmittedFields = []string{
			"domain_name",
		}
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]interface{}) (string, error) {
			// expected id format: arn:aws:es:us-east-1:123456789123:domain/example-domain
			w := strings.Split(tfstate["id"].(string), "/")
			if len(w) != 2 {
				return "", errors.New("terraform ID should be the ARN of the es domain")
			}
			return w[len(w)-1], nil
		}
		r.References = config.References{
			"vpc_options.security_group_ids": config.Reference{
				Type:              "github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha2.SecurityGroup",
				RefFieldName:      "SecurityGroupRefs",
				SelectorFieldName: "SecurityGroupSelector",
			},
			"vpc_options.subnet_ids": config.Reference{
				Type:              "github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha2.Subnet",
				RefFieldName:      "SubnetRefs",
				SelectorFieldName: "SubnetSelector",
			},
			"encrypt_at_rest.kms_key_id": config.Reference{
				Type:      "github.com/crossplane-contrib/provider-jet-aws/apis/kms/v1alpha2.Key",
				Extractor: common.PathARNExtractor,
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_elasticsearch_domain_policy", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider

		r.References = map[string]config.Reference{
			"domain_name": {
				Type:              "Domain",
				RefFieldName:      "DomainRefs",
				SelectorFieldName: "DomainSelector",
			},
		}
	})

	p.AddResourceConfigurator("aws_elasticsearch_domain_saml_options", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider

		r.References = map[string]config.Reference{
			"domain_name": {
				Type:              "Domain",
				RefFieldName:      "DomainRefs",
				SelectorFieldName: "DomainSelector",
			},
		}
	})
}
