# codelab 作成環境

この devcontainer は、vscode で codelab を作成するための環境を提供します。

この環境では、保存時に自動で HTML に変換しブラウザに反映するよう各種機能の導入・設定がされています。

## 環境構築の前提条件

- docker を使用可能であること (`docker version` が正常に動作すること)
- vscode がインストールされていること
- vscode に拡張機能 [ms-vscode-remote.vscode-remote-extensionpack](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack) がインストールされていること

## vscode 起動手順

1. vscode でこのディレクトリを開く
2. コマンドパレットを開く
    - (Windows の場合) [Ctrl] + [Shift] + [P] を入力する
3. コマンドパレットで「Dev Conteiners: Reopen in Container」をクリックする
    - 見つからない場合、パレットに「reopen in」のようにコマンドの一部を入力すると絞り込みできる
4. ウィンドウが自動で閉じてすぐにまた自動で開く
    - 初回起動時はウィンドウが閉じるまでに時間がかかる
5. 上で自動で開いたウィンドウが使用可能になる
    - 初回起動時は使用可能になるまで時間がかかる

次回以降は、vscode のアイコンを右クリックして「(ディレクトリ名) [Dev Conttainer]」をクリックするなどでも起動できます。

## ディレクトリ構成

下記のディレクトリ構成で動作するように、自動ビルドなどを設定しています。

```plain
.
+-- .devcontainer
|       : Dev Container + vscode の設定ファイル等
|
+-- .vscode
|       : vscode の設定ファイル等
|
+-- dist
|       : 生成したHTMLファイルの出力先（出力時に自動生成）
|
+-- src
        : codelab の元となる markdown ファイルを置く場所
```

## 作業手順 (執筆時)

1. `./src/` ディレクトリに拡張子が `.md` のファイルを作成・編集する
    - [markdown codelab の書式](https://github.com/googlecodelabs/tools/tree/main/claat/parser/md) に沿って記述し、保存する
    - 保存時、`./dist/` ディレクトリ下に自動で HTML が生成される
2. vscode のウィンドウの右下の「Go Live」をクリックすることで、Live Server を起動する
3. Live Server を起動すると自動でブラウザが起動して Live Server のページを表示するので、HTML が生成されたディレクトリを選ぶことで、出力されたHTMLをブラウザで表示する
4. `.md` ファイルを編集して保存すると、自動でHTMLが再生成される。HTMLが再生成されるとブラウザが自動でリロードして、最新の状態をブラウザで確認できる。

## 作業手順 (完成品の取り出し)

前提として、上記の執筆手順で執筆していて、`./dist/` ディレクトリ内に HTML が生成されていることを想定します。

1. `./dist/` から目的の HTML を `./dist/` の外へコピーする。

    - `./dist/` 内で下記の作業をすると、HTML自動生成が作動した際に作業内容が失われてしまう恐れがあるため

2. WEBサーバに配備せずファイルシステム上でも使えるようにする場合は次の手順を実施する

    1. HTML ファイルに次の置換を実施する：「href="//」⇒「href="http://」

3. Google Analytics が必要である場合を除き、HTMLファイル内の `<google-codelab-analytics>...</google-codelab-analytics>` タグを取り除く

4. 左上の×ボタン (Close ボタン) や右下の Done ボタンのリンク先が `/` になってしまう (2024年7月現在) ため、リンク先を変更する必要があれば下記のコードを各 HTML の `<head>...</head>` 内に記載する

    ```html
    <script>
        window.addEventListener('DOMContentLoaded', (event) => {
            document.getElementById("arrow-back").setAttribute("href", "/url-on-close");
            document.getElementById("done").setAttribute("href", "/url-on-done");
        });
    </script>
    ```

    MEMO: いかにもナンセンスな解決方法だが、リンク先をあらかじめ指定する方法が無いらしい (2024年7月現在)。こちらの issue も参照: <https://github.com/googlecodelabs/tools/issues/535>

## 自動更新の仕組み

HTMLの自動生成とブラウザの自動リロードはそれぞれ下記の拡張機能で実施しています。動作を変更したい場合は該当する設定を編集してください。

### HTML の自動生成

HTML の自動生成は、保存時に自動でコマンドを実行できるようにする拡張機能 [Run on Save](https://marketplace.visualstudio.com/items?itemName=pucelle.run-on-save) で実現しています。この拡張機能を使用することで、`./src/**/*.md` に該当するファイルを保存した際に HTML 生成コマンドである `claat export` が自動で実行されるようにしています。

Run on Save は単に任意のコマンドを保存時に実行するだけの拡張機能です。そのため、HTML生成のパラメータは `claat export` 実行時の引数として指定します。

### ブラウザの自動更新

ブラウザの自動更新には拡張機能 [Live Server](https://marketplace.visualstudio.com/items?itemName=ritwickdey.LiveServer) を使用しています。この拡張機能を使用すると、静的にファイルを提供するWEBサーバが立ち上がり、それをブラウザで開くことで HTML ファイルを閲覧できます。Live Server が提供する HTML ファイルを開いているとき、vscode でファイルが更新されるとブラウザが自動で更新されます。
