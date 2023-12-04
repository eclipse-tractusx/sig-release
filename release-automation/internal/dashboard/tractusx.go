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
	"context"
	"log"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
	"tractusx-release-automation/internal/container"
	"tractusx-release-automation/internal/docs"
	"tractusx-release-automation/internal/helm"
	"tractusx-release-automation/internal/repo"
	"tractusx-release-automation/internal/tractusx"
)

const gitHubOrg = "eclipse-tractusx"

var gitHubClient *github.Client

func CheckProducts() ([]CheckedProduct, []Repository, []Repository) {
	repoInfoByRepoUrl := make(map[string]repoInfo)
	var unhandledRepos, archivedRepos []Repository

	repos := getOrgRepos()

	for _, repo := range repos {
		metadata := getMetadataForRepo(repo)

		if repo.Archived {
			archivedRepos = append(archivedRepos, repo)
		} else if metadata == nil {
			unhandledRepos = append(unhandledRepos, repo)
		} else {
			repoInfoByRepoUrl[repo.URL] = repoInfo{metadata: *metadata, repoName: repo.Name, repoUrl: repo.URL}
		}
	}

	var checkedProducts []CheckedProduct
	for _, p := range getProductsFromMetadata(repoInfoByRepoUrl) {
		checkedProduct := CheckedProduct{Name: p.Name, LeadingRepo: p.LeadingRepo, RepoCategory: p.RepoCategory, OverallPassed: true}
		for _, r := range p.Repositories {
			dir := repoClone(r)
			defer os.RemoveAll(dir)
			if strings.EqualFold(r.URL, p.LeadingRepo) {
				checkedProduct.ChartsDetails = getChartsDetails(dir)
			}
			checkedRepo := runQualityChecks(r, dir)
			checkedProduct.OverallPassed = checkedProduct.OverallPassed && checkedRepo.PassedAllGuidelines
			checkedProduct.CheckedRepositories = append(checkedProduct.CheckedRepositories, checkedRepo)
		}

		checkedProducts = append(checkedProducts, checkedProduct)
	}

	return checkedProducts, unhandledRepos, archivedRepos
}

func getChartsDetails(repoDir string) []ChartDetails {
	log.Println("Collecting charts info.")
	var chartsDetails []ChartDetails

	chartDir := path.Join(repoDir, "charts")
	helmCharts, err := os.ReadDir(chartDir)
	if err != nil || len(helmCharts) == 0 {
		return chartsDetails
	}

	for _, hc := range helmCharts {
		if !helm.IsChartDirectory(path.Join(chartDir, hc.Name())) {
			continue
		}
		chartValues := helm.ChartYamlFromFile(path.Join(chartDir, hc.Name(), "Chart.yaml"))
		chartDetails := ChartDetails{ChartName: chartValues.Name, ChartVersion: chartValues.Version, AppVersion: chartValues.AppVersion}
		chartsDetails = append(chartsDetails, chartDetails)
	}

	return chartsDetails
}

func repoClone(repo Repository) string {
	dir, err := cloneRepo(repo)
	if err != nil {
		log.Printf("Could not clone repo %s. Error: %s", repo.URL, err)
		return ""
	}
	return dir
}

func runQualityChecks(repo Repository, dir string) CheckedRepository {
	log.Printf("Starting checks for repo: %s", repo.Name)
	checkedRepo := CheckedRepository{RepoUrl: repo.URL, RepoName: repo.Name, PassedAllGuidelines: true}

	if _, err := os.Stat(dir); dir == "" || err != nil {
		return CheckedRepository{}
	}

	for _, check := range initializeChecksForDirectory(dir) {
		testResult := check.Test()
		checkedRepo.PassedAllGuidelines = checkedRepo.PassedAllGuidelines && (testResult.Passed || check.IsOptional())

		guidelineCheck := GuidelineCheck{
			Passed:           testResult.Passed,
			Optional:         check.IsOptional(),
			ErrorDescription: testResult.ErrorDescription,
			GuidelineUrl:     check.ExternalDescription(),
			GuidelineName:    check.Name(),
		}
		checkedRepo.GuidelineChecks = append(checkedRepo.GuidelineChecks, guidelineCheck)
	}

	return checkedRepo
}

