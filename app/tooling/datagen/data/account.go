package data

type BillingAccount struct {
	Type string `json:"@type" faker:"oneof: BillingAccount"`
	//BaseType        *string          `json:"@baseType" faker:"-"`
	//SchemaLocation  *string          `json:"@schemaLocation" faker:"-"`
	//Href            *string          `json:"href" faker:"-"`
	ID               int              `json:"id" faker:"boundary_start=100000000, boundary_end=999999999"`
	AccountType      *string          `json:"accountType" faker:"oneof: private, business"`
	Description      string           `json:"description" faker:"name"`
	LastModified     *string          `json:"lastModified" faker:"timestamp"`
	Name             int              `json:"name" faker:"boundary_start=7890000520, boundary_end=7990000520"`
	PaymentStatus    *string          `json:"paymentStatus" faker:"oneof: paid, in arrears"`
	RatingType       *string          `json:"ratingType" faker:"oneof: prepaid, postpaid"`
	State            string           `json:"state" faker:"oneof: ordered, active, canceled, terminated"`
	AccountBalance   AccountBalance   `json:"accountBalance"`                //remove array - only need one
	AccountRelation  *string          `json:"accountRelationship" faker:"-"` //todo
	BillStructure    *string          `json:"billStructure" faker:"-"`       //todo
	CreditLimit      int              `json:"creditLimit" faker:"boundary_start=500, boundary_end=20000"`
	PaymentMethod    PaymentMethod    `json:"defaultPaymentMethod"`
	FinancialAccount *string          `json:"financialAccount" faker:"-"`
	PaymentPlan      *string          `json:"paymentPlan" faker:"-"`
	Contacts         []Contact        `json:"contact" faker:"-"`
	RelatedParty     []RelatedParty   `json:"relatedParty" faker:"-"`
	TaxExemption     *string          `json:"taxExemption" faker:"-"`
	Characteristic   []Characteristic `json:"characteristic" faker:"-"`
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
	ContactName   string          `json:"contactName"`
	ContactType   string          `json:"contactType"`
	PartyRoleType interface{}     `json:"partyRoleType"`
	RelatedParty  interface{}     `json:"relatedParty"`
	ValidFor      ValidFor        `json:"validFor"`
	ContactMedium []ContactMedium `json:"contactMedium"`
}

type ValidFor struct {
	StartDateTime string `json:"startDateTime" faker:"timestamp"`
	EndDateTime   string `json:"endDateTime" faker:"timestamp"`
}

type ContactMedium struct {
	Type           string      `json:"@type"`
	BaseType       interface{} `json:"@baseType"`
	SchemaLoc      interface{} `json:"@schemaLocation"`
	MediumType     string      `json:"mediumType"`
	Preferred      bool        `json:"preferred"`
	Characteristic MediumChar  `json:"characteristic"`
	ValidFor       interface{} `json:"validFor"`
}

type MediumChar struct {
	City            interface{} `json:"city"`
	Country         interface{} `json:"country"`
	EmailAddress    string      `json:"emailAddress"`
	FaxNumber       interface{} `json:"faxNumber"`
	PhoneNumber     interface{} `json:"phoneNumber"`
	PostCode        interface{} `json:"postCode"`
	SocialNetworkId interface{} `json:"socialNetworkId"`
	StateOrProvince interface{} `json:"stateOrProvince"`
	Street1         interface{} `json:"street1"`
	Street2         interface{} `json:"street2"`
	BaseType        interface{} `json:"@baseType"`
	SchemaLocation  interface{} `json:"@schemaLocation"`
	Type            string      `json:"@type"`
}

type RelatedParty struct {
	Type           string      `json:"@type"`
	ReferredType   string      `json:"@referredType"`
	BaseType       interface{} `json:"@baseType"`
	SchemaLocation interface{} `json:"@schemaLocation"`
	Href           interface{} `json:"href"`
	ID             string      `json:"id"`
	Name           string      `json:"name"`
	Role           string      `json:"role"`
}

type Characteristic struct {
	Type      string `json:"@type"`
	BaseType  string `json:"@baseType"`
	SchemaLoc string `json:"@schemaLocation"`
	Name      string `json:"name"`
	ValueType string `json:"valueType"`
	Value     string `json:"value"`
}
