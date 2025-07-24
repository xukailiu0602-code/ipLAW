package models

type QARecord struct {
	ID        string      `bson:"_id" json:"id"`
	UserID    string      `bson:"user_id" json:"user_id"`
	Query     string      `bson:"query" json:"query"`
	Contexts  []Context   `bson:"contexts" json:"contexts"`
	Answer    string      `bson:"answer" json:"answer"`
	Citations []string    `bson:"citations" json:"citations"`
	Laws      []Law       `bson:"laws" json:"laws"`
	Cases     []Case      `bson:"cases" json:"cases"`
	RiskScore int         `bson:"risk_score" json:"risk_score"`
	Rationale string      `bson:"rationale" json:"rationale"`
	CreatedAt int64       `bson:"created_at" json:"created_at"`
}

type Context struct {
	ID   string `bson:"id" json:"id"`
	Text string `bson:"text" json:"text"`
}

type Law struct {
	Article     string `bson:"article" json:"article"`
	Description string `bson:"description" json:"description"`
}

type Case struct {
	CaseID     string  `bson:"case_id" json:"caseId"`
	Court      string  `bson:"court" json:"court"`
	Year       int     `bson:"year" json:"year"`
	Summary    string  `bson:"summary" json:"summary"`
	Similarity float32 `bson:"similarity" json:"similarity"`
}