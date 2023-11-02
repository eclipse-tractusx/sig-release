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

package dashboard

import (
	"errors"
	"fmt"
	"log"
	"os"

	"gopkg.in/src-d/go-git.v4"
)

// cloneRepo creates a temporary directory and clones the given repo into it.
// If successful, the directory, that the repo was cloned into is returned.
// Error in case the temp dir cannot be created or the repo not cloned
func cloneRepo(repo Repository) (string, error) {
	dir, err := os.MkdirTemp("", fmt.Sprintf("%s-*", repo.Name))
	if err != nil {
		return "", errors.New("could not create temp dir to clone repo into")
	}
	log.Printf("Created temp dir for repo check: %s; Cloning %s", dir, repo.URL)

	if _, err := git.PlainClone(dir, false, &git.CloneOptions{URL: repo.URL, Depth: 0}); err != nil {
		return "", errors.New(fmt.Sprintf("Could not clone repo %s (%s)", repo.Name, repo.URL))
	}
	return dir, nil
}
