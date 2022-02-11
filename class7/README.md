# Introduction to operating systems

## Setup

I am on an M1 mac, so I will use a Canonical Multipass VM:

```bash
multipass launch --name ssba bionic
multipass exec ssba -- /bin/bash
```

## 1. Programs

How many programs exist on this image?

Can use find from the root, test for executables, then count returned lines.

```bash
find / -executable -type f | wc -l
```

Returns 573.

What types of files are executable on this system?

```
find / -executable -type f | xargs file -b | \
cut -d ',' -f -4 | sort | uniq -c | sort -nr | head
```

[Helpful link for explaining shell commands](explainshell.com explanation).

The `strings` command tries to pull printable ascii strings from a file (which may be a binary file with snippets of ascii, for example).

Executables typically have a few bytes dedicated as a "magic number". For ELF files, this is `0x7F` followed by `ELF` in ASCII. I can confirm this for `/bin/true` using the following command:

```bash
xxd /bin/true | head -1
``` 

## 2. Processes

How many processes are running?

```bash
ps aux | wc -l
```

What states are they in, and what do those states mean?

```bash
ps aux | awk '{ print $8 }' | sort | uniq -c | sort -nr
```

Output:
- 28 I<
- 18 S
- 13 Ss
- 8 I
- 7 Ssl
- 4 S+
- 2 Ss+
- 2 SN
- 1 STAT
- 1 S<s
- 1 R+

And code meanings:

```
PROCESS STATE CODES
       Here are the different values that the s, stat and state output
       specifiers (header "STAT" or "S") will display to describe the
       state of a process:

               D    uninterruptible sleep (usually IO)
               R    running or runnable (on run queue)
               S    interruptible sleep (waiting for an event to
                    complete)
               T    stopped by job control signal
               t    stopped by debugger during the tracing
               W    paging (not valid since the 2.6.xx kernel)
               X    dead (should never be seen)
               Z    defunct ("zombie") process, terminated but not
                    reaped by its parent

       For BSD formats and when the stat keyword is used, additional
       characters may be displayed:

               <    high-priority (not nice to other users)
               N    low-priority (nice to other users)
               L    has pages locked into memory (for real-time and
                    custom IO)
               s    is a session leader
               l    is multi-threaded (using CLONE_THREAD, like NPTL
                    pthreads do)
               +    is in the foreground process group
```

## 3. Virtual Memory

Viewing `pmap` for simple `malloc` loop:

The output shows the map start address, size of the map, mode (permissions) on the map, and mapping (file backing the map or [ anon ] for allocated memory with [ stack ] for the program stack).

```
9985:   ./a.out
0000aaaaaf4c3000      4K r-x-- a.out
0000aaaaaf4d3000      4K r---- a.out
0000aaaaaf4d4000      4K rw--- a.out
0000aaaab6274000    132K rw---   [ anon ]
0000ffff9db38000   1276K r-x-- libc-2.27.so
0000ffff9dc77000     64K ----- libc-2.27.so
0000ffff9dc87000     16K r---- libc-2.27.so
0000ffff9dc8b000      8K rw--- libc-2.27.so
0000ffff9dc8d000     16K rw---   [ anon ]
0000ffff9dc91000    116K r-x-- ld-2.27.so
0000ffff9dcb4000      8K rw---   [ anon ]
0000ffff9dcbc000      4K r----   [ anon ]
0000ffff9dcbd000      4K r-x--   [ anon ]
0000ffff9dcbe000      4K r---- ld-2.27.so
0000ffff9dcbf000      8K rw--- ld-2.27.so
0000ffffce2b0000    132K rw---   [ stack ]
 total             1800K
```

When I run this with `watch pmap ${pid}`, the first [ anon ] entry grows while the rest of the `pmap` remains static.

The ordering here agrees with what is described in CS: APP. From lowest to highest virtual address, we have:
1. The program (`a.out`)
2. The heap for dynamic memory allocation (first `[ anon ]`)
3. Shared libraries (`libc` and `ld`)
4. User stacj (`[ stack ]`)

## 4. Files and Devices

Claim "everything in Unix is a file".

`stat` can be used see information about a file.

`stat /bin/true`

```
  Size: 26768           Blocks: 56         IO Block: 4096   regular file
Device: 801h/2049d      Inode: 110         Links: 1
Access: (0755/-rwxr-xr-x)  Uid: (    0/    root)   Gid: (    0/    root)
Access: 2022-01-21 15:32:36.300000003 -0800
Modify: 2018-01-18 01:43:49.000000000 -0800
Change: 2022-01-04 08:04:35.351985228 -0800
```

Read 100 bytes from `/dev/urandom` to a file in your home directory, twice. Confirm both files are 100 bytes and that they are different files. 

`xxd -l 100 /dev/urandom | xxd -r > ~/dump3`