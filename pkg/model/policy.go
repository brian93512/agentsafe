package model

// Action is the gateway enforcement decision for a tool.
type Action string

const (
	ActionAllow           Action = "ALLOW"
	ActionRequireApproval Action = "REQUIRE_APPROVAL"
	ActionBlock           Action = "BLOCK"
)

// ActionFromGrade derives the gateway Action from a risk Grade.
func ActionFromGrade(g Grade) Action {
	switch g {
	case GradeA, GradeB:
		return ActionAllow
	case GradeC, GradeD:
		return ActionRequireApproval
	case GradeF:
		return ActionBlock
	default:
		return ActionBlock
	}
}

// RateLimit specifies optional throughput constraints applied to an allowed tool.
type RateLimit struct {
	RequestsPerMinute int
	BurstSize         int
}

// GatewayPolicy is the deployment-ready enforcement record for one tool.
type GatewayPolicy struct {
	ToolName  string
	Action    Action
	RateLimit *RateLimit
	Reason    string
	Score     RiskScore
}

// NewGatewayPolicy constructs a GatewayPolicy from a tool name and its RiskScore.
func NewGatewayPolicy(toolName string, score RiskScore, rateLimit *RateLimit) GatewayPolicy {
	return GatewayPolicy{
		ToolName:  toolName,
		Action:    ActionFromGrade(score.Grade),
		RateLimit: rateLimit,
		Score:     score,
	}
}
