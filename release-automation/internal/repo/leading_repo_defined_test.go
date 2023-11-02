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

package repo

import (
	"os"
	"path"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
	"tractusx-release-automation/internal/tractusx"
)

// Explicitly remove .tractusx file before and after running tests to ensure clean preconditions for each test
func TestMain(m *testing.M) {
	_ = os.Remove(".tractusx")
	code := m.Run()
	_ = os.Remove(".tractusx")
	os.Exit(code)
}

func TestShouldFailIfMetadataFileIsMissing(t *testing.T) {
	result := NewLeadingRepositoryDefined("./").Test()

	if result.Passed {
		t.Errorf("LeadingRepoDefined should faild, if there is no repository metadata")
	}
}

func TestShouldFailIfLeadingRepositoryMetadataPropertyIsUndefined(t *testing.T) {
	metadata := tractusx.Metadata{
		ProductName:       "ProductWithoutLeadingRepo",
		LeadingRepository: "",
		Repositories:      nil,
	}
	givenProductMetadata(t, metadata)

	result := NewLeadingRepositoryDefined("./").Test()
	if result.Passed {
		t.Errorf("Check should fail if metadata file exists, but does not have leading repo defined")
	}
	if !strings.Contains(result.ErrorDescription, "leadingRepository property must be defined in .tractusx metadata file") {
		t.Errorf("Leading repository check should provide error description if failing")
	}
}

func TestShouldSucceedIfLeadingRepoIsDefinedInMetadata(t *testing.T) {
	metadata := tractusx.Metadata{
		ProductName:       "ProductWithoutLeadingRepo",
		LeadingRepository: "https://github.com/eclipse-tractusx/sig-infra",
		Repositories:      nil,
	}
	givenProductMetadata(t, metadata)

	result := NewLeadingRepositoryDefined("./").Test()

	if !result.Passed {
		t.Errorf("Leading Repo check should pass, if configured in metadata")
	}
}

func TestShouldUseMetadataAtSpecifiedDirectory(t *testing.T) {
	testDir := t.TempDir()
	metadata := tractusx.Metadata{
		ProductName:       "ProductWithoutLeadingRepo",
		LeadingRepository: "https://github.com/eclipse-tractusx/sig-infra",
		Repositories:      nil,
	}
	givenProductMetadataAtDir(t, metadata, testDir)

	result := NewLeadingRepositoryDefined(testDir).Test()

	if !result.Passed {
		t.Errorf("Leading Repo check should use metadata from a given path")
	}
}

func givenProductMetadata(t *testing.T, metadata tractusx.Metadata) {
	givenProductMetadataAtDir(t, metadata, "./")
}

func givenProductMetadataAtDir(t *testing.T, metadata tractusx.Metadata, dir string) {
	yamlContent, err := yaml.Marshal(&metadata)
	if err != nil {
		t.Errorf("Could not serialize metadata test content")
	}

	err = os.WriteFile(path.Join(dir, ".tractusx"), yamlContent, 0600)
	if err != nil {
		t.Errorf("Could not write test metadata file")
	}
}
