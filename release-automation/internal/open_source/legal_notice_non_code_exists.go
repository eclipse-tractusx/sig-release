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
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"tractusx-release-automation/internal/tractusx"
)

var consideredFileRegexes = []string{
	".md",
	".adoc",
}

var excludedFileGlobs = []string{
	".github/",
	"CHANGELOG.md",
	"CODE_OF_CONDUCT.md",
	"SECURITY.md",
	"NOTICE.md",
	"**/NOTICE.md",
	"**/DOCKER_NOTICE.md",
	"node_modules/",
	"**/node_modules/",
}

type NoticeForNonCodeExists struct {
	baseDir string
}

func NewNoticeForNonCodeExists(baseDir string) tractusx.QualityGuideline {
	return &NoticeForNonCodeExists{baseDir}
}

func (r *NoticeForNonCodeExists) IsOptional() bool {
	return false
}

func (r *NoticeForNonCodeExists) Name() string {
	return "TRG 7.07 - Notice Section For Non-Code (Documentation)"
}

func (r *NoticeForNonCodeExists) Description() string {
	return "For text files, like files in the markdown format, the attribution is done directly in the file as described in this section. The attribution is shown with an example for a CC-BY-4.0 licensed markdown file. For other formats like slides, pdf, and others adapt the information in an adequate way."
}

func (r *NoticeForNonCodeExists) ExternalDescription() string {
	return "https://eclipse-tractusx.github.io/docs/release/trg-7/trg-7-07#documentation"
}

func (r *NoticeForNonCodeExists) BaseDir() string {
	return r.baseDir
}

func (r *NoticeForNonCodeExists) Test() *tractusx.QualityResult {
	checkPassed := true

	excludedGlobs := getExcludedGlobs(r.baseDir)
	log.Printf("Going to exclude the following files for notice section check: %v", excludedGlobs)

	files, err := collectDocumentationFiles(r.baseDir, excludedGlobs)

	if err != nil {
		log.Printf("Could not collect documentation files due ot error: %s", err)
		return &tractusx.QualityResult{ErrorDescription: "Unable to find all markdown or adoc files!", Passed: false}
	}

	invalidFiles := []string{}
	for _, file := range files {
		_, err := checkNoticeSection(file)

		if err != nil {
			checkPassed = false
			invalidFiles = append(invalidFiles, file)
			log.Printf("Could not find notice section with heading 2, SPDX-License-Identifier, SPDX-File-CopyrightText, Source URL for file: %s", file)
		}
	}

	if checkPassed == false {
		return &tractusx.QualityResult{ErrorDescription: fmt.Sprintf("Could not find notice section with heading 2, SPDX-License-Identifier, SPDX-File-CopyrightText, Source URL for following files: %v", invalidFiles), Passed: false}
	}
	return &tractusx.QualityResult{Passed: true}
}

// Performs the following checks:
// - file ends with heading 2 "Notice" (.adoc == NOTICE)
// - license is contained (- SPDX-License-Identifier: CC-BY-4.0)
// - at least one copyright holder (- SPDX-FileCopyrightText: YYYY Name)
// - url is contained (- Source URL)
func checkNoticeSection(filePath string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lastHeading string
	var hasSPDX bool
	var hasCopyright bool
	var hasSourceURL bool

	// Regular expressions for matching SPDX and Source URL
	spdxRegex := regexp.MustCompile(`^- SPDX-License-Identifier: .+`)
	copyrightRegex := regexp.MustCompile(`^- SPDX-FileCopyrightText: \d{4} .+`)
	sourceURLRegex := regexp.MustCompile(`^- Source URL: .+`)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Check for last heading based on file type
		if strings.HasSuffix(filePath, ".md") && strings.HasPrefix(line, "## ") {
			lastHeading = line[3:] // Get the heading text
		} else if strings.HasSuffix(filePath, ".adoc") && strings.HasPrefix(line, "== ") {
			lastHeading = line[3:] // Get the heading text
		}

		// Check for SPDX License Identifier
		if spdxRegex.MatchString(line) {
			hasSPDX = true
		}

		// Check for Copyright Holder
		if copyrightRegex.MatchString(line) {
			hasCopyright = true
		}

		// Check for Source URL
		if sourceURLRegex.MatchString(line) {
			hasSourceURL = true
		}
	}

	if err := scanner.Err(); err != nil {
		return false, err
	}

	// Validate the conditions
	isValid := false
	var errorMessages = []string{}

	//log.Printf("File %s has lastHeading '%s'.", filePath, lastHeading)
	// we could use strings.ToUpper(lastHeading) if we would like to not explicitly check the UPPERCASENES
	if lastHeading == "NOTICE" {
		isValid = hasSPDX && hasCopyright && hasSourceURL
	} else {
		isValid = false
		errorMessages = append(errorMessages, fmt.Sprintf("File %s misses Level 2 NOTICE (uppercase) header.", filePath))
	}
	if !hasSPDX {
		isValid = false
		errorMessages = append(errorMessages, fmt.Sprintf("File %s misses SPDX-License-Identifier.", filePath))
	}
	if !hasCopyright {
		isValid = false
		errorMessages = append(errorMessages, fmt.Sprintf("File %s misses SPDX-FileCopyrightText.", filePath))
	}
	if !hasSourceURL {
		isValid = false
		errorMessages = append(errorMessages, fmt.Sprintf("File %s misses Source URL .", filePath))
	}

	if errorMessages != nil && len(errorMessages) > 0 {
		return false, errors.New(strings.Join(errorMessages, "\n"))
	}

	return isValid, nil
}

