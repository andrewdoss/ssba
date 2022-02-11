# Virtual memory and file systems

## Pre-class questions

- What is the page size for your operating system's virtual memory system?

```bash
getconf PAGESIZE
```

Output: 4096 bytes, 4 KiB.

- Approximate the size (in bytes) of the page table for your operating system right now, based on the size of programs that are running.

Based on `top`, my VM is currently using 115,800 Kib of memory. If each page is 4 KiB, that results in about 28,950 pages of memory. Each page table entry requires 8 bytes in a 64-bit system (is that really all used?) resulting in about 231,600 bytes or 1.8 MiB for the page table. In contrast, my host OS is using over 10 GiB, which would require > 180 MiB to store page tables. Probably more is required to signal addresses that are not in use?

- Write a program which consumes more and more physical memory at a predictable rate. Use top or a similar program to observe its execution. What happens as your memory utilization approaches total available memory? What happens when it reaches it?

I wrote a loop in C that used `malloc` to allocate memory for a 5M element int array every 1 ms. 

First, the "free" memory decreased until it go near zero, then the "buff/cache" memory decreased and finally when buff/cache was also exhausted, `malloc` began returning 0 values instead of memory addresses. I did not see any memory getting swapped to disk. 

- Suppose the designers of your operating systems propose quadrupling the page size. What would be the trade-offs?

Advantages of a larger page size:
- Fewer entries required in a page table, less memory.
- Fewer cache misses when looking up entries in page table (because it is smaller).
- Better locality when reading sequentially from a larger page.

Disadvantages of a larger page size:
- More memory fragmentation/waste (a program cannot use partial pages).