// Copyright 2023 Greptime Team
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

type Frontend struct {
	Replicas int `yaml:"replicas" validate:"gt=0"`

	GRPCAddr     string `yaml:"grpcAddr" validate:"omitempty,hostname_port"`
	HTTPAddr     string `yaml:"httpAddr" validate:"omitempty,hostname_port"`
	PostgresAddr string `yaml:"postgresAddr" validate:"omitempty,hostname_port"`
	MetaAddr     string `yaml:"metaAddr" validate:"omitempty,hostname_port"`

	LogLevel string `yaml:"logLevel"`
}
