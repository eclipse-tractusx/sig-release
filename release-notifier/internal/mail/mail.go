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

package mail

import (
	"bytes"
	"fmt"
	"net/smtp"
	"os"
	"text/template"
)

const smtpServer = "smtp.gmail.com"
const smtpPort = "587"
const mailTemplate = "templates/mail.html.tmpl"
const recipentMail = "TRACTUSX_MAILINGLIST"
const senderMailEnv = "DEVSECOPS_NOTIFICATION_EMAIL"
const senderPassEnv = "DEVSECOPS_NOTIFICATION_EMAIL_PASSWORD"

func SendPSQLRelNotification(release string) {
	var buff bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	buff.Write([]byte(fmt.Sprintf("Subject: Action Required: %s Update your Helm Chart dependencies.\n%s\n\n", release, mimeHeaders)))

	t, _ := template.ParseFiles(mailTemplate)
	t.Execute(&buff, struct {
		PSQLRelease string
	}{
		PSQLRelease: release,
	})

	sender := os.Getenv(senderMailEnv)
	password := os.Getenv(senderPassEnv)
	recipents := []string{os.Getenv(recipentMail)}
	sendMail(sender, password, recipents, buff.Bytes())
}

func sendMail(sender string, password string, recipents []string, body []byte) {
	auth := smtp.PlainAuth("", sender, password, smtpServer)
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, sender, recipents, body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Email Sent!")
}
