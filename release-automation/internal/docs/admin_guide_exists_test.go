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
)

func TestPassIfAnyMarkdownFileIsPresent(t *testing.T) {
	dir := t.TempDir()
	filePath := path.Join(dir, "docs", "admin", "README.md")

	err := os.MkdirAll(path.Dir(filePath), os.ModePerm)
	if err != nil {
		t.Fatalf("failed to create directories: %v", err)
	}

	fmt.Printf("%s\n", filePath)
	_, err = os.Create(filePath)
	if err != nil {
		t.Fatalf("failed to create file: %v", err)
	}

	fmt.Printf("%s\n", filePath)
	_, _ = os.Create(filePath)

	result := NewAdminGuideExists(dir).Test()

	if !result.Passed {
		t.Error("Admin Guide is not in directory")
	}
}

func TestFailIfMarkdownFileIsMissing(t *testing.T) {
	dir := t.TempDir()
	filePath := path.Join(dir, "docs", "admin", "README.txt")
	err := os.MkdirAll(path.Dir(filePath), os.ModePerm)
	if err != nil {
		t.Fatalf("failed to create directories: %v", err)
	}

	fmt.Printf("%s\n", filePath)
	_, err = os.Create(filePath)
	if err != nil {
		t.Fatalf("failed to create file: %v", err)
	}

	fmt.Printf("%s\n", filePath)
	_, _ = os.Create(filePath)

	result := NewAdminGuideExists(dir).Test()

	if result.Passed {
		t.Error("AdminGuideExists should fail if no Markdown is present in docs/admin.")
	}
}

func TestProvidesErrorDescriptionIfFailing(t *testing.T) {
	result := NewChangelogExists("./").Test()

	if result.ErrorDescription == "" {
		t.Errorf("Failing tests should provide an error description")
	}
}
