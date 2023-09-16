FROM golang:1.21-bullseye

# [Optional] Uncomment this section to install additional OS packages.
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends jq

RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install golang.org/x/tools/gopls@latest
RUN go install honnef.co/go/tools/cmd/staticcheck@latest

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1
