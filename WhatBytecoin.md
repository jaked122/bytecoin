FILE IS OUT OF DATE!

Table of Contents
=================

1. Proof-of-Contribution
	* how it is different from proof-of-work
	* proof-of-contribution is effectively a reputation system
	* tethering the currency to a network with real value

2. Conditional Spending
	* enabling risk free interactions
	* limited language to protect parsers

3. Subchains
	* unspendable wallets
	* (?) multiple options for validating subblocks
	* keynodes
	* (?) confirming via file hosting
	* (?) confirming via hashcash
	* (?) confirming via primes?

3. Justice Transactions
	* pay extra to punish the provider
	* importance of anonymity to minimize malice

4. Definitions

5. Progression of the block chain
	* what is a block chain and why one is needed
	(!) * what is stored in the bytecoin block chain
	* finding blocks
	* handling failed blocks
	* time alotted for mining
	* which transactions should be included in the block

6. Hosting Files
	* cheating is possible
	* contributions must be *new*
	* contributions must be *available to everyone*
	* bandwidth modifications here

7. Mining While Hosting
	* a constant mining rate
	* timestamping
	* distribution of mined coins

8. Payment for Hosting
	* selling the file storage commodity
	* inherent value leads to stability
	* payment is put into a pool that is distributed evenly among hosts

9. Preserving Files
	* erasure coding
	* bytecoin reliability and file loss potential
	* picking redundacny
	* hosts provide a death deposit
	* replacing hosts when they disappear

10. Load Balancing
	* reddit hug problem
	* min bandwidth solution
	* dynamically adjusting the min bandwidth
	* min min set by paying client
	* other clients can take over payment
	* files literally broken into ~1000 pieces for high recovery performance, and risiliance to DOS attacks

11. Dishonest Hosts

12. Network Panic
	* determining that the network is in a panicked state
	* intentional strong forking
	* permanent fragmenting (forcing a new 'initial' block)

13. Merging Networks
	* merging an unpanicked network with a panicked network
	* merging two panicked networks

15. Wallets
	* just like bitcoin wallets?
	* no messages, just transactions


Proof-of-Contribution
=====================

Proof-of-contribution is like proof-of-work, except that instead of performing computations to achieve a hash, the computer in question provides resources ("contributions") to a network that can then be used by other participants in the network. The particular commodity of the bytecoin network is disk storage. Your contribution makes up a certain percent of the network, and that contribution directly relates to the amount of coins that you mine.

If you abstract proof-of-contribution to a higher level view, you realize that it is effectively a reputation system. Initially, there are a set of nodes each with a volume of reputation (their confirmed contribution). When new nodes come and add contributions to the network, these nodes recieve reputation in accordance to their contribution. It is stronger than a typical reputation system because contribution is a difficult task requiring the deployment of real world resources (IE contributing file storage space to the network).


Conditional Spending
====================

The network supports a very basic conditional language that can place limitations on a transaction. On example of a conditional transaction is "Send wallet A 300 coins if wallet A sends 1000 coins to wallet B." A full language specification will be available, but the language will have inherent limitations so that the cost of processing conditional spends is very low. Conditions will relate to specific network states, and will be completely verifyable by peering back into the block chain. Conditional spending may have implications about breaking the blockchain into large volumes of parallel chains.


Definitions
===========

A 'network identifier' is the hash of the most recent nondeterministic block in the network. This is actually a linked list containing all of the nondeterministic blocks in the network back to the genesis block. Most networks will have a very short linked list of nondeterministic blocks.

A 'host' is someone who is contributing file storage to the network and also mining coins.

A 'consumer' is someone who is purchasing file storage from the network.

An 'indictment' is a motion by the network to reject a particular block in the blockchain, thus rewiding the network and starting from the block before the indicted block.

'Chain hosting' is buying stroage from the network, and then using the storage you purchased to host files on the network. This seemly pointless exercise allows you to host files for free (consumers cover all of your costs) while you get to claim mining rewards from the network. Chain hosting is an illegal activity, and must be prevented by the network.

'Bytecoin reliability' is a measurement of how good the network is at keeping files online, and can be used to derive a reasonable amount of redundancy when storing files.

