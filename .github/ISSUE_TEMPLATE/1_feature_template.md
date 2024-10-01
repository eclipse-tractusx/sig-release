---
name: Propose a Feature
about: propose a new functionality
---

# Description
<!-- 
A clear and concise description of what the desired feature will look like. 
What is the purpose, what´s the expected result. 
Please describe.
-->

## Feature Team

### Contributor
<!-- names are already needed for open planning -->
- Contributor 1
- Contributor 2

### Committer
<!-- names are already needed for open planning -->
- Committer 1
- Committer 2
<!--  can be delivered after Planning Phase 1 -->

## Explain the topic in 2 sentences
<!-- provide a short overview -->

## What's the benefit?
<!--  do we increase security/efficiency/ reduce complexity? -->

## What are the Risks/Dependencies ?
<!--  breaking change ? do you have interface partner you need to align with ? -->

## Detailed explanation
<!-- explain the idea in detail , what is the current situation , what needs to be changed  -->

### Current implementation

### Proposed improvements

## User Stories
<!--  define the sub-issues that lead to the full developement of this feature -->
- Issue 1, linked to specific repository
- Issue 2, linked to another specific repository

## Acceptance Criteria
<!--
Add all Acceptence Criteria. These criteria are important for the definition of done
-->
- [ ] Criteria 1
- [ ] Criteria 2
- [ ] Criteria 3

## Test cases
<!-- add testcases - proposed structure (Description/Steps/Expected Result) -->

### TestCase 1
<!--  description -->

#### Steps

1. Do something
2. Click something
3. Add something

#### Expected Result

1. Expectation
2. Expectation
3. Expectation

### Architecture Management Committee Check
<!--
The Architecture Management Committee monitors and controls the overarching architecture. It is essential that all applications and documentations follows a baseline set of standards and guidelines. These small checks ensure that the proposed change does not compromise our general principles.
-->
The following items are ensured (answer: yes) after this issue is implemented:

- [ ] **Architecture:** This issue is compliant with the [overarching architecture](https://eclipse-tractusx.github.io/docs/tutorials/e2e/inform/architecture) and its existing standards and protocols
- [ ] **Data Sovereignty:** All data sharing activities across company boundaries follow the [Dataspace Protocol](https://docs.internationaldataspaces.org/dataspace-protocol/overview/readme) via a compliant Connector (like the [tractusx-edc](https://github.com/eclipse-tractusx/tractusx-edc) or similar, see [Connector KIT](https://eclipse-tractusx.github.io/docs-kits/next/category/connector-kit))
- [ ] **Interoperability:** Digital Twins are used (compliant to the [Digital Twin KIT](https://eclipse-tractusx.github.io/docs-kits/next/category/digital-twin-kit) and the [Industry Core KIT](https://eclipse-tractusx.github.io/docs-kits/next/category/industry-core-kit))
- [ ] **Data Format:** The data model is based on a [published Semantic Model](https://github.com/eclipse-tractusx/sldt-semantic-models)

**Justification:** _(Only needs to be filled out if at lease one checkbox item could not be answered with "yes")_

## Additional information
<!-- this is only needed, if contributors and committers are not known during feature creation -->
- [ ] I'm willing to contribute to this feature
- [ ] I provide necessary developer resources

PLease see here an [example feature](https://github.com/eclipse-tractusx/sig-release/issues/882)
