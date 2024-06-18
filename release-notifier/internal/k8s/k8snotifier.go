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

package k8s

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"release-notifier/internal/mail"
	"strings"
	"text/template"

	"github.com/Masterminds/semver/v3"
	"github.com/gocolly/colly"
)

const mailTemplate = "templates/mail-k8s.html.tmpl"
const artifactName = "k8s_release"

func getLatestRel() string {
	release, _ := semver.NewVersion("0.0.0")
	website := "https://kubernetes.io/releases/"

	log.Println("Querying", website)
	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Can't load the page: ", err)
	})

	c.OnHTML("span.release-inline-value", func(e *colly.HTMLElement) {
		if !strings.Contains(e.Text, "release") {
			return
		}
		r, err := semver.NewVersion(strings.Split(e.Text, " ")[0])
		if err != nil {
			return
		}
		if r.Compare(release) == 1 {
			release = r
		}
	})

	c.Visit(website)
	return release.String()
}

func getRelFromArtifact() string {
	data, err := os.ReadFile(artifactName)

	if err != nil {
		return ""
	}

	release := string(data)
	return release
}

func IsNewRelease() bool {
	latestRelease := getLatestRel2("kubernetes")
	prevRelease := getRelFromArtifact()
	if latestRelease != prevRelease {
		log.Printf("New release is out: %v\n", latestRelease)
		if err := os.WriteFile(artifactName, []byte(latestRelease), 0644); err != nil {
			log.Fatalln(err)
			return false
		}
		return true
	}
	log.Println("No new release :(")
	return false
}

func buildContent(mailTemplate string) ([]byte, error) {
	var buff bytes.Buffer
	t, err := template.ParseFiles(mailTemplate)
	if err != nil {
		return nil, err
	}
	t.Execute(&buff, struct {
		NewK8SRelease     string
		AlignedK8SRelease string
	}{
		NewK8SRelease:     getRelFromArtifact(),
		AlignedK8SRelease: os.Getenv("CURRENT_ALIGNED_K8S_VER"),
	})
	return buff.Bytes(), nil
}

func Notify() error {
	body, err := buildContent(mailTemplate)
	if err != nil {
		return err
	}
	subject := fmt.Sprintf("Action Required: Kubernetes New Release (%s)", getRelFromArtifact())
	return mail.SendMail(subject, body)
}

type Project struct {
	Id string `yaml:"id"`
	Name string `yaml:"name"`
}

type ProjectSet struct {
	Projects []Project
}

type Release struct {
	Version string `yaml:"version"`
}

type ReleasesSet struct {
	Releases []Release
}

func getLatestRel2(product string) string {
    url := "https://api.newreleases.io/v1/projects"

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Fatalln("Error creating request:", err)
    }
	apiKey := os.Getenv("NEWRELEASESIO_APIKEY")
    req.Header.Add("X-Key", apiKey)
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatalln("Error making request:", err)
    }
    defer resp.Body.Close()

	projects := ProjectSet{}
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln("Error reading response body:", err)
    }
	err = json.Unmarshal(body, &projects)
    if err != nil {
        log.Fatalf("Error unmarshaling JSON: %v", err)
    }
	for _, p := range projects.Projects {
		if strings.EqualFold(p.Name, fmt.Sprintf("%s/%s", product, product)) {
			fmt.Println(p)
			url = fmt.Sprintf("%s/%s/releases", url, p.Id)
		}
	}

	fmt.Println(url)
	req, err = http.NewRequest("GET", url, nil)
	req.Header.Add("X-Key", apiKey)
    if err != nil {
        log.Fatalln("Error creating request:", err)
    }
	resp, err = client.Do(req)
    if err != nil {
        log.Fatalln("Error making request:", err)
    }
    defer resp.Body.Close()

	releases := ReleasesSet{}
    body, err = io.ReadAll(resp.Body)
	if err != nil {
        log.Fatalln("Error reading response body:", err)
    }
	err = json.Unmarshal(body, &releases)
    if err != nil {
        log.Fatalf("Error unmarshaling JSON: %v", err)
    }
	fmt.Println(releases)
	os.Exit(0)
	return ""
}