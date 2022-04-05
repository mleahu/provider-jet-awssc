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

package servicecatalog

import (
	"github.com/crossplane-contrib/provider-jet-awssc/config/common"
	"github.com/crossplane/terrajet/pkg/config"
)

// Configure adds configurations for servicecatalog group.
func Configure(p *config.Provider) {

	p.AddResourceConfigurator("aws_servicecatalog_provisioning_artifact", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		// r.ExternalName = config.IdentifierFromProvider
	})

	p.AddResourceConfigurator("aws_servicecatalog_provisioned_product", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		// r.ExternalName = config.IdentifierFromProvider
		r.UseAsync = true

		/*		r.References["subnet_ids"] = config.Reference{
					//Type:              "github.com/crossplane/provider-aws/apis/ec2/v1beta1.Subnet",
					//RefFieldName:      "SubnetIdRefs",
					//SelectorFieldName: "SubnetIdSelector",
					Type:              "github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha2.Subnet",
					RefFieldName:      "SubnetIdRefs",
					SelectorFieldName: "SubnetIdSelector",
				}
		*/
	})

}
