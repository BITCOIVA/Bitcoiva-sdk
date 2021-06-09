package client

import (
	"github.com/BITCOIVA/Bitcoiva-sdk/x/distribution/client/cli"
	"github.com/BITCOIVA/Bitcoiva-sdk/x/distribution/client/rest"
	govclient "github.com/BITCOIVA/Bitcoiva-sdk/x/gov/client"
)

// param change proposal handler
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
)
