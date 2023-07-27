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

package tractusx

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
)

const gitHubOrg = "eclipse-tractusx"

type Product struct {
	Name         string
	LeadingRepo  string
	Repositories []string
}

type ReleaseGuidelineCheck struct {
	Repository string
	Guideline  string
	Passed     bool
}

func QueryProducts() []Product {
	opt := &github.RepositoryListByOrgOptions{Type: "public"}
	client := getGitHubClient()
	repos, _, err := client.Repositories.ListByOrg(context.Background(), gitHubOrg, opt)

	for _, repo := range repos {
		log.Printf("Getting tractusx metadata for repository: %s", *repo.Name)
		contents, _, _, err := client.Repositories.GetContents(context.Background(), gitHubOrg, *repo.Name, ".tractusx", nil)
		if err != nil {
			log.Printf("Could not get .tractusx metadata for repository: %s", *repo.Name)
		} else {
			content, _ := contents.GetContent()
			log.Printf("Got filecontent: %s", content)
		}
	}

	if err != nil {
		log.Printf("Could not query repositories for GitHub organization: %v", err)
	}

	return []Product{}
}

func getGitHubClient() *github.Client {
	return github.NewClient(getGitHubHttpClient())
}

func getGitHubHttpClient() *http.Client {
	if os.Getenv("GITHUB_ACCESS_TOKEN") == "" {
		return nil
	}

	return oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_ACCESS_TOKEN")},
	))
}
