---
name: FOSS Documentation Release 24.03
about: Use this template to track all Documentation-relevant topics for your component with regards to the upcoming Milestone.
title: "[FOSS NAME] Release Documentation"
labels: documentation, foss
assignees: kelaja
---



<!-- 
Thanks for your contribution! Please fill out this template as good as possible. 
Important: Contributing Guidelines can be found here: https://eclipse-tractusx.github.io/docs/oss/how-to-contribute
Checkout the repository README for process description. 
-->

# Release Documentation 24.03

Source in Catena-X Confluence and Expert Contacts [here](https://confluence.catena-x.net/x/DOZkBQ) (Source only accessible for Catena-X Consortia members in current transition phase).


- [ ] **Source Code**
- [ ] **Architecture Documents**
- [ ] **Administrator`s Guide** (User assistance)
- [ ] **End-User Manual** (User assistance)
- [ ] **Interfaces Documentation**
- [ ] **UX consistency** Style Guideline for User Interfaces


### Source Code

  - Code is centrally managed in Eclipse Tractus-X repository.
  - All active and relevant [Release Guidelines](https://eclipse-tractusx.github.io/docs/release) are fulfilled (with confirmation by DevSecOps SystemTeam).

  _Artefact Repo:_
    - [Eclipse Tractus-X](https://github.com/eclipse-tractusx)
    - FOSS [Infopage](https://confluence.catena-x.net/x/_AZHAw)(Source only accessible for Catena-X Consortia members in current transition phase)
  
  _Best Practice_
    - Consult the [regular office hours](https://catenax-ng.github.io/) as early as possible; expect optimization loops based on feedback.

  _Contact:_
    - DevSecOps System Team
  
### Architecture Documents

  - Arc42 documentation completed for relevant product version.
  - In MarkDown format. Link to document available. Must point to leading repository within Tractus-X. Process the JIRA ticket and schedule an appointment with Expert to obtain approval prior to Gate review.
  - Provide link to documentation as early as possible and _mark changes_ to previous version (if applicable); expect optimization loops based on feedback.
  
  _Artefact Repo:_
    - [GitHub Readme how-to](https://confluence.catena-x.net/x/iVIAAQ) (Source only accessible for Catena-X Consortia members in current transition phase)

  _Contact:_
    - SYSTEM ARCHITECT

### Administrator`s Guide (User assistance)

  - Admin Guide is available with the software at the same time
  - Admin Guide is correct and up to date; english is a must
  - The documentation is of appropriate maturity for any Operating Company with global business practice intentions.
  - In MarkDown format. Link to document available. Must point to leading repository within Tractus-X.
  
  _Best Practice_:

    - Process the issue and schedule an appointment with SYSTEM ARCHITECT to obtain approval prior to Gate review.
    - Provide link to documentation as early as possible and mark changes to previous version (if applicable); expect optimization loops based on feedback.
    - User assistance ensures that for example administrators get all the information they need to accomplish their tasks with the software. Refer to a administration guide, which covers "install/deploy, configure the software" – as appropriate for the type of software and the information needs of the target group(s)
    - UI text and embedded help complete the User assistance.

  _Contact:_
    - SYSTEM ARCHITECT

### End-User Manual (User assistance)
  - End-User Manual is available with the software at the same time
  - End-User Manual is correct and up to date; english is a must
  - The documentation is of appropriate maturity to be handed over from the CX Consortia to any Operations Company with global business practice intentions.

  _Best Practice:_

    - Process the issue and schedule an appointment with SYSTEM ARCHITECT to obtain approval prior to Gate review.
    - Provide link to documentation as early as possible and mark changes to previous version (if applicable); expect optimization loops based on feedback.
    - User assistance ensures that end users and others get all the information they need to accomplish their tasks with the software. Refer to a user guide, which covers "install/deploy, configure, and use the software" – as appropriate for the type of software and the information needs of the target group(s)
    - UI text and embedded help complete the User assistance.

  _Contact:_
    - SYSTEM ARCHITECT

### Interfaces Documentation
  - API documentation contains all relevant interfaces for integration testing and is completed for relevant product version.
  - Link to document available.
  - Interface contract signed by all involved parties.

  _Best Practice:_
  
    - Process the issue and schedule an appointment with SYSTEM ARCHITECT to obtain approval prior to Gate review.
    - Provide link to documentation as early as possible and mark changes to previous version (if applicable); expect optimization loops based on feedback.
    - see [Open APIs](https://www.openapis.org/)

  _Contact:_ 
    - SYSTEM ARCHITECT

### UX consistency Style Guideline for User Interfaces
  
    - Mandatory for FrontEnd modules where the IP is Open Source or owned by Catena-X.
    - User Interfaces are in line with the Catena-X Style Guidelines
      - user interface style review has been executed  (review owner: SYSTEM ARCHITECT5 UX)
      - review feedback (if existing) got incorporated
      - all findings are assessed
      - all findings (high/very high) are fixed or cleaned up (evidence by re-review)
      - approval of the application CX Style conformity is available (given by the review owner)

  _Best Practice:_
    - Obtain approval from Style Guideline Owner, prior to Gate review
    - use issue, include app URL & TestUser, assign to SYSTEM ARCHITECT5, expect review loop
    - [Style Components](https://portal.dev.demo.catena-x.net/_storybook/?path=/story)
    - [LINK](https://confluence.catena-x.net/x/DVIAAQ) to Style Guideline (Source only accessible for Catena-X Consortia members in current transition phase)
    - LINK to FrontEnd validations (will be added asap)
    - Please note, you can use the [official public available CX shared component library](https://www.npmjs.com/package/cx-portal-shared-components?activeTab=readme) (react supported) to easily develop applications that are in-line with the CX style guidelines
