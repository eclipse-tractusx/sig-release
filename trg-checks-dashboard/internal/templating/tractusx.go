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

package templating

import (
	"context"
	"log"
	"os"
	"path"
	"strings"

	"github.com/eclipse-tractusx/tractusx-quality-checks/pkg/product"
	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
)

const gitHubOrg = "eclipse-tractusx"

var gitHubClient *github.Client

var checks = []func(string) bool{
	func(dir string) bool {
		_, err := os.Stat(path.Join(dir, "README.md"))
		return err != nil
	},
}

func CheckProducts() ([]CheckedProduct, []Repository) {
	repoInfoByRepoUrl := make(map[string]repoInfo)
	var unhandledRepos []Repository

	repos := getOrgRepos()

	for _, repo := range repos {
		metadata := getMetadataForRepo(repo)

		if metadata == nil {
			unhandledRepos = append(unhandledRepos, Repository{Name: repo.Name, URL: repo.Url})
		} else {
			repoInfoByRepoUrl[repo.Url] = repoInfo{metadata: *metadata, repoName: repo.Name, repoUrl: repo.Url}
		}
	}

	var checkedProducts []CheckedProduct
	for _, p := range getProductsFromMetadata(repoInfoByRepoUrl) {
		checkedProduct := CheckedProduct{Name: p.Name, LeadingRepo: p.LeadingRepo, OverallPassed: true}
		for _, repo := range p.Repositories {
			checkedRepo := runQualityChecks(repo)
			checkedProduct.OverallPassed = checkedProduct.OverallPassed && checkedRepo.PassedAllGuidelines
			checkedProduct.CheckedRepositories = append(checkedProduct.CheckedRepositories, checkedRepo)
		}

		checkedProducts = append(checkedProducts, checkedProduct)
	}

	return checkedProducts, unhandledRepos
}

func runQualityChecks(repo Repository) CheckedRepository {
	log.Printf("Starting checks for repo: %s", repo.Name)
	checkedRepo := CheckedRepository{RepoUrl: repo.URL, RepoName: repo.Name, PassedAllGuidelines: true}

	dir, err := cloneRepo(repo)
	if err != nil {
		log.Printf("Could not clone repo %s. Error: %s", repo.URL, err)
		return CheckedRepository{}
	}

	for _, check := range checks {
		log.Println("Running checks...")
		passed := check(dir)
		checkedRepo.PassedAllGuidelines = checkedRepo.PassedAllGuidelines && passed
		checkedRepo.GuidelineChecks = append(checkedRepo.GuidelineChecks, GuidelineCheck{Passed: passed, GuidelineUrl: "https://trg.com", GuidelineName: "TRG 1"})
	}

	return checkedRepo
}

func getProductsFromMetadata(metadataForRepo map[string]repoInfo) []Product {
	log.Println("Forming products from repo metadata")

	leadingRepoToProduct := make(map[string]*Product)
	for url, info := range metadataForRepo {
		log.Printf("Merging metadata for %s", url)
		if _, containsProductForLeadingRepo := leadingRepoToProduct[info.metadata.LeadingRepository]; !containsProductForLeadingRepo {
			log.Printf("No product for leading repo %s yet. Adding empty one", info.metadata.LeadingRepository)
			leadingRepoToProduct[info.metadata.LeadingRepository] = &Product{}
		}

		p := leadingRepoToProduct[info.metadata.LeadingRepository]
		log.Printf("Adding repository %s (URL: %s) to product %s (Name: %s)", info.repoName, info.repoUrl, p.Name, info.metadata.LeadingRepository)
		p.Repositories = append(p.Repositories, Repository{Name: info.repoName, URL: info.repoUrl})

		if strings.ToLower(url) == strings.ToLower(info.metadata.LeadingRepository) {
			log.Printf("Repo %s is leading, addign name (%s) + repo URL (%s) to product", url, info.metadata.ProductName, info.metadata.LeadingRepository)
			p.Name = info.metadata.ProductName
			p.LeadingRepo = info.metadata.LeadingRepository
		}
	}

	var products []Product
	for _, p := range leadingRepoToProduct {
		products = append(products, *p)
	}
	return products
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

func init() {
	if os.Getenv("GITHUB_ACCESS_TOKEN") == "" {
		gitHubClient = github.NewClient(nil)
	}

	httpClient := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_ACCESS_TOKEN")},
	))
	gitHubClient = github.NewClient(httpClient)
}
