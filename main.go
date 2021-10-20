package main

import (
	"bytes"
	"crypto/rsa"
	"fmt"
	"math/rand"
	"time"

	"./config"
	"./core"
	"./role"
	"./util"
)

var chain core.Blockchain
var users []*role.User
var miner *role.Miner

const userCount = 4

func boostNetwork() {
	// 1. create the initial user of blockchain
	firstUser := role.CreateBoostUser()

	// 2. boost the blockchain with initial user
	diff := core.CreateMADifficulty(10000, 0.2, 16)
	chain = core.InitializeBlockchainWithDiff(&firstUser.Address, diff)
}

func boostUsers() {
	// create 10 users
	for i := 0; i < userCount; i++ {
		user := role.CreateUser(chain)
		users = append(users, user)
		util.GetMainLogger().Infof("User %v created\n", user.GetShortIdentity())
	}
}

func startTrading() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	block := miner.GetBlockChain().GetLatestBlock()
	var usage [userCount + 1]int // miner is the last one
	for i := 0; i <= userCount; i++ {
		usage[userCount] = 0
	}

	var from, to int
	for i := 0; true; i++ {
		from = r1.Intn(userCount)
		if from < userCount/2 {
			to = userCount/2 + r1.Intn(userCount/2)
		} else {
			to = r1.Intn(userCount / 2)
		}

		if util.Hash(*block) != util.Hash(miner.GetBlockChain().GetLatestBlock()) {
			util.GetMainLogger().Infof("Chain confirmed a new block. Clean the usage\n")

			block = miner.GetBlockChain().GetLatestBlock()
		}

		if block != nil {
			time.Sleep(500 * time.Millisecond)
		}

		amount := r1.Intn(config.MinerRewardBase / 1000)
		fee := r1.Intn(10)
		if couldUserPostTransaction(miner.Address) && int(miner.GetBlockChain().BalanceOf(&miner.Address)) > amount {
			miner.SendTo(users[to], uint64(amount), uint64(fee))
			time.Sleep(1 * time.Second)
		}

		amount = r1.Intn(config.MinerRewardBase / 1000)
		fee = r1.Intn(userCount)
		if couldUserPostTransaction(users[from].Address) && int(miner.GetBlockChain().BalanceOf(&users[from].Address)) > amount {
			users[from].SendTo(users[to], uint64(amount), uint64(fee))
			time.Sleep(1 * time.Second)
		}
	}
}

func initializeOneMinerAndStartMining() {
	miner = role.CreateMiner(chain)
	miner.StartMining()
}

func printStatus() {
	for {
		var buffer bytes.Buffer
		buffer.WriteString(fmt.Sprintf("Miner[%s:%d]] ", miner.GetShortIdentity(), miner.GetBlockChain().BalanceOf(&miner.Address)))

		for i := 0; i < userCount; i++ {
			buffer.WriteString(fmt.Sprintf("User[%s:%d]] ", users[i].GetShortIdentity(), miner.GetBlockChain().BalanceOf(&users[i].Address)))
		}

		util.GetMainLogger().Debugf("Account Status: %s\n", buffer.String())

		time.Sleep(1 * time.Second)
	}
}

func couldUserPostTransaction(sender rsa.PublicKey) bool {
	for _, tran := range miner.GetBlockChain().TransactionPool {
		if tran.Sender == sender {
			return false
		}
	}
	return true
}

func runSimulator() {
	// 1. boost the blochchain
	util.GetMainLogger().Infof("Start to boost blockchain \n")
	boostNetwork()
	util.GetMainLogger().Infof("Finished boosting blockchain \n")

	util.GetMainLogger().Infof("Start to boost miner \n")
	go initializeOneMinerAndStartMining()

	util.GetMainLogger().Infof("Start to boost users \n")
	boostUsers()
	util.GetMainLogger().Infof("Finished boosting users \n")

	go printStatus()

	util.GetMainLogger().Infof("Start to boost trading \n")
	go startTrading()

	for {
		time.Sleep(10 * time.Second)
	}
}

func main() {
	runSimulator()
}