// Function to collect markdown and asciidoc files considering exlusions
func collectDocumentationFiles(baseDir string, excludedGlobs []string) ([]string, error) {
	var collectedFiles []string

	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relativePath, err := filepath.Rel(baseDir, path)
		if err != nil {
			return err
		}

		// Check if directory should be excluded
		if info.IsDir() {
			if shouldExcludeDir(relativePath, excludedGlobs) {
				log.Printf("Skipping directory: %s", relativePath)
				return filepath.SkipDir
			}
			return nil
		}

		// Check if the file has the desired extensions
		if strings.HasSuffix(info.Name(), ".md") || strings.HasSuffix(info.Name(), ".adoc") {
			// Check against excluded globs
			if !matchesExcludedGlob(relativePath, excludedGlobs) {
				collectedFiles = append(collectedFiles, path)
			} else {
				log.Printf("Excluding file: %s", relativePath)
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return collectedFiles, nil
}

// Function to check if a directory should be excluded
func shouldExcludeDir(dirPath string, excludedGlobs []string) bool {
	// Normalize path separators
	normalizedPath := filepath.ToSlash(dirPath)

	for _, glob := range excludedGlobs {
		// Remove trailing slash from glob for comparison
		globPattern := strings.TrimSuffix(glob, "/")

		// Check for exact directory name match (e.g., "node_modules/")
		if normalizedPath == globPattern || strings.HasSuffix(normalizedPath, "/"+globPattern) {
			return true
		}

		// Check for glob patterns (e.g., "**/node_modules/" or ".github/")
		if strings.Contains(globPattern, "**") {
			// Convert glob to path pattern
			// "**/node_modules" should match any path ending with "/node_modules"
			if strings.HasPrefix(globPattern, "**/") {
				suffix := strings.TrimPrefix(globPattern, "**/")
				if normalizedPath == suffix || strings.HasSuffix(normalizedPath, "/"+suffix) {
					return true
				}
			}
			// Try filepath.Match as fallback
			matched, err := filepath.Match(globPattern, normalizedPath)
			if err == nil && matched {
				return true
			}
		} else {
			// For non-glob patterns, check if path starts with the pattern
			if strings.HasPrefix(normalizedPath, globPattern) {
				return true
			}
		}
	}
	return false
}

// Function to check if a file path matches any of the excluded globs
func matchesExcludedGlob(path string, excludedGlobs []string) bool {
	normalizedPath := filepath.ToSlash(path)

	for _, glob := range excludedGlobs {
		// Try direct match
		match, err := filepath.Match(glob, normalizedPath)
		if err == nil && match {
			return true
		}

		// Try matching against the basename
		match, err = filepath.Match(glob, filepath.Base(path))
		if err == nil && match {
			return true
		}

		// For patterns like "**/something", check if path contains the pattern
		if strings.Contains(glob, "**/") {
			suffix := strings.TrimPrefix(glob, "**/")
			if normalizedPath == suffix || strings.Contains(normalizedPath, "/"+suffix) {
				return true
			}
		}
	}
	return false
}

func getExcludedGlobs(dir string) []string {
	file, err := tractusx.MetadataFromLocalFile(dir)

	if err != nil {
		return []string{}
	}

	return append(excludedFileGlobs, file.LegalNoticesNonCode...)
}

func (a *NoticeForNonCodeExists) IsApplicableToCategory(category tractusx.RepoCategory) bool {
	return category == tractusx.RepoCategoryProduct || category == tractusx.RepoCategorySupport || category == tractusx.RepoCategorySpecial
}
