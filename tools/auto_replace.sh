#!/bin/bash

# 監視対象のディレクトリ
DIR_TO_WATCH=/workspace_local/dist

# 左上の×ボタン (Close ボタン) のリンク先
URL_CLOSE='/url-on-close'
# 右下の Done ボタンのリンク先
URL_DONE='/url-on-done'

# 追加する <script> タグ
script_tag=`tr -d '\n' <<EOF
<script>
    window.addEventListener('DOMContentLoaded', (event) => {
        document.getElementById("arrow-back").setAttribute("href", "${URL_CLOSE}");
        document.getElementById("done").setAttribute("href", "${URL_DONE}");
    });
</script>
EOF
`

# <style> タグ内に追加するスタイル指定
style_part=`tr -d '\n' <<EOF
google-codelab #drawer .metadata {
  display: none;
}
EOF
`

# ディレクトリ内のすべてのファイルの変更を監視し、変更があった場合に文字列置換を実行
inotifywait -m -e modify --format '%w%f' -r "$DIR_TO_WATCH" | while read -r FILE; do
    if [[ "$FILE" == *.html ]]; then
        # プロトコルのないリンクを修正

        sed -i \
            -e 's|href="//|href="http://|g' \
            -e "s/google-codelab-analytics/!--/g" \
            -e "s|^</head>|${script_tag}</head>|" \
            -e "s| </style>|${style_part}</style>|" \
            "$FILE"
        echo "$(date --iso=seconds) Replaced in $FILE"
    fi
done
