package accounting

//ContactPerson is a Contact Person
type ContactPerson struct {

	// First name of person
	FirstName string `json:"FirstName,omitempty" xml:"FirstName,omitempty"`

	// Last name of person
	LastName string `json:"LastName,omitempty" xml:"LastName,omitempty"`

	// Email address of person
	EmailAddress string `json:"EmailAddress,omitempty" xml:"EmailAddress,omitempty"`

	// boolean to indicate whether contact should be included on emails with invoices etc.
	IncludeInEmails bool `json:"IncludeInEmails,omitempty" xml:"IncludeInEmails,omitempty"`
}
