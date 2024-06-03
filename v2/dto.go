package v2

type Request struct {
	CompanyCode      string                 `json:"companyCode"`
	ProgramCode      string                 `json:"programCode"`
	MembershipNumber string                 `json:"membershipNumber"`
	ActivityCode     string                 `json:"activityCode"`
	TxnHeader        map[string]interface{} `json:"txnHeader"`
	AcceptPayment    map[string]interface{} `json:"acceptPayment"`
}

type PaymentDetail struct {
	Details map[string]interface{}
}

type PaymentDocuments struct {
	Documents                map[string]interface{} `json:"paymentDocuments"`
	DocumentDynamicAttribute []DocumentDynamicAttribute
}

type DocumentDynamicAttribute struct {
	DynamicAttribute map[string]interface{} `json:"documentDynamicAttribute"`
}
