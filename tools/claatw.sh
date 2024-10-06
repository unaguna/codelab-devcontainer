#!/bin/bash

set -ue

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

file_list=($@)
if [ ${#file_list[@]} -eq 0 ]; then
    file_list=`find $SRC_DIR -name '*.md'`
fi

claat export -o "$tmp_dir" ${file_list[@]}

find "$tmp_dir" -name index.html | xargs patch_dist.sh
if [ -f "${INDEX_SRC_PATH:-/}" ]; then
    # run by `go run` because the shebang of make_index.go is not effective in GitHub Actions
    DIST_DIR="$tmp_dir" go run $(which make_index.go) "$DIST_DIR" "$tmp_dir"
    echo $(date --iso=seconds) 'The index page is generated'
fi

cp -r "$tmp_dir/." "$DIST_DIR"
