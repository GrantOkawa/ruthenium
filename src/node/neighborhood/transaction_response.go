package neighborhood

type TransactionResponse struct {
	RecipientAddress string
	SenderAddress    string
	SenderPublicKey  string
	Signature        string
	Timestamp        int64
	Value            uint64
}