'File loss potential' is bytecoin reliability, but includes variance.

A 'subchain' is a block chain that comprises a different set of nodes than the global bytecoin network, (and therefore has a different blockchain) but can still perform transactions on the bytecoin network.


The Block Chain
===============

The block chain is a database of all the historical transactions of bytecoin. From the block chain, you can figure out which files are on the network, who has the files on the network, how much currency each wallet has, and each transaction that a wallet has ever made. A block chain is needed because you have a very large network trying to agree on a single database that can be updated from any part of the network. Blocks are discreet updates to the database so that the network can be in agreement of the current state.

(!move this)
Each time a new block is announced to the network, it contains a list of: 
	* all the files that were removed from the network since the previous block
	* all the files that were added to the network since the previous block
	* all of the hosts that have left the network
	* all of the hosts that have joined the network as well as how much disk space each new host is contributing
	* any changes in capacity that hosts have announced
	* all unconditional transaction that have been announced
	* all conditional transtactions that have been announced
	* all conditional transactions that have been fulfilled
	* any subchains that have been formed
	* any subchains that have been removed
	* all wallets associated to subchains
	* all superblocks from subchains

In bytecoin, the blockchain is managed by file hosts. Each block, a host is chosen to aggregate the next block in the blockchain. Hosts have a probability of being chosen equal to the percent of the network that their disk space contributions make up. A host contributing 15% of the total network has a 15% chance of being chosen as the aggregator each block. The way that a specific host is chosen is by using the hash of a previous block and mapping it to a particular host. This gives a deterministic model for picking hosts that can be verfied for past blocks as well as the current block.

Because hosts are chosen ahead of time, there is a chance that the host will fail to produce a block. Because of this, a chosen host must produce a block for the block chain by a certain time. If the host does not, the network will vote to go to the next chosen host (a process called 'indictment'). The network votes by each host in the network adding a signature to an indictment. When 51% of the network (as measured by disk space) has signed an indictment, the host and any potential blocks that it may produce are rejected, and the next host is chosen. The next host is chosen by taking the hash that chose the previous host and hashing it again. The next host now has a set amount of time to produce a block, or it will be indicted through the same process. When a host is indicted, the signed indictment is included into the block chain as proof that the network chose to move onto the next host. The indictment is announced to the network, and the network continues with the new host. Hosts can be indicted both for either being too slow or for being offline.

The amount of time allotted to a host for finding a block is adjusted with the speed of the network. If a host gets indicted for being too slow, the network extends the amount of time that the next host has to find a block. If a host releases a block on time or early, the network decreases the amount of time allotted for finding blocks. On a network with sufficiently fast hosts, a new block could be theoretically released every few seconds. This will probably be the case during the early days of the network, however as the network grows to global scale the block speed will probably decrease dramatically.

Because hosts are known ahead of time, all transactions can be forwarded to the host compiling a block, decreasing the amount of chatter on the network. To further clean up the network, the host that should recieve the current network transactions should be the host that is responsible for mining the *next* block, not the current block. By the time the previous block has been released, the next host should already have all the transactions for the next block, and can release the block faster. Any transactions sent to the host late (because of network lag and inconsistency) will be forwarded to the next host.

The block chain is a collection of wallets, and the properties of the wallets contain the remaining information necessary to derive everything else. A wallet is represented in the block chain by its public key. The hash of the public key is the address of the wallet. Each wallet has these sets of things: available currency, files that it is hosting, files that it is paying for, free space that it is advertising (as a host), and a list of conditional statements active on the wallet. As a struct, it would look like this:

// You will notice that this is not actually C code, the structs will have to be parsed as data streams
struct Wallet {
	uint64_t currencyHigh;
	uint64_t currencyLow; // 128 bits of precision
	uint64_t numFilesSupporting;
	struct filesSupporting[numFilesSupporting];
	uint64_t numFilesHosting;
	struct filesHosting[numFilesHosting];
	uint64_t numFreeSpaces;
	uint64_t numConditionalStatements;
	struct conditionalStatements[numConditionalStatements];
}

