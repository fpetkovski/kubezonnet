package agent

//go:generate sh -c "go run github.com/cilium/ebpf/cmd/bpf2go -cflags '-I'$(go list -m -f '{{.Dir}}' github.com/cilium/ebpf)'/examples/headers' kubezonnet kubezonnet.c"
