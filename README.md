# my-database

Build your own database

Feymann once said, “what I can’t build, I don’t understand”

**Conventional database definition:**

- Durability & Automicity
  - They can persist data to disk.
  - They are disk-based, can work with larger-than-memory data.
  - They are implemented from scratch, not as wrappers over other databases.

typical filesystem use (writing to files) has no durability guarantee, resulting in data loss or corruption after a power loss, while typical database use guarantees durability.

atomicity is a core concept in database systems. It means **a transaction must either succeed completely or not run at all.**

- Indexing data structures
  - OLAP can involve large amounts of data, with aggregations and joins. Indexing can be limited or non-existent. They are mostly column-based data stores. Used to execute ad hoc, offline, “analytical” queries that are not sensitive to latency.
  - OLTP touches small amounts of data using indexes. Low latency and cost. Based on either B+tree or LSM-tree data structures. Used to execute preprogrammed, user-facing queries that require immediate results.

**Challenges**

Challenge 1: In memory data structures vs in disk data structures
studying these characteristic of RAM and disk & indexing limit the data structure choice to B+Tree & LSM-Tree

Challenge 2: Persisting data on disk —> put them on disk and incrementally update them while maintaining atomicity.

Challenge 3: Concurrency —> For in-memory data, it’s usually OK to serialize the data structure access with a single mutex. For disk-based data, the IO latency makes this impractical and requires more advanced concurrency control.
