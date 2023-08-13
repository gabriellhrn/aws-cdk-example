ARG GO_VERSION=1.21
ARG AWS_CDK_VERSION=2.91.0

FROM golang:${GO_VERSION}-alpine

RUN apk -v --no-cache --update add \
        git \
        nodejs \
        npm \
    && git config --global init.defaultBranch main \
    && git config --global user.email "fulano@example.com" \
    && git config --global user.name "Fulano de Tal" \
    && npm install -g aws-cdk@${AWS_CDK_VERSION}

WORKDIR "/app"

ENTRYPOINT ["cdk"]
CMD ["--version"]
