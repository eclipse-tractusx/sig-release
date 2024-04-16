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
	"testing"
	"tractusx-release-automation/internal/filesystem"
)

var listOfFilesToBeCreated = []string{
	"CODE_OF_CONDUCT.md",
	"CONTRIBUTING.md",
	"LICENSE",
	"NOTICE.md",
	"README.md",
	"SECURITY.md",
	"DEPENDENCIES",
}

var listOfDirsToBeCreated = []string{
	"docs",
	"charts",
}

const metadataTestFile = "./test/metadata_test_template.yaml"

func TestShouldPassIfRepoStructureExistsWithoutOptional(t *testing.T) {
	dir := t.TempDir()
	setEnv(dir, t)

	filesystem.CreateFiles(createFilesPaths(dir, listOfFilesToBeCreated))
	filesystem.CreateDirs(createFilesPaths(dir, listOfDirsToBeCreated))

	repoStructureTest := NewRepoStructureExists(dir)
	result := repoStructureTest.Test()

	if !result.Passed {
		t.Errorf("Structure exists with optional files, but test still fails.")
	}
}

func TestShouldPassIfRepoStructureExistsWithOptional(t *testing.T) {
	dir := t.TempDir()
	setEnv(dir, t)

	listOfFilesToBeCreated = append(listOfFilesToBeCreated, []string{"INSTALL.md", "AUTHORS.md"}...)

	filesystem.CreateFiles(createFilesPaths(dir, listOfFilesToBeCreated))
	filesystem.CreateDirs(createFilesPaths(dir, listOfDirsToBeCreated))

	repoStructureTest := NewRepoStructureExists(dir)
	result := repoStructureTest.Test()

	if !result.Passed {
		t.Errorf("Structure exists without optional files, but test still fails.")
	}

}

func TestShouldFailIfRepoStructureIsMissing(t *testing.T) {
	dir := t.TempDir()
	setEnv(dir, t)

	repoStructureTest := NewRepoStructureExists(dir)

	result := repoStructureTest.Test()

	if result.Passed {
		t.Errorf("RepoStructureExist should fail if repo structure exists.")
	}
}

func TestShouldPassWithMultipleDependenciesFiles(t *testing.T) {
	dir := t.TempDir()
	setEnv(dir, t)

	newListOfFilesToBeCreated := append(listOfFilesToBeCreated[:len(listOfFilesToBeCreated)-1], []string{"DEPENDENCIES_FRONTEND", "DEPENDENCIES_BACKEND"}...)
	filesystem.CreateFiles(createFilesPaths(dir, newListOfFilesToBeCreated))
	filesystem.CreateDirs(createFilesPaths(dir, listOfDirsToBeCreated))

	repoStructureTest := NewRepoStructureExists(dir)
	result := repoStructureTest.Test()

	if !result.Passed {
		t.Errorf("There is multiple DEPENDENCIES files, test should pass.")
	}
}

func setEnv(dir string, t *testing.T) {
	copyTemplateFileTo(path.Join(dir, ".tractusx"), t)
	_ = os.Setenv("GITHUB_REPOSITORY", "eclipse-tractusx/sig-infra")
	_ = os.Setenv("GITHUB_REPOSITORY_OWNER", "tester")
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

func createFilesPaths(dir string, files []string) []string {
	fullPaths := []string{}
	for _, file := range files {
		fullPaths = append(fullPaths, path.Join(dir, file))
	}
	return fullPaths
}
