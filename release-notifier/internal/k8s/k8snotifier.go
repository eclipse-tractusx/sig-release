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
	"fmt"
	"github.com/Masterminds/semver/v3"
	"github.com/gocolly/colly"
	"log"
	"os"
	"release-notifier/internal/mail"
	"strings"
	"text/template"
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
	latestRelease := getLatestRel()
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
