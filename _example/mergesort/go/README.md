# Merge sort in Go

## BenchHub usage

```bash
bh serve
bh ping
bh run
```

## Test

- test expects the input got sorted in place ... which is a bit strange for merge sort

## Implementations

- `mergeOnlySort` creates tiny slices on the fly and creates a fresh new copy and copy back to source
