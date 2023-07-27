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
	"os"
	"strings"

	"github.com/eclipse-tractusx/tractusx-quality-checks/pkg/product"
	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
)

const gitHubOrg = "eclipse-tractusx"

var gitHubClient *github.Client

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

type Repository struct {
	Name, URL string
}

// QueryProducts loops through all public repositories found in our GitHub organization.
// Using the `.tractusx` metadata file, products and there related repositories are gathered and returned as []Product.
// As second return value, a slice of repository URLs is returned, with all the repositories, that either do not have
// the metadata file present, have an un-parseable one or one, that does not contain
// product information, are returned as second return value.
func QueryProducts() ([]Product, []Repository) {
	productMap := make(map[string]Product)
	var unhandled []Repository

	repos := getOrgRepos()

	for _, repo := range repos {
		metadata := getMetadataForRepo(repo)

		if metadata == nil {
			unhandled = append(unhandled, Repository{Name: repo.Name, URL: repo.Url})
		} else {
			// Only take leading repo metadata into account, since it should be the complete one
			if metadata.LeadingRepository == repo.Url {
				productMap[metadata.ProductName] = productFromMetadata(metadata)
			}
		}
	}
	var products []Product
	for _, p := range productMap {
		products = append(products, p)
	}
	return products, unhandled
}

func getOrgRepos() []product.Repository {
	repos, _, err := gitHubClient.Repositories.ListByOrg(context.Background(), gitHubOrg, &github.RepositoryListByOrgOptions{Type: "public"})
	if err != nil {
		log.Printf("Could not query repositories for GitHub organization: %v", err)
	}

	var result []product.Repository
	for _, r := range repos {
		result = append(result, product.Repository{Name: *r.Name, Url: *r.HTMLURL})
	}
	return result
}

func getMetadataForRepo(repo product.Repository) *product.Metadata {
	log.Printf("Getting tractusx metadata for repository: %s", repo.Name)
	contents, _, _, err := gitHubClient.Repositories.GetContents(context.Background(), gitHubOrg, repo.Name, ".tractusx", nil)
	if err != nil {
		log.Printf("Could not get .tractusx metadata for repository: %s", repo.Name)
		return nil
	}

	content, _ := contents.GetContent()
	metadata, err := product.MetadataFromFile([]byte(content))
	if err != nil {
		log.Printf("Could not parse .tractusx metadata for repository: %s", repo.Name)
		return nil
	}
	return metadata
}

func productFromMetadata(metadata *product.Metadata) Product {
	p := Product{
		Name:         metadata.ProductName,
		LeadingRepo:  metadata.LeadingRepository,
		Repositories: repoURLs(metadata.Repositories),
	}

	if !containsLeadingRepo(p.Repositories, p.LeadingRepo) {
		p.Repositories = append(p.Repositories, p.LeadingRepo)
	}

	return p
}

func containsLeadingRepo(repos []string, leadingRepo string) bool {
	for _, r := range repos {
		if strings.ToLower(r) == strings.ToLower(leadingRepo) {
			return true
		}
	}
	return false
}

func repoURLs(repositories []product.Repository) []string {
	var urls []string
	for _, r := range repositories {
		urls = append(urls, r.Url)
	}
	return urls
}

func init() {
	if os.Getenv("GITHUB_ACCESS_TOKEN") == "" {
		gitHubClient = github.NewClient(nil)
	}

	httpClient := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_ACCESS_TOKEN")},
	))
	gitHubClient = github.NewClient(httpClient)
}
