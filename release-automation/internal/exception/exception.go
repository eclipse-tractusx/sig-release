/*******************************************************************************
 * Copyright (c) 2024 Contributors to the Eclipse Foundation
 *
 * See the NOTICE file(s) distributed with this work for additional
 * information regarding copyright ownership.
 *
 * This program and the accompanying materials are made available under the
 * terms of the Apache License, Version 2.0 which is available at
 * https://www.apache.org/licenses/LICENSE-2.0.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 ******************************************************************************/

package exception

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

const ExceptionsData = "internal/exception/exceptions.yaml"

 type Exception struct {
	Trg string `yaml:"trg"`
	Repositories []string `yaml:"repository"`
 }

 type Config struct {
	Exceptions []Exception `yaml:"exceptions"`
 }

 func GetExceptionsFromFile(filepath string) (*Config, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("unable to read %v", filepath)
	}

	var c Config
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return nil, fmt.Errorf("unable to parse YAML file: %v", filepath)
	}

	return &c, nil
 }

// Checks if repository has exception for specific TRG
func (c *Config) IsExceptioned(trg string, repository string) bool {
	for _, e := range c.Exceptions {
		if e.Trg == trg {
			for _, r := range e.Repositories {
				if strings.EqualFold(r, repository) {
					return true
				}
			}
		}
	}
	return false
}
