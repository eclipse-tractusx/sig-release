# The Tractus-X metadata file

The `tractusx-quality-checks` rely on a metadata file to read product information and configuration for the checks.
The following section describe details like location, filename and the quality check configuration options, you can declare.

## General metadata

- Filename: `.tractusx`
- Location: root of your repository

Every [eclipse-tractusx](https://github.com/eclipse-tractusx/) repository is expected to have a `.tractusx` metadata file
present. General metadata is only mandatory in the [leading product repository](https://eclipse-tractusx.github.io/docs/release/trg-2/trg-2-4) though.

The structure for leading repository metadata is as follows:

```yaml
# Will be used on dashboard etc. to refer to your product; only mandatory in the leading repo
product: "your-product-name"
# mandatory info in every repo
leadingRepository: "https://github.com/eclipse-tractusx/<your-leading-repo>"
# optional section to refer to all of your teams repositories
repositories:
  - name: "your-product-repo"
    usage: "Main documentation and release repo for <product>"
    url: "https://github.com/eclipse-tractusx/<your-leading-repo>"
  - name: "your-product-library"
    usage: "A library supporting <product>, that is released separately"
    url: "https://github.com/eclipse-tractusx/<your-product-lib>"
```

## Repository category

There are release guidelines, that target specific repositories, like the leading one for example.
In order to differentiate the repositories, a `repoCategory` can be set in the metadata file to clarify the role of the repository.
The category is set like shown in the following snippet:

```yaml
repoCategory: "special" # default: product; available options: "special", "support", "product" 
```

## Quality check configuration

The `.tractusx` metadata file does contain configuration options to specify treatment in regards to guideline checks.

### Exclude Dockerfiles from aligned base image check (TRG 4.02)

As defined in [TRG 4.02](https://eclipse-tractusx.github.io/docs/release/trg-4/trg-4-02), all of our Docker images, that
are published to [DockerHub](https://hub.docker.com/u/tractusx/), have to be based on an aligned base image.

To mark a Docker image as non-published and therefor skipping the aligned base image check, you can add the path to the
respective `Dockerfile` in a `skipReleaseChecks.alignedBaseImage` section.

Example:

```yaml
# section to explicitly skip certain release guideline checks
skipReleaseChecks:
  # Skip TRG 4.02 aligned base images; Specify the complete path and filename to the dockerfile that is used to build
  # a non published Docker image. Path is treated relative to your repository root.
  alignedBaseImage:
    - "path/to/non-published/Dockerfile"
    - "path/to/Dockerfile/only/used/in/testing-pipeline/Dockerfile.ui-tests"
```

### Define Deviating Architecture Doc Folder for Built Documentation (TRG 1.05)

[TRG 1.05](https://eclipse-tractusx.github.io/docs/release/trg-1/trg-1-05/) foresees the architecture documentation in `docs/architecture`. If you use e.g. AsciiDoc you might want to specify a deviating directory as you commonly have `src` folder that needs to be compiled / built.

```yaml
# section to explicitly reconfigure certain release guideline checks
configureReleaseChecks:
  # Define a deviation from docs/architecture/ for your arc42
  ArchitectureDocEntryPath: docs/src/architecture/
```

### Exclude Markdown / AsciiDoc Files from Notice Section Check (TRG 7.07)

As defined in [TRG 7.07](https://eclipse-tractusx.github.io/docs/release/trg-7/trg-7-07#how-to-include-legal-notices), all documents need to have a notice section at the end of the file specifying in heading 2 a notice file.

By default, the `.github` directory is excluded from this check, but all other files are validated.

```yaml
# section to explicitly skip certain release guideline checks
skipReleaseChecks:
  # Skip the Notice section for non-distributed artefacts that don't need a notice section
  LegalNoticeNonCode:
    # exclude the directory src/assets
    # this makes sense e.g. if you have a workflow that replaces a user_guide during release / distribution in the frontend that includes a notice section but you place a dummy file there.
    - frontend/src/assets/
```

## OpenAPI Specification

It is possible and highly recommended to list OpenAPI specifications related to your product in the respective section `openApiSpecs` of metadata file as in the example below:

```yaml
openApiSpecs:
- "https://raw.githubusercontent.com/eclipse-tractus/<your-product-repo>/product_version_openapi.yaml"
```

Provided specifications are used by automation to generate and publish Swagger UI documentations.

## Full example

The following snippet shows a complete example of a `.tractusx` metadata file with all its options.

```yaml
# Will be used on dashboard etc. to refer to your product; only mandatory in the leading repo
product: "your-product-name"
# mandatory info in every repo
leadingRepository: "https://github.com/eclipse-tractusx/<your-leading-repo>"
# default: product; available options: "special", "support", "product"
repoCategory: "special"
# optional section to refer to all of your teams repositories
repositories:
  - name: "your-product-repo"
    usage: "Main documentation and release repo for <product>"
    url: "https://github.com/eclipse-tractusx/<your-leading-repo>"
  - name: "your-product-library"
    usage: "A library supporting <product>, that is released separately"
    url: "https://github.com/eclipse-tractusx/<your-product-lib>"
# section to explicitly skip certain release guideline checks
skipReleaseChecks:
  alignedBaseImage:
    - "/path/to/non-published/Dockerfile"
openApiSpecs:
- "https://raw.githubusercontent.com/eclipse-tractus/<your-product-repo>/product_version_openapi.yaml"

```
