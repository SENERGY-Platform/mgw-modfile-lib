/*
 * Copyright 2023 InfAI (CC SES)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package model

import (
	"github.com/SENERGY-Platform/mgw-modfile-lib/modfile"
	"io/fs"
	"time"
)

type Port string

type ByteFmt uint64

type ModFile struct {
	modfile.Base
	// url without schema (e.g. github.com/user/repo)
	ID string `yaml:"id" json:"id"`
	// module name
	Name string `yaml:"name" json:"name,omitempty"`
	// short text describing the module
	Description string `yaml:"description" json:"description,omitempty"`
	// module tags
	Tags []string `yaml:"tags" json:"tags,omitempty"`
	// module license name (e.g. Apache License 2.0)
	License string `yaml:"license" json:"license"`
	// module author
	Author string `yaml:"author" json:"author"`
	// module version (must be prefixed with 'v' and adhere to the semantic versioning guidelines, see https://semver.org/ for details)
	Version string `yaml:"version" json:"version"`
	// module type (e.g. device-connector specifies a module for integrating devices)
	Type string `yaml:"type" json:"type"`
	// specifies whether a module can only be deployed once or multiple times
	DeploymentType string `yaml:"deploymentType" json:"deploymentType"`
	// supported cpu architectures
	Architectures []string `yaml:"architectures" json:"architectures,omitempty"`
	// map depicting the services the module consists of (keys serve as unique identifiers and can be reused elsewhere in the modfile to reference a service)
	Services map[string]Service `yaml:"services" json:"services,omitempty"`
	// map linking module services to reference variables (identifiers as defined in ModFile.Services serve as keys)
	ServiceReferences map[string][]DependencyTarget `yaml:"serviceReferences" json:"serviceReferences,omitempty"`
	// map linking volumes to mount points (keys represent volume names)
	Volumes map[string][]VolumeTarget `yaml:"volumes" json:"volumes,omitempty"`
	// external modules required by the module (keys represent module IDs)
	Dependencies map[string]ModuleDependency `yaml:"dependencies" json:"dependencies,omitempty"`
	// host resources required by services (e.g. devices, sockets, ...)
	HostResources map[string]HostResource `yaml:"hostResources" json:"hostResources,omitempty"`
	// secrets required by services (e.g. certs, keys, ...)
	Secrets map[string]Secret `yaml:"secrets" json:"secrets,omitempty"`
	// configuration values required by services
	Configs map[string]ConfigValue `yaml:"configs" json:"configs,omitempty"`
	// map of groups for categorising user inputs (keys serve as unique identifiers and can be reused elsewhere in the modfile to reference a group)
	InputGroups map[string]InputGroup `yaml:"inputGroups" json:"inputGroups,omitempty"`
}

type Service struct {
	// service name
	Name string `yaml:"name" json:"name"`
	// container image (must be versioned via tag or digest, e.g. srv-image:v1.0.0)
	Image string `yaml:"image" json:"image"`
	// configurations for running the service container (e.g. restart strategy, stop timeout, ...)
	RunConfig RunConfig `yaml:"runConfig" json:"runConfig,omitempty"`
	// files or dictionaries to be mounted from module repository
	Include []BindMount `yaml:"include" json:"include,omitempty"`
	// temporary file systems (in memory) required by the service
	Tmpfs []TmpfsMount `yaml:"tmpfs" json:"tmpfs,omitempty"`
	// http endpoints of the service to be exposed via the api gateway
	HttpEndpoints []HttpEndpoint `yaml:"httpEndpoints" json:"httpEndpoints,omitempty"`
	// service ports to be published on the host
	Ports []SrvPort `yaml:"ports" json:"ports,omitempty"`
	// identifiers of internal services that must be running before this service is started
	RequiredServices []string `yaml:"requiredServices" json:"requiredServices,omitempty"`
}

type Duration time.Duration

type RunConfig struct {
	// defaults to 3 if nil
	MaxRetries *int `yaml:"maxRetries" json:"maxRetries,omitempty"`
	RunOnce    bool `yaml:"runOnce" json:"runOnce"`
	// defaults to 5s if nil
	StopTimeout *Duration `yaml:"stopTimeout" json:"stopTimeout,omitempty"`
	StopSignal  *string   `yaml:"stopSignal" json:"stopSignal,omitempty"`
	PseudoTTY   bool      `yaml:"pseudoTTY" json:"pseudoTTY,omitempty"`
}

type BindMount struct {
	// absolute path in container
	MountPoint string `yaml:"mountPoint" json:"mountPoint"`
	// relative path in module repo
	Source   string `yaml:"source" json:"source"`
	ReadOnly bool   `yaml:"readOnly" json:"readOnly"`
}

type FileMode fs.FileMode

type TmpfsMount struct {
	// absolute path in container
	MountPoint string `yaml:"mountPoint" json:"mountPoint"`
	// tmpfs size in bytes provided as integer or in human-readable form (e.g. 64Mb)
	Size ByteFmt `yaml:"size" json:"size"`
	// linux file mode to be used for the tmpfs provided as string (e.g. 777, 0777; defaults to 770 if nil)
	Mode *FileMode `yaml:"mode" json:"mode"`
}

type HttpEndpoint struct {
	// endpoint name
	Name string `yaml:"name" json:"name,omitempty"`
	// internal endpoint path
	Path string `yaml:"path" json:"path"`
	// port the service is listening on
	Port *int `yaml:"port" json:"port"`
	// optional external path to be used by the api gateway
	ExtPath *string `yaml:"extPath" json:"extPath,omitempty"`
}

type SrvPort struct {
	// port name
	Name *string `yaml:"name" json:"name,omitempty"`
	// port number or port range (e.g. 8080-8081)
	Port Port `yaml:"port" json:"port"`
	// port number or port range (e.g. 8080-8081), can be overridden during deployment to avoid collisions (arbitrary ports are used if nil)
	HostPort *Port `yaml:"hostPort" json:"hostPort,omitempty"`
	// specify port protocol (defaults to tcp if nil)
	Protocol *string `yaml:"protocol" json:"protocol,omitempty"`
}

type VolumeTarget struct {
	// absolute path in container
	MountPoint string `yaml:"mountPoint" json:"mountPoint"`
	// service identifiers as used in ModFile.Services to map the mount point to a number of services
	Services []string `yaml:"services" json:"services"`
}

type ModuleDependency struct {
	// version of required module (e.g. =v1.0.2, >v1.0.2., >=v1.0.2, >v1.0.2;<v2.1.3, ...)
	Version string `yaml:"version" json:"version"`
	// map linking required services to reference variables (identifiers as defined in ModFile.Services of the required module are used as keys)
	RequiredServices map[string][]DependencyTarget `yaml:"requiredServices" json:"requiredServices"`
}

type DependencyTarget struct {
	// container environment variable to hold the addressable reference of the service
	RefVar string `yaml:"refVar" json:"refVar"`
	// service identifiers as used in ModFile.Services to map the reference variable to a number of services
	Services []string `yaml:"services" json:"services"`
}

type ResourceMountTarget struct {
	// absolute path in container
	MountPoint string `yaml:"mountPoint" json:"mountPoint"`
	// service identifiers as used in ModFile.Services to map the mount point to a number of services
	Services []string `yaml:"services" json:"services"`
}

type HostResourceTarget struct {
	ResourceMountTarget `yaml:",inline"`
	// if true resource will be mounted as read only
	ReadOnly bool `yaml:"readOnly" json:"readOnly"`
}

type Resource struct {
	// tags for aiding resource identification (e.g. a vendor), unique type and tag combinations can be used to select resources without requiring user interaction
	Tags []string `yaml:"tags" json:"tags,omitempty"`
	// meta info for user input via gui (if nil and not optional the tag combination must yield a unique resource)
	UserInput *UserInput `yaml:"userInput" json:"userInput,omitempty"`
	Optional  bool       `yaml:"optional" json:"optional"`
}

type HostResource struct {
	Resource `yaml:",inline"`
	// mount points for the resource
	Targets []HostResourceTarget `yaml:"targets" json:"targets"`
}

type Secret struct {
	Resource `yaml:",inline"`
	// resource type as defined by external services managing resources (e.g. serial-device, certificate, ...)
	Type string `yaml:"type" json:"type"`
	// mount points for the secret
	Targets []ResourceMountTarget `yaml:"targets" json:"targets"`
}

type ConfigValue struct {
	// default configuration value or nil
	Value any `yaml:"value" json:"value,omitempty"`
	// list of possible configuration values
	Options []any `yaml:"options" json:"options,omitempty"`
	// if true a value not defined in options can be set (only required if options are provided)
	OptionsExt bool `yaml:"optionsExt" json:"optionsExt,omitempty"`
	// type of the configuration value (e.g. text, number, date, ...)
	Type string `yaml:"type" json:"type"`
	// type specific options (e.g. number supports min, max values or step)
	TypeOptions map[string]any `yaml:"typeOptions" json:"typeOptions,omitempty"`
	// data type of the configuration value (e.g. string, int, ...)
	DataType string `yaml:"dataType" json:"dataType"`
	// set to true if multiple configuration values are required
	IsList bool `yaml:"isList" json:"isList,omitempty"`
	// delimiter to be used for marshalling multiple configuration values (defaults to "," if nil)
	Delimiter *string `yaml:"delimiter" json:"delimiter,omitempty"`
	// meta info for user input via gui (if nil a default value must be set)
	UserInput *UserInput `yaml:"userInput" json:"userInput,omitempty"`
	// reference variables for the configuration value
	Targets  []ConfigTarget `yaml:"targets" json:"targets,omitempty"`
	Optional bool           `yaml:"optional" json:"optional,omitempty"`
}

type ConfigTarget struct {
	// container environment variable to hold the configuration value
	RefVar string `yaml:"refVar" json:"refVar"`
	// service identifiers as used in ModFile.Services to map the reference variable to a number of services
	Services []string `yaml:"services" json:"services"`
}

type UserInput struct {
	// input name (e.g. used as a label for input field)
	Name string `yaml:"name" json:"name"`
	// short text describing the input
	Description *string `yaml:"description" json:"description,omitempty"`
	// group identifier as used in ModFile.InputGroups to assign the user input to an input group
	Group *string `yaml:"group" json:"group,omitempty"`
}

type InputGroup struct {
	// input group name
	Name string `yaml:"name" json:"name"`
	// short text describing the input group
	Description *string `yaml:"description" json:"description,omitempty"`
	// group identifier as used in ModFile.InputGroups to assign the input group to a parent group
	Group *string `yaml:"group" json:"group,omitempty"`
}
