# TRG checks dashboard

The TRG checks dashboard provides a product focussed view of [Tractus-X Release Guidelines](https://eclipse-tractusx.github.io/docs/release) (TRGs) compliance.
Using the `.tractsux` metadata file contained in our repositories, they are grouped to products and checks on repository level
are gathered to provided an overview on product basis.

## Getting started

## Development

To run the go program locally, do as follows:

```bash
# build the program
cd release-automation
go build -o release-automation
PROGRAM_PATH=$(pwd)/release-automation
# executable is created in release-automation/release-automation
# you can now run the program with $PROGRAM_PATH

# Example: run trg checks locally
cd <root-of-your-repo>
$PROGRAM_PATH checkLocal ./
```

> [!Tip]
> Please refer to the [cobra command files](./cmd) to learn more about to running the commands.

### Local Testing

You need to define the GITHUB_ACCESS_TOKEN to create the dashboard as you'll else will run into a 403 RATE LIMIT. Having a the GitHub CLI installed, you can use the following command to set it:

```bash
export GITHUB_ACCESS_TOKEN=$(gh auth token)
```

### Styling

We are using [Materialize](https://materializecss.com) as frontend framework to style this page.
The required CSS and JS is statically included in the `assets` directory.

To upgrade Materialize, see the CDN links provided in the [Getting started guide](https://materializecss.com/getting-started.html) and use them to download the
minified version of materialize. Example of getting v1.0.0:

```shell
curl -o web/assets/css/materialize.min.css https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/css/materialize.min.css
curl -o web/assets/js/materialize.min.js https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/js/materialize.min.js
```

## Exceptions

Please follow [EXCEPTIONS.md](https://github.com/eclipse-tractusx/sig-release/blob/main/release-automation/EXCEPTIONS.md) document to manage exceptions.
