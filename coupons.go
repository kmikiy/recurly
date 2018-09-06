package recurly

import "encoding/xml"

// Coupon represents an individual coupon on your site.
type Coupon struct {
	XMLName                  xml.Name          `xml:"coupon"`
	Code                     string            `xml:"coupon_code"`
	Name                     string            `xml:"name"`
	Description              string            `xml:"description,omitempty"`
	DiscountType             string            `xml:"discount_type"`
	DiscountInCents          *DiscountInCents  `xml:"discount_in_cents,omitempty"`
	DiscountPercent          int               `xml:"discount_percent,omitempty"`
	InvoiceDescription       string            `xml:"invoice_description,omitempty"`
	RedeemByDate             NullTime          `xml:"redeem_by_date,omitempty"`
	MaxRedemptions           NullInt           `xml:"max_redemptions,omitempty"`
	AppliesToAllPlans        NullBool          `xml:"applies_to_all_plans,omitempty"`
	Duration                 string            `xml:"duration,omitempty"`
	TemporalUnit             string            `xml:"temporal_unit,omitempty"`
	TemporalAmount           NullInt           `xml:"temporal_amount,omitempty"`
	AppliesToNonPlanCharges  NullBool          `xml:"applies_to_non_plan_charges,omitempty"`
	RedemptionResource       string            `xml:"redemption_resource,omitempty"`
	MaxRedemptionsPerAccount NullInt           `xml:"max_redemptions_per_account,omitempty"`
	CouponType               string            `xml:"coupon_type,omitempty"`
	UniqueCodeTemplate       string            `xml:"unique_code_template,omitempty"`
	PlanCodes                *[]CouponPlanCode `xml:"plan_codes>plan_code,omitempty"`
	FreeTrialAmount          NullInt           `xml:"free_trial_amount,omitempty"`
	FreeTrialUnit            string            `xml:"free_trial_unit,omitempty"`
	CreatedAt                NullTime          `xml:"created_at,omitempty"`
	State                    string            `xml:"state,omitempty"`

	// Deprecated: SingleUse          NullBool          `xml:"single_use,omitempty"`
	// Deprecated: AppliesForMonths   NullInt           `xml:"applies_for_months,omitempty"`
}

// CouponPlanCode holds an xml array of plan_code items that this coupon
// will work with.
type CouponPlanCode struct {
	Code string `xml:",innerxml"`
}

type DiscountInCents struct {
	USD int `xml:"USD,omitempty"`
}
