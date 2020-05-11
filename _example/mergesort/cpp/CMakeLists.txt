cmake_minimum_required(VERSION 3.15 FATAL_ERROR)

project(mergesort LANGUAGES CXX)
set(CMAKE_CXX_STANDARD 17)

set(BENCHMARK_ENABLE_TESTING OFF CACHE BOOL "Suppressing benchmark's tests" FORCE)
add_subdirectory(vendor/google/benchmark)

add_subdirectory(vendor/catchorg/Catch2)

add_executable(test_mergesort test/test_mergesort.cpp)
target_link_libraries(test_mergesort Catch2::Catch2)

add_executable(bench_mergesort bench/bench_mergesort.cpp)
target_link_libraries(bench_mergesort benchmark)