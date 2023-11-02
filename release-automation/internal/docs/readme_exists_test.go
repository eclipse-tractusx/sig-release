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

func TestShouldFailIfReadmeDoesNotExist(t *testing.T) {
	result := NewReadmeExists("./").Test()

	if result.Passed {
		t.Errorf("Readme check should fail, if no README file present")
	}
}

func TestProvideErrorDescriptionOnFailingTest(t *testing.T) {
	expectedError := "Did not find a README.md file in current directory!"

	result := NewReadmeExists("./").Test()

	if result.ErrorDescription != expectedError {
		t.Errorf("Readme check does not provide correct error description on failing check! \nexprected: %s, \ngot: %s", expectedError, result.ErrorDescription)
	}
}

func TestShouldPassIfReadmeExists(t *testing.T) {
	if _, err := os.Create("README.md"); err != nil {
		t.Errorf("Could not create README.md for test")
	}
	defer os.Remove("README.md")

	result := NewReadmeExists("./").Test()

	if !result.Passed {
		t.Errorf("README exists, but test still fails")
	}

}

func TestShouldFindReadmeAtGivenBasePath(t *testing.T) {
	dir := t.TempDir()
	if _, err := os.Create(path.Join(dir, "README.md")); err != nil {
		t.Errorf("Could not create README.md for test")
	}

	result := NewReadmeExists(dir).Test()

	if !result.Passed {
		t.Errorf("Could not find README.md at given base path")
	}
}
