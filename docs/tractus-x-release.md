# Tractus-X Release

The Tractus-X release process aims to ensure compatibility of Tractus-X components and defines
and defines common quality criteria. Releases are initiated on a quarterly basis.
The desired enhancement of the dataspace functionality is planned in the [open refinement and planning](open-refinement-and-planning.md) sessions.
Past release documention can be found in [eclipse-tractusx/tractus-x-release](https://github.com/eclipse-tractusx/tractus-x-release).

To be included in a Tractus-X release, products __must__ opt-in to the process.
This is done by requesting a release review 4 weeks prior to the planned release date, using the "Product Release Check" [issue template](https://github.com/eclipse-tractusx/sig-release/issues/new/choose).

<!--
TODO:
- technical steps to verify compatibility. Who is doing what?
- TRG checks
- e2e?
- CHANGELOG sounding (are dataspace feature mentionend and properly linked)
-->

## Artifacts

- Changelog (overarching features)
- Umbrella Helm Chart(s)
- Overarching architecture documentation (if available)

## Versioning

- Calendar versioning

## Process

<!-- How are products (considered to be) integrated in the Tractus-X release -->
- Prior to release, a Tractus-X umbrella Chart release candidate branch is created
- Producs issue PRs with their release candidate

### Quality Assurance

- User Journey tests (e2e) defined in e2e-testing repo
- User Journey tests automatically executed via `helm test` (nightly + on every product rc PR)
