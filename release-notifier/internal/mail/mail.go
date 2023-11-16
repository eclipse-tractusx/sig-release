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
	"log"
	"net/smtp"
	"os"
)

const smtpServer = "smtp.gmail.com"
const smtpPort = "587"
const recipentMailEnv = "TRACTUSX_MAILINGLIST"
const senderMailEnv = "DEVSECOPS_NOTIFICATION_EMAIL"
const senderPassEnv = "DEVSECOPS_NOTIFICATION_EMAIL_PASSWORD"


func SendMail(body []byte) {
	sender := os.Getenv(senderMailEnv)
	password := os.Getenv(senderPassEnv)
	recipent := os.Getenv(recipentMailEnv)

	auth := smtp.PlainAuth("", sender, password, smtpServer)
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, sender, []string{recipent}, body)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println("Email Sent!")
}
