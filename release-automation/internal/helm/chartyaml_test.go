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

import "testing"

func TestShouldPassIfCanReadChartYaml(t *testing.T) {
	c := ChartYamlFromFile("test/TestChartValid.yaml")
	if c.Name != "test-application" {
		t.Errorf("Test should pass but name of the helm chart is not 'test-application'.")
	}
}

func TestShouldBeValidForSemanticVersion(t *testing.T) {
	c := newChartYaml()
	c.Version = "1.2.3"
	if !c.IsVersionValid() {
		t.Errorf("Test should pass since version is valid according Semantic Version schema.")
	}
}

func TestShouldBeInValidForSemanticVersion(t *testing.T) {
	c := newChartYaml()
	c.Version = "1.2.ABC"
	if c.IsVersionValid() {
		t.Errorf("Test should fail since version is invalid according to Semantic Version schema.")
	}
}

func TestShouldPassIfConfigurationSettingsArePresent(t *testing.T) {
	c := ChartYamlFromFile("test/TestChartValid.yaml")
	missingFields := c.GetMissingMandatoryFields()

	if len(missingFields) > 0 {
		t.Errorf("Test should pass since all configuration settings are set.")
	}
}

func TestShouldFailIfConfigurationSettingsAreMissing(t *testing.T) {
	c := ChartYamlFromFile("test/TestChartInValid.yaml")
	missingFields := c.GetMissingMandatoryFields()

	if len(missingFields) == 0 {
		t.Errorf("Test should fail since one configuration setting is missing.")
	}
}
