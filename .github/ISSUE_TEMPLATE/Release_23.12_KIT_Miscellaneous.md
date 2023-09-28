name: Release MISCELLANEOUS 23.12
about: Use this template to track miscellaneous topics for your KIT with regards to the upcoming Milestone.
title: "[KIT NAME] Release Miscellaneous"
labels: miscellaneous
assignees: kelaja
---
<!-- 
Thanks for your contribution! Please fill out this template as good as possible. 
Important: Contributing Guidelines can be found here: https://eclipse-tractusx.github.io/docs/oss/how-to-contribute
Checkout the repository README for process description. 
-->

# Release 23.12 Miscellaneous

- [ ] **1 Week prior to QG Identify Version under Review**
  - Pull request for KIT version under review must be placed min one week prior to your reserved Quality-Gate timeslot (to enable formal reviews by relevant BOs, BDAs, DevSecOps System Team) 
  - Example checkpoints: 
    - Aligned basic navigation structure 
    - Core naming schema for critical page like the Changelog 
    - Lincs / references between Eclipse Tractus-X vs. Catena-X 
    - Aligned FOSS critera 
    - ... 

  _Artifact Repository:_
    - Eclipse Tractus-X repo of KIT 

  _Contacts:_
    - Responsible: PO 
    - Review by: 
      - T-X Committer / Project Lead 
      - DevSecOps System Team
     
