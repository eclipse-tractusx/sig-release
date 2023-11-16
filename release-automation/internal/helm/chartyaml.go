/*******************************************************************************
 * Copyright (c) 2023 Contributors to the Eclipse Foundation
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

package helm

import (
	"fmt"
	"os"
	"reflect"
	"regexp"

	"gopkg.in/yaml.v3"
)

type Chartyaml struct {
	ApiVersion  string `yaml:"apiVersion"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	AppVersion  string `yaml:"appVersion"`
	Version     string `yaml:"version"`
}

func newChartYaml() *Chartyaml {
	return &Chartyaml{
		ApiVersion:  "",
		Name:        "",
		Description: "",
		AppVersion:  "",
		Version:     "",
	}
}

func ChartYamlFromFile(ymlfile string) *Chartyaml {
	data, err := os.ReadFile(ymlfile)
	if err != nil {
		fmt.Printf("Unable to read %v.\n", ymlfile)
		return nil
	}

	var c Chartyaml
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		fmt.Printf("Unable to parse YAML file: %v.\n", ymlfile)
		return nil
	}

	return &c
}

func (c *Chartyaml) IsVersionValid() bool {
	/*
		Below regular expresion is used to verify version string according to Semantic Versioning schema (https://semver.org).
		Following examples match:
		- "1.2.3"
		- "1.0.0-alpha"
		- "1.2.3-alpha.1+ef365"
		Following don't match:
		- "1.2.abc"
		- "0.1-alpha"
	*/
	regexPattern := `^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`

	match, err := regexp.MatchString(regexPattern, c.Version)
	if err != nil {
		fmt.Println("Error occured when validating semantic version.")
		return false
	}
	return match
}

func (c *Chartyaml) GetMissingMandatoryFields() []string {
	chartValues := reflect.ValueOf(*c)
	chartType := chartValues.Type()
	numChartFields := chartValues.NumField()

	var missingFields []string
	for i := 0; i < numChartFields; i++ {
		chartField := chartType.Field(i)
		chartFieldValue := chartValues.Field(i)

		if chartFieldValue.Len() == 0 {
			missingFields = append(missingFields, chartField.Name)
		}
	}
	return missingFields
}
