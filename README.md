This is my attempt at making a readme while I develop, meaning that as I add
features and change things the documentation will keep up. For bytecoin I feel
that this is particularly necessary because there is a massive volume of new
cryptocurrency concepts.

Right now there is a client and a server.

The client takes the argument 'store' followed by a filename, opens a
connection to the server on port 8080 and sends the file to the server.

The server takes the file and prints it to storage.ytc

In order to store the block chain, it was decided that a full database would
be necessary. The database chosen is sqlite. For that reason, you will need to
run the command: 'go get github.com/mattn/go-sqlite3'

the database currently contains a single table of hosts and files. each file
must be unique. each file has an associated host, multiple files can have
the same host. The database is saved to a file for persistence. Currently hosts
and files are always stored with the same key.


TL;DR
==

A cryptocurrency that is backed by a commodity that has real market value; disk
storage on the cloud. The value of the cryptocurrency should be more or less
tethered to the value of disk storage on the cloud.

Also introduced is compatibility with/awareness of other cryptocurrencies, as
well as conditional spending, as well as a psychological game based on game
theory and additional revelations brought about by the ultimatum game.

Furthermore, instead of having a predetermined number of coins get mined, 
a forever constant number of coins are mined per day, however this is 
turned into an equilibrium by a very small transaction tax. At some point,
the number of coins entering the system per day will equal the number of
coins being taxed from the system per day.

There are many justifications for all these additions to the cryptocurrency,
part of it being an excercise in exploration of the possibilities. But I do
believe that all of the proposed changes have the potential to bring
cryptocurrency appreciable steps forward, and in fact I am surprised that more
of these concepts have not been implemented already.

Though I have not mentioned them yet, there will also be considerations for
adding tools to bytecoin such as Zerocoin, distributed exchanges, and
potentially anonymous file storage and retrieval.

The Concept
===

Note: I have spent multiple months thinking this through, and that makes it
difficult to explain clearly. If someone else figures out how everything
works and writes an alternative summary, I would be happy to feature it.
I am confident (but not perfectly so) in the security of this model,
however I am not a professional and it should be audited.

Bitcoin as a currency has shown us that cryptocurrency is possible. Bytecoin
aims to take cryptocurrency to yet another echelon of viability. Though the
name may suggest an arrogance and hint at an evolution beyond bitcoin, bytecoin
actually derives its name from its source of mining. Bitcoins are mined through
the hashcash algorithm. Bytecoins are mined by contributing long term file
storage to the network. 

Such a dramatic shift required a lot of tinkering with the setup of the network
to prevent cheating and keep the network secure from cryptographic attack.

The biggest shift is in the way that the block chain fucntions and maintains
stability. The sacrifice that had to be made was the ability to shift between
forked chains. With bitcoin, a fork (or potential fork) is solved by looking at
the depth on each block chain. The longest chain (or deepest) is chosen as the
one true block to be incorporated into the block chain. With bytecoin, a
deterministic model is used to determine which blocks are potentially valid.

With bitcoin, that might be dangerous because the act of mining bitcoins is not
deterministic at all. Any node could find the next block and luck determines
which node gets the bitcoins. With bytecoins, the nodes in charge of finding
the next block are predetermined - the luck portion of the equation is sorted
out multiple blocks ahead of time.

When it is time to push a transaction to the network, a client will submit
the transaction to the node that will be solving the block that the transaction
appears in. It can do this, because unlike bitcoin, the node is known ahead of
time. The node responsible for solving the block will sign the transaction and
return the signed transaction to the client. The client now has a promise from
the node that the transaction will make it into the block chain. If the node
does not put the transaction into the blockchain, the client can cry foul,
providing the signed transaction as proof, and then the node will be thrown
from the network.

Each block, three nodes are responsible for compiling the transaction and the
blcok data. They will have a list of all transactions that occured between 
their block and the previous block, and they will compile the transactions
and announce a block to the network. The three nodes may produce different
results. This is okay, because just as the nodes have been chosen ahead of
time, a pecking order has also been chosen. What happens is that the 3 blocks
are merged together into a single block, all unique transactions included
into the network. In the event of a double spend, the spend included by the
node of highest priority will be included into the blockchain. In the event of
dishonesty (IE a node includes a double spend in a compiled block) the block
from that node is rejected and the node is thrown out of the network.

