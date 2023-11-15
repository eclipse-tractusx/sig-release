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

package psql

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"text/template"
	"release-notifier/internal/mail"
	"github.com/gocolly/colly"
)

const mailTemplate = "templates/mail-psql.html.tmpl"

// Semantic Versioning schema regex
const regexPattern = `^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`

func GetLatestRel() string {
	var release string
	website := "https://bitnami.com/stack/postgresql"

	log.Println("Quering", website)
	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Can't load the page: ", err)
	})

	c.OnHTML("div.stack__header__properties__property", func(e *colly.HTMLElement) {
		if match, _ := regexp.MatchString(regexPattern, e.ChildText("p")); match {
			release = e.ChildText("p")
		}
	})

	c.Visit(website)
	return release
}

func GetPrevRelFromArtifact(artifact string) string {
	data, err := os.ReadFile(artifact)

	if err != nil {
		return ""
	}

	release := string(data)
	return release
}

func SaveLatestRel(release string) {
	err := os.WriteFile("psql_release", []byte(release), 0644)

	if err != nil {
		log.Fatalln(err)
	}
}

func Notify(newRelease string, alignedRelease string) {
	var buff bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	buff.Write([]byte(fmt.Sprintf("Subject: Action Required: PostgreSQL New Release (%s) and Update Process\n%s\n\n", newRelease, mimeHeaders)))

	t, _ := template.ParseFiles(mailTemplate)
	t.Execute(&buff, struct {
		NewPSQLRelease string
		AlignedPSQLRelease string
	}{
		NewPSQLRelease: newRelease,
		AlignedPSQLRelease: alignedRelease,
	})

	mail.SendMail(buff.Bytes())
}