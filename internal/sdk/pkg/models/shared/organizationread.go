// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type OrganizationRead struct {
	OrganizationID   string  `json:"organizationId"`
	OrganizationName string  `json:"organizationName"`
	Email            string  `json:"email"`
	Pba              bool    `json:"pba"`
	OrgLevelBilling  bool    `json:"orgLevelBilling"`
	SsoRealm         *string `json:"ssoRealm,omitempty"`
}

func (o *OrganizationRead) GetOrganizationID() string {
	if o == nil {
		return ""
	}
	return o.OrganizationID
}

func (o *OrganizationRead) GetOrganizationName() string {
	if o == nil {
		return ""
	}
	return o.OrganizationName
}

func (o *OrganizationRead) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *OrganizationRead) GetPba() bool {
	if o == nil {
		return false
	}
	return o.Pba
}

func (o *OrganizationRead) GetOrgLevelBilling() bool {
	if o == nil {
		return false
	}
	return o.OrgLevelBilling
}

func (o *OrganizationRead) GetSsoRealm() *string {
	if o == nil {
		return nil
	}
	return o.SsoRealm
}