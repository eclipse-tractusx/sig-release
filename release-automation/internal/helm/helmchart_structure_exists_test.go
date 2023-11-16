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
	"os"
	"path"
	"testing"

	"tractusx-release-automation/internal/filesystem"
)

var validChartYmlTestFile = "test/StructureTestTemplate.yaml"

func TestShouldPassIfHelmDirIsMissing(t *testing.T) {
	result := NewHelmStructureExists("./").Test()

	if !result.Passed {
		t.Errorf("Helm directory doesn't exist hence test should pass.")
	}
}

func TestShouldPassForEmptyChartsDir(t *testing.T) {
	_ = os.Mkdir("charts", 0750)
	defer os.Remove("charts")

	result := NewHelmStructureExists("./").Test()

	if !result.Passed {
		t.Errorf("Helm directory doesn't contain any charts hence test should pass.")
	}
}

func TestShouldFailIfHelmStructureIsMissing(t *testing.T) {
	_ = os.MkdirAll("charts/exampleChart", 0750)
	filesystem.CreateFiles([]string{"charts/exampleChart/Chart.yaml"})
	defer os.RemoveAll("charts")

	result := NewHelmStructureExists("./").Test()

	if result.Passed {
		t.Errorf("Helm structure is missing hence test should fail.")
	}
}

func TestShouldPassIfHelmStructureExist(t *testing.T) {
	_ = os.MkdirAll("charts/exampleChart/templates", 0750)
	filesystem.CreateFiles([]string{
		"charts/exampleChart/.helmignore",
		"charts/exampleChart/Chart.yaml",
		"charts/exampleChart/LICENSE",
		"charts/exampleChart/README.md",
		"charts/exampleChart/values.yaml",
	})
	defer os.RemoveAll("charts")
	copyTemplateFileTo("charts/exampleChart/Chart.yaml", t)

	result := NewHelmStructureExists("./").Test()

	if !result.Passed {
		t.Errorf("Helm structure exists hence test should pass.")
	}
}

func TestShouldPassIfChartStructureExistsAtGivenBaseDir(t *testing.T) {
	dir := t.TempDir()
	_ = os.MkdirAll(path.Join(dir, "charts", "exampleChart"), 0770)
	filesystem.CreateFiles([]string{
		path.Join(dir, "charts/exampleChart/.helmignore"),
		path.Join(dir, "charts/exampleChart/LICENSE"),
		path.Join(dir, "charts/exampleChart/README.md"),
		path.Join(dir, "charts/exampleChart/values.yaml"),
	})
	copyTemplateFileTo(path.Join(dir, "charts", "exampleChart", "Chart.yaml"), t)

	result := NewHelmStructureExists(dir).Test()

	if !result.Passed {
		t.Errorf("Should pass, if helm chart file and dir structure exists at given base dir")
	}
}

func copyTemplateFileTo(path string, t *testing.T) {
	templateFile, err := os.ReadFile(validChartYmlTestFile)
	if err != nil {
		t.Errorf("Could not read template file necessary for this test")
	}
	err = os.WriteFile(path, templateFile, 0770)
	if err != nil {
		t.Errorf("Could not copy template file to designated path")
	}
}
