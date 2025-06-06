FROM golang:1.22
WORKDIR /app

# Necessary for Gpg Sign in Devcontainer
RUN apt update && apt install --no-install-recommends -y \
        git gpg gnupg gpg-agent socat

CMD ["tail", "-f", "/dev/null"]