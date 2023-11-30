//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
  Copyright The Kubernetes Authors.

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

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package openapi

import (
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/srl-labs/clabernetes/apis/v1alpha1.Definition": schema_srl_labs_clabernetes_apis_v1alpha1_Definition(
			ref,
		),
		"github.com/srl-labs/clabernetes/apis/v1alpha1.Deployment": schema_srl_labs_clabernetes_apis_v1alpha1_Deployment(
			ref,
		),
		"github.com/srl-labs/clabernetes/apis/v1alpha1.Expose": schema_srl_labs_clabernetes_apis_v1alpha1_Expose(
			ref,
		),
		"github.com/srl-labs/clabernetes/apis/v1alpha1.ExposedPorts": schema_srl_labs_clabernetes_apis_v1alpha1_ExposedPorts(
			ref,
		),
		"github.com/srl-labs/clabernetes/apis/v1alpha1.FileFromConfigMap": schema_srl_labs_clabernetes_apis_v1alpha1_FileFromConfigMap(
			ref,
		),
		"github.com/srl-labs/clabernetes/apis/v1alpha1.FileFromURL": schema_srl_labs_clabernetes_apis_v1alpha1_FileFromURL(
			ref,
		),
		"github.com/srl-labs/clabernetes/apis/v1alpha1.ImagePull": schema_srl_labs_clabernetes_apis_v1alpha1_ImagePull(
			ref,
		),
		"github.com/srl-labs/clabernetes/apis/v1alpha1.LinkEndpoint": schema_srl_labs_clabernetes_apis_v1alpha1_LinkEndpoint(
			ref,
		),
		"github.com/srl-labs/clabernetes/apis/v1alpha1.Persistence": schema_srl_labs_clabernetes_apis_v1alpha1_Persistence(
			ref,
		),
		"github.com/srl-labs/clabernetes/apis/v1alpha1.Topology": schema_srl_labs_clabernetes_apis_v1alpha1_Topology(
			ref,
		),
		"github.com/srl-labs/clabernetes/apis/v1alpha1.TopologyList": schema_srl_labs_clabernetes_apis_v1alpha1_TopologyList(
			ref,
		),
		"github.com/srl-labs/clabernetes/apis/v1alpha1.TopologySpec": schema_srl_labs_clabernetes_apis_v1alpha1_TopologySpec(
			ref,
		),
		"github.com/srl-labs/clabernetes/apis/v1alpha1.TopologyStatus": schema_srl_labs_clabernetes_apis_v1alpha1_TopologyStatus(
			ref,
		),
		"github.com/srl-labs/clabernetes/apis/v1alpha1.Tunnel": schema_srl_labs_clabernetes_apis_v1alpha1_Tunnel(
			ref,
		),
	}
}

