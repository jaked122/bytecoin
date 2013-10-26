Colin's Simplifying Proposal
----------------------------

The blockchain contains two sets of records
    Host Record
        - Space Available (Unchecked)
        - Files Stored
        - Credit Balance
        - Address
    File Record
        - Lists money left
        - Lists proofs
        - States host
        - Storage metadata
    When a block is published it publishes the new set, as well as the signed
    updates

Update Types
    Balance Transfer
        - This transfers balances between arbitrary hosts
        - Must be signed by source host
    File Addition / Update
        - Has balance added (done by file source)
        - Lists where/how to get file (Done by file source)
        - Lists hosts where file is assigned (Done by block)
        - Has file puzzles (created by file source)
        - has payout rate (set by block)
        - Removes balance from hosts where file is
          stored for recreation bonus (Done by block)
    File Balance payout
        - Triggered by proof of storage from server (block signature)
    File Balance Update
        - Adds balance / replication rate
        - Must not lower eta
            - Adding balance always okay
            - Adding rate requires cur_balance/copies per additional copy
    File Shuffle
        - Describes files being moved, generally attached to a file addition
        - Lowers Host Balance
        - Adds Other Host Balance
    Host Addition
        - Adds a possible host storage

Trusted Curve
25, 519
