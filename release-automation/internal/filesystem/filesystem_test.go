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

package filesystem

import (
	"os"
	"testing"
)

var testFile = "_testfile"
var testDir = "_testdir"

func TestShouldPassIfFileGetsCreated(t *testing.T) {
	CreateFiles([]string{testFile})
	defer os.Remove(testFile)

	if _, err := os.Stat(testFile); err != nil {
		t.Errorf("File was supposed to be created but there is none.")
	}
}

func TestShouldPassIfDirGetsCreated(t *testing.T) {
	CreateDirs([]string{testDir})
	defer os.Remove(testDir)

	if finfo, err := os.Stat(testDir); err != nil || !finfo.IsDir() {
		t.Errorf("Directory was supposed to be created but there is none.")
	}
}

func TestShouldPassIfFileGetsDeleted(t *testing.T) {
	os.Create(testFile)
	CleanFiles([]string{testFile})

	if _, err := os.Stat(testFile); err == nil {
		t.Errorf("File was to be deleted but is still exists.")
	}
}

func TestShouldPassIfReturnsMissingFile(t *testing.T) {
	missingFiles := CheckMissingFiles([]string{testFile})

	if missingFiles[0] != testFile {
		t.Errorf("Missing file is not the one expected.")
	}
}