Mining
======

Bytecoin has a static amount of coins that get mined each day, just like bitcoin. Blocks however are not mined at a specific rate. This means that the amount of coins released into the network per block varies based on the amount of time it takes to release each block. The time between blocks is the distance in time between the previous host seeing his previous block, and the current host seeing the previous hosts block. Illustrated:

[Older Block] -> [Previous Host] -- [Previous Block] -> [Current Host]

The time between blocks is the distance between the moment that [Previous Host] saw [Older Block] and the moment that [Current Host] saw [Previous Block]. That distance is how long it took to mine [Previous Block], and that time is used to determine the number of coins mined in the current block.

Mined coins are distributed evenly among hosts each block according to contributions. The reward for the host that successfully mines a block comes from a different source, which will be discussed later in the writeup.

Payment for Hosting
===================

Cloud file storage is a commodity with market value. The file storage provided by hosts on the bytecoin network is essentially a cloud storage service, and it is cloud storage that can be purchased. The bytecoin network will act as a market maker, using a combination of supply, demand, and variance to set the price of renting filestorage on the network. Filestorage is *rented* not purchased, which means recurring payments for file storage. The market maker on the network will try to set a price high enough such that there is a little bit of filestorage always available (to accomodate random spikes in demand - this is why variance is one of the variables in the equation) on the network, but low enough such that the vast majority of the network is always being rented out.

This gives inherent value to the network. If there are thousands of hosts providing exabytes of data to the network, the price of bytecoin has a minimum - if the price falls greatly, the price of cloud storage will also fall greatly and people will take advantage of the extremely cheap prices. If the price of bytecoin spikes greatly, the value of mining will be great and this will encourage hosts to flock to the network. Hosts joining the network will increase the amount of data storage on the network and that will increase the inherent value of bytecoin. The simple act of bytecoin spiking in value will (over time) raise the actual value of bytecoin to match the inflation.

This market mechanism is not perfect, and as the bytecoin network grows our perceived value of cloud storage may change dramatically. But at the very least bytecoin should be more stable than bitcoin, because the value is tethered to a real world commodity that people are willing to pay real money for (example: dropbox. But at a corporate scale, you may even find services like YouTube starting to lean on bytecoin to host large files).

(filler) why there should be mining at all.

Hosting Files
=============

Hosts contribute disk storage to the network and then mine coins proportional to the amount of disk space that they host. Bytecoin has a (needs auditing) comprehensive set of algorithms for detecting cheaters.

	1. Contributions must be *new*. If a contribution is already on the network, it cannot be recontributed. I am specifically talking about an attack called 'chain hosting', where a node will pretend to host space on the network, and then when given a file to host will rent space off of the network to host the file. This allows a node to pretend that it is contributing to the network, when it is actually recursively hosting files on the network. While the host will not get any money for selling the disk storage, the host will recieve payment for mining, thus creating arbitrage for the host. I do not believe there is a way to make this inherently unprofitable for the hosts without removing the mining feature altogether. For this reason, chain hosting is something that needs to be detected. Chain hosting is detected using a mediator: when a client downloads a file, a small and randomly distributed portion of the file must be recieved through a mediator. An honest host will always use a mediator, and an honest client will always expect the file from a mediator. The mediator works by holding onto the mediated portion of the file for a set period of time, after which it will pass the data on. A chain hoster sandwhiched between two honest clients will have two use 2 mediators, thus doubling the amount of time that the held portion of the file is kept. When the client does not recieve the full mediated portion of the file after 1 wait cycle, the client can report the previous host for chain hosting.
		
		In all other cases, contributions are assumed to be new to the network.

	2. Contributions must be *useful*. Bytecoin defines useful contributions as disk storage that can be provided to anybody. The worry is that because hosts are able to mine coins from nowhere in proportion to the amount of disk storage that they provide, some hosts will attempt to lie about the size of their contributions. I try to assert that if disk storage is distributed randomly throughout the network, hosts will not be able to pretend that they are contributing a great volume because they will need to prove that they actually have the files they are randomly hosting. They cannot host a computable file (such as a list of natural numbers) and then store it on one of their own nodes, because they have no control over which hosts recieve their file, and similarly they have no control over which files the end up hosting.

		This randomness is achieved through hashing. When a file is announced to be hosted, the hash of the file is used to determine where on the network the file will land. Except then a file can be manipulated such that it hashes to a favorable place, therefore a file must be hashed again against information that is not available to the uploader of the file. The hash of the next block is not known to the uploader of the file, therefore the hash of the next block to be found will be used to determine the final destination of the file. It should be noted that this is a random place on the *entire network*, not just a random place amongst the free space of the network. There is a high probability that the new file will replace an existing file, but since the existing file has been randomly chosen, the existing file can be moved into a location with free space.

		As long as this method ensures that a host cannot control which files they are hosting, and as long as point one holds, the network should be safe from dishonesty.

