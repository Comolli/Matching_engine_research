# Overall Design
* ## User
* ## gateway 
* ## Matching Engine
* ## Database
> ------------------------
## Matching Algorithm
* ### Oders queue
## Memory matching
> --------------
## Multistage hot backup
> -------------------------------
## Memory state replication

* ##  Key technical points
        >  Deterministic algorithmic state machine services are deployed to multiple independent matching engines
        >  Accept the gateway order as input to the deterministic matching algorithm state machine
        >  According to the demand of matching algorithm, select a kind of order ordering way
        >  Each matching engine matches orders that have been sorted in a sorted way
        >  Take the transaction record output from the deterministic matching algorithm state machine as a response to the user or database
        >  Monitor the status or output differences of the match engine replicas
* ## implementation scheme
     ### Implement the matching system based on memory state machine replication
        >  Atomic multicast is used to solve the problem of reliable multicast and global ordering for matching engine orders
        >  The matching technology of pipeline based on unlocked order queue is adopted to provide fast order matching
        >  Asynchronous consistent persistence technology is used to realize the interaction with the database
        > Failure reserve is used to monitor the status of the matching engine cluster and ensure the fault tolerance of the system