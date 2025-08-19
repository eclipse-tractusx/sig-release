#!/bin/bash
#
# Copyright (c) 2023 Contributors to the Eclipse Foundation
#
# See the NOTICE file(s) distributed with this work for additional
# information regarding copyright ownership.
#
# This program and the accompanying materials are made available under the
# terms of the Apache License, Version 2.0 which is available at
# https://www.apache.org/licenses/LICENSE-2.0.
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
# WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
# License for the specific language governing permissions and limitations
# under the License.
#
# SPDX-License-Identifier: Apache-2.0
#

#
# Script used to generate concatenated dependencies from TRG dashboard & release notifier.
# Ensure you have latest Eclipse Dash License tool:
# https://repo.eclipse.org/service/local/artifact/maven/redirect?r=dash-licenses&g=org.eclipse.dash&a=org.eclipse.dash.licenses&v=LATEST
# Export location of the dash license tool to a environment variable $DASH_LICENSE_TOOL:
# i.e: $ export DASH_LICENSE_TOOL="/Users/xyz/Downloads/org.eclipse.dash.licenses-1.0.3-20231116.065040-178.jar"

cat release-automation/go.sum release-notifier/go.sum > dash-input.sum && java -jar $DASH_LICENSE_TOOL dash-input.sum -summary DEPENDENCIES
rm dash-input.sum
