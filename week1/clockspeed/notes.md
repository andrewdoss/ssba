## Week 1 Exercise: Exploring your architecture

### 1. Fetch-decode-execute cycle

We need a *model* of CPU operation to reason about the execution of our code. Box's aphorism that "all models are wrong, but some are useful" applies here, as modern CPUs operate with many tricky optimizations and special cases that do not fit into a simple model. In our case, "useful" is building general intuition about how computers execute the code we write as programmers, *not* to claim we can accurately benchmark our code analytically.  

In our case, a useful model is assuming that the CPU executes one instruction per clock-cycle. *fetch-decode*execute* refers to a cycle where the CPU fetches an instruction from memory, decodes the instruction into a CPU operation, and executes the operation. 

We can attempt to crudely test this model by comparing the number of "basic" operations our computer executes per second to the advertised clock speed. *basic* is quoted because what might appear to be a basic operation in a programming language can translate to more (or fewer) assembly instructions, and even then, assembly instructions aren't necessarily one-to-one with CPU operations. 

For my crude approximation, I wrote an "empty" for loop in Go (`clockspeed.testLoop` in `clock_speed.go`) with 1B iterations, and then used the `go test` benchmarking tool to measure the execution time. I measured an average of ~313 ns over 4 trials, which means roughly 3.19B loop iterations per second.

I am running on an M1 Mac, which has an [unofficial](https://www.anandtech.com/show/16252/mac-mini-apple-m1-tested) (Apple does not release clock-speed for their chips) clock-speed of 3.2 GHz, which is astonishingly close to my measurement of 3.19 GHz. However, there are a few caveats/nuances to mention: 
- I actually don't understand how the iteration frequency could get to 3.19 GHz. An empty Go loop is not a single instruction, as it requires at least both an increment and a comparison at the assembly level (checked in ARM64 with the Compiler Explorer at [godbolt.org](https://godbolt.org/)). However, I don't know what the ARM translation looks like here and, again, assembly instructions are not necessarily 1:1 with CPU instructions.
- I also tried running with the Go garbage collecter explicitly disabled and got the same result. This makes sense to me, as the empty for loop program should only use memory on the stack anyways.
- I later explored adding different arithmetic operations inside the loop (e.g. integer addition, multiplication, and division). I stumbled on one interesting example where "special case" optimizations break simplified execution models. I created a local variable `x` initilized to 0, and incremented it to 1B using x += 1 inside the loop. When I looked at the AMD64 instructions, I noticed that the compiler didn't even run the addition inside the loop, instead it appeared to just dual-purpose the existing loop counter/increment as the local variable. I also didn't see an execution speed different with multiplication and division, somehow. 

I then reran the "benchmarks" using Python with ipython `%timeit` and found about 0.1B loop iterations per second, about 30x slower than the Go equivalent. This is interesting for this specific benchmark, but also a fairly contrived comparison and not necessarily representaive of the performance difference in real-world applications. 

### 2. Registers

Arithmetic operations and memory load/store operations occur over register values. Therefore, a 32-bit system can only address 2^32 = 4 Gibibytes or RAM. 

Note: this was useful for configuring Go debugging with VS Code.

### 3. Memory access and caches

M1 RAM latency ~100ns, so ~300 instructions can run in the time required to access memory. 

Caching helps with this, as is demonstrated by looping over an array of arrays of ints in the `cache_demo` directory. Looping row-wise (contiguous memory) is 5x faster than looping column-wise. 