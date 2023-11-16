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

package container

import (
	"os"
	"path"
	"testing"
)

func TestShouldReturnEmptyStringIfNoDockerfilePresent(t *testing.T) {
	foundFiles := findDockerfilesAt("./")
	if len(foundFiles) != 0 {
		t.Errorf("#findDockerfilesAt should return empty string if no file could be found")
	}
}

func TestShouldFindSingleDockerfile(t *testing.T) {
	tempDir := t.TempDir()
	if err := newDockerfile().writeTo(tempDir); err != nil {
		t.Errorf("Failed to write Dockerfile for test in temp dire")
	}

	foundFiles := findDockerfilesAt(tempDir)
	if len(foundFiles) != 1 {
		t.Errorf("Dockerfile written to temp dir could not be found by #findDockerfilesAt")
	}
}

func TestShouldFindDockerfileInSubdirectory(t *testing.T) {
	tempDir := t.TempDir()
	desiredPath := tempDir + "/abc/def"
	if err := os.MkdirAll(desiredPath, 0777); err != nil {
		t.Errorf("Could not create folder structure for test; err %s", err)
	}

	if err := newDockerfile().writeTo(desiredPath); err != nil {
		t.Errorf("Could not write test Dockerfile to desired temp subdirectory")
	}

	foundFiles := findDockerfilesAt(tempDir)

	if len(foundFiles) == 0 {
		t.Errorf("Could not find Dockerfile in subdirectory")
	}
}

func TestShouldFindMultipleDockerfiles(t *testing.T) {
	tempDir := t.TempDir()
	secondTempDir := t.TempDir()

	_ = newDockerfile().writeTo(tempDir)
	_ = newDockerfile().writeTo(secondTempDir)

	foundFiles := findDockerfilesAt(path.Join(tempDir, "../"))

	if len(foundFiles) != 2 {
		t.Errorf("Did not find all Dockerfiles")
	}
}

func TestShouldFindDockerfilesWithUnusualNames(t *testing.T) {
	tempDir := t.TempDir()
	dockerfile := newDockerfile()
	dockerfile.filename = "Dockerfile.full"
	_ = dockerfile.writeTo(tempDir)
	dockerfile.filename = "db.Dockerfile"
	_ = dockerfile.writeTo(tempDir)
	dockerfile.filename = "Dockerfile-dev"
	_ = dockerfile.writeTo(tempDir)

	foundFiles := findDockerfilesAt(tempDir)

	if len(foundFiles) != 3 {
		t.Errorf("Could not find all uncommonly named Dockerfiles! Found %s", foundFiles)
	}
}

func TestShouldReturnErrorIfNoDockerfileFoundAtPath(t *testing.T) {
	if _, err := dockerfileFromPath(t.TempDir()); err == nil {
		t.Errorf("#dockerfileFromPath should return an error if the file to read does not exist!")
	}
}

func TestShouldReadDockerfileIfPresent(t *testing.T) {
	tempDir := t.TempDir()
	testDockerfilePath := path.Join(tempDir, "Dockerfile")
	if err := os.WriteFile(testDockerfilePath, []byte("FROM ubuntu"), 0777); err != nil {
		t.Errorf("Could not write test Dockerfile")
	}

	dockerfile, err := dockerfileFromPath(testDockerfilePath)
	if err != nil {
		t.Errorf("Reading a dockerifle from path should not result in an error")
	}

	if dockerfile.commands[0] != "FROM ubuntu" {
		t.Errorf("Did not read dockerfile contents from path")
	}
}
