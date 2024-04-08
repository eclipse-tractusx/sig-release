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
	"gopkg.in/yaml.v3"
	"io"
	"net/http"
	"strings"
)

const ExceptionsData = "https://raw.githubusercontent.com/eclipse-tractusx/sig-release/main/release-automation/exceptions.yaml"
type Exception struct {
	Trg          string   `yaml:"trg"`
	Repositories []string `yaml:"repository"`
}

type config struct {
	Exceptions []Exception `yaml:"exceptions"`
}

type ExceptionsMap map[string][]string

func GetData() (ExceptionsMap, error) {
	data, err := fetchYaml(ExceptionsData)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch YAML data from url: %v", ExceptionsData)
	}
	config, err := parseData(data)
	if err != nil {
		return nil, fmt.Errorf("unable to parse exceptions")
	}
	var exMap = make(map[string][]string)
	for _, e := range config.Exceptions {
		exMap[e.Trg] = e.Repositories
	}
	return exMap, nil
}

func fetchYaml(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil || response.StatusCode > 299 {
		return nil, fmt.Errorf("unable to get %v", url)
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func parseData(data []byte) (*config, error) {
	var c config
	err := yaml.Unmarshal(data, &c)
	if err != nil {
		return nil, fmt.Errorf("unable to parse YAML data")
	}
	return &c, nil
}

// Checks if repository has exception for specific TRG
func (m ExceptionsMap) IsExceptioned(trg string, repository string) bool {
	repos := m[trg]
	for _, r := range repos {
		if strings.EqualFold(r, repository) {
			return true
		}
	}
	return false
}
