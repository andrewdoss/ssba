# Networking Overview

## Pre-class exercise

### The `pcap` file format

Reconstruct an image from a packet capture.

The packet capture has been dumped using `libpcap` with the following format:
```ascii
              +------------------------------+
              |        Magic number          |
              +--------------+---------------+
              |Major version | Minor version |
              +--------------+---------------+
              |      Time zone offset        |
              +------------------------------+
              |     Time stamp accuracy      |
              +------------------------------+
              |       Snapshot length        |
              +------------------------------+
              |   Link-layer header type     |
              +------------------------------+
       The per-file header length is 24 octets.
```

```bash
xxd -l 24 -c 4 < net.cap
```

Which returns this hex dump:
```bash
00000000: d4c3 b2a1  ....
00000004: 0200 0400  ....
00000008: 0000 0000  ....
0000000c: 0000 0000  ....
00000010: ea05 0000  ....
00000014: 0100 0000  .... 
```

### Questions
- What's the magic number?
    - The magic number is `d4c3b2a1`, which indicates that my machine uses the opposite byte ordering, for the pcap-specific aspects of the file, relative to the machine that stored the data.
- What are the major and minor versions?
    - Major version is 2
    - Minor version is 4
- Are the values that ought to be zero in fact zero?
    - Yes, the values that ought to be zero are in fact zero (the time zone offset and time stamp accuracy).
- What is the snapshot length?
    - `0x05ea` or 1514 bytes
- What is the link layer header type?
    - 1, which maps to `LINKTYPE_ETHERNET`, which aligns with our expectation. 

### Per-packet Headers

Next, come "per-packet" headers:
```
              +----------------------------------------------+
              |          Time stamp, seconds value           |
              +----------------------------------------------+
              |Time stamp, microseconds or nanoseconds value |
              +----------------------------------------------+
              |       Length of captured packet data         |
              +----------------------------------------------+
              |   Un-truncated length of the packet data     |
              +----------------------------------------------+
       The per-packet header length is 16 octets.
```

We can look at the first packet:

```bash
xxd -l 40 -c 4 < net.cap | tail -4
```

And get:
```bash
00000018: 4098 d057  @..W
0000001c: 0a1f 0300  ....
00000020: 4e00 0000  N...
00000024: 4e00 0000  N...
```

### Questions
- What is the size of the first packet?
    - 78 bytes
- Was any data truncated?
    - No

This appears to show that the next packet is 0x4e or 4*16 + 14 = 78 bytes long. If, that's correct, I should be able to add 78 to my previous command and get another per-packet header back:

```bash
xxd -l 118 -c 4 < net.cap | tail -4
```

At this point, I wrote a basic program to parse the capture config and the number of packets.

### Parsing the Ethernet Headers