func schema_srl_labs_clabernetes_apis_v1alpha1_Definition(
	ref common.ReferenceCallback,
) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Definition holds the underlying topology definition for the Topology CR. A Topology *must* have one -- and only one -- definition type defined.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"containerlab": {
						SchemaProps: spec.SchemaProps{
							Description: "Containerlab holds a valid containerlab topology.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"kne": {
						SchemaProps: spec.SchemaProps{
							Description: "Kne holds a valid kne topology.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}

func schema_srl_labs_clabernetes_apis_v1alpha1_Deployment(
	ref common.ReferenceCallback,
) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Deployment holds configurations relevant to how clabernetes configures deployments that make up a given topology.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"resources": {
						SchemaProps: spec.SchemaProps{
							Description: "Resources is a mapping of nodeName (or \"default\") to kubernetes resource requirements -- any value set here overrides the \"global\" config resource definitions. If a key \"default\" is set, those resource values will be preferred over *all global settings* for this topology -- meaning, the \"global\" resource settings will never be looked up for this topology, and any kind/type that is *not* in this resources map will have the \"default\" resources from this mapping applied.",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/api/core/v1.ResourceRequirements"),
									},
								},
							},
						},
					},
					"privilegedLauncher": {
						SchemaProps: spec.SchemaProps{
							Description: "PrivilegedLauncher, when true, sets the launcher containers to privileged. By default, we do our best to *not* need this/set this, and instead set only the capabilities we need, however its possible that some containers launched by the launcher may need/want more capabilities, so this flag exists for users to bypass the default settings and enable fully privileged launcher pods.",
							Default:     false,
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"filesFromConfigMap": {
						SchemaProps: spec.SchemaProps{
							Description: "FilesFromConfigMap is a slice of FileFromConfigMap that define the configmap/path and node and path on a launcher node that the file should be mounted to. If the path is not provided the configmap is mounted in its entirety (like normal k8s things), so you *probably* want to specify the sub path unless you are sure what you're doing!",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type: []string{"array"},
										Items: &spec.SchemaOrArray{
											Schema: &spec.Schema{
												SchemaProps: spec.SchemaProps{
													Default: map[string]interface{}{},
													Ref: ref(
														"github.com/srl-labs/clabernetes/apis/v1alpha1.FileFromConfigMap",
													),
												},
											},
										},
									},
								},
							},
						},
					},
					"filesFromURL": {
						SchemaProps: spec.SchemaProps{
							Description: "FilesFromURL is a mapping of FileFromURL that define a URL at which to fetch a file, and path on a launcher node that the file should be downloaded to. This is useful for configs that are larger than the ConfigMap (etcd) 1Mb size limit.",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type: []string{"array"},
										Items: &spec.SchemaOrArray{
											Schema: &spec.Schema{
												SchemaProps: spec.SchemaProps{
													Default: map[string]interface{}{},
													Ref: ref(
														"github.com/srl-labs/clabernetes/apis/v1alpha1.FileFromURL",
													),
												},
											},
										},
									},
								},
							},
						},
					},
					"persistence": {
						SchemaProps: spec.SchemaProps{
							Description: "Persistence holds configurations relating to persisting each nodes working containerlab directory.",
							Default:     map[string]interface{}{},
							Ref: ref(
								"github.com/srl-labs/clabernetes/apis/v1alpha1.Persistence",
							),
						},
					},
					"containerlabDebug": {
						SchemaProps: spec.SchemaProps{
							Description: "ContainerlabDebug sets the `--debug` flag when invoking containerlab in the launcher pods. This is disabled by default.",
							Default:     false,
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"launcherLogLevel": {
						SchemaProps: spec.SchemaProps{
							Description: "LauncherLogLevel sets the launcher clabernetes worker log level -- this overrides whatever is set on the controllers env vars for this topology. Note: omitempty because empty str does not satisfy enum of course.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/srl-labs/clabernetes/apis/v1alpha1.FileFromConfigMap", "github.com/srl-labs/clabernetes/apis/v1alpha1.FileFromURL", "github.com/srl-labs/clabernetes/apis/v1alpha1.Persistence", "k8s.io/api/core/v1.ResourceRequirements"},
	}
}

func schema_srl_labs_clabernetes_apis_v1alpha1_Expose(
	ref common.ReferenceCallback,
) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Expose holds configurations relevant to how clabernetes exposes a topology.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"disableNodeAliasService": {
						SchemaProps: spec.SchemaProps{
							Description: "DisableNodeAliasService indicates if headless services for each node in a containerlab topology should *not* be created. By default, clabernetes creates these headless services for each node so that \"normal\" docker and containerlab service discovery works -- this means you can simply resolve \"my-neat-node\" from within the namespace of a topology like you would in docker locally. You may wish to disable this feature though if you have no need of it and just don't want the extra services around. Additionally, you may want to disable this feature if you are running multiple labs in the same namespace (which is not generally recommended by the way!) as you may end up in a situation where a name (i.e. \"leaf1\") is duplicated in more than one topology -- this will cause some problems for clabernetes!",
							Default:     false,
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"disableExpose": {
						SchemaProps: spec.SchemaProps{
							Description: "DisableExpose indicates if exposing nodes via LoadBalancer service should be disabled, by default any mapped ports in a containerlab topology will be exposed.",
							Default:     false,
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"disableAutoExpose": {
						SchemaProps: spec.SchemaProps{
							Description: "DisableAutoExpose disables the automagic exposing of ports for a given topology. When this setting is disabled clabernetes will not auto add ports so if you want to expose (via a load balancer service) you will need to have ports outlined in your containerlab config (or equivalent for kne). When this is `false` (default), clabernetes will add and expose the following list of ports to whatever ports you have already defined:\n\n21    - tcp - ftp 22    - tcp - ssh 23    - tcp - telnet 80    - tcp - http 161   - udp - snmp 443   - tcp - https 830   - tcp - netconf (over ssh) 5000  - tcp - telnet for vrnetlab qemu host 5900  - tcp - vnc 6030  - tcp - gnmi (arista default) 9339  - tcp - gnmi/gnoi 9340  - tcp - gribi 9559  - tcp - p4rt 57400 - tcp - gnmi (nokia srl/sros default)\n\nThis setting is *ignored completely* if `DisableExpose` is true!",
							Default:     false,
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}

func schema_srl_labs_clabernetes_apis_v1alpha1_ExposedPorts(
	ref common.ReferenceCallback,
) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ExposedPorts holds information about exposed ports.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"loadBalancerAddress": {
						SchemaProps: spec.SchemaProps{
							Description: "LoadBalancerAddress holds the address assigned to the load balancer exposing ports for a given node.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"tcpPorts": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "TCPPorts is a list of TCP ports exposed on the LoadBalancer service.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: 0,
										Type:    []string{"integer"},
										Format:  "int32",
									},
								},
							},
						},
					},
					"udpPorts": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "UDPPorts is a list of UDP ports exposed on the LoadBalancer service.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: 0,
										Type:    []string{"integer"},
										Format:  "int32",
									},
								},
							},
						},
					},
				},
				Required: []string{"loadBalancerAddress", "tcpPorts", "udpPorts"},
			},
		},
	}
}

