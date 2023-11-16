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

package docs

import (
	"os"
	"path"
	"testing"
)

func TestShouldFailIfFileDoesNotExist(t *testing.T) {
	result := NewInstallExists("./").Test()

	if result.Passed {
		t.Errorf("InstallExists should fail if no INSTALL.md present")
	}
}

func TestShouldProvideErrorDescriptionOnFailure(t *testing.T) {
	expectedError := "Optional file INSTALL.md not found in current directory."

	result := NewInstallExists("./").Test()

	if result.ErrorDescription != expectedError {
		t.Errorf("Install.md check does not provide correct error description fon failing check!\n"+
			"expected: %s\ngot: %s", expectedError, result.ErrorDescription)
	}
}

func TestShouldPassIfFileExists(t *testing.T) {
	if _, err := os.Create("INSTALL.md"); err != nil {
		t.Errorf("Error creating test preconditions!")
	}
	defer os.Remove("INSTALL.md")

	result := NewInstallExists("./").Test()

	if !result.Passed {
		t.Errorf("INSTALL exists, but test still fails")
	}
}

func TestShouldFindFileAtGivenBaseDir(t *testing.T) {
	tempDir := t.TempDir()
	if _, err := os.Create(path.Join(tempDir, "INSTALL.md")); err != nil {
		t.Errorf("Error creating test preconditions!")
	}

	result := NewInstallExists(tempDir).Test()

	if !result.Passed {
		t.Errorf("INSTALL exists, but test still fails")
	}
}
