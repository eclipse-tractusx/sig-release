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
	"context"
	"fmt"

	"github.com/google/go-github/v53/github"
	"tractusx-release-automation/internal/tractusx"
)

type defaultBranch struct {
}

func NewDefaultBranch() tractusx.QualityGuideline {
	return &defaultBranch{}
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

func (d defaultBranch) Test() *tractusx.QualityResult {
	repoInfo := getRepoInfo(GetRepoBaseInfo())

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
