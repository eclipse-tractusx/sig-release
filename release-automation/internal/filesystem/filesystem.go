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
	"log"
	"os"
	"strings"
)

func CreateFiles(files []string) {
	for _, file := range files {
		if _, err := os.Create(file); err != nil {
			log.Fatalf("could not create file %s. Error %s", file, err)
		}
	}
}

func CreateDirs(dirs []string) {
	for _, dir := range dirs {
		os.MkdirAll(dir, 0750)
	}
}

func CleanFiles(files []string) {
	for _, file := range files {
		os.Remove(file)
	}
}

func CheckMissingFiles(listOfFiles []string) []string {
	var missingFiles []string

	for _, file := range listOfFiles {

		if _, err := os.Stat(file); os.IsNotExist(err) {
			missingFiles = append(missingFiles, file)
		}
	}
	return missingFiles
}

func FindPrefixedFiles(dir string, prefix string) []string {
	var foundFiles []string
	fs, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}
	for _, f := range fs {
		if !f.IsDir() && strings.HasPrefix(f.Name(), prefix) {
			foundFiles = append(foundFiles, f.Name())
		}
	}
	return foundFiles
}
