#!/bin/bash

set -eu


################################################################################
# Constant values
################################################################################

readonly DIST_DIR=${DIST_DIR}
readonly BUILD_DIR_NAME=build
readonly zip_path="${WORKSPACE}/${BUILD_DIR_NAME}/${PROJECT_NAME}-dist.zip"


################################################################################
# Script information
################################################################################

# Readlink recursively
# 
# This can be achieved with `readlink -f` in the GNU command environment,
# but we implement it independently for mac support.
#
# Arguments
#   $1 - target path
#
# Standard Output
#   the absolute real path
function itr_readlink() {
    local target_path=$1

    (
        cd "$(dirname "$target_path")"
        target_path=$(basename "$target_path")

        # Iterate down a (possible) chain of symlinks
        while [ -L "$target_path" ]
        do
            target_path=$(readlink "$target_path")
            cd "$(dirname "$target_path")"
            target_path=$(basename "$target_path")
        done

        echo "$(pwd -P)/$target_path"
    )
}

# The current directory when this script started.
ORIGINAL_PWD=$(pwd)
readonly ORIGINAL_PWD
# The path of this script file
SCRIPT_PATH=$(itr_readlink "$0")
readonly SCRIPT_PATH
# The directory path of this script file
SCRIPT_DIR=$(cd "$(dirname "$SCRIPT_PATH")"; pwd)
readonly SCRIPT_DIR
# The path of this script file
SCRIPT_NAME=$(basename "$SCRIPT_PATH")
readonly SCRIPT_NAME


################################################################################
# main
################################################################################

cd "$WORKSPACE"

# generate codelabs
"$SCRIPT_DIR/claatw.sh"

# create a directory in which the zip created
mkdir -p "$BUILD_DIR_NAME"

# create list of files to add the zip
cd "$DIST_DIR"
targets=$(find . -mindepth 1 -maxdepth 1 -type f,d -print0 | xargs -0 echo)
readonly targets

# create zip file
rm -f "$zip_path"
zip -r "$zip_path" $targets
