package data

type Agreement struct {
	Type                   string                   `json:"@type" faker:"oneof: Agreement"`
	ID                     int                      `json:"id" faker:"boundary_start=2000000000, boundary_end=2999999999"`
	AgreementType          string                   `json:"agreementType" faker:"oneof: commercial"`
	Description            string                   `json:"description" faker:"oneof: This agreement ..."`
	DocumentNumber         int                      `json:"documentNumber" faker:"boundary_start=300000000, boundary_end=400000000"`
	InitialDate            string                   `json:"initialDate" faker:"timestamp"`
	Name                   string                   `json:"name" faker:"oneof: Digital Subscriber Line 100, Digital Subscriber Line 300"`
	StatementOfIntent      string                   `json:"statementOfIntent" faker:"oneof: Agreement on minimum prices"`
	Status                 string                   `json:"status" faker:"oneof: in process, approved, rejected"`
	Version                string                   `json:"version" faker:"oneof: 1.0, 1.3, 1.5"`
	AgreementAuthorization []AgreementAuthorization `json:"agreementAuthorization"`
	AgreementItem          []AgreementItem          `json:"agreementItem"`
	AgreementPeriod        AgreementPeriod          `json:"agreementPeriod"`
	AgreementSpecification AgreementSpecification   `json:"agreementSpecification"`
	AssociatedAgreement    AssociatedAgreement      `json:"associatedAgreement"`
	CompletionDate         string                   `json:"completionDate" faker:"timestamp"`
	EngagedParty           []RelatedParty           `json:"engagedParty"`
}

type AgreementAuthorization struct {
	Date                    string `json:"date" faker:"timestamp"`
	SignatureRepresentation string `json:"signatureRepresentation" faker:"name"`
	State                   string `json:"state" faker:"oneof: ordered, active, canceled, terminated"`
}

type AgreementItem struct {
	Product         []ProductRef    `json:"product"`
	ProductOffering ProductOffering `json:"productOffering"`
	TermOrCondition TermOrCondition `json:"termOrCondition"`
}

type ProductOffering struct {
	Name string `json:"name" faker:"oneof: Telco Special"`
}

type AssociatedAgreement struct {
	Name string `json:"name" faker:"oneof: General Partnership Agreement"`
}

type TermOrCondition struct {
	Id          string `json:"id" faker:"oneof: 1,2,3"`
	Description string `json:"description" faker:"oneof: delivery should be done in Germany"`
}

type ProductRef struct {
	Type         string `json:"@type" faker:"oneof: ProductRef"`
	ReferredType string `json:"@referredType" faker:"oneof: Product"`
	Name         string `json:"name" faker:"oneof: Digital Subscriber Line 100, Digital Subscriber Line 300"`
}

type AgreementPeriod struct {
	StartDateTime string `json:"startDateTime" faker:"timestamp"`
	EndDateTime   string `json:"endDateTime" faker:"timestamp"`
}

type AgreementSpecification struct {
	Name string `json:"name" faker:"oneof: Moon Agreement Template"`
}
