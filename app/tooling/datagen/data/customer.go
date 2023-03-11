package data

type Customer struct {
	Type          string            `json:"@type" faker:"oneof: Customer"`
	ID            int               `json:"id" faker:"boundary_start=100000000, boundary_end=999999999"`
	Name          string            `json:"name" faker:"name"`
	Status        string            `json:"status" faker:"oneof: active, rejected, canceled, not producible, in realization, terminated, in termination"`
	Account       []AccountRef      `json:"account"`
	Agreement     []AgreementRef    `json:"agreement"`
	ContactMedium []ContactMedium   `json:"contactMedium"`
	CreditProfile []CreditProfile   `json:"creditProfile"`
	EngagedParty  []RelatedPartyRef `json:"engagedParty"`
	PaymentMethod []PaymentMethod   `json:"paymentMethod"`
	RelatedParty  []RelatedPartyRef `json:"relatedParty"`
	ValidFor      ValidFor          `json:"validFor" faker:"validFor"`
}

type AccountRef struct {
	Type         string `json:"@type" faker:"oneof: AccountRef"`
	ReferredType string `json:"@referredType" faker:"oneof: BillingAccount"`
	ID           int    `json:"id" faker:"boundary_start=100000000, boundary_end=999999999"`
	Description  string `json:"description" faker:"name"` //
	Name         int    `json:"name" faker:"boundary_start=7890000520, boundary_end=7990000520"`
}

type AgreementRef struct {
	Type         string `json:"@type" faker:"oneof: AgreementRef"`
	ReferredType string `json:"@referredType" faker:"oneof: Agreement"`
	ID           int    `json:"id" faker:"boundary_start=100000000, boundary_end=999999999"`
	Name         string `json:"name" faker:"oneof: Digital Subscriber Line 100, Digital Subscriber Line 300"`
}

type RelatedPartyRef struct {
	Type         string `json:"@type" faker:"oneof: RelatedParty"`
	ReferredType string `json:"@referredType" faker:"oneof: Individual"`
	ID           int    `json:"id" faker:"boundary_start=100000000, boundary_end=999999999"`
	Name         string `json:"name" faker:"name"`
}

type CreditProfile struct {
	CreditProfileDate string   `json:"creditProfileDate" faker:"timestamp"`
	CreditRiskRating  int      `json:"creditRiskRating" faker:"boundary_start=1, boundary_end=10"`
	CreditScore       int      `json:"creditScore" faker:"boundary_start=1, boundary_end=15"`
	ValidFor          ValidFor `json:"validFor" faker:"validFor"`
}
