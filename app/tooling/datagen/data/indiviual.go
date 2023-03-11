package data

type Individual struct {
	Type               string            `json:"@type" faker:"oneof: Individual"`
	ID                 int               `json:"id" faker:"boundary_start=100000000, boundary_end=999999999"`
	BirthDate          string            `json:"birthDate" faker:"timestamp"`
	CountryOfBirth     string            `json:"countryOfBirth" faker:"oneof: Deutschland"`
	FamilyName         string            `json:"familyName" faker:"last_name"`
	FormattedName      string            `json:"formattedName" faker:"name"`
	FullName           string            `json:"fullName" faker:"name"`
	Gender             string            `json:"gender" faker:"gender"`
	GivenName          string            `json:"givenName" faker:"first_name"`
	LanguageAbility    []LanguageAbility `json:"languageAbility"`
	PreferredGivenName string            `json:"preferredGivenName" faker:"first_name"`
	Status             string            `json:"status" faker:"oneof: initialized, validated, deceased"`
}

type LanguageAbility struct {
	Type                string   `json:"@type" faker:"oneof: LanguageAbility"`
	IsFavouriteLanguage bool     `json:"isFavouriteLanguage"`
	LanguageCode        string   `json:"languageCode" faker:"oneof: de"`
	LanguageName        string   `json:"languageName" faker:"oneof: German"`
	ValidFor            ValidFor `json:"validFor" faker:"validFor"`
}
