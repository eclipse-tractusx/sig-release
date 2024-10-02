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
	"fmt"
	"net/smtp"
	"os"
)

const smtpServer = "smtp.gmail.com"
const smtpPort = "587"
const recipientMailEnv = "TRACTUSX_MAILINGLIST"
const senderMailEnv = "DEVSECOPS_NOTIFICATION_EMAIL"
const senderPassEnv = "DEVSECOPS_NOTIFICATION_EMAIL_PASSWORD"

func SendMail(subject string, body []byte) error {
	sender := os.Getenv(senderMailEnv)
	password := os.Getenv(senderPassEnv)
	recipient := os.Getenv(recipientMailEnv)

	msg := []byte(fmt.Sprintf("To: %s\nSubject: %s\nMIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n\n\n", recipient, subject))
	msg = append(msg, body...)
	auth := smtp.PlainAuth("", sender, password, smtpServer)
	return smtp.SendMail(smtpServer+":"+smtpPort, auth, sender, []string{recipient}, msg)
}
