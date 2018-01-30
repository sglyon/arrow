GO_BUILD=go build
GO_TEST=go test
GOPATH=$(realpath ../../../..)

# this converts rotate instructions from "ro[lr] <reg>" -> "ro[lr] <reg>, 1" for yasm compatibility
PERL_FIXUP_ROTATE=perl -i -pe 's/(ro[rl]\s+\w{2,3})$$/\1, 1/'

C2GOASM=c2goasm -a -f
CC=clang
C_FLAGS=-target x86_64-unknown-none -masm=intel -mno-red-zone -mstackrealign -mllvm -inline-threshold=1000 -fno-asynchronous-unwind-tables \
	-fno-exceptions -fno-rtti -O3 -fno-builtin -ffast-math
ASM_FLAGS_AVX2=-mavx2 -mfma
ASM_FLAGS_SSE3=-msse3
ASM_FLAGS_SSE4=-msse4

GO_SOURCES  := $(shell find . -path ./_lib -prune -o -name '*.go' -not -name '*_test.go')
ALL_SOURCES := $(shell find . -path ./_lib -prune -o -name '*.go' -name '*.s' -not -name '*_test.go')

INTEL_SOURCES := \
	memory/memory_avx2_amd64.s memory/memory_sse4_amd64.s \
	math/float64_avx2_amd64.s math/float64_sse4_amd64.s

.PHONEY: test bench generate

generate: $(INTEL_SOURCES)

bench: $(GO_SOURCES) $(INTEL_SOURCES)
	$(GO_TEST) $(GO_TEST_ARGS) -bench=. -run=- ./...

bench-noasm: $(GO_SOURCES) $(INTEL_SOURCES)
	$(GO_TEST) $(GO_TEST_ARGS) -tags='noasm' -bench=. -run=- ./...

test: $(GO_SOURCES) $(INTEL_SOURCES)
	$(GO_TEST) $(GO_TEST_ARGS) ./...

test-noasm: $(GO_SOURCES) $(INTEL_SOURCES)
	$(GO_TEST) $(GO_TEST_ARGS) -tags='noasm' ./...

_lib/memory_avx2.s: _lib/memory.c
	$(CC) -S $(C_FLAGS) $(ASM_FLAGS_AVX2) $^ -o $@ ; $(PERL_FIXUP_ROTATE) $@

_lib/memory_sse4.s: _lib/memory.c
	$(CC) -S $(C_FLAGS) $(ASM_FLAGS_SSE4) $^ -o $@ ; $(PERL_FIXUP_ROTATE) $@

memory/memory_avx2_amd64.s: _lib/memory_avx2.s
	$(C2GOASM) -a -f $^ $@

memory/memory_sse4_amd64.s: _lib/memory_sse4.s
	$(C2GOASM) -a -f $^ $@

_lib/math/float64_avx2.s: _lib/math/float64.c
	$(CC) -S $(C_FLAGS) $(ASM_FLAGS_AVX2) $^ -o $@ ; $(PERL_FIXUP_ROTATE) $@

_lib/math/float64_sse4.s: _lib/math/float64.c
	$(CC) -S $(C_FLAGS) $(ASM_FLAGS_SSE4) $^ -o $@ ; $(PERL_FIXUP_ROTATE) $@

math/float64_avx2_amd64.s: _lib/math/float64_avx2.s
	$(C2GOASM) -a -f $^ $@

math/float64_sse4_amd64.s: _lib/math/float64_sse4.s
	$(C2GOASM) -a -f $^ $@

