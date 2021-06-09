package client

import (
	govclient "github.com/BITCOIVA/Bitcoiva-sdk/x/gov/client"
	"github.com/BITCOIVA/Bitcoiva-sdk/x/params/client/cli"
	"github.com/BITCOIVA/Bitcoiva-sdk/x/params/client/rest"
)

// param change proposal handler
var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
