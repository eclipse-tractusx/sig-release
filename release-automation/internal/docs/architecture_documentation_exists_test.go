/*******************************************************************************
 * Copyright (c) 2025 Fraunhofer-Gesellschaft zur Foerderung der angewandten Forschung e.V. (represented by Fraunhofer ISST)
 * Copyright (c) 2025 Contributors to the Eclipse Foundation
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
package docs

import (
	"fmt"
	"os"
	"path"
	"testing"

	"tractusx-release-automation/internal/tractusx"

	"gopkg.in/yaml.v3"
)

func TestShouldFailIfArchitectureDocumentationDoesNotExist(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := path.Join(tmpDir, "docs", "architecture", "WrongFile.yaml")

	err := os.MkdirAll(path.Dir(filePath), os.ModePerm)
	if err != nil {
		t.Fatalf("failed to create directories: %v", err)
	}

	_, err = os.Create(filePath)
	if err != nil {
		t.Fatalf("failed to create file: %v", err)
	}
	fmt.Printf("%s\n", filePath)

	result := NewArchitectureDocumentationExists(tmpDir).Test()

	if result.Passed {
		t.Errorf("Architecture documentation check should fail, if no architecture documentation is present")
	}
}

func TestShouldPassIfArchitectureDocumentationFromTractusxFileExists(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := path.Join(tmpDir, "docs", "architecture", "arc42", "markdown.md")

	err := os.MkdirAll(path.Dir(filePath), os.ModePerm)
	if err != nil {
		t.Fatalf("failed to create directories: %v", err)
	}

	_, err = os.Create(filePath)
	if err != nil {
		t.Fatalf("failed to create file: %v", err)
	}

	saveMetadataConfigToSkip(path.Join("docs", "architecture", "arc42", "markdown.md"), tmpDir)
	fmt.Printf("%s\n", filePath)

	result := NewArchitectureDocumentationExists(tmpDir).Test()

	if !result.Passed {
		t.Errorf("Architecture documentation check should pass, because the file has been specified via .tractusx.")
	}
}

func TestShouldFailIfArchitectureDocumentationFromTractusxFileDoesNotExist(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := path.Join(tmpDir, "docs", "architecture", "arc42", "WRONG.md")

	err := os.MkdirAll(path.Dir(filePath), os.ModePerm)
	if err != nil {
		t.Fatalf("failed to create directories: %v", err)
	}

	_, err = os.Create(filePath)
	if err != nil {
		t.Fatalf("failed to create file: %v", err)
	}

	saveMetadataConfigToSkip(path.Join("docs", "architecture", "arc42", "markdown.md"), tmpDir)
	fmt.Printf("%s\n", filePath)

	result := NewArchitectureDocumentationExists(tmpDir).Test()

	if result.Passed {
		t.Errorf("Architecture documentation check should fail, if no architecture documentation is present in .tractux file location")
	}
}

func TestShouldPassIfArchitectureDocumentationMdExists(t *testing.T) {

	tmpDir := t.TempDir()
	filePath := path.Join(tmpDir, "docs", "architecture", "README.md")

	err := os.MkdirAll(path.Dir(filePath), os.ModePerm)
	if err != nil {
		t.Fatalf("failed to create directories: %v", err)
	}

	_, err = os.Create(filePath)
	if err != nil {
		t.Fatalf("failed to create file: %v", err)
	}

	result := NewArchitectureDocumentationExists(tmpDir).Test()

	if !result.Passed {
		t.Errorf("Architecture documentation exists, but test still fails")
	}

}

func TestShouldPassIfArchitectureDocumentationAdocExists(t *testing.T) {

	tmpDir := t.TempDir()
	filePath := path.Join(tmpDir, "docs", "architecture", "README.adoc")

	err := os.MkdirAll(path.Dir(filePath), os.ModePerm)
	if err != nil {
		t.Fatalf("failed to create directories: %v", err)
	}

	_, err = os.Create(filePath)
	if err != nil {
		t.Fatalf("failed to create file: %v", err)
	}

	result := NewArchitectureDocumentationExists(tmpDir).Test()

	if !result.Passed {
		t.Errorf("Architecture documentation exists, but test still fails")
	}

}

func saveMetadataConfigToSkip(architectureDocPath string, dir string) {
	metadata := tractusx.Metadata{
		ConfigureReleaseChecks: tractusx.ConfigureReleaseChecks{
			ArchitectureDocEntryPath: architectureDocPath,
		},
	}

	bytes, _ := yaml.Marshal(&metadata)
	_ = os.WriteFile(path.Join(dir, ".tractusx"), bytes, 0644)
}