func initializeChecksForDirectory(dir string) []tractusx.QualityGuideline {
	var checks []tractusx.QualityGuideline
	tractusx.ErrorOutputFormat = tractusx.WebErrOutputFormat

	checks = append(checks, docs.NewReadmeExists(dir))
	checks = append(checks, docs.NewInstallExists(dir))
	checks = append(checks, docs.NewChangelogExists(dir))
	//checks = append(checks, repo.NewRepoStructureExists(dir))
	checks = append(checks, repo.NewLeadingRepositoryDefined(dir))
	checks = append(checks, container.NewAllowedBaseImage(dir))
	checks = append(checks, helm.NewHelmStructureExists(dir))

	return checks
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

		if strings.EqualFold(url, info.metadata.LeadingRepository) {
			log.Printf("Repo %s is leading, addign name (%s) + repo URL (%s) to product", url, info.metadata.ProductName, info.metadata.LeadingRepository)
			p.Name = info.metadata.ProductName
			p.LeadingRepo = info.metadata.LeadingRepository
			p.RepoCategory = info.metadata.RepoCategory
		}
	}

	var products []Product
	for _, p := range leadingRepoToProduct {
		products = append(products, *p)
	}
	sort.Slice(products, func(i, j int) bool {
		return strings.ToLower(products[i].Name) < strings.ToLower(products[j].Name)
	})
	return products
}

type listFunc[T any] func(ctx context.Context, options *github.ListOptions) ([]T, *github.Response, error)

func paginate[T any](ctx context.Context, listFunc listFunc[T], listOps *github.ListOptions) ([]T, error) {
	var allItems []T

	for {
		items, resp, err := listFunc(ctx, listOps)
		if err != nil {
			return allItems, err
		}

		allItems = append(allItems, items...)

		if resp.NextPage == 0 {
			break
		}

		listOps.Page = resp.NextPage
	}

	return allItems, nil
}

func listOrgRepos(ctx context.Context, listOps *github.ListOptions) ([]*github.Repository, *github.Response, error) {
	repos, response, err := gitHubClient.Repositories.ListByOrg(ctx, gitHubOrg, &github.RepositoryListByOrgOptions{
		Type:        "public",
		ListOptions: *listOps})
	return repos, response, err
}

func getOrgRepos() []Repository {
	repos, err := paginate(context.Background(), listOrgRepos, &github.ListOptions{
		Page:    0,
		PerPage: 100,
	})

	if err != nil {
		log.Printf("Could not query repositories for GitHub organization: %v", err)
	}

	var result []Repository
	for _, r := range repos {
		result = append(result, Repository{Name: *r.Name, URL: *r.HTMLURL, Archived: *r.Archived})
	}
	return result
}

func getMetadataForRepo(repo Repository) *tractusx.Metadata {
	log.Printf("Getting tractusx metadata for repository: %s", repo.Name)
	contents, _, _, err := gitHubClient.Repositories.GetContents(context.Background(), gitHubOrg, repo.Name, ".tractusx", nil)
	if err != nil {
		log.Printf("Could not get .tractusx metadata for repository %s: %v", repo.Name, err)
		return nil
	}

	content, _ := contents.GetContent()
	metadata, err := tractusx.MetadataFromFile([]byte(content))
	if err != nil {
		log.Printf("Could not parse .tractusx metadata for repository: %s", repo.Name)
		return nil
	}
	return metadata
}

func init() {
	if os.Getenv("GITHUB_ACCESS_TOKEN") == "" {
		gitHubClient = github.NewClient(nil)
	} else {
		httpClient := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: os.Getenv("GITHUB_ACCESS_TOKEN")},
		))
		gitHubClient = github.NewClient(httpClient)
	}
}
