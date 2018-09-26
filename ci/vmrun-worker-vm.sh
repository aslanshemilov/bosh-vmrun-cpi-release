#!/bin/bash
set -o errexit
set -o pipefail
set -o nounset

cd $(dirname $0)
RELEASE_DIR=../
: ${STATE_DIR:="$PWD/state"}
if ! [ -f $STATE_DIR/env.sh ]; then
  echo "no $STATE_DIR/env.sh file. Create and fill with required fields"
  exit 1
fi

source $STATE_DIR/env.sh
: ${VMRUN_BIN_PATH?"!"}
: ${OVFTOOL_BIN_PATH?"!"}
: ${VDISKMANAGER_BIN_PATH?"!"}
: ${VMRUN_NETWORK:?"!"}
: ${VMRUN_WORKER_IP?"!"}
: ${NETWORK_CIDR:?"!"}
: ${NETWORK_GW:?"!"}
: ${NETWORK_DNS:?"!"}
: ${LINUX_STEMCELL:?"!"}
: ${VMWARE_SERIAL_NUMBER:?"!"}
: ${VMRUN_WORKER_RELEASE_URL:?"!"}

if [ -n ${RESET:-""} ]; then
  RECREATE_VM="y"
  RECREATE_RELEASE="y"
fi

bosh_cli_linux_url="https://s3.amazonaws.com/bosh-cli-artifacts/bosh-cli-5.1.1-linux-amd64"
bosh_cli_darwin_url="https://s3.amazonaws.com/bosh-cli-artifacts/bosh-cli-5.1.1-darwin-amd64"
bosh_bin="bin/bosh-$OSTYPE"
if ! [ -f $bosh_bin ]; then
  curl -L $bosh_cli_linux_url > bin/bosh-linux-gnu
  curl -L $bosh_cli_darwin_url > bin/bosh-darwin17
  chmod +x bin/bosh*
fi

if ! [ -f $LINUX_STEMCELL ]; then
	echo "missing stemcell: $LINUX_STEMCELL"
	exit 1
fi

if [ -n ${RECREATE_RELEASE:-""} -o ! -f $STATE_DIR/cpi.tgz ] ; then
  echo "-----> `date`: Create dev release"
  HOME=$STATE_DIR/bosh_home \
    $bosh_bin create-release --sha2 --force --dir $RELEASE_DIR --tarball $STATE_DIR/cpi.tgz
fi

echo "-----> `date`: Deploy Start"

vm_store_path=$STATE_DIR/vm-store-path
if ! [ -d $vm_store_path ]; then
  mkdir -p $vm_store_path
fi

cpi_url=file://$STATE_DIR/cpi.tgz
cpi_sha1=$(shasum -a1 < $LINUX_STEMCELL | awk '{print $1}')
stemcell_sha1=$(shasum -a1 < $LINUX_STEMCELL | awk '{print $1}')
vmrun_worker_release_sha1=$(curl $VMRUN_WORKER_RELEASE_URL | shasum -a1 | awk '{print $1}')

HOME=$STATE_DIR/bosh_home \
$bosh_bin ${BOSH_COMMAND:-"create-env"} vmrun-worker-vm.yml \
  --vars-store ./$STATE_DIR/vmrun-worker-vm-creds.yml \
  --state ./$STATE_DIR/vmrun-worker-vm-state.json \
  -v cpi_url=$cpi_url \
  -v cpi_url=$cpi_sha1 \
  -v internal_ip="$VMRUN_WORKER_IP"  \
  -v internal_cidr="$NETWORK_CIDR" \
  -v internal_gw="$NETWORK_GW" \
  -v network_name="$VMRUN_NETWORK" \
  -v stemcell_url=file://$LINUX_STEMCELL \
  -v stemcell_sha1=$stemcell_sha1 \
  -v vmrun_bin_path="$VMRUN_BIN_PATH" \
  -v ovftool_bin_path="$OVFTOOL_BIN_PATH" \
  -v vdiskmanager_bin_path="$VDISKMANAGER_BIN_PATH" \
  -v vm_store_path="$vm_store_path" \
  -v vmrun_worker_release_url="$VMRUN_WORKER_RELEASE_URL" \
  -v vmrun_worker_release_sha1="$vmrun_worker_release_sha1" \
  -v vmware_serial_number="$VMWARE_SERIAL_NUMBER" \
  ${RECREATE_VM:+"--recreate"} \
  ;

echo "-----> `date`: Deploy Done"