Preserving Files
================

Hosts cannot be trusted to stay online, so data must be uploaded with redundancy. Some varient of erasure coding will be used to provide redundancy to the network. This gives the network the ability to replace hosts that go offline. The metric 'bytecoin reliability' is the percentage of hosts that are offline at any given moment without replacement. Bytecoin reliability can be used to determine optimal erasure coding settings such that a file has a near guarantee of remaining on the network without being overly redundant or expensive. I estimate that optimal settings on a floursing bytecoin network will typically result in less than 1.25 redunancy.

Clients get to pick their own redundancy settings when uploading a file. In erasure coding, the two relevant numbers are the number of total hosts owning a file and the number of hosts that are required to be online in order for the file to be fully recoverable.

Host Disappearance 
==================

Hosts must pay forward a deposit to the network before hosting files. This is their fee for going offline, and the fee is directly proportional to the cost of replacing the host (there is a bandwidth cost associated with replacing a host). This fee is never returned, because it is assumed that the host will eventually go offline. When a host does go offline, they will need to repay the fee in order to come back online.

Hosts must prove they are online through some sort of ping magic every so often. In the event that a host stops pinging, the network will raise a flag on their files. If the host fails to come back online after a certain amount of time, the host is considered lost, and the deposit is taken and used to recover the lost files and give them to a different host.

This allows malicious parties to DDOS major hosts and cost them large amounts of bytecoins, almost in a terrorist sort of way. Other than typical prevention measures such as Cloudflare, I do not know how to stop this. Ideally, the cost would be put on the person responsible for the loss of the host. This might actually be possible in a distributed network, but such a network does not yet exist.

(filler) method for becoming a host with no initial bytecoins

Dishonest Hosts
===============

There are two ways for a host to be dishonest. The first is for a host to release a dishonest block. The second is for a host to claim that they are hosting files when they have actually lost the files.

Detecting the first type of dishonesty is easy; just check the block against the existing database. If the block allows a double spend or has any other inconsistency, an indictment of that host can be started, and the block can be rejected.

The second type of dishonesty requires proof of hosting. The easiest way for a host to prove that they have a file is for someone to download the entire file and verify that the hashes match. There also may be a way to prove hosting because of the erasure coding; as long as the erasure coding remains consistent across the network I believe you can assume that the files are consistent. This may have bandwidth costs.

The alternative way to verify files is for a client to create a cryptographic puzzle on the file. By hashing the file against a random string, the client can produce a seed that can then be used to deterministically create a public and private key pair. The client creates this pairing before uploading the file to the network. When the client wishes for the host to prove that it still has the file, the client will reveal the public key and random string, and the host must produce the private key. This can only be done once, so the client will need multiple random strings and public/private key pairs - one for each time the client wishes to challenge the host for the file.

Hosts will be required to prove hosting every so often, and they will not be paid for the time period until they prove hosting. If the client is missing such that the client cannot challenge the host, then the host does not need to prove itself in order to get paid. Hopefully though the erasure coding plan works, then the clients involvement will not be necessary.

Ultimatum Game
==============