func schema_srl_labs_clabernetes_apis_v1alpha1_FileFromConfigMap(
	ref common.ReferenceCallback,
) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "FileFromConfigMap represents a file that you would like to mount (from a configmap) in the launcher pod for a given node.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"filePath": {
						SchemaProps: spec.SchemaProps{
							Description: "FilePath is the path to mount the file.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"configMapName": {
						SchemaProps: spec.SchemaProps{
							Description: "ConfigMapName is the name of the configmap to mount.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"configMapPath": {
						SchemaProps: spec.SchemaProps{
							Description: "ConfigMapPath is the path/key in the configmap to mount, if not specified the configmap will be mounted without a sub-path.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"filePath", "configMapName"},
			},
		},
	}
}

func schema_srl_labs_clabernetes_apis_v1alpha1_FileFromURL(
	ref common.ReferenceCallback,
) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "FileFromURL represents a file that you would like to mount from a URL in the launcher pod for a given node.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"filePath": {
						SchemaProps: spec.SchemaProps{
							Description: "FilePath is the path to mount the file.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"url": {
						SchemaProps: spec.SchemaProps{
							Description: "URL is the url to fetch and mount at the provided FilePath. This URL must be a url that can be simply downloaded and dumped to disk -- meaning a normal file server type endpoint or if using GitHub or similar a \"raw\" path.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"filePath", "url"},
			},
		},
	}
}

func schema_srl_labs_clabernetes_apis_v1alpha1_ImagePull(
	ref common.ReferenceCallback,
) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ImagePull holds configurations relevant to how clabernetes launcher pods handle pulling images.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"insecureRegistries": {
						SchemaProps: spec.SchemaProps{
							Description: "InsecureRegistries is a slice of strings of insecure registries to configure in the launcher pods.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"pullThroughOverride": {
						SchemaProps: spec.SchemaProps{
							Description: "PullThroughOverride allows for overriding the image pull through mode for this particular topology.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"pullSecrets": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "PullSecrets allows for providing secret(s) to use when pulling the image. This is only applicable *if* ImagePullThrough mode is auto or always. The secret is used by the launcher pod to pull the image via the cluster CRI. The secret is *not* mounted to the pod, but instead is used in conjunction with a job that spawns a pod using the specified secret. The job will kill the pod as soon as the image has been pulled -- we do this because we don't care if the pod runs, we only care that the image gets pulled on a specific node. Note that just like \"normal\" pull secrets, the secret needs to be in the namespace that the topology is in.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func schema_srl_labs_clabernetes_apis_v1alpha1_LinkEndpoint(
	ref common.ReferenceCallback,
) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "LinkEndpoint is a simple struct to hold node/interface name info for a given link.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"nodeName": {
						SchemaProps: spec.SchemaProps{
							Description: "NodeName is the name of the node this link resides on.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"interfaceName": {
						SchemaProps: spec.SchemaProps{
							Description: "InterfaceName is the name of the interface on the node this link is on.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"nodeName", "interfaceName"},
			},
		},
	}
}

