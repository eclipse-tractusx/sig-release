Product releases are done by creating a GitHub release. Every release is following the [semantic versioning](https://semver.org/) pattern and must be marked in the git history via tag.
If suitable for the product teams development process, release candidate versions can help to provide early previews of a product for integration testing.

Tractus-X products do follow their own release process. They are individually developed and therefore no strict workflow is enforced. To ensure a consistent view on releases, the following aspects should still be met:
A CHANGELOG.md file, that follows keep a changelog recommendations is maintained and updated with descriptions for the current product release. Links to GitHub external artifacts (i.e. mvncentral or Docker Hub) are referenced. GitHub releases are used to publish artifacts. Git tags are added accordingly (usually done automatically). The changelog content of the current release is documented in the GitHub release. Helm charts are released into the aligned helm chart repository.
Please check the leading product repositories for details.
Independent of the technical details of product release processes, the released artifacts and minimum level of quality assurance is the same across all of them.
The release contents for individual products are planned in the [open refinement and planning](./open-refinement-and-planning.md) sessions. These sessions define the scope for all dataspace products, and therefore paint a clear picture of enhancements necessary for every single product. The product teams themselves are then responsible to define fine-grained task breakdowns as suiting their own needs.
> **HINT:**
> If product enhancements contribute to an overarching dataspace feature, please link the dataspace feature defined in `sig-release` in your product issue.
> These links can be done in either the description or any comment on your issue. Use `eclipse-tractusx/sig-release#issue-number`
Once planning activities for a certain iteration are completed, the implementation phase is starting and product teams can coordinate, where necessary through our [communication channels](https://eclipse-tractusx.github.io/docs/oss/getting-started#communication-channels).

Every product needs to adhere to the [Tractus-X release guidelines (TRGs)](https://eclipse-tractusx.github.io/docs/release).
> **HINT:**
> this is in addition to the integration- and unit-test suites the teams maintain
> Take special care about legal requirements defined in TRG 7. They **must** be adhered to on every PR.
Compliance to these guidelines **should** be verified on every PR, but **must** be complied with, as soon as a new product version is planned to be included in the Tractus-X release.

Integration tests are performed by maintaining a Product Umbrella Chart, that deploys all relevant dataspace products to spin up an artificial dataspace based on Tractus-X products.
There is a dedicated repository [eclipse-tractusx/tractus-x-umbrella](https://github.com/eclipse-tractusx/tractus-x-umbrella), that is maintained collectively to reflect the latest product release versions.

Products **must** opt-in to the process in order to be included in a Tractus-X release.
This is done by requesting a release review 4 weeks prior to the planned release date, using the "Product Release Check" [issue template](https://github.com/eclipse-tractusx/sig-release/issues/new/choose).


## Artifacts

- GitHub release (Changelog + Sourcecode)
  Most Tractus-X products are software products. In these cases, packaging the sourcecode together with a list of used dependencies is a good starting point and should be included as artifact of a release in almost any case.

- Container images
  All container images provided by tractus-x are only provided for development, testing etc. without any guarantee on license or security. Feel free to use them on your own risk, all images can be build by yourself through provided Dockerfiles.

- Product Umbrella Helm Chart
  Applications developed in the Tractus-X context typically provide a Helm chart for easy deployment on kubernetes.
  To add a Helm chart as a release artifact it has to be packaged. There are multiple tools, that help packaging charts. We recommend using the chart-releaser-action GitHub action though, since together with activated GitHub pages, it can transform your repository to function as Helm chart repository on its own.

Additionally, Tractus-X offers a central Helm chart repository. It supports two channels for released Helm charts - dev and stable.
The dev-channel is used to publish the most recently released charts. It is updated nightly and automatically pulls in the latest chart releases of the eclipse-tractusx GitHub organization.
The stable channel is used by the release management group, to publish all helm charts, that were successfully tested and included in an overarching release. This means, that the stable channel only includes specific versions of product charts, that are tested to the best of our knowledge to work together with other stable charts.

Patching strategy
tbd

Helm Repository
For information about using the Tractus-X Helm repository, please refer to the charts repository.


## Versioning

- Semantic Versioning


## Quality Assurance

- Defined in TRGs
- Responsibility of every contributor
- Currently still verified in a Quality Gate process. TBD if this process will continue.