In bytecoin, you pay for a service, and sometimes the service may not be consistent with what was promised. Here we use the example of bandwidth. Say you request a file from a host promising to deliver at 10mpbs. When the host gives you the file at 5mbps, this host has lied. You paid top dollar for your 10mbps service and got measly 5mbps service. In a nonflexible system, the host would get paid for delivering 10mbps and you would have paid for 10mbps service. In bytecoin, you can pay an extra 50% to punish the host. Your complaint will remove the amount you paid the host, and then punish the host an additional 50%. The end result is that you paid 1.5x 10mbps price, but the host loses money instead of makes profit. This makes it highly unprofitable for the host to advertise a 10mbps service when they cannot fulfill the commitment.

The game theory idea is that no satisfied customer would ever punish the host for delivering the advertised price, and that an angry customer would be angry enough to spend more and punish the host, even though they need to spend more in order to punish the host.

Bandwidth Costs
===============

Hosts are required to maintain a minimum bandwidth speed. Some files may also have minimum bandwidth requirements, which will be achieved through increasing redundancy until all of the requirements are met. When a host is found lying about unloading capabilities, the host can be punished through using the bytecoin ultimatum punishment.

Network Panic
=============

The network will fail if 51% of hosts cannot be collected together. Imagine a disaster situation such as a country withdrawing from the internet, which eliminates 60% of the nodes on the network. There is a strong chance the host chosen to find the next block has been destroyed in the disaster. The network will motion to move past the host, however since only 40% of the network remains, the hosts will never reach the 51% needed to finish the indictment. The network needs to detect and manage this type of situation, so that the currency can exist even when disaster strikes. The network detects this state by putting a time limit on an indictment. Either the indictment will be resolved by the indicted host producing a satisfactory block or the indictment will be resolved by reaching 51% of the votes from the network. If neither of these events happen after a certain duration (to be derived later), then the network will enter a panicked state.

When a host realizes that the current indictment has not passed in time, a panic block will be announced. The panic block contains no information except for the indictment that failed to pass, from which a new list of hosts can be derived. Multiple hosts are likely to announce a panic block, these conflicts are resoved by merging the two indictment lists and creating a new block that contains signatures compiled from the two conflicting blocks. Eventually every node that announced a panic block will have merged and a complete list of nodes available to the network will have been compiled.

This new block will be the new network identifier, chained on top of the previous network identifiers. Network identifiers are used to identify divergent networks.


Merging Networks
================

Networks can split up (through panicking) into sizes that are smaller than 51% of the network, which means that it is possible for multiple panicked networks to exist that do not see each other. If they at some point reconnect, they will have conflicting block chains. They will attempt to merge.

The first step to conflict resolution is to find the latest common block. If there is no common block in the entire block chain of the two networks, then they are completely independent networks and will not merge. Otherwise there is a single point of divergence for the networks. Divergence will have happened under three types of circumstances:
	* one network proceeded like normal and one network entered a panic state 
	* both networks entered a panic satate
	* one of the networks has made a mistake

One network panicked and hard forked, while the other network did not need to panic:
In the event that only one network panicked, a merge is fairly easy. All transcations in the panicked network that are compatible with the unpanicked network (which is to be considered the real network) will be merged into the real network, and all transactions that are not compatible with the real network will be rejected outright. Transactions will be merged in the order that they were announced to the panicked network. Data in the merging network will be intially rejected, because of risk of conspiracy. It will have to bootstrap to the network the same way a new host will bootstrap to the network. Wallets and hosts in the panicked network are likely to lose a lot of money, but that is to be considered a casualty of being separated from the greater network. Network panic should be something that never happens except in moments of extreme crisis, so this should not be a problem very often.

Two panicked networks need to merge together:
In the event that two panicked networks meet, the network that had more signitures on the indictment on the panic block recieves precendence, and then the two merge like normal.

Neither network hard forked:
Either one of the networks made a mistake at some point in the block chain, or the two networks are actually completely different networks.

Subchains
=========

The final feature of bytecoin is subchains, something that really unlocks the network. Subchains will be the method for integrating anonymity into the network, for integrating with other cryptocurrencies, all of it allowing bytecoin to be the 'master' blockchain from which the others can derive stability.

Conditional Spending
====================

Wallets
=======

Wallets are public keys, which can be used to sign transactions. The block chain tracks a bunch of information related to each wallet.
