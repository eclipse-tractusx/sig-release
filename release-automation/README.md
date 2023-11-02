# TRG checks dashboard

The TRG checks dashboard provides a product focussed view of [Tractus-X Release Guidelines](https://eclipse-tractusx.github.io/docs/release) (TRGs) compliance.
Using the `.tractsux` metadata file contained in our repositories, they are grouped to products and checks on repository level
are gathered to provided an overview on product basis.

## Getting started

## Development

### Styling

We are using [Materialize](https://materializecss.com) as frontend framework to style this page.
The required CSS and JS is statically included in the `assets` directory.

To upgrade Materialize, see the CDN links provided in the [Getting started guide](https://materializecss.com/getting-started.html) and use them to download the
minified version of materialize. Example of getting v1.0.0:

```shell
curl -o web/assets/css/materialize.min.css https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/css/materialize.min.css
curl -o web/assets/js/materialize.min.js https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/js/materialize.min.js
```

To 