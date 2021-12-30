package amazon

import (
	"github.com/gocql/gocql"
)

type AmazonQuote struct {
	EntityID                   gocql.UUID `json:"entity_id"`
	QuoteID                    gocql.UUID `json:"quote_id"`
	AmazonOrderReferenceId     string     `json:"amazon_order_reference_id"`
	SandboxSimulationReference string     `json:"sandbox_simulation_reference"`
	Confirmed                  int        `json:"confirmed"`
}
