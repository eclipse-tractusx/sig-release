---
name: Release Compliance 23.12
about: Use this template for tracking Catena-X Release 23.12.
title: "[FOSS NAME] Release Compliance 23.12"
labels: compliance
assignees: @kelaja
milestones: 23.12
---

<!---
 * Copyright (c) 2021,2023 Contributors to the Eclipse Foundation
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
--->

<!-- 
Thanks for your contribution! Please fill out this template as good as possible. 
Important: Contributing Guidelines can be found here: https://eclipse-tractusx.github.io/docs/oss/how-to-contribute
Checkout the repository README for process description. 
-->



## Release Compliance 23.12
Source in Catena-X Confluence + Expert Contacts [here](https://confluence.catena-x.net/x/DOZkBQ).


 - [ ] **Gaia-X compliance** confirmed by Catena-X expert.
	
    - Obtain approval from CX Gaia-X expert(s), prior to Gate review.

    - Provide info as early as possible; expect optimization loops based on feedback

    => check [linked PDF](https://confluence.catena-x.net/download/attachments/90498572/SD_Data_for_Onboarding.pdf?version=1&modificationDate=1690457195303&api=v2) summary for requirements

- [ ] **GDPR compliance** must be ensured
  (personal data; Data Protection & Privacy DPP)
  - assessment must be completed via Catena-X [GDPR questionnaire](https://confluence.catena-x.net/download/attachments/90498572/Catena-X%20GDPR%20Declaration%20and%20Requirements_V3.xlsx?version=1&modificationDate=1690457195339&api=v2)
  - either the product/service doesn`t process any relevant data,
  - or - in case the assessment reveals potential issues - a more severe "Data Privacy Impact Assessment" shall be conducted (if applicable: robust mitigation actions must be completed)
  Provide evidence for GDPR compliance prior to Gate review. Previous assessments remain valid, as long as approach to handling of personal data was not changed.

- [ ] **Interoperability Check**

  Acceptance criteria are fulfilled as described in in the following area on confluence: [Interoperability Quality Gate Requirements](https://confluence.catena-x.net/x/DkwjB)

  - Approval will be documented centrally by Expert(s)

  - Obtain approval from Interoperability experts prior to Gate review.

  Therefore consult the [regular office hours](https://confluence.catena-x.net/x/fzkAAQ) as early as possible; expect optimization loops based on feedback. 

  **Contact:** SYSTEM ARCHITECT2

- [ ] **Data Sovereignty Check**

  Acceptance criteria are fulfilled as described in in the following area on confluence:

   - [Data Sovereignty Guardrails for Release 3.2](https://confluence.catena-x.net/x/qPTeB)

  - Approval will be documented centrally by Expert(s)

  - Obtain approval from Data Sovereignty experts prior to Gate review: Experts will schedule a review with your team.

  Expect optimization loops based on feedback.

  **Contact:** SYSTEM ARCHITECT3, SYSTEM ARCHITECT4

- [ ] **Confirmation of _published_ CX Standards**
  
  Product Owner confirms:

  - The FOSS product under review is compliant to all relevant Catena-X Standards which are already published on the [Association website](https://catena-x.net/de/standard-library).
  > [!NOTE]  
  > Failure to do so may ultimately make the reference implementation useless, as the mandatory CX Certification cannot be accomplished 

- [ ] **Verification of _foreseen_ CX Standards**
  
  Product Owner confirms:

  - The content of upcoming CX Standardization Requests (STAN) has been reviewed and feedback was provided to the responsible PCA/BDA. As per current status, the FOSS product under review will be compliant as soon as the CX Standard is published. ([Info here](https://confluence.catena-x.net/x/XtyAAQ))

  - Minimum: PCA/BDA have finalized all release-relevant Standardization Requests (STAN). Those have passed the CX Association`s technical committee (TC4S) and are pre-published on the website. Opt-out period in progress.

   - Best case: Association Board approval available and CX Standard is published on website.
