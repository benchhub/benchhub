# Overview

NOTE: this is WIP

The overall workflow is there is an adapter for different benchmark frameworks. Adapter calls BenchHub before running the benchmark to allocate resources.
After the benchmark is finished, adapter upload the result to BenchHub.

In its simplest form where the benchmark runs locally, the only resource get allocated is an id for the benchmark e.g. [gobench](gobench.md).