# In memory store

The in memory store is mainly for testing.
Each 'tale' is a slice of decoded proto message and there is extra index for common query,
e.g. `map[string]int` as secondary index for a string column to avoid scan all the protos.