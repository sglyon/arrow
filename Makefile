GO_ARGS=-tags '$(GO_TAGS)'
GO_BUILD=go build $(GO_ARGS)
GOPATH=$(realpath ../../../..)

C2GOASM=c2goasm -a -f
CC=clang
C_FLAGS=-target x86_64-unknown-none -masm=intel -mno-red-zone -mstackrealign -mllvm -inline-threshold=1000 -fno-asynchronous-unwind-tables \
	-fno-exceptions -fno-rtti -O3 -fno-builtin
ASM_FLAGS_AVX2=-mavx2 -mfma
ASM_FLAGS_SSE3=-msse3

GO_SOURCES := $(shell find . -path ./_lib -prune -o -name '*.go' -not -name '*_test.go')
ALL_SOURCES := $(shell find . -path ./_lib -prune -o -name '*.go' -name '*.s' -not -name '*_test.go')

INTEL_SOURCES := memory/memory_avx2_amd64.s memory/memory_sse3_amd64.s

.PHONEY: test bench generate

generate: $(INTEL_SOURCES)

bench: $(GO_SOURCES) $(INTEL_SOURCES)
	go test -bench=. -run=- ./...

test: $(GO_SOURCES) $(INTEL_SOURCES)
	go test ./...

_lib/memory_avx2.s: _lib/memory.c
	$(CC) -S $(C_FLAGS) $(ASM_FLAGS_AVX2) $^ -o $@

_lib/memory_sse3.s: _lib/memory.c
	$(CC) -S $(C_FLAGS) $(ASM_FLAGS_SSE3) $^ -o $@

memory/memory_avx2_amd64.s: _lib/memory_avx2.s
	$(C2GOASM) -a -f $^ $@

memory/memory_sse3_amd64.s: _lib/memory_sse3.s
	$(C2GOASM) -a -f $^ $@

