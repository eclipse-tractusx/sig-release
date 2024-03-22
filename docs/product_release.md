# Product Release

Tractus-X products do follow their own release process. Please check the leading product repositories for details.
Independent of the technical details of product release processes, the released artifacts and minimum level of quality ensurance
is the same across all of them.

The release contents for individual products are are planned in the [open refinement and planning](./open-refinement-and-planning.md) sessions.
These sessions define the scope for all dataspace components, but therefore paint a clear picture of enhancements necessary for every single product.
The product teams themselfes are then responsible to define fine-grained task breakdowns as suiting their own needs.

> __HINT:__
> If product enhancements contribute to an overarching dataspace feature, please link the dataspace feature defined in `sig-release` in your product issue.
> These links can be done in either the description or any comment on your issue. Use `eclipse-tractusx/sig-release#issue-number`

Once planning activities for a certain iteration are completed, the implementation phase is starting and product teams can coordinate, where necessary
through our [communication channels](https://eclipse-tractusx.github.io/docs/oss/getting-started#communication-channels).

Integration tests are performed by maintaining an umbrella Chart, that deploys all relevant dataspace components to spin up an artificial dataspace based on Tractus-X components.
There is a dedicated repository [eclipse-tractusx/tractus-x-umbrella](https://github.com/eclipse-tractusx/tractus-x-umbrella), that is maintained collectively to reflect the latest product release versions.

Additional to the integration- and unit-test suites the teams maintain, every product also needs to adhere to the [Tractus-X release guidelines (TRGs)](https://eclipse-tractusx.github.io/docs/release).
Compliance to these guidelines __should__ be verified on every PR, but __must__ be complied with, as soon as a new product version is planned to be included in the Tractus-X release.

> __HINT:__
> Take special care about legal requirements defined in TRG 7. They __must__ be adhered to on every PR.

Product releases are done by creating a GitHub release. Every release is following the [semantic versioning](https://semver.org/) pattern and marked in the git history via tag.
If suitable for the product teams development process, release candidate versions can help to provide early previews of a product for integration testing.

## Artifacts

- GitHub release (Changelog + Sourcecode)
- Container images
- Helm Chart

## Versioning

- Semantic Versioning

## Quality Assurance

- Defined in TRGs
- Responsibility of every contributor
- (Additonally verified in quality gate process. TBD, if kept after consortia)
