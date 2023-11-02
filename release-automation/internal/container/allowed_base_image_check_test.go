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
	"math/rand"
	"os"
	"path"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
	"tractusx-release-automation/internal/tractusx"
)

func TestShouldPassIfNoDockerfilePresent(t *testing.T) {
	result := NewAllowedBaseImage("./").Test()

	if result == nil || result.Passed == false {
		t.Errorf("Allowed base image check should pass, if there is no Dockerfile found")
	}
}

func TestShouldFailIfDockerfileWithUnapprovedBaseImagePresent(t *testing.T) {
	file := correttoDockerfile()
	_ = file.writeTo("./")
	defer os.Remove(file.filename)

	result := NewAllowedBaseImage("./").Test()

	if result.Passed {
		t.Errorf("Allowed based image check should fail, if Dockerfile with unapproved base image found")
	}
}

func TestShouldPassIfTemurinIsUsedAsBaseImage(t *testing.T) {
	file := newDockerfile().appendCommand("FROM eclipse-temurin:17")
	_ = file.writeTo("./")
	defer os.Remove(file.filename)

	if !(NewAllowedBaseImage("./").Test().Passed) {
		t.Errorf("eclipse/temurin should be recognized as approved base image")
	}
}

func TestShouldNotFailIfOnlyBuildLayerDeviatesFromTemurin(t *testing.T) {
	file := newDockerfile().
		appendCommand("FROM amazoncorretto:8 as builder").
		appendCommand("COPY . .").
		appendCommand("FROM eclipse-temurin:17")
	_ = file.writeTo("./")
	defer os.Remove(file.filename)

	if !(NewAllowedBaseImage("./").Test().Passed) {
		t.Errorf("Unapproved images in build layers should not let the check fail")
	}
}

func TestShouldFailForUnallowedBaseImageInSubdirectory(t *testing.T) {
	subdirectory := randomSubDir(t)
	dockerfile := correttoDockerfile()
	_ = dockerfile.writeTo(subdirectory)
	defer os.RemoveAll(subdirectory)

	result := NewAllowedBaseImage("./").Test()

	if result.Passed {
		t.Errorf("Unapproved base images should also be detected in sub directories")
	}
}

func TestShouldFailIfAtLeastOneDockerfileWithUnallowedBaseImageIsFound(t *testing.T) {
	firstSubdir := randomSubDir(t)
	secondSubDir := randomSubDir(t)
	corretto := correttoDockerfile()
	temurin := temurinDockerfile()

	_ = corretto.writeTo(firstSubdir)
	_ = temurin.writeTo(secondSubDir)
	defer os.RemoveAll(firstSubdir)
	defer os.RemoveAll(secondSubDir)

	result := NewAllowedBaseImage("./").Test()

	if result.Passed {
		t.Errorf("Base image check should fail, if at least one Dockerfile with unallowed base image is found!")
	}
}

func TestShouldAllowBaseImagesFromWhitelist(t *testing.T) {
	baseImageAllowList = []string{"my/baseimage", "my/other/baseimage"}

	firstSubdir := randomSubDir(t)
	secondSubDir := randomSubDir(t)
	allowedBaseImage := dockerFileWithBaseImage(baseImageAllowList[0])
	otherAllowedBaseImage := dockerFileWithBaseImage(baseImageAllowList[1])
	_ = allowedBaseImage.writeTo(firstSubdir)
	_ = otherAllowedBaseImage.writeTo(secondSubDir)
	defer os.RemoveAll(firstSubdir)
	defer os.RemoveAll(secondSubDir)

	result := NewAllowedBaseImage("./").Test()

	if !result.Passed {
		t.Errorf("Should allow base images from whitelist")
	}
}

func TestShouldIncludeAllUnallowedBaseImagesInErrorDescription(t *testing.T) {
	firstSubdir := randomSubDir(t)
	secondSubDir := randomSubDir(t)
	_ = dockerFileWithBaseImage("badBaseImage").writeTo(firstSubdir)
	_ = dockerFileWithBaseImage("another/unallowed").writeTo(secondSubDir)
	defer os.RemoveAll(firstSubdir)
	defer os.RemoveAll(secondSubDir)

	result := NewAllowedBaseImage("./").Test()

	if !strings.Contains(result.ErrorDescription, "badBaseImage") || !strings.Contains(result.ErrorDescription, "another/unallowed") {
		t.Errorf("Error message should contain all denied base images")
	}

}

func TestShouldSkipDockerfilesThatAreExcludedByMetadataConfig(t *testing.T) {
	tempDir := t.TempDir()
	baseImageAllowList = []string{"only-this"}
	_ = dockerFileWithBaseImage("an-unaligned-one").writeTo(tempDir)
	saveMetadataConfigToSkip(path.Join(tempDir, "Dockerfile"), tempDir)

	result := NewAllowedBaseImage(tempDir).Test()

	if !result.Passed {
		t.Errorf("Aligned base image check should succeed, if non-aligned image is configured to be ignored")
	}
}

func dockerFileWithBaseImage(baseImage string) *dockerfile {
	return newDockerfile().appendCommand("FROM " + baseImage)
}

func randomSubDir(t *testing.T) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 10)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	subDirPath := string(b)
	if err := os.MkdirAll(subDirPath, 0770); err != nil {
		t.Errorf("Could not create random subdirectory at path: " + subDirPath)
	}

	return subDirPath
}

func saveMetadataConfigToSkip(dockerfilePath string, dir string) {
	metadata := tractusx.Metadata{
		SkipReleaseChecks: tractusx.SkipReleaseChecks{
			AlignedBaseImages: []string{
				dockerfilePath,
			},
		},
	}

	bytes, _ := yaml.Marshal(&metadata)
	_ = os.WriteFile(path.Join(dir, ".tractusx"), bytes, 0644)
}
