#!/bin/bash
#
# Script used to generate concatenated dependencies from TRG dashboard & release notifier.
# Ensure you have latest Eclipse Dash License tool:
# https://repo.eclipse.org/service/local/artifact/maven/redirect?r=dash-licenses&g=org.eclipse.dash&a=org.eclipse.dash.licenses&v=LATEST
# Export location of the dash license tool to a environment variable $DASH_LICENSE_TOOL:
# i.e: $ export DASH_LICENSE_TOOL="/Users/xyz/Downloads/org.eclipse.dash.licenses-1.0.3-20231116.065040-178.jar"

cat release-automation/go.sum release-notifier/go.sum > dash-input.sum && java -jar $DASH_LICENSE_TOOL dash-input.sum -summary DEPENDENCIES
rm dash-input.sum
