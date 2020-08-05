# 2020-05-21 Storage Requirement

## Background

It's obvious a project like BenchHub will need some novel database in the end.
However, we need BenchHub to develop a good database, so we can focus on writing database instead of drawing bar chart from csv files.
For early stage of BenchHub, aggregated result should fit into a single node RDBMS and has a reasonable performance.
Metrics need to use a TSDB and profiling data can be stored as it is with general purpose compression?

Though if we consider the multi dimension nature of benchmark data and the query pattern, RDBMS may not work well ...
But if the data is small enough ... everything should work well unless json encoded and one file per tuple.

## Reference

- GWP Google-Wide Profiling: A Continuous Profiling Infrastructure
