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
	ID                string                        `yaml:"id"`                // url without schema (e.g. github.com/user/repo)
	Name              string                        `yaml:"name"`              // module name
	Description       string                        `yaml:"description"`       // short text describing the module
	Tags              []string                      `yaml:"tags"`              // module tags
	License           string                        `yaml:"license"`           // module license name (e.g. Apache License 2.0)
	Author            string                        `yaml:"author"`            // module author
	Version           string                        `yaml:"version"`           // module version (must be prefixed with 'v' and adhere to the semantic versioning guidelines, see https://semver.org/ for details)
	Type              string                        `yaml:"type"`              // module type (e.g. device-connector specifies a module for integrating devices)
	DeploymentType    string                        `yaml:"deploymentType"`    // specifies whether a module can only be deployed once or multiple times
	Architectures     []string                      `yaml:"architectures"`     // supported cpu architectures
	Services          map[string]Service            `yaml:"services"`          // map depicting the services the module consists of (keys serve as unique identifiers and can be reused elsewhere in the modfile to reference a service)
	ServiceReferences map[string][]DependencyTarget `yaml:"serviceReferences"` // map linking module services to reference variables (identifiers as defined in ModFile.Services serve as keys)
	Volumes           map[string][]VolumeTarget     `yaml:"volumes"`           // map linking volumes to mount points (keys represent volume names)
	Dependencies      map[string]ModuleDependency   `yaml:"dependencies"`      // external modules required by the module (keys represent module IDs)
	HostResources     map[string]HostResource       `yaml:"hostResources"`     // host resources required by services (e.g. devices, sockets, ...)
	Secrets           map[string]Secret             `yaml:"secrets"`           // secrets required by services (e.g. certs, keys, ...)
	Configs           map[string]ConfigValue        `yaml:"configs"`           // configuration values required by services
	InputGroups       map[string]InputGroup         `yaml:"inputGroups"`       // map of groups for categorising user inputs (keys serve as unique identifiers and can be reused elsewhere in the modfile to reference a group)
}

type Service struct {
	Name             string         `yaml:"name"`             // service name
	Image            string         `yaml:"image"`            // container image (must be versioned via tag or digest, e.g. srv-image:v1.0.0)
	RunConfig        RunConfig      `yaml:"runConfig"`        // configurations for running the service container (e.g. restart strategy, stop timeout, ...)
	Include          []BindMount    `yaml:"include"`          // files or dictionaries to be mounted from module repository
	Tmpfs            []TmpfsMount   `yaml:"tmpfs"`            // temporary file systems (in memory) required by the service
	HttpEndpoints    []HttpEndpoint `yaml:"httpEndpoints"`    // http endpoints of the service to be exposed via the api gateway
	Ports            []SrvPort      `yaml:"ports"`            // service ports to be published on the host
	RequiredServices []string       `yaml:"requiredServices"` // identifiers of internal services that must be running before this service is started
}

type Duration time.Duration

type RunConfig struct {
	MaxRetries  *int      `yaml:"maxRetries"` // defaults to 3 if nil
	RunOnce     bool      `yaml:"runOnce"`
	StopTimeout *Duration `yaml:"stopTimeout"` // defaults to 5s if nil
	StopSignal  *string   `yaml:"stopSignal"`
	PseudoTTY   bool      `yaml:"pseudoTTY"`
}

type BindMount struct {
	MountPoint string `yaml:"mountPoint"` // absolute path in container
	Source     string `yaml:"source"`     // relative path in module repo
	ReadOnly   bool   `yaml:"readOnly"`
}

type FileMode fs.FileMode

type TmpfsMount struct {
	MountPoint string    `yaml:"mountPoint"` // absolute path in container
	Size       ByteFmt   `yaml:"size"`       // tmpfs size in bytes provided as integer or in human-readable form (e.g. 64Mb)
	Mode       *FileMode `yaml:"mode"`       // linux file mode to be used for the tmpfs provided as string (e.g. 777, 0777; defaults to 770 if nil)
}

type HttpEndpoint struct {
	Name    string  `yaml:"name"`    // endpoint name
	Path    string  `yaml:"path"`    // internal endpoint path
	Port    *int    `yaml:"port"`    // port the service is listening on
	ExtPath *string `yaml:"extPath"` // optional external path to be used by the api gateway
}

