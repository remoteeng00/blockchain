
**User -** anyone who as an account(hash address) is a user. User holds the coins in their account and is able to send/receive coins to/from other users.

**Transaction -** any coin transfer between two users is a transaction. The activities of users will generate transaction, which is the source of transaction pool.

**Miner -** a special role, who doesn't generate transaction, but collect/validate transaction from the pool.

**Block -** a collection of validated transactions. Every miner can propose a block, but only the one acknowledged by most of the miners will be the official block in the chian.

**Difficulty -**  a measure of how difficult it is to find a hash below a given target. Valid blocks must have a hash below this target. Mining pools also have a pool-specific share difficulty setting a lower limit for shares. In bitcoin, the network difficulty changes every 2016 blocks. For my implementation, the difficulty changes every block.

**Nonce -** a 32-bit (4-byte) field to random the hash generation. Any change to the block data (such as the nonce) will make the block hash completely different. The resulting hash has to be a value less than the current difficulty and so will have to have a certain number of leading zero bits to be less than that. As this iterative calculation requires time and resources, the presentation of the block with the correct nonce value constitutes proof of work.

**Reward -** when a block is discovered, the miner may award themselves a certain number of bitcoins, which is agreed-upon by everyone in the network. Normally the rewarding transaction is the first transaction in a block proposed by the miner.

**Fee -** The miner is also awarded the fees paid by users sending transactions. The fee is an incentive for the miner to include the transaction in their block. In the future, as the number of new bitcoins miners are allowed to create in each block dwindles, the fees will make up a much more important percentage of mining income. Ethereum is a good example of fee usage.

## How does the simulated workflow work?

## Building the code

First, you need to install Golang
if Mac

	brew install go 

if Ubuntu

	sudo apt-get update && sudo apt-get -y upgrade && sudo apt-get install -y golang-go

Once you have installed and rebooted, log in, then open up the program “terminal.” Now run the command…

	sudo apt-get update && sudo apt-get -y upgrade && sudo apt-get -y install git

Do not forget to download the missing library:

	go get -d -v .

Go to the directory and build the code 

	go build

Then run it with fun

	go run main.go

