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
	"context"
	"fmt"

	"tractusx-release-automation/internal/tractusx"

	"github.com/google/go-github/v53/github"
)

type defaultBranch struct {
	baseDir string
}

func NewDefaultBranch(baseDir string) tractusx.QualityGuideline {
	return &defaultBranch{baseDir}
}

func (d defaultBranch) Name() string {
	return "TRG 2.01 - Default Branch"
}

func (d defaultBranch) Description() string {
	return "Default branch must be main."
}

func (d defaultBranch) ExternalDescription() string {
	return "https://eclipse-tractusx.github.io/docs/release/trg-2/trg-2-1"
}

func (d *defaultBranch) BaseDir() string {
	return d.baseDir
}

func (d defaultBranch) Test() *tractusx.QualityResult {
	r := GetRepoBaseInfo(d.baseDir)
	repoInfo := getRepoInfo(r)

	if *repoInfo.Fork {
		// There is no need to enforce default branches on forks
		// Since all the other checks should be executable on forks, we cannot let this single check break a workflow
		return &tractusx.QualityResult{Passed: true}
	}

	if *repoInfo.DefaultBranch != "main" {
		return &tractusx.QualityResult{
			Passed:           false,
			ErrorDescription: "Default branch not set to 'main'.",
		}
	}

	return &tractusx.QualityResult{Passed: true}
}

func (d defaultBranch) IsOptional() bool {
	return false
}

func getRepoInfo(repo *RepoInfo) *github.Repository {
	ctx := context.Background()
	client := *github.NewClient(nil)

	repoInfo, _, err := client.Repositories.Get(ctx, repo.Owner, repo.Reponame)
	if err != nil {
		fmt.Printf("Error querying GitHub API:\n%v\n", err)
	}

	return repoInfo
}

func (a *defaultBranch) IsApplicableToCategory(category tractusx.RepoCategory) bool {
	return category == tractusx.RepoCategoryProduct || category == tractusx.RepoCategorySupport || category == tractusx.RepoCategorySpecial
}
