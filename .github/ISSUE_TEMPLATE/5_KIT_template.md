---
name: Propose KIT
labels: kit
---

# Overview

## Description
<!-- Please provide a short description -->

## Impact
<!-- Please name the impact -->
- Impact 1

## KIT Team

### Contributor
<!-- names are already needed for open planning -->
- Contributor 1
- Contributor 2

### Committer
<!-- names are already needed for open planning -->
- Committer 1
- Committer 2

## Architectural Relevance
<!--
The Architecture Management Committee monitors and controls the overarching architecture. It is essential that all applications and documentations follows a baseline set of standards and guidelines. These small checks ensure that the proposed change does not compromise our general principles.
-->
The following items are ensured (answer: yes) after this issue is implemented.

In the context of the standards [126](https://catenax-ev.github.io/docs/standards/CX-0126-IndustryCorePartType) and [127](https://catenax-ev.github.io/docs/standards/CX-0127-IndustryCorePartInstance), typically only one is applicable, depending on the specific use case. Please cross out one of the two standards that does not apply.

- [ ] This feature aligns with our current architectural guidelines
  - **Data Sovereignty:** All data sharing activities across company boundaries follow the [Catena-X Regulatory Framework](https://catenax-ev.github.io/docs/next/regulatory-framework/governance-framework), in particular the [Data Exchange Governance](https://catenax-ev.github.io/docs/next/regulatory-framework/20000ft/data-exchange-governance), and the [Dataspace Protocol](https://docs.internationaldataspaces.org/dataspace-protocol/overview/readme) via a compliant Connector (like the [tractusx-edc](https://github.com/eclipse-tractusx/tractusx-edc) or similar, see [Connector KIT](https://eclipse-tractusx.github.io/docs-kits/next/category/connector-kit))
    - [ ] [CX-0010 Business Partner Number](https://catenax-ev.github.io/docs/next/standards/CX-0010-BusinessPartnerNumber)
    - [ ] [CX-0013 Identity of Member Companies](https://catenax-ev.github.io/docs/next/standards/CX-0013-IdentityOfMemberCompanies)
    - [ ] [CX-0018 Data Space Connectivity (EDC)](https://catenax-ev.github.io/docs/next/standards/CX-0018-DataspaceConnectivity)
    - [ ] [CX-0049 DID Document Schema](https://catenax-ev.github.io/docs/next/standards/CX-0049-DIDDocumentSchema)
    - [ ] [CX-0050 Framework Agreement Credential](https://catenax-ev.github.io/docs/next/standards/CX-0050-FrameworkAgreementCredential)
    - [ ] [CX-0149 Verified Company Identity](https://catenax-ev.github.io/docs/next/standards/CX-0149-Dataspaceidentityandidentification)
  - **Interoperability:** Digital Twins are used (compliant to the [Digital Twin KIT](https://eclipse-tractusx.github.io/docs-kits/next/category/digital-twin-kit) and the [Industry Core KIT](https://eclipse-tractusx.github.io/docs-kits/next/category/industry-core-kit))
    - [ ] [CX-0001 EDC Discovery API](https://catenax-ev.github.io/docs/next/standards/CX-0001-EDCDiscoveryAPI)
    - [ ] [CX-0002 Digital Twins in Catena-X](https://catenax-ev.github.io/docs/next/standards/CX-0002-DigitalTwinsInCatenaX)
    - [ ] [CX-0018 Data Space Connectivity (EDC)](https://catenax-ev.github.io/docs/next/standards/CX-0018-DataspaceConnectivity)
    - [ ] [CX-0126 Industry Core: Part Type 2.0.0](https://catenax-ev.github.io/docs/standards/CX-0126-IndustryCorePartType)
    - [ ] [CX-0127 Industry Core: Part Instance 2.0.0](https://catenax-ev.github.io/docs/standards/CX-0127-IndustryCorePartInstance)
  - **Data Format:**
    - [ ] The data model is based on a [published Semantic Model](https://github.com/eclipse-tractusx/sldt-semantic-models)
- [ ] The impact on the overall system architecture has been assessed. The Feature does not require changes to the architecture or any existing standard? Please have a look here on the [overarching architecture](https://eclipse-tractusx.github.io/docs/tutorials/e2e/inform/architecture)
- [ ] Potential risks or conflicts with existing architecture has been assessed

**Justification:** _(Fill this out, if at least one of the checkboxes above cannot be ticked. Contact the Architecture Management Committee to get an approval for the justification)_

## Helpfull links

- [Catena-X standard library](https://catenax-ev.github.io/docs/next/standards/overview)