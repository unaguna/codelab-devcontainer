#!/bin/bash

set -u

################################################################################
# Temporally files
################################################################################

# All temporally files which should be deleted on exit
tmpfile_list=( )

function remove_tmpfile {
    set +e
    for tmpfile in "${tmpfile_list[@]}"
    do
        if [ -e "$tmpfile" ]; then
            rm -fR "$tmpfile"
        fi
    done
    set -e
}
trap remove_tmpfile EXIT
trap 'trap - EXIT; remove_tmpfile; exit -1' INT PIPE TERM

# the output of `gradle projects`
tmp_dir=$(mktemp -d)
readonly tmp_dir
tmpfile_list+=( "$tmp_dir" )


################################################################################
# main
################################################################################

claat export -o "$tmp_dir" $@

find "$tmp_dir" -name index.html | xargs patch_dist.sh
DIST_DIR="$tmp_dir" make_index.go

cp -r "$tmp_dir/." "$DIST_DIR"
