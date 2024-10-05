FROM mcr.microsoft.com/devcontainers/go:1-1.22-bookworm

RUN go install github.com/googlecodelabs/tools/claat@latest

RUN apt update \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# volumes によるマウント領域では正常に動作しない処理をするためのスペース。
# 必要に応じて、/workspace_local 内のファイルへのシンボリックリンクを /workspace に作成する
RUN mkdir /workspace_local && chown vscode:vscode /workspace_local \
    && mkdir /workspace_local/dist && chown vscode:vscode /workspace_local/dist

COPY bashrc_suffix /home/vscode/bashrc_suffix
RUN cat /home/vscode/bashrc_suffix >> /home/vscode/.bashrc
RUN rm /home/vscode/bashrc_suffix