- [ ] **Approval for all Solution Acceptance Criteria**
  - Product Manager / Business Owner confirms: 
    - All acceptance criteria met (as defined in PI documentation). Relevant evidence is available. 

  _Best Practice:_
    - PO / PCA to obtain BO / BDM OR PM / SA approval prior to Gate review. 

  _Artifact Repository:_
    - [JIRA solution Level](https://jira.catena-x.net/projects/CXSOLUTION/summary)

  _Contacts:_ 
    - Responsible: PO/PCA 
    - Consult: KIT-CONTACT1, KIT-CONTACT2 
    - Review by: BO / BDM (from relevant Domain) OR PM / SA
  
- [ ] **Test and Approval of Developer Journey**
  - PREFERABLE
  - Business Owner / Product Owner should conduct a case study with at least 1 provider (ideally outside) of the consortium to test the developer journey. 

  _Best Practice:_
    - PO / BO to obtain feedback from case studies prior to Gate review. 

  _Artifact Repository:_
    - Case Study 

  _Contacts:_ 
    - Responsible: PO/BDA 
    - Consult: KIT-CONTACT1, KIT-CONTACT2 
    - Review by: Case Study Team

- [ ] **Provide change log page in KIT**

  - Each KIT team must provide a KIT-specific release note/change log for developers
  - +++ Note: Add non-software explanation +++ 

  _Best Practice:_
    - http://changelog.md/ 
    - For good practice visit: [keepachangelog](https://keepachangelog.com/en/1.0.0/)

  _Artifact Repository:_
    - Provide change log page in KIT - [example](https://eclipse-tractusx.github.io/docs/kits/Traceability%20Kit/Traceability%20Kit%20Changelog)

  _Contacts:_ 
    - Responsible: PO/BDA 
    - Review by: Release Management



# Release_23.12_KIT_Adoption_View.md

- [ ] **KIT Repo in Tractus-X**

  - A KIT must have an own repository in Eclipse Tractus-X. The naming of the repository must have a postfix "-kit" 
  - e.g., "traceability-kit" 

  _Artifact Repository:_
    - [eclipse-tractusx](https://github.com/eclipse-tractusx)

  _Contacts:_ 
    - Responsible: PCA/BDA 
    - Consult: KIT-CONTACT1, KIT-CONTACT2 
    - Review by: Tractus-X Project Leads: KIT-CONTACT2, SYSTEM_TEAM3, SYSTEM_TEAM2

- [ ] **Vision and Mission Documents in Repository**

  - Documentation must be available in Eclipse Tractus-X as part of the adoption view MD files. 

  _Best Practice:_
    - [eclipse tractusx-docs-artefacts](https://eclipse-tractusx.github.io/docs/artefacts/)
    - [How-to-create-a-KIT-website](https://github.com/eclipse-tractusx/eclipse-tractusx.github.io/wiki/How-to-create-a-KIT-website%3F)

  _Artifact Repository:_
    - [eclipse-tractusx](https://github.com/eclipse-tractusx) 

  _Contacts:_ 
    - Responsible: PO/BO (from relevant Domain) 
    - Consult: KIT-CONTACT1, KIT-CONTACT2 
    - Review by: PCA/BDA

- [ ] **Business Value Documents in Repository**

  - Documentation must be available in Eclipse Tractus-X as part of the adoption view MD files. 

  _Best Practice:_
    - [eclipse tractusx-docs-artefacts](https://eclipse-tractusx.github.io/docs/artefacts/) 
    - [How-to-create-a-KIT-website](https://github.com/eclipse-tractusx/eclipse-tractusx.github.io/wiki/How-to-create-a-KIT-website%3F) 

  _Artifact Repository:_
    - [eclipse-tractusx](https://github.com/eclipse-tractusx) 

  _Contacts:_ 
    - Responsible: PO/BO (from relevant Domain) 
    - Consult: KIT-CONTACT1, KIT-CONTACT2 
    - Review by: PCA/BDA

- [ ] **Semantic Model standards submitted and published**

  - KIT Product Owner confirms: 
    - The KIT under review is compliant to all relevant Semantic Model Standards which are already published on the Association website. 
    - The content of upcoming Semantic Model Standardization Requests (STAN) has been reviewed and feedback was provided to the responsible PCA/BDA. 

  - Note: 
     - Please create and submit standards in time according to process and timeline > The [standardization process](https://confluence.catena-x.net/display/CXSTAN/The+standardization+process) (Source only accessible for Catena-X Consortia members in current transition phase.)

     - An appendix must be created with a list of certification criteria (e.g., test cases) for the conformity assessment process. 

  _Best Practice:_
    - A KIT must link to the published semantic model standards on the Catena-X Association website.
    - https://catena-x.net/de/standard-library 
    - [eclipse tractusx-docs-artefacts](https://eclipse-tractusx.github.io/docs/artefacts/) 
    - [How-to-create-a-KIT-website](https://github.com/eclipse-tractusx/eclipse-tractusx.github.io/wiki/How-to-create-a-KIT-website%3F) 

    - Note: 
      - failure to do so may ultimately make the KIT useless, as the mandatory CX Certification cannot be accomplished 

  _Artifact Repository:_
    - [eclipse-tractusx](https://github.com/eclipse-tractusx) 

  _Contacts:_ 
    - Responsible: PO/BO (from relevant Domain) 
    - Responsible for Process: Semantic Team 
    - Consult: KIT-CONTACT1, KIT-CONTACT2 
    - Review by: PCA/BDA


- [ ] **Logic / Schema standards submitted and published**

  - KIT Product Owner confirms: 
    - The KIT under review is compliant to all relevant Logic / Schema Standards which are already published on the Association website. 
    - The content of upcoming Logic / Schema Standardization Requests (STAN) has been reviewed and feedback was provided to the responsible PCA/BDA. 

  - Note: 
    - Please create and submit standards in time according to process and timeline > The [standardization process](https://confluence.catena-x.net/display/CXSTAN/The+standardization+process) (Source only accessible for Catena-X Consortia members in current transition phase.)
    - An appendix must be created with a list of certification criteria (e.g., test cases) for the conformity assessment process. 

  _Best Practice:_
    - A KIT must link to the published logic/schema standards on the CX Website.   
    - https://catena-x.net/de/standard-library
    - [eclipse tractusx-docs-artefacts](https://eclipse-tractusx.github.io/docs/artefacts/)
    - [How-to-create-a-KIT-website](https://github.com/eclipse-tractusx/eclipse-tractusx.github.io/wiki/How-to-create-a-KIT-website%3F) 
    - Minimum: PCA/BDA have finalized all release-relevant Standardization Requests (STAN). Those have passed the CX Association`s technical committee (TC4S) and are pre-published on the website. Opt-out period in progress. 
    - Best case: Association Board approval is available and CX Standard is published on website. 

  _Contacts:_ 
    - Responsible: PO/BO (from relevant Domain) 
    - Responsible for Process: PCA/BDA 
    - Consult: KIT-CONTACT1, KIT-CONTACT2 
    - Review by: PCA/BDA


- [ ] **Business Process Documentation**

  - Documentation must be available in Eclipse Tractus-X as part of the adoption view MD files. 

  _Best Practice:_
    - [eclipse tractusx-docs-artefacts](https://eclipse-tractusx.github.io/docs/artefacts/) 
    - [How-to-create-a-KIT-website](https://github.com/eclipse-tractusx/eclipse-tractusx.github.io/wiki/How-to-create-a-KIT-website%3F) 

  _Contacts:_ 
    - Responsible: PO/BO (from relevant Domain) 
    - Consult: KIT-CONTACT1, KIT-CONTACT2 
    - Review by: PCA/BDA


- [ ] **Architecture Documentation**

  - Documentation must be available in Eclipse Tractus-X as part of the adoption view MD files. 

  _Best Practice:_
    - [eclipse tractusx-docs-artefacts](https://eclipse-tractusx.github.io/docs/artefacts/) 
    - [How-to-create-a-KIT-website](https://github.com/eclipse-tractusx/eclipse-tractusx.github.io/wiki/How-to-create-a-KIT-website%3F) 

  _Contacts:_ 
    - Responsible: PO/BO (from relevant Domain) 
    - Responsible for Process: EAs 
    - Consult: KIT-CONTACT1, KIT-CONTACT2 
    - Review by: PCA/BDA

- [ ] **Access and Usage Policies Documentation**

  - Mandatory if a use case needs access & usage policies. 

  - Documentation must be available in Eclipse Tractus-X as part of the adoption view MD files. 

  - Policies must contain: 
    - a human-readable text (aligned with legal) 
    - a technical policy that can be executed by the EDC 

  _Best Practice:_
    - [eclipse tractusx-docs-artefacts](https://eclipse-tractusx.github.io/docs/artefacts/) 
    - [How-to-create-a-KIT-website](https://github.com/eclipse-tractusx/eclipse-tractusx.github.io/wiki/How-to-create-a-KIT-website%3F) 
    - https://github.com/eclipse-tractusx/edc-legal-policies 

  _Artifact Repository:_
    - [eclipse-tractusx](https://github.com/eclipse-tractusx) 

  _Contacts:_ 
    - Responsible: PO/BO (from relevant Domain) 
    - Support for Data Sovereignty: SYSTEM ARCHITECT3
    - Consult: KIT-CONTACT1, KIT-CONTACT2 
    - Review by: PO / BO (from relevant Domain)


- [ ] **Tutorials and Documenation**

  - PREFERABLE
  - Documentation must be available in Eclipse Tractus-X as part of the adoption view MD files. 

  _Best Practice:_
    - [eclipse tractusx-docs-artefacts](https://eclipse-tractusx.github.io/docs/artefacts/) 
    - [How-to-create-a-KIT-website](https://github.com/eclipse-tractusx/eclipse-tractusx.github.io/wiki/How-to-create-a-KIT-website%3F) 


  _Contacts:_ 
    - Responsible: PO/BO (from relevant Domain) 
    - Consult: KIT-CONTACT1, KIT-CONTACT2 
    - Review by: PCA/BDA


- [ ] **White Paper**

  - PREFERABLE
  - Documentation must be available in Eclipse Tractus-X as part of the adoption view MD files. 

  _Best Practice:_
    - [eclipse tractusx-docs-artefacts](https://eclipse-tractusx.github.io/docs/artefacts/) 
    - [How-to-create-a-KIT-website](https://github.com/eclipse-tractusx/eclipse-tractusx.github.io/wiki/How-to-create-a-KIT-website%3F) 

  _Contacts:_ 
    - Responsible: PO/BO (from relevant Domain) 
    - Consult: KIT-CONTACT1, KIT-CONTACT2 
    - Review by: PCA/BDA

# Release_23.12_KIT_Development_View.md

- [ ] **API Specification / Protocols**

  - Documentation for the development view must be available in Eclipse Tractus-X as part of the development view MD files. 

  _Best Practice:_
    - [eclipse tractusx-docs-artefacts](https://eclipse-tractusx.github.io/docs/artefacts/) 
    - [How-to-create-a-KIT-website](https://github.com/eclipse-tractusx/eclipse-tractusx.github.io/wiki/How-to-create-a-KIT-website%3F) 


  _Artifact Repository:_
    - [eclipse-tractusx](https://github.com/eclipse-tractusx) 

  _Contacts:_ 
    - Responsible: PO/BO (from relevant Domain) 
    - Responsible for Process: PCA/BDA 
    - Consult: KIT-CONTACT1, KIT-CONTACT2 
    - Review by: PCA/BDA


- [ ] **Sample Data Available in Repository**

  - Sample data must be available in Eclipse Tractus-X as part of the development view MD files. 

  _Best Practice:_
    - [eclipse tractusx-docs-artefacts](https://eclipse-tractusx.github.io/docs/artefacts/) 
    - [How-to-create-a-KIT-website](https://github.com/eclipse-tractusx/eclipse-tractusx.github.io/wiki/How-to-create-a-KIT-website%3F) 

  _Contacts:_ 
    - Responsible: PO/BO (from relevant Domain) 
     - Responsible for Process: PO / BO (from relevant Domain) 
    - Review by: PCA/BDA


- [ ] **Reference Implementation**

  - Mandatory, if reference implementation is provided 
  - A reference implementation must follow the FOSS release gates and must be available in the corresponding repository in the Eclipse Tractus-X project. 

  _Best Practice:_
    - https://confluence.catena-x.net/x/IpSmB (Source only accessible for Catena-X Consortia members in current transition phase.)

  _Artifact Repository:_
    - [eclipse-tractusx](https://github.com/eclipse-tractusx) 

  _Contacts:_ 
    - Responsible: PO/BO (from relevant Domain) 
    - Responsible for Process: PO / BO (from relevant Domain) 
    - Consult: SYSTEM_TEAM1
    - Review by: PCA/BDA
 
- [ ] **Development-View Architecture Documents in Repository**

  - Mandatory (for Platform Domain / CX ART) 
  - Documentation must be available in Eclipse Tractus-X as part of the development view MD files.

  _Artifact Repository:_
    - [eclipse-tractusx](https://github.com/eclipse-tractusx) 

  _Contacts:_ 
    - Responsible: PCA / BDA 
    - Responsible for Process: EAs 
    - Review by: PCA/BDA

# Release_23.12_KIT_Operations_View.md

- [ ] **Quick Setup Guide and Deployment Scripts in Repository**

  - Documentation for the operations view must be available in Eclipse Tractus-X as MD files. 
  - A reference implementation must be deployable via versioned and released HelmCharts with ArgoCD and as local deployment. In addition, a quick setup guide must be provided. 
    - Quick Setup Guides 
    - Deployment Scripts (e.g., local, helm charts, terraform) 

  _Best Practice:_
    - [eclipse tractusx-docs-artefacts](https://eclipse-tractusx.github.io/docs/artefacts/) 
    - [How-to-create-a-KIT-website](https://github.com/eclipse-tractusx/eclipse-tractusx.github.io/wiki/How-to-create-a-KIT-website%3F) 


  _Artifact Repository:_
    - [eclipse-tractusx](https://github.com/eclipse-tractusx) 

  _Contacts:_ 
    - Responsible: PO/BO (from relevant Domain) 
    - Responsible for Process: PO / BO (from relevant Domain) 
    - Consult: KIT-CONTACT1, KIT-CONTACT2 + System Team 
    - Review by: PCA/BDA
