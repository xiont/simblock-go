package simulator

import "simblock-go/settings"

func SimulatorInit() {
	(*settings.FUNCS)["BitcoinCoreTable"] = NewBitcoinCoreTable
	(*settings.FUNCS)["ProofOfWork"] = NewProofOfWork
}
