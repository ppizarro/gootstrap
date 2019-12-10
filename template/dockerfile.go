package template

const Dockerfile = `# ---------------------------------------------------------------------
#  The first stage container, for image dev
# ---------------------------------------------------------------------
#FROM golang:{{.GoVersion}}-stretch AS dev
FROM golang@{{.GoDigest}} AS dev

ENV GOLANG_CI_LINT_VERSION=v{{.CILintVersion}}

RUN cd /usr && \
    wget -O - -q https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s ${GOLANG_CI_LINT_VERSION}

ARG USER_ID
ARG GROUP_ID

RUN groupadd -f -g ${GROUP_ID} devuser && \
    useradd -m -g ${GROUP_ID} -u ${USER_ID} devuser || echo "user already exists"

USER ${USER_ID}:${GROUP_ID}

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

# ---------------------------------------------------------------------
#  The second stage container, for building the application
# ---------------------------------------------------------------------
FROM dev AS builder

RUN apt-get update && \
    apt-get dist-upgrade -y && \
    apt-get install -y --no-install-recommends ca-certificates tzdata && \
	    update-ca-certificates

COPY . .

ARG LDFLAGS

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="$LDFLAGS" -o /app/{{.Project}} ./cmd/{{.Project}}

# ---------------------------------------------------------------------
#  The third stage container, for running the application
# --------------------------------------------------------------------
FROM scratch

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /app/{{.Project}} /bin/{{.Project}}

# Use an unprivileged user.
USER nobody

ENTRYPOINT ["/bin/{{.Project}}"]
`