func schema_srl_labs_clabernetes_apis_v1alpha1_Persistence(
	ref common.ReferenceCallback,
) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Persistence holds information about how to persist the containlerab lab directory for each node in a topology.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"enabled": {
						SchemaProps: spec.SchemaProps{
							Description: "Enabled indicates if persistence of hte containerlab lab/working directory will be placed in a mounted PVC.",
							Default:     false,
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"claimSize": {
						SchemaProps: spec.SchemaProps{
							Description: "ClaimSize is the size of the PVC for this topology -- if not provided this defaults to 5Gi. If provided, the string value must be a valid kubernetes storage requests style string. Note the claim size *cannot be made smaller* once created, but it *can* be expanded. If you need to make the claim smaller you must delete the topology (or the node from the topology) and re-add it.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"storageClassName": {
						SchemaProps: spec.SchemaProps{
							Description: "StorageClassName is the storage class to set in the PVC -- if not provided this will be left empty which will end up using your default storage class. Note that currently we assume you have (as default) or provide a dynamically provisionable storage class, hence no selector.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"enabled"},
			},
		},
	}
}

func schema_srl_labs_clabernetes_apis_v1alpha1_Topology(
	ref common.ReferenceCallback,
) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Topology is an object that holds information about a clabernetes Topology -- that is, a valid topology file (ex: containerlab topology), and any associated configurations.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref: ref(
								"github.com/srl-labs/clabernetes/apis/v1alpha1.TopologySpec",
							),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref: ref(
								"github.com/srl-labs/clabernetes/apis/v1alpha1.TopologyStatus",
							),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/srl-labs/clabernetes/apis/v1alpha1.TopologySpec", "github.com/srl-labs/clabernetes/apis/v1alpha1.TopologyStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_srl_labs_clabernetes_apis_v1alpha1_TopologyList(
	ref common.ReferenceCallback,
) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TopologyList is a list of Topology objects.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref: ref(
											"github.com/srl-labs/clabernetes/apis/v1alpha1.Topology",
										),
									},
								},
							},
						},
					},
				},
				Required: []string{"items"},
			},
		},
		Dependencies: []string{
			"github.com/srl-labs/clabernetes/apis/v1alpha1.Topology", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_srl_labs_clabernetes_apis_v1alpha1_TopologySpec(
	ref common.ReferenceCallback,
) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TopologySpec is the spec for a Topology resource.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"definition": {
						SchemaProps: spec.SchemaProps{
							Description: "Definition defines the actual set of nodes (network ones, not k8s ones!) that this Topology CR represents. Historically, and probably most often, this means Topology holds a \"normal\" containerlab topology file that will be \"clabernetsified\", however this could also be a \"kne\" config, or perhaps others in the future.",
							Default:     map[string]interface{}{},
							Ref: ref(
								"github.com/srl-labs/clabernetes/apis/v1alpha1.Definition",
							),
						},
					},
					"expose": {
						SchemaProps: spec.SchemaProps{
							Description: "Expose holds configurations relevant to how clabernetes exposes a topology.",
							Default:     map[string]interface{}{},
							Ref: ref(
								"github.com/srl-labs/clabernetes/apis/v1alpha1.Expose",
							),
						},
					},
					"deployment": {
						SchemaProps: spec.SchemaProps{
							Description: "Deployment holds configurations relevant to how clabernetes configures deployments that make up a given topology.",
							Default:     map[string]interface{}{},
							Ref: ref(
								"github.com/srl-labs/clabernetes/apis/v1alpha1.Deployment",
							),
						},
					},
					"imagePull": {
						SchemaProps: spec.SchemaProps{
							Description: "ImagePull holds configurations relevant to how clabernetes launcher pods handle pulling images.",
							Default:     map[string]interface{}{},
							Ref: ref(
								"github.com/srl-labs/clabernetes/apis/v1alpha1.ImagePull",
							),
						},
					},
				},
				Required: []string{"definition"},
			},
		},
		Dependencies: []string{
			"github.com/srl-labs/clabernetes/apis/v1alpha1.Definition", "github.com/srl-labs/clabernetes/apis/v1alpha1.Deployment", "github.com/srl-labs/clabernetes/apis/v1alpha1.Expose", "github.com/srl-labs/clabernetes/apis/v1alpha1.ImagePull"},
	}
}

