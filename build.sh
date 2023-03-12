 #!/bin/bash

CGO_ENABLED=0 go build -o ./build/server ./cmd
podman build -t "ghcr.io/kialakun/forward-proxy" .
podman push "ghcr.io/kialakun/forward-proxy"