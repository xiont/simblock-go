package main

type IAbstractConsensusAlgo interface {
	GenesisBlock() IBlock
	Minting() IAbstractMintingTask
	SetNumConnection(connection int)
	IsReceivedBlockValid(receivedBlock IBlock, currentBlock IBlock) bool
}

type AbstractConsensusAlgo struct {
	IACA IAbstractConsensusAlgo
}
