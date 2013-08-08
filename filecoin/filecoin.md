### Filecoin ###

Transactions on the nimbus network have a few requirements that are not met by the bitcoin currency. One key concept is 'conditional spending,' or the ability to incorporate clauses into transactions you submit to the network. An example is "I will give Bob $10,000 if gives at least $25,000 to Alice."

Another problem is that the 10 minute breaks for block discovery is too long. Transactions need to be near-instant, and filecoin aims to have a safe solution for confirming transactions into a ledger in seconds as opposed to minutes. A design goal is to completely elimate the block chain, replacing it with a simple ledger.

Filecoin also aims to be more lightweight and scalable, to be capable of a transaction volume several orders of magnitude beyond bitcoin. It is unclear whether this is a reasonable goal.

Finally, bitcoin prices are very unstable. Even if part of the bitcoin population is certain that bitcoin will only go up and to the right from here forward, the lay business person is not as certain. Filecoin aims to be perfectly reliable for the business of providing cloud storage through a distributed network.

## Equlibrium Volume ##

A major criticism of bitcoin is that the volume is a set value. Coins that get lost are never replaced. If the network grows to inifinty size, then the value of the lost coins will grow to inifinty as well. The founder of bitcoin has an estimated 600k to 1.2m bitcoins. If bitcoin becomes a globally cherished currency with a market cap over 10 trillion, this man will be the richest man in the world by an order of magnitude. Early bitcoin adopters will be the emperors. This seems unfair, and also dangerous. We don't know what they will  do with the money or the power.

Filecoin is tightly coupled against the nimbus network, to achieve an equilibrium volume according to the size and needs of the nimbus network. Filecoin exists to drive the nimbus network.

Coins are produced according to the number of files in the nimbus network. As the nimbus network stores more files, more coins are produced. As the nimbus network stores less files, the number of coins produced decreases.

Coins are removed from the system through a transaction tax. All transactions have a tax of 0.1%. The tax is elimiated from the network. Nobody gets this tax, instead the volume of coins in the network goes down.

This creates an equilibrium. If there are a constant amount of files on the network, a point will be reached where the % transaction fee for storing the files is equal to the amount of currency added to the network by stroing the files. If lots of files are being stored, then the equilibrium volume of the network will be high. If few files are being stored, the equilibruim volume of the network will be low.

If lots of transactions happen outside of the nimbus network (IE personal trades instead of using the currency to store files), then the volume will decrease and the value should raise. Mined coins go to the people seeding disk space and bandwidth on the network, so the incentive to host will go up as the value raises.

This is because, unlike Bitcoin, filecoin is backed by the nimbus network. Bitcoin has no inherent value, only the value that people give it so that it may be traded. Filecoin has inherent value: a filecoin can always be spent to obtain disk space on the cloud storage network. This makes filecoin 'good' money. The medium that it's printed on has value, and the network is set to an equilbrium that keeps the value of the currency close to the value of data that it can be traded for on the nimbus network.

## Fast Transaction Confirmation ##

The idea is not polished yet, and has known vulnerabilities. I believe however that these vulnerabilities will be closed as the idae evolves.

On the bitcoin network, blocks allow a particular miner to become the final arbiter of which spends are confirmed into the network. If a double spend occurs, only one spend can be chosen to go into the block chain. Blocks are used to solve the double spending problem. And a block chain only works as long as a single entity cannot get control of a large portion of the mining process.

Block mining also has the disadvantage of wasting CPU cycles. Money is spent on hardware that serves little purpose, other than protecting the transaction ledger of bitcoin. I believe a better scheme can be found.

It's time for lunch but I have a [faulty] framework that's serving as a rough draft. I will post it later. Perhaps I will postpone until I have ironed out the faults more.

