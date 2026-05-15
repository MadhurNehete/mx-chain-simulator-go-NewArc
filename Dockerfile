FROM golang:1.23.6 AS builder


WORKDIR /multiversx
COPY . .

RUN go mod tidy

WORKDIR /multiversx/cmd/chainsimulator

RUN go build -o chainsimulator

RUN mkdir -p /lib_amd64 /lib_arm64

RUN cp $(go list -m -f '{{.Dir}}' github.com/multiversx/mx-chain-vm-v1_4-go)/wasmer/libwasmer_linux_amd64.so /lib_amd64/
RUN cp $(go list -m -f '{{.Dir}}' github.com/multiversx/mx-chain-vm-go)/wasmer2/libvmexeccapi.so /lib_amd64/

RUN cp $(go list -m -f '{{.Dir}}' github.com/multiversx/mx-chain-vm-v1_4-go)/wasmer/libwasmer_linux_arm64_shim.so /lib_arm64/
RUN cp $(go list -m -f '{{.Dir}}' github.com/multiversx/mx-chain-vm-go)/wasmer2/libvmexeccapi_arm.so /lib_arm64/

FROM ubuntu:22.04
ARG TARGETARCH
RUN apt-get update && apt-get install -y git curl socat

COPY --from=builder /multiversx/cmd/chainsimulator /multiversx/chainsimulator

EXPOSE 8085

WORKDIR /multiversx/chainsimulator

# Copy architecture-specific files
COPY --from=builder "/lib_${TARGETARCH}/*" "/lib/"

# Create a non-root user to run the simulator
RUN groupadd -r simulator && useradd -r -g simulator simulator
RUN chown -R simulator:simulator /multiversx
USER simulator

ENTRYPOINT ["sh", "-c", "socat TCP-LISTEN:8085,fork,reuseaddr TCP:127.0.0.1:8085 & ./chainsimulator --rest-api-interface=127.0.0.1 --skip-configs-download"]


