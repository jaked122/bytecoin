# DHT #
Nodes connect through a distributed hash table. Each node gets a random key and situates in a particular corner of the network.

Files have a key according to the same convention. Files are stored on whatever node is closest to their key.

When a user uploads a file, it is broken up into many pieces according to the RAID scheme picked by the user. For security reasons, all schemes must be striped such that no two nodes get the same piece of data. For cheap, unreliable storage (which will probably not be permitted/supported by most clients), data is unprotected by parity or mirroring, and is simply striped across a number of nodes. For extreme protection, data can be stored in a series of unlimited mirrored RAID 6 arrays. Applications will be responsible for choosing defaults and making the settings user friendly, but the recommendation for unsafe is unmirrored RAID 5 across 10 nodes, and the recommendation for safe is RAID 51 across 20 nodes. (public files will automatically map to safe settings...?)

When a file is stored on the network, that file is added to the public ledger with some metadata. Information about how the file is protected (number of nodes, RAID scheme, and a hash for each node) is included in the header. The next section contains instructions for payment processing. How many months, how fast, any pricing limits. The final section of the file contains all needed cryptography to insure that the various files are still on the network.

Based on the .nimbus file, the uploaded file will spread out on the network to a volume of nodes. Which nodes depend on iterated hashes of the .nimbus file, which is kept in the public ledger. (? good idea - these files could be large ?)

Things to think about:

1. who's in charge of tracking the .nimbus file? I'm thinking that merely a hash of the file goes into the public ledger, and that hash is associated with a key on the network, a node that's responsible for holding the .nimbus file.

2. is there a minimum or maximum number of nodes that a file is allowed to split between? A small file split between 1000 nodes will incur lots of overhead. This should be okay though because there are microfees for every piece of overhead used. Splitting up a file into many pieces will have tradeoffs, but we can leave it to the application to figure out what those tradeoffs best are.

3. what happens to a popular file that's on an unstable setup? I guess the key replacement scheme will keep it spread out along many nodes, which will keep the file safe, because topographic closeness is unrelated to geographic closeness for hops.