If a node fails to deliver a block in time, that block is inserted into the
network as an empty block. If all three nodes fail to produce a block into
the network, then that entire block is lost. This should not be a common
occurance, and will hopefully not inconvenience too many people. Nodes will
be considered as failed if they are dishonest, however it is also possible
that they will be in some sort of physical trouble (like a power outage) that
renders them unable to provide a block.

Such a method for tracking transactions and handling a blockchain completely
removes the need for mining. The only point of vulnerability is in being
severed from the network, or being conned into believing you are in the real
network when you are actually in a fork of the network.

The problem of tracking transactions has been solved, but mining is still a
conundrum. This is solved by the contribution of disk storage to the bytecoin
network. 

People can contribute to the bytecoin network by contributing disk space - 
they host files on their machines in return for access to newly minted coins.
But it's actually more than that: because bytecoin contributes disk space as a
commodity with actual market value, contributors to the network are able to
charge for their services, in effect double dipping - getting mining rewards
from the bytecoin network as well as getting fees for hosting other people's
files on their machines.

This means that just like bitcoin, anybody willing to contribute to the network
has a fair chance at mining the currency. It also means that there is a huge
potentially for very cheap disk storage on a cloud network, so cheap that it
may make services like dropbox and google drive irrelevant.

There is a cryptographic problem here: how do we know that a person is actually
contributing all of the disk space that they are claiming, how do we know that
files are not being sent immediately to /dev/null?

Before uploading a file to the network, a client will make a series of
cryptographic puzzles that can only be solved if the file is available. The
nieve implementation of these cryptographic puzzles involves the use of the
public key system. To produce one puzzle, the client takes a cryptographic
string and hashes it against the entire file, creating a private key. The
private key is then used to create a public key. Finally, the client encrypts
a random string using the public key and then stores the public key, the 
unencrypted string, and the encrypted string. The puzzle is then to find the
private key that can be combined with the public key and the cryptographic text
to produce the randomly chosen string. The client must also store the random
string that was originally chosen to hash against the whole file.

When the client wishes for a host to prove that they have the file, the client
need only present the 4 pieces and ask the host to produce the private key that
solves the puzzle. The host will be able to produce the private key by using the
same random string as the client and hashing it against the file. One the host
has the private key, the host will announce it and the network can verify
collectively that this private key solves the puzzle. The host can then be
confirmed to still have the file in question.

This challenge only works one time of course, because the host can then cache
the private key and use it for all future challenges. That's why the client
must make a series of puzzles - a different starting random string will be
needed each time to insure that the host cannot just cache the private keys.
Furthermore, the random strings must be kept secret from the network until the
client is ready for the hosts to prove that they are still holding the file.

This means that the client must still store some information, even when they
are storing things on the network. This is okay, because the client will need
to be responsible for the coin wallet anyway, and the size of the puzzles
should be substantially smaller than the size of the files being stored on the
network.

The next problem is the problem of conspiracy. If two people collaborate, they
can pretend to store an infinite amount of data on the network and get all of
the newly mined coins. The works by one person storing perfectly predictable
data, such as a series of 0s. The other person does not need to keep that
informaiton on disk, they will be able to solve cryptographic puzzles all day
while pretending to store exabytes of data on the network from their dinky
laptop.

To prevent this, measures must be taken to insure that clients have no control
over which hosts recieve the data, and hosts have no control over which clients
they service. The first measure is that newly added filespace to the network
cannot be inserted to directly. This is to prevent one person from adding a
large amount of space and then having their friend fill it right away with
predictable data. New files must be added into the disk space that has been
available on the network for a long time, in a random fashion that is heavily
weighted towards the disk space that has had the same file for the longest.

When a new file is added to space that has been filled for a long time, it
bumps an old file out of the way. The old file then moves into the new space.
This means that the file being added to the new space is essentially some
random file from the network, instead of the new file being added by the
potential consiprators. 

Finally, which node the client's file ends up on depends on the hash of the
file, however to further remove control from the client's hands, the hash of
the file is then hashed against the block released after the client announces
an intent to store the file. Because the client cannot predict what the hash
of the block will be, the client cannot manipulate the hash such that a
predictable file (like all 0s) ends up on a friend's machine.

I do not think that this security model is perfect,
but I believe that the most major hurdels have been cleared, and that
perfecting the model is as simple as figuring out what other attacks have not
yet been considered.

