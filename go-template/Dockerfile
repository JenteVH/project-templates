FROM golang:1.25-alpine

RUN apk add --no-cache \
    neovim \
    git \
    curl \
    wget \
    build-base \
    ripgrep \
    fd \
    tree \
    bash \
    nodejs \
    npm \
    gcompat \
    libstdc++

# Install air for hot reloading during development
RUN go install github.com/air-verse/air@latest

# Install gopls (Go language server)
RUN go install golang.org/x/tools/gopls@latest

# Install Go formatting tools
RUN go install mvdan.cc/gofumpt@latest
RUN go install golang.org/x/tools/cmd/goimports@latest

# Install Delve (Go debugger)
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Set working directory
WORKDIR /app

# Copy go mod files (if they exist) and download dependencies
COPY go.* ./
RUN go mod download || true

COPY . .

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
