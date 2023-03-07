package data

type BillingAccount struct {
	Type string `json:"@type" faker:"oneof: BillingAccount"`
	//BaseType        *string          `json:"@baseType" faker:"-"`
	//SchemaLocation  *string          `json:"@schemaLocation" faker:"-"`
	//Href            *string          `json:"href" faker:"-"`
	ID             int            `json:"id" faker:"boundary_start=100000000, boundary_end=999999999"`
	AccountType    *string        `json:"accountType" faker:"oneof: private, business"`
	Description    string         `json:"description" faker:"name"`
	LastModified   *string        `json:"lastModified" faker:"timestamp"`
	Name           int            `json:"name" faker:"boundary_start=7890000520, boundary_end=7990000520"`
	PaymentStatus  *string        `json:"paymentStatus" faker:"oneof: paid, in arrears"`
	RatingType     *string        `json:"ratingType" faker:"oneof: prepaid, postpaid"`
	State          string         `json:"state" faker:"oneof: ordered, active, canceled, terminated"`
	AccountBalance AccountBalance `json:"accountBalance"` //remove array - only need one
	//AccountRelation  *string          `json:"accountRelationship" faker:"-"` //todo
	//BillStructure    *string          `json:"billStructure" faker:"-"`       //todo
	CreditLimit      int              `json:"creditLimit" faker:"boundary_start=500, boundary_end=2000"`
	PaymentMethod    PaymentMethod    `json:"defaultPaymentMethod"`
	FinancialAccount FinancialAccount `json:"financialAccount"`
	PaymentPlan      PaymentPlan      `json:"paymentPlan"`  //remove array - only need one
	Contacts         Contact          `json:"contact"`      //remove array - only need one
	RelatedParty     RelatedParty     `json:"relatedParty"` //remove array - only need one
	TaxExemption     TaxExemption     `json:"taxExemption"`
	//Characteristic   Characteristic   `json:"characteristic" faker:"-"` //remove array - only need one
}

type TaxExemption struct {
	CertificateNumber   int      `json:"certificateNumber" faker:"boundary_start=45678909876, boundary_end=56789098765"`
	IssuingJurisdiction string   `json:"issuingJurisdiction" faker:"oneof: Embassy"`
	Reason              string   `json:"reason" faker:"oneof: VIP"`
	ValidFor            ValidFor `json:"validFor"`
}

type PaymentPlan struct {
	NumberOfPayments int           `json:"numberOfPayments" faker:"boundary_start=1, boundary_end=3"`
	Priority         int           `json:"priority" faker:"oneof: 1,2"`
	Status           string        `json:"status" faker:"oneof: Effective"`
	TotalAmount      int           `json:"totalAmount" faker:"boundary_start=50, boundary_end=100"`
	PlanType         string        `json:"planType" faker:"oneof: regular"`
	ValidFor         ValidFor      `json:"validFor"`
	PaymentMethod    PaymentMethod `json:"paymentMethod"`
}

type FinancialAccount struct {
	Name           string         `json:"name" faker:"oneof: partnership account"`
	AccountBalance AccountBalance `json:"accountBalance"`
}

type AccountBalance struct {
	Amount      float64  `json:"amount" faker:"amount"`
	BalanceType string   `json:"balanceType" faker:"oneof: ReceivableBalance"`
	ValidFor    ValidFor `json:"validFor"`
}

type PaymentMethod struct {
	//Href *string `json:"href" faker:"-"`
	//ID   int     `json:"id" faker:"-"`
	Name string `json:"name" faker:"oneof: direct debit, professional payment, family payment"`
}

type Contact struct {
	ContactName   string        `json:"contactName" faker:"name"`
	ContactType   string        `json:"contactType" faker:"oneof: primary"  `
	PartyRoleType string        `json:"partyRoleType" faker:"oneof: publisher"  `
	RelatedParty  RelatedParty  `json:"relatedParty"`
	ValidFor      ValidFor      `json:"validFor"`
	ContactMedium ContactMedium `json:"contactMedium"` //remove array - only need one
}

type ValidFor struct {
	StartDateTime string `json:"startDateTime" faker:"timestamp"`
	EndDateTime   string `json:"endDateTime" faker:"timestamp"`
}

type ContactMedium struct {
	Type string `json:"@type"`
	//BaseType       string     `json:"@baseType"`
	//SchemaLoc      string     `json:"@schemaLocation"`
	MediumType     string     `json:"mediumType" faker:"oneof: PostalAddress"`
	Preferred      bool       `json:"preferred"`
	Characteristic MediumChar `json:"characteristic"`
	ValidFor       ValidFor   `json:"validFor"`
}

type MediumChar struct {
	City         string `json:"city" faker:"oneof: Oldenburg"`
	Country      string `json:"country" faker:"oneof: Germany"`
	EmailAddress string `json:"emailAddress" faker:"email"`
	FaxNumber    string `json:"faxNumber" faker:"e_164_phone_number"`
	PhoneNumber  string `json:"phoneNumber" faker:"e_164_phone_number"`
	PostCode     string `json:"postCode" faker:"oneof: 26133, 26121, 26131"`
	//SocialNetworkId string `json:"socialNetworkId"`
	StateOrProvince string `json:"stateOrProvince" faker:"oneof: Niedersachsen"`
	Street1         string `json:"street1" faker:"oneof: Peterstr. 12, Gaststr. 44"`
	//Street2         string `json:"street2"`
	//BaseType        string `json:"@baseType"`
	//SchemaLocation  string `json:"@schemaLocation"`
	Type string `json:"@type"`
}

type RelatedParty struct {
	//Type           string `json:"@type"`
	//ReferredType   string `json:"@referredType"`
	//BaseType       string `json:"@baseType"`
	//SchemaLocation string `json:"@schemaLocation"`
	//Href           string `json:"href"`
	//ID             string `json:"id"`
	Name string `json:"name" faker:"name"`
	Role string `json:"role" faker:"oneof: owner"`
}

type Characteristic struct {
	Type string `json:"@type"`
	//BaseType  string `json:"@baseType"`
	//SchemaLoc string `json:"@schemaLocation"`
	Name      string `json:"name"`
	ValueType string `json:"valueType"`
	Value     string `json:"value"`
}
