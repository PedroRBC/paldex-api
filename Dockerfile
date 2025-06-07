FROM ubuntu:latest
WORKDIR /app

# Necessary for Gpg Sign in Devcontainer
RUN apt update && apt install --no-install-recommends -y \
        git gpg gnupg gpg-agent socat

# Install Go
COPY --from=golang:latest /usr/local/go /usr/local/go
ENV GOROOT=/usr/local/go
ENV PATH=$GOROOT/bin:$PATH

RUN go install golang.org/x/tools/gopls@latest && \
    go install honnef.co/go/tools/cmd/staticcheck && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install golang.org/x/tools/cmd/goimports@latest && \
    go install github.com/haya14busa/goplay/cmd/goplay@latest && \
    go install github.com/josharian/impl@latest && \
    go install github.com/fatih/gomodifytags@latest && \
    go install github.com/cweill/gotests/gotests@latest && \
    go install github.com/godoctor/godoctor@latest


CMD ["tail", "-f", "/dev/null"]