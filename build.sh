set -euxo pipefail
GOBIN=$(pwd)/functions go install ./..