The entire network will act as a market maker for the commodity of file
storage. Hosts will contribute filespace to the network and in return get an
explicit amount of money, decided by the current block. Clients will also pay
an explicit amount of money decided by the current block. These prices exclude
the rewards that hosts get for mining data on the network. An equilibrium will
be established based on variance that sets a price such that as much disk space
is being sold to consumers at once as possible without ever actually running
out of disk space. Clients must always be able to buy more space, so the price
must be set high enough that there is never a moment where there is no free
space on the network for a client to buy. The exact equation for deciding the
prices is still up for debate.

The amount of currency added to the network per day is constant. The block rate
is not constant, so the amount of currency added to the network for a specific
block depends on the amount of time that it took to compute the previous block.
This is not quite true; the network actually bases all decisions two blocks
into the past. Therefore, the amount of currency added for a given block
depends on the amount of time between the second-most-recent block and the
thrid-most-recent-block.

The hash of the third most recent block gives random information about which 
three nodes are in charge of finding the next block. Nodes are chosen randomly
by the hash according to how much disk space they are hosting. Hosting twice
the disk space gives twice the chance of being chosen to mine the block. The
reward for being the person to mine the block is actually rather small by
comparison to bitcoin - perhaps a 2% fee as opposed to the full 100%. The rest
of the mined coins are distributed to all the nodes hosting filespace according
to the volume of space they are hosting. Compiling blocks is viewed as a chore
and not as a rewarding thing. There is a reward but it is small. I do not think
that the small reward will discourage people from computing the blocks,
especially because there will be a penalty for failing to compute a block -
though the penalty is yet to be decided.

I now bring up the idea of conditional spending plus awareness of external
networks. This provides an environment for trades like: If bitcoin wallet A
sends X coins to bitcoin wallet B, then I will send Y bytecoins to bytecoin
wallet C. Conditional spending also allows clients to say something like
"Host file A on the network until the price of hosting exceeds X or until
the total amount spent hosting this file exceeds Y." When somebody makes
a conditional spend, they 'spend' the coins up front, meaning the network
removes the coins from the clients wallet. Then, if the condiiton is filled
the network will send the coins to their appropriate destination. If the 
transaction expires before the condition is filled, then the network will
return the coins to the original wallet that spent them.

The final problem is the problem of bandwidth. The network will enforce a
minimum bandwidth for hosts based on how much data they are hosting. As a
temporary suggestion, clients downloading a file from a host should expect to
get a full 1 mbps per 1 GB that they download, meaning that all files should
finish downloading within about 3 hours. (This does seem fast, maybe the
minimums will be relaxed). Since I could find no way to cryptographically
verify that someone had sent an appropriate amount of bandwidth to a client at
an appropriate speed, I chose instead to rely on a social game.

When downloading something, a client will pay a markey price for the bandwidth,
probably determined by the current price of file space. When the file is done
downloading, the client has the option to report the host to the network and
complain about the host being to slow or just simply not delivering the file at
all. The key dynamic is that the client has already paid for the file, and will
not be refunded under any cirucumstances. Additionally, the client must pay
even more money to have the host punished. In a game theory sense, there is no
reason for the client to punish the host because the client can only lost more
money. But from a societal perspective, it is like the ultimatum game. If the
client feels slighted, they can extend a small amount of additional harm upon
themselves to punish the host. If the client chooses to punish the host, the
host will not recieve the coins from the transaction (the coins are instead
donated back to the network in the form of decimation) and additionally the
host will be fined a little bit beyond that (rather, the host will lose a
deductible that they paid before the bandwidth transfer started). The cost of
the fine and severity of the punishemnt are determined by the network, not by
the client or host.

I said finally, but then I forgot about consistency and redundancy. Hosts
cannot be trusted to hold onto files - websites go down and so do hosts.
Because of this, a file is not uploaded to the network to a single host in its
original form. The client first chooses an amount of redundancy (following a
system like Tahoe-LAFS) establishing how much redundancy is desired and what
type of RAID array should be used when storing the file on hosts. The network
will then keep track of when a host goes down, and will react immediately to
rebuild the file from the redundant pieces held on other hosts.


I have not gone into specific details like what should be contained in a block
or how everything should look exactly, but I believe that in this README I have
covered all of the major points, enough that one should be able to construct a
functional protocol. I intend to construct such a protocol over the next few
months, along with a first implementation of that protocol. In order to improve
the health of the cryptocurrency ecosystem, bytecoin shoud rely on hashing and
cryptographic algorithms that are not currently found in the cryptocurrency
ecosystem. My proposal is SHA3 for all hashing and whatever public key cipher
is currently regarded as the most trusted (because currently no cryptocurrency
uses public keys).
