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

package tractusx

import (
	"os"
	"path"
	"reflect"
	"testing"
)

var metadataFromTestTemplate = Metadata{
	ProductName:       "sig-infra",
	LeadingRepository: "https://github.com/eclipse-tractusx/sig-infra",
	RepoCategory:      "special",
	Repositories: []Repository{
		{
			Name:             "sig-infra",
			UsageDescription: "Contains issue templates and workflows",
			Url:              "https://github.com/eclipse-tractusx/sig-infra",
		},
		{
			Name:             "charts",
			UsageDescription: "Central Helm repository containing all released charts of Tractus-X",
			Url:              "https://github.com/eclipse-tractusx/charts",
		},
	},
	SkipReleaseChecks: SkipReleaseChecks{
		AlignedBaseImages: []string{
			"/path/to/non-published/Dockerfile",
		},
	},
}

const metadataTestFile = "./test/metadata_test_template.yaml"

func TestShouldReturnErrorWhenReadingNonExistentDefaultMetadataFile(t *testing.T) {
	_, err := MetadataFromLocalFile("./")

	if err == nil {
		t.Errorf("Reading the product metadata file from default location should fail if not existing")
	}
}

// This test works with a predefined metadata file located at ./test/metadata_test_template.yaml (see metadataTestFile const).
// This is a well known file, so we can confidently assert property names
func TestShouldReadProductMetadataFromDefaultFile(t *testing.T) {
	copyTemplateFileTo(".tractusx", t)
	defer os.Remove(".tractusx")

	metadata, err := MetadataFromLocalFile("./")

	if err != nil {
		t.Errorf("Should be able to read metadata file after copying to default location")
	}

	if !reflect.DeepEqual(metadata, &metadataFromTestTemplate) {
		t.Errorf("Metadata read from file does not match test template values")
	}
}

func TestShouldReadMetadataFromGivenBaseDir(t *testing.T) {
	tempDir := t.TempDir()
	copyTemplateFileTo(path.Join(tempDir, ".tractusx"), t)

	metadata, _ := MetadataFromLocalFile(tempDir)

	if !reflect.DeepEqual(metadata, &metadataFromTestTemplate) {
		t.Errorf("Metadata read from given base directory does not match test template values")
	}
}

func copyTemplateFileTo(path string, t *testing.T) {
	templateFile, err := os.ReadFile(metadataTestFile)
	if err != nil {
		t.Errorf("Could not read template file necessary for this test")
	}
	err = os.WriteFile(path, templateFile, 0644)
	if err != nil {
		t.Errorf("Could not copy template file to designated path")
	}
}
