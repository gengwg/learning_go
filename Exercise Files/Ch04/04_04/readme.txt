This directory is intentionally empty.

memory is manged at run time.
memory allocated and deallocated automatically.
deallocated by garbage collector.
objects out of scope or set to nil are eligible.

new() allocates but not initialize memory.
results in zeroed storage but returns a memory address.

new() allocates and initialize memory.
allocates non-zeroed storage and returns a memory address.


m := make(map[string]int)
m["key"] = 42
fmt.Println(m)

