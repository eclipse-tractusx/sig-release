## Intro

Tractus-X operates on a quarterly release cycle. Each quarter, a new Tractus-X release following CalVer is released. Each Tractus-X release contains multiple products product_release.md.

Once per year, Tractus-X releases a major release, making the remaining three releases minor releases. A major release may contain critical breaking changes that have a major impact on data space participants, such as changes to enablement services.

A minor release contains backward compatible functionality. In addition, patch versions provide backward compatible bug and security fixes.

The overarching Tractus-X releases are coordinated by the release management group (SIG Release). This team collaborates with the currently released versions of Tractus-X products to formulate the overarching release.

As mentioned in [planning.md](planning.md) the Tractus-X project follows an overarching qualification process that all official quarterly Tractus-X releases adhere to. This process applies E2E activities for testing and security for all relevant release participants in compound, which stakeholders can build on. Each product can/shall be released on demand on a higher cadence, see [product_release.md].


## Scope / Content

The Tractus-X release process aims to ensure compatibility of Tractus-X products and defines common quality criteria.

The desired enhancement of the dataspace functionality is planned in the [open refinement and planning](open-refinement-and-planning.md) sessions.

Past release documentation can be found in [eclipse-tractusx/tractus-x-release](https://github.com/eclipse-tractusx/tractus-x-release).

Note: In every quarterly TX release, we shall expect evolutions (new versions) of some products (see [product_release.md]). But also some other products may NOT face a change/up-revision in the given iteration. Nevertheless, those products (“carry-over”) are still essential parts of the release bundle.

Products **must** opt-in to the process in order to be included in a Tractus-X release. (see product release.md)


## Process abstract (future state; or “to-be”)

- Planning
- Feature Work and Product Release
- Quality assurance in parallel
- Opt-In for participation in TX Release
- Completion of Q-Assurance
- TX Release

Once feature work on product level is initiated, the quality assurance of the foreseen release package runs in parallel. The vehicle of Umbrella helm charts is used to continuously test pre-defined product bundles. Automatically and environment-less.

Reference onion rings: (add URL)

Example for cxOS (https://catena-x.net/fileadmin/_online_media_/231006_Whitepaper_Tractus-X.pdf (https://catena-x.net/fileadmin/_online_media_/231006_Whitepaper_Tractus-X.pdf):

A configured umbrella helm chart starts-up all relevant products. No matter if carry-over versions of a previous TX release or evolutions of the current iteration. User Journey tests (e2e) are automatically performed - customized for the respective bundle of products and within given boundary conditions.

Any findings are looped back to the product development teams. Product teams are obliged to fix any defects which prevent successful testing of the bundle.

A bundle is ready for Tractus-X release, once product feature work is completed for all products and e2e-testing was successful without critical findings.
Additionally, manual e2e tests may apply. Environment and execution tbd.


### Operational steps to conclude the TX release

Pre-requisites
!!! If the last release, approved by the Eclipse Foundation, is more than a year old, you have to initiate a new release on the Eclipse Tractus-X Project page at least a week prior to the planned release date (see the Release steps)
•	Helm chart versions and app versions for all products intended for the release
•	An official announcement text highlighting release features, vital for the announcement email
•	An optional official announcement image usable for the Eclipse Tractus-X Project page

1. Verify the presence of all helm chart versions in the release helm repository
1. Verify all referenced links are working (documentation, kits, etc.)
1. Create a Pull Request with the new changelog entry
1.	Gather feedback and proceed with the merge
1.	Create a GitHub release and incorporate the new changelog entry
1.	Insert the new changelog entry under versions on the eclipse-tractusx changelog page
1.	Register a new Tractus-X Release on the Eclipse Tractus-X Project page (Navigate to the Create new release button/link on the right side)
1.	Modify the Download section on the Eclipse Tractus-X Project page via the Downloads, Software Repositories, and Marketplace -> Downloads Message section
1. Compose an email to the tractusx-dev mailing list announcing the new release
      Example:

```
Subject line: [tractusx-dev] Announcing Tractus-X 24.03 - Latest Release Now Available 🎉
Hello Tractus-X Community,

We're excited to announce the latest release of Tractus-X 24.03.

🔗 Release and Helm Charts
You can view the full changelog and Helm chart versions for each component here: Tractus-X 24.03 Release Notes.

🙏 Acknowledgments
We'd like to thank all contributors, users, and community members who have played a role in this release. Your feedback, contributions, and ongoing support fuel the advancement of this project.

💬 We Value Your Feedback
We invite you to test out the new suite release and share your experiences with us. If you encounter any issues or have suggestions for future improvements, please submit them on our GitHub repository.

Thank you for your continued support and enthusiasm for the Tractus-X project. Together, we're driving the future of the ecosystem.

Best regards,

Björn, Daniel, Sigi and Stephan
Eclipse Tractus-X Project Leads
```


## Artifacts

Note: Tractus-X products are released in multiple formats. The number of artifacts and its format varies between individual product and the overarching releases.
- Changelog (overarching features)
- Umbrella Helm Chart(s)
- Overarching architecture documentation (if available)

Tractus-X offers a central Helm chart repository. It supports two channels for released Helm charts - dev and stable.
The dev-channel is used to publish the most recently released charts. It is updated nightly and automatically pulls in the latest chart releases of the eclipse-tractusx GitHub organization.
The stable channel is used by the release management group, to publish all helm charts, that were successfully tested and included in an overarching release. This means, that the stable channel only includes specific versions of product charts, that are tested to the best of our knowledge to work together with other stable charts.


## Versioning

- Calendar versioning


### Quality Assurance

- User Journey tests automatically executed via `helm test` (nightly + on every product rc PR)
- User Journey tests (e2e) defined in e2e-testing repo


Release Management Acceptance Criteria (TBD if this process will continue):
The release participation can be initiated by creating issues for the acceptance criteria check from the Issue templates. Each release the templates for the acceptance criteria may be updated. There are two paths for processing the acceptance criteria for a product to participate in a TX-release.

Guidance on how to use the issue Templates
Select issue type "release_ac" to make the issue visible in the Release Management Kanban Board. Furthermore select the appropriate milestone to distinguish the ticket from other releases.

The Release Happy Path for the Acceptance Criteria
The three process steps to get to the status you need to pass the Q-Gate are shown in the happy path process flow.
Each acceptance criteria issue in GitHub contains a note with the prime contacts so that it is clear who is the assigned expert or release manager.
Happy Path graph

The other Release Path
If the evidence is not sufficient so that the criterium can not be accepted in the quality gate (QG), obligations for the product team will be defined to make a reassessment.
