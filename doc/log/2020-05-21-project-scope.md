# 2020-05-21 Project Scope

## Background

Like most of my projects, the scope changes frequently ... so better keep a log for it.

## Design

BenchHub is designed for comparing the performance of:

- same system, different version of code
- same benchmark, different system

It provides the following (stage 1, still not fully implemented after 2 years XD)

- adapter for different frameworks
- storage for benchmark result aggregated result, metrics and profiling data.

It will provide the following (stage 2)

- managed runtime distributed systems
- tuning parameter to meet a performance and/or budget goal

It will provide the following (stage 3)

- tracing and fault injection (a must for validate correctness of system that can'be formally verified)
- validate correctness of distributed systems (what's the point of being fast if you are wrong `int fast = (int)"slow"`)

It will eventually provide the following (stage 4)

- shared managed runtime offer to open source projects to further reduce cost
- long term storage for benchmark result