GOOGLE_BENCHMARK_VERSION=0ab2c2906b95f41dfa7b692ca4ab856e20ebd3d4
CATCH2_VERSION=f0e596e2522f64451f6444edab5c59b8b9e8e082

.PHONY: build
# TODO: compile is a bit slow when it comes to catch ... because it is static linked? and how to improve compile speed ...
build:
	cd build && make

build-init:
	mkdir -p build
	cd build && cmake ..

.PHONY: bench test
bench:
	./build/bench_mergesort

test:
	./build/test_mergesort

download-vendor: download-catch2 download-google-benchmark

download-catch2:
	mkdir -p vendor/catchorg
	wget https://github.com/catchorg/Catch2/archive/${CATCH2_VERSION}.zip
	unzip ${CATCH2_VERSION}.zip -d vendor/catchorg
	mv vendor/catchorg/Catch2-${CATCH2_VERSION} vendor/catchorg/Catch2

download-google-benchmark:
	mkdir -p vendor/google
	wget https://github.com/google/benchmark/archive/${GOOGLE_BENCHMARK_VERSION}.zip
	unzip ${GOOGLE_BENCHMARK_VERSION}.zip -d vendor/google
	mv vendor/google/benchmark-${GOOGLE_BENCHMARK_VERSION} vendor/google/benchmark