func schema_srl_labs_clabernetes_apis_v1alpha1_TopologyStatus(
	ref common.ReferenceCallback,
) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TopologyStatus is the status for a Containerlab topology resource.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is the topology kind this CR represents -- for example \"containerlab\".",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"configs": {
						SchemaProps: spec.SchemaProps{
							Description: "Configs is a map of node name -> clab config -- in other words, this is the original containerlab configuration broken up and modified to use multi-node topology setup (via host links+vxlan). This is stored as a raw message so we don't have any weirdness w/ yaml tags instead of json tags in clab things, and so we kube builder doesnt poop itself.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"configsHash": {
						SchemaProps: spec.SchemaProps{
							Description: "ConfigsHash is a hash of the last stored Configs data.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"tunnels": {
						SchemaProps: spec.SchemaProps{
							Description: "Tunnels is a mapping of tunnels that need to be configured between nodes (nodes:[]tunnels).",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type: []string{"array"},
										Items: &spec.SchemaOrArray{
											Schema: &spec.Schema{
												SchemaProps: spec.SchemaProps{
													Ref: ref(
														"github.com/srl-labs/clabernetes/apis/v1alpha1.Tunnel",
													),
												},
											},
										},
									},
								},
							},
						},
					},
					"tunnelsHash": {
						SchemaProps: spec.SchemaProps{
							Description: "TunnelsHash is a hash of the last stored Tunnels data. As this can change due to the dns suffix changing and not just the config changing we need to independently track this state.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"filesFromURLHashes": {
						SchemaProps: spec.SchemaProps{
							Description: "FilesFromURLHashes is a mapping of node FilesFromURL hashes stored so we can easily identify which nodes had changes in their FilesFromURL data so we can know to restart them.",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"nodeExposedPorts": {
						SchemaProps: spec.SchemaProps{
							Description: "NodeExposedPorts holds a map of (containerlab) nodes and their exposed ports (via load balancer).",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref(
											"github.com/srl-labs/clabernetes/apis/v1alpha1.ExposedPorts",
										),
									},
								},
							},
						},
					},
					"nodeExposedPortsHash": {
						SchemaProps: spec.SchemaProps{
							Description: "NodeExposedPortsHash is a hash of the last stored NodeExposedPorts data.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"imagePullSecretsHash": {
						SchemaProps: spec.SchemaProps{
							Description: "ImagePullSecretsHash is a hash of the last stored ImagePullSecrets data.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{
					"kind",
					"configs",
					"configsHash",
					"tunnels",
					"tunnelsHash",
					"filesFromURLHashes",
					"nodeExposedPorts",
					"nodeExposedPortsHash",
					"imagePullSecretsHash",
				},
			},
		},
		Dependencies: []string{
			"github.com/srl-labs/clabernetes/apis/v1alpha1.ExposedPorts", "github.com/srl-labs/clabernetes/apis/v1alpha1.Tunnel"},
	}
}

func schema_srl_labs_clabernetes_apis_v1alpha1_Tunnel(
	ref common.ReferenceCallback,
) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Tunnel represents a VXLAN tunnel between clabernetes nodes (as configured by containerlab).",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"id": {
						SchemaProps: spec.SchemaProps{
							Description: "ID is the VXLAN ID (vnid) for the tunnel.",
							Default:     0,
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"localNodeName": {
						SchemaProps: spec.SchemaProps{
							Description: "LocalNodeName is the name of the local node for this tunnel.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"remoteName": {
						SchemaProps: spec.SchemaProps{
							Description: "RemoteName is the name of the service to contact the remote end of the tunnel.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"remoteNodeName": {
						SchemaProps: spec.SchemaProps{
							Description: "RemoteNodeName is the name of the remote node.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"localLinkName": {
						SchemaProps: spec.SchemaProps{
							Description: "LocalLinkName is the local link name for the local end of the tunnel.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"remoteLinkName": {
						SchemaProps: spec.SchemaProps{
							Description: "RemoteLinkName is the remote link name for the remote end of the tunnel.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{
					"id",
					"localNodeName",
					"remoteName",
					"remoteNodeName",
					"localLinkName",
					"remoteLinkName",
				},
			},
		},
	}
}
