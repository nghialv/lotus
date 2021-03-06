#!/usr/bin/env bash

# Copyright (c) 2018 Lotus Load
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.

# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.

set -o errexit
set -o nounset
set -o pipefail

OPTION=${1:-""}
SUFFIX=""
if [ ! -z "$OPTION" ]; then
  SUFFIX="-$OPTION"
fi

ROOT=$(dirname ${BASH_SOURCE})/..
MANIFESTS_DIR="${ROOT}/install/manifests${SUFFIX}"
VALUE_FILE="${ROOT}/install/manifest-generate-values${SUFFIX}.yaml"
HELM_CHART_DIR="${ROOT}/install/helm"

echo "Generating manifests to tmp..."
helm template --name lotus -f $VALUE_FILE $HELM_CHART_DIR --output-dir /tmp

echo "Deleting all old manifests..."
mkdir -p ${MANIFESTS_DIR}
rm -rf ${MANIFESTS_DIR}/*

echo "Copying generated manifests to manifests folder..."
cp /tmp/lotus/templates/* ${MANIFESTS_DIR}
for f in $(find /tmp/lotus/charts/grafana/templates -type f); do
  cp $f ${MANIFESTS_DIR}/grafana-${f##*/};
done

echo "Deleting tmp data..."
rm -rf /tmp/lotus

echo "Done"