type SrvPort struct {
	Name     *string `yaml:"name"`     // port name
	Port     Port    `yaml:"port"`     // port number or port range (e.g. 8080-8081)
	HostPort *Port   `yaml:"hostPort"` // port number or port range (e.g. 8080-8081), can be overridden during deployment to avoid collisions (arbitrary ports are used if nil)
	Protocol *string `yaml:"protocol"` // specify port protocol (defaults to tcp if nil)
}

type VolumeTarget struct {
	MountPoint string   `yaml:"mountPoint"` // absolute path in container
	Services   []string `yaml:"services"`   // service identifiers as used in ModFile.Services to map the mount point to a number of services
}

type ModuleDependency struct {
	Version          string                        `yaml:"version"`          // version of required module (e.g. =v1.0.2, >v1.0.2., >=v1.0.2, >v1.0.2;<v2.1.3, ...)
	RequiredServices map[string][]DependencyTarget `yaml:"requiredServices"` // map linking required services to reference variables (identifiers as defined in ModFile.Services of the required module are used as keys)
}

type DependencyTarget struct {
	RefVar   string   `yaml:"refVar"`   // container environment variable to hold the addressable reference of the service
	Services []string `yaml:"services"` // service identifiers as used in ModFile.Services to map the reference variable to a number of services
}

type ResourceMountTarget struct {
	MountPoint string   `yaml:"mountPoint"` // absolute path in container
	Services   []string `yaml:"services"`   // service identifiers as used in ModFile.Services to map the mount point to a number of services
}

type HostResourceTarget struct {
	ResourceMountTarget `yaml:",inline"`
	ReadOnly            bool `yaml:"readOnly"` // if true resource will be mounted as read only
}

type Resource struct {
	Tags      []string   `yaml:"tags"`      // tags for aiding resource identification (e.g. a vendor), unique type and tag combinations can be used to select resources without requiring user interaction
	UserInput *UserInput `yaml:"userInput"` // meta info for user input via gui (if nil and not optional the tag combination must yield a unique resource)
	Optional  bool       `yaml:"optional"`
}

type HostResource struct {
	Resource `yaml:",inline"`
	Targets  []HostResourceTarget `yaml:"targets"` // mount points for the resource
}

type Secret struct {
	Resource `yaml:",inline"`
	Type     string                `yaml:"type"`    // resource type as defined by external services managing resources (e.g. serial-device, certificate, ...)
	Targets  []ResourceMountTarget `yaml:"targets"` // mount points for the secret
}

type ConfigValue struct {
	Value       any            `yaml:"value"`       // default configuration value or nil
	Options     []any          `yaml:"options"`     // list of possible configuration values
	OptionsExt  bool           `yaml:"optionsExt"`  // if true a value not defined in options can be set (only required if options are provided)
	Type        string         `yaml:"type"`        // type of the configuration value (e.g. text, number, date, ...)
	TypeOptions map[string]any `yaml:"typeOptions"` // type specific options (e.g. number supports min, max values or step)
	DataType    string         `yaml:"dataType"`    // data type of the configuration value (e.g. string, int, ...)
	IsList      bool           `yaml:"isList"`      // set to true if multiple configuration values are required
	Delimiter   *string        `yaml:"delimiter"`   // delimiter to be used for marshalling multiple configuration values (defaults to "," if nil)
	UserInput   *UserInput     `yaml:"userInput"`   // meta info for user input via gui (if nil a default value must be set)
	Targets     []ConfigTarget `yaml:"targets"`     // reference variables for the configuration value
	Optional    bool           `yaml:"optional"`
}

type ConfigTarget struct {
	RefVar   string   `yaml:"refVar"`   // container environment variable to hold the configuration value
	Services []string `yaml:"services"` // service identifiers as used in ModFile.Services to map the reference variable to a number of services
}

type UserInput struct {
	Name        string  `yaml:"name"`        // input name (e.g. used as a label for input field)
	Description *string `yaml:"description"` // short text describing the input
	Group       *string `yaml:"group"`       // group identifier as used in ModFile.InputGroups to assign the user input to an input group
}

type InputGroup struct {
	Name        string  `yaml:"name"`        // input group name
	Description *string `yaml:"description"` // short text describing the input group
	Group       *string `yaml:"group"`       // group identifier as used in ModFile.InputGroups to assign the input group to a parent group
}
