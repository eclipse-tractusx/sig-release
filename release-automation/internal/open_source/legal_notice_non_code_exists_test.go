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

package open_source

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"testing"

	"tractusx-release-automation/internal/tractusx"

	"gopkg.in/yaml.v3"
)

type TestStruct struct {
	content      string
	dirPath      string
	fileName     string
	expected     bool
	expectingErr bool
}

func TestShouldPassIfInvalidFilesAreExcluded(t *testing.T) {
	tests := []TestStruct{
		{
			content: `## NOTICE
- SPDX-License-Identifier: CC-BY-4.0
- SPDX-FileCopyrightText: 2023 John Doe
- Source URL: https://example.com`,
			dirPath:  "./",
			fileName: "valid.md",
			expected: true,
		},
		{
			content: `## notice
- SPDX-License-Identifier: CC-BY-4.0
- SPDX-FileCopyrightText: 2023 John Doe
- Source URL: https://example.com`,
			dirPath:  "subdir",
			fileName: "invalid_notice.adoc",
			expected: true,
		},
		{
			content: `## NOTICE
- SPDX-License-Identifier: CC-BY-4.0
- Source URL: https://example.com`,
			dirPath:  "subdir",
			fileName: "missing_copyright_and_excluded.md",
			expected: false,
		},
		{
			content: `## NOTICE
- SPDX-FileCopyrightText: 2023 John Doe
- Source URL: https://example.com`,
			dirPath:  "subdir",
			fileName: "wrong_filetype.yaml",
			expected: false,
		},
		{
			content: `## NOTICE
- SPDX-FileCopyrightText: 2023 John Doe
- Source URL: https://example.com`,
			dirPath:  ".github",
			fileName: "excluded.md",
			expected: false,
		},
		{
			content: `## NOTICE
- SPDX-FileCopyrightText: 2023 John Doe
- Source URL: https://example.com`,
			dirPath:  ".github/ISSUE_TEMPLATE",
			fileName: "excluded.md",
			expected: false,
		},
	}

	dir := t.TempDir()

	exclusionSlice := []string{
		"subdir/missing_copyright_and_excluded.md",
	}

	saveMetadataConfigToExclude(exclusionSlice, dir)

	for _, test := range tests {

		filePath := path.Join(dir, test.dirPath, test.fileName)

		err := os.MkdirAll(path.Dir(filePath), os.ModePerm)
		// Create a temporary file
		tmpFile, err := os.Create(filePath)
		if err != nil {
			t.Fatalf("Could not create temp file: %v", err)
		} else {
			log.Printf("Created file %s.", filePath)
		}
		defer os.Remove(tmpFile.Name()) // Clean up

		// Write test content to the file
		if _, err := tmpFile.WriteString(test.content); err != nil {
			t.Fatalf("Could not write to temp file: %v", err)
		}
		tmpFile.Close()
	}

	files, err := collectDocumentationFiles(dir, getExcludedGlobs(dir))

	if err != nil {
		t.Errorf("Unexpected error during collection of files: %v", err)
	}

	if len(files) != 2 {
		t.Errorf("Two files should have been collected, but collected: %d", len(files))
	}

	matchingTests := []*TestStruct{}

	for _, file := range files {
		matchingTestCase := getMatchingTest(file, tests)
		matchingTests = append(matchingTests, matchingTestCase)

		if matchingTestCase == nil {
			t.Errorf("Unexpected error as a test should be found for %s", file)
		} else {
			if matchingTestCase.expected == false {
				t.Errorf("File %s for test case %v should have not been collected.", file, matchingTestCase)
			} else {
				log.Printf("File %s for test case %v has been collected successfully.", file, matchingTestCase)
			}
		}
	}
}

func getMatchingTest(fileName string, tests []TestStruct) *TestStruct {
	for _, testCase := range tests {
		_, err := filepath.Match(fileName, path.Join(testCase.dirPath, testCase.fileName))
		if err == nil {
			return &testCase
		}
	}
	return nil
}

func saveMetadataConfigToExclude(LegalNoticesNonCode []string, dir string) {
	metadata := tractusx.Metadata{
		SkipReleaseChecks: tractusx.SkipReleaseChecks{
			LegalNoticesNonCode: LegalNoticesNonCode,
		},
	}

	bytes, _ := yaml.Marshal(&metadata)
	_ = os.WriteFile(path.Join(dir, ".tractusx"), bytes, 0644)
}

func TestCheckNoticeSection(t *testing.T) {
	tests := []struct {
		content      string
		fileName     string
		expected     bool
		expectingErr bool
	}{
		{
			content: `## NOTICE
- SPDX-License-Identifier: CC-BY-4.0
- SPDX-FileCopyrightText: 2023 John Doe
- Source URL: https://example.com`,
			fileName:     "valid.md",
			expected:     true,
			expectingErr: false,
		},
		{
			content: `## notice
- SPDX-License-Identifier: CC-BY-4.0
- SPDX-FileCopyrightText: 2023 John Doe
- Source URL: https://example.com`,
			fileName:     "invalid_notice.md",
			expected:     false,
			expectingErr: false,
		},
		{
			content: `## NOTICE
- SPDX-License-Identifier: CC-BY-4.0
- Source URL: https://example.com`,
			fileName:     "missing_copyright.md",
			expected:     false,
			expectingErr: false,
		},
		{
			content: `## NOTICE
- SPDX-FileCopyrightText: 2023 John Doe
- Source URL: https://example.com`,
			fileName:     "missing_license.md",
			expected:     false,
			expectingErr: false,
		},
		{
			content: `== NOTICE
- SPDX-License-Identifier: CC-BY-4.0
- SPDX-FileCopyrightText: 2023 John Doe
- Source URL: https://example.com`,
			fileName:     "valid.adoc",
			expected:     true,
			expectingErr: false,
		},
	}

	for _, test := range tests {
		// Create a temporary file
		tmpFile, err := os.Create(test.fileName)
		if err != nil {
			t.Fatalf("Could not create temp file: %v", err)
		}
		defer os.Remove(tmpFile.Name()) // Clean up

		// Write test content to the file
		if _, err := tmpFile.WriteString(test.content); err != nil {
			t.Fatalf("Could not write to temp file: %v", err)
		}
		tmpFile.Close()

		// Check file requirements
		isValid, err := checkNoticeSection(tmpFile.Name())

		if (err != nil) != test.expectingErr {
			t.Errorf("Unexpected error for %s: %v", test.fileName, err)
		}

		if isValid != test.expected {
			t.Errorf("Expected %v for %s, got %v", test.expected, test.fileName, isValid)
		}
	}
}
