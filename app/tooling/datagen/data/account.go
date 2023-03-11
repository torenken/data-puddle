package data

type BillingAccount struct {
	Type             string           `json:"@type" faker:"oneof: BillingAccount"`
	ID               int              `json:"id" faker:"boundary_start=100000000, boundary_end=999999999"`
	AccountType      *string          `json:"accountType" faker:"oneof: private, business"`
	Description      string           `json:"description" faker:"name"`
	LastModified     *string          `json:"lastModified" faker:"timestamp"`
	Name             int              `json:"name" faker:"boundary_start=7890000520, boundary_end=7990000520"`
	PaymentStatus    *string          `json:"paymentStatus" faker:"oneof: paid, in arrears"`
	RatingType       *string          `json:"ratingType" faker:"oneof: prepaid, postpaid"`
	State            string           `json:"state" faker:"oneof: ordered, active, canceled, terminated"`
	AccountBalance   []AccountBalance `json:"accountBalance" faker:""`
	CreditLimit      int              `json:"creditLimit" faker:"boundary_start=500, boundary_end=2000"`
	PaymentMethod    PaymentMethod    `json:"defaultPaymentMethod"`
	FinancialAccount FinancialAccount `json:"financialAccount"`
	PaymentPlan      []PaymentPlan    `json:"paymentPlan"`
	Contacts         []Contact        `json:"contact"`
	RelatedParty     []RelatedParty   `json:"relatedParty"`
	TaxExemption     TaxExemption     `json:"taxExemption"`
}

type TaxExemption struct {
	CertificateNumber   int      `json:"certificateNumber" faker:"boundary_start=45678909876, boundary_end=56789098765"`
	IssuingJurisdiction string   `json:"issuingJurisdiction" faker:"oneof: Embassy"`
	Reason              string   `json:"reason" faker:"oneof: VIP"`
	ValidFor            ValidFor `json:"validFor" faker:"validFor"`
}

type PaymentPlan struct {
	NumberOfPayments int           `json:"numberOfPayments" faker:"boundary_start=1, boundary_end=3"`
	Priority         int           `json:"priority" faker:"oneof: 1,2"`
	Status           string        `json:"status" faker:"oneof: Effective"`
	TotalAmount      int           `json:"totalAmount" faker:"boundary_start=50, boundary_end=100"`
	PlanType         string        `json:"planType" faker:"oneof: regular"`
	ValidFor         ValidFor      `json:"validFor" faker:"validFor"`
	PaymentMethod    PaymentMethod `json:"paymentMethod"`
}

type FinancialAccount struct {
	Name           string         `json:"name" faker:"oneof: partnership account"`
	AccountBalance AccountBalance `json:"accountBalance"`
}

type AccountBalance struct {
	Amount      float64  `json:"amount" faker:"amount"`
	BalanceType string   `json:"balanceType" faker:"oneof: ReceivableBalance"`
	ValidFor    ValidFor `json:"validFor" faker:"validFor"`
}

type PaymentMethod struct {
	Name string `json:"name" faker:"oneof: direct debit, professional payment, family payment"`
}

type Contact struct {
	ContactName   string          `json:"contactName" faker:"name"`
	ContactType   string          `json:"contactType" faker:"oneof: primary"`
	PartyRoleType string          `json:"partyRoleType" faker:"oneof: publisher"`
	RelatedParty  RelatedParty    `json:"relatedParty"`
	ValidFor      ValidFor        `json:"validFor" faker:"validFor"`
	ContactMedium []ContactMedium `json:"contactMedium"`
}

type ValidFor struct {
	StartDateTime string `json:"startDateTime"`
	EndDateTime   string `json:"endDateTime"`
}

type ContactMedium struct {
	MediumType     string     `json:"mediumType" faker:"oneof: PostalAddress"`
	Preferred      bool       `json:"preferred"`
	Characteristic MediumChar `json:"characteristic"`
	ValidFor       ValidFor   `json:"validFor" faker:"validFor"`
}

type MediumChar struct {
	City            string `json:"city" faker:"oneof: Oldenburg"`
	Country         string `json:"country" faker:"oneof: Deutschland"`
	EmailAddress    string `json:"emailAddress" faker:"oneof: jane.doe@example.com, john.doe@example.com, peter.parker@example.com"`
	FaxNumber       string `json:"faxNumber" faker:"oneof: +49 441 45910, +49 441 32500, +49 441 9123-99"`
	PhoneNumber     string `json:"phoneNumber" faker:"oneof: +49 441 45912, +49 441 32509, +49 441 9123-0"`
	PostCode        string `json:"postCode" faker:"oneof: 26121, 26122, 26135"`
	StateOrProvince string `json:"stateOrProvince" faker:"oneof: Niedersachsen"`
	Street1         string `json:"street1" faker:"oneof: Lange Straße, Schlossplatz, Elisabethstraße, Heiligengeiststraße"`
}

type RelatedParty struct {
	Name string `json:"name" faker:"name"`
	Role string `json:"role" faker:"oneof: owner"`
}

type Characteristic struct {
	Name      string `json:"name"`
	ValueType string `json:"valueType"`
	Value     string `json:"value"`
}
