if [ -f "${CSS_SRC_PATH:-/}" ]; then
    cp ${CSS_SRC_PATH} ${DIST_DIR}
elif [ -n "${CSS_SRC_PATH}" ]; then
    echo -n '' > ${DIST_DIR}/$(basename ${CSS_SRC_PATH})
fi
