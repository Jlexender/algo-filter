# Bloom Value Storage

Nothing but an implementation of a bloom filter for A&DS course.

## How to use

* `./bvs [new|load] [file]` to create a new bloom filter or load an existing one from a file 

* CLI arguments:
  * `insert` - insert a value into the bloom filter
  * `check` - check if a value is in the bloom filter
  * `exit` - exit the program
* Program automatically saves the bloom filter to a file when exiting properly


## Features
- Bloom filter with 2^32 bits of storage
- Compressed storage using bitsets
- Alternative hashing algorithms provided

