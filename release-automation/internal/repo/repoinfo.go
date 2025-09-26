/*******************************************************************************
 * Copyright (c) 2023 Contributors to the Eclipse Foundation
 * Copyright (c) 2025 Fraunhofer-Gesellschaft zur Foerderung der angewandten Forschung e.V. (represented by Fraunhofer ISST)
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
	"fmt"
	"os"
	"path"
	"strings"

	"tractusx-release-automation/internal/tractusx"

	"github.com/go-ini/ini"
)

// RepoInfo is a struct to keep basic information about a repository
type RepoInfo struct {
	Owner    string
	Reponame string
}

// GetRepoBaseInfo gathers information about repo owner and name.
// It leverages environment variables typically available in GitHub workflows.
// As fallback option, the local git config (.git/config) file is used.
// Results are returned as *RepoInfo
func GetRepoBaseInfo(repoDir string) *RepoInfo {
	const (
		BASEURL = "https://github.com/"
		SSHBASE = "git@github.com:"
		SECTION = `remote "origin"`
		SUFFIX  = ".git"
	)

	if os.Getenv("GITHUB_REPOSITORY") != "" && os.Getenv("GITHUB_REPOSITORY_OWNER") != "" {
		// env variable is available when executed as GH Action/Workflow/Check
		result := RepoInfo{
			Owner:    os.Getenv("GITHUB_REPOSITORY_OWNER"),
			Reponame: strings.Split(os.Getenv("GITHUB_REPOSITORY"), "/")[1],
		}

		return &result
	}

	// Parse local git configuration when executing locally
	cfg, err := ini.Load(path.Join(repoDir, ".git/config"))
	if err != nil {
		fmt.Printf("Failed to read file: %v", err)
		return &RepoInfo{Owner: "unknown", Reponame: "unknown"}
	}

	url := cfg.Section(SECTION).Key("url").String()

	var repoSplitInfo []string
	if strings.Contains(url, BASEURL) {
		repoSplitInfo = strings.Split(strings.TrimSuffix(strings.TrimPrefix(url, BASEURL), SUFFIX), "/")
	} else if strings.Contains(url, SSHBASE) {
		repoSplitInfo = strings.Split(strings.TrimSuffix(strings.TrimPrefix(url, SSHBASE), SUFFIX), "/")
	}

	result := RepoInfo{
		Owner:    repoSplitInfo[0],
		Reponame: repoSplitInfo[1],
	}

	return &result
}

func isLeadingRepo(repoDir string) bool {
	metadata, err := tractusx.MetadataFromLocalFile(repoDir)
	repoInfo := GetRepoBaseInfo(repoDir)
	fullRepoName := "https://github.com/eclipse-tractusx/" + (*repoInfo).Reponame

	if err != nil || metadata.LeadingRepository != fullRepoName {
		return false
	}

	return true
}

func isProductRepo(repoDir string) bool {
	metadata, err := tractusx.MetadataFromLocalFile(repoDir)

	if err != nil || metadata.RepoCategory != tractusx.RepoCategoryProduct {
		return false
	}

	return true
}
