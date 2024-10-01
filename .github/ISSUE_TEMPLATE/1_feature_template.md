---
name: Propose a Feature
about: propose a new functionality
---

## Description
<!-- 
A clear and concise description of what the desired feature will look like. 
What is the purpose, what´s the expected result. 
Please describe.
-->

## Impact/Dependencies
<!--
If possible, please provide insight in what and how many components and teams might be affected or give a hint, if the feature implementation might result in breaking changes.
If there are dependencies to other features please name it.
-->

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
## Acceptance Criteria
<!--
Add all Acceptence Criteria. These criteria are important for the definition of done
-->
- [ ] Criteria 1
- [ ] Criteria 2
- [ ] Criteria 3

## Linked Subtasks/Product specific tasks
<!--
PLease link all related subtasks which are need to fullfill the requirements for this featue. If you add a URL from a different issue (of a different repository) this will be interpreted by GitHub automaticly
-->
- [ ] Issue 1 to product repository
- [ ] Issue 2 to product repository
- [ ] Issue 3 to product repository

## Additional information 

- [ ] I'm willing to contribute to this feature
- [ ] I provide necessary developer resources
