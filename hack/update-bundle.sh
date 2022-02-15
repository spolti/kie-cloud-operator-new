#!/bin/bash
# Copyright 2021 Red Hat, Inc. and/or its affiliates
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


set -e

source ./hack/go-mod-env.sh
VERSION=$(go run getversion.go)
CSV_VERSION=$(go run getversion.go -csv)
CSV_PRIOR_VERSION=$(go run getversion.go -csvPrior)

echo "Operator csv version: ${CSV_VERSION} -- Prior version: ${CSV_PRIOR_VERSION}"

# The script will update the containerImage inside CSV and remove the replaces field
#if [[ -z ${1} ]]; then
#    echo "No image given. Please provide the image to use in CSV"
#    exit 1
#fi

#echo "Will update the bundle CSV with ${1} image"

## update the image
#sed -i  "s|quay.io/kiegroup/kogito-operator.*|${1}|g" bundle/app/manifests/kogito-operator.clusterserviceversion.yaml

# update opreator name.version

sed -i "s/businessautomation-operator.v${VERSION}/businessautomation-operator.${CSV_VERSION}/" bundle/manifests/businessautomation-operator.clusterserviceversion.yaml

## remove the replaces field
echo ${CSV_PRIOR_VERSION}
sed -i "s/replaces: businessautomation-operator.0.0.0/replaces: businessautomation-operator.${CSV_PRIOR_VERSION}/" bundle/manifests/businessautomation-operator.clusterserviceversion.yaml

# update selector operated-by
sed -i "s/operated-by: businessautomation-operator.0.0.0/operated-by: businessautomation-operator.${CSV_VERSION}/" bundle/manifests/businessautomation-operator.clusterserviceversion.yaml

# update csv version
sed -i "s/version: ${VERSION}/version: ${CSV_VERSION}/" bundle/manifests/businessautomation-operator.clusterserviceversion.yaml
echo "Bundle CSV updated"