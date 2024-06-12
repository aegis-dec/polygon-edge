package server

import (
	"github.com/aegis-dec/polygon-edge/chain"
	"github.com/aegis-dec/polygon-edge/consensus"
	consensusDev "github.com/aegis-dec/polygon-edge/consensus/dev"
	consensusDummy "github.com/aegis-dec/polygon-edge/consensus/dummy"
	consensusIBFT "github.com/aegis-dec/polygon-edge/consensus/ibft"
	consensusPolyBFT "github.com/aegis-dec/polygon-edge/consensus/polybft"
	"github.com/aegis-dec/polygon-edge/forkmanager"
	"github.com/aegis-dec/polygon-edge/secrets"
	"github.com/aegis-dec/polygon-edge/secrets/awsssm"
	"github.com/aegis-dec/polygon-edge/secrets/gcpssm"
	"github.com/aegis-dec/polygon-edge/secrets/hashicorpvault"
	"github.com/aegis-dec/polygon-edge/secrets/local"
	"github.com/aegis-dec/polygon-edge/state"
)

type GenesisFactoryHook func(config *chain.Chain, engineName string) func(*state.Transition) error

type ConsensusType string

type ForkManagerFactory func(forks *chain.Forks) error

type ForkManagerInitialParamsFactory func(config *chain.Chain) (*forkmanager.ForkParams, error)

const (
	DevConsensus     ConsensusType = "dev"
	IBFTConsensus    ConsensusType = "ibft"
	PolyBFTConsensus ConsensusType = consensusPolyBFT.ConsensusName
	DummyConsensus   ConsensusType = "dummy"
)

var consensusBackends = map[ConsensusType]consensus.Factory{
	DevConsensus:     consensusDev.Factory,
	IBFTConsensus:    consensusIBFT.Factory,
	PolyBFTConsensus: consensusPolyBFT.Factory,
	DummyConsensus:   consensusDummy.Factory,
}

// secretsManagerBackends defines the SecretManager factories for different
// secret management solutions
var secretsManagerBackends = map[secrets.SecretsManagerType]secrets.SecretsManagerFactory{
	secrets.Local:          local.SecretsManagerFactory,
	secrets.HashicorpVault: hashicorpvault.SecretsManagerFactory,
	secrets.AWSSSM:         awsssm.SecretsManagerFactory,
	secrets.GCPSSM:         gcpssm.SecretsManagerFactory,
}

var genesisCreationFactory = map[ConsensusType]GenesisFactoryHook{
	PolyBFTConsensus: consensusPolyBFT.GenesisPostHookFactory,
}

var forkManagerFactory = map[ConsensusType]ForkManagerFactory{
	PolyBFTConsensus: consensusPolyBFT.ForkManagerFactory,
}

var forkManagerInitialParamsFactory = map[ConsensusType]ForkManagerInitialParamsFactory{
	PolyBFTConsensus: consensusPolyBFT.ForkManagerInitialParamsFactory,
}

func ConsensusSupported(value string) bool {
	_, ok := consensusBackends[ConsensusType(value)]

	return ok
}
