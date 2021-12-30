package catalog

import "github.com/gocql/gocql"

type CatalogEavAttribute struct {
	AttributeID               gocql.UUID `json:"attribute_id"`
	FrontendInputRenderer     string     `json:"frontend_input_renderer"`
	IsGlobal                  bool       `json:"is_global"`
	IsVisible                 bool       `json:"is_visible"`
	IsSearchable              bool       `json:"is_searchable"`
	IsFilterable              bool       `json:"is_filterable"`
	IsComparable              bool       `json:"is_comparable"`
	IsVisibleOnFront          bool       `json:"is_visible_on_front"`
	IsHtmlAllowedOnFront      bool       `json:"is_html_allowed_on_front"`
	IsUsedForPriceRules       bool       `json:"is_used_for_price_rules"`
	IsFilterableInSearch      bool       `json:"is_filterable_in_search"`
	UsedInProductListing      bool       `json:"used_in_product_listing"`
	UsedForSortBy             bool       `json:"used_for_sort_by"`
	ApplyTo                   string     `json:"apply_to"`
	IsVisibleInAdvancedSearch bool       `json:"is_visible_in_advanced_search"`
	Position                  int        `json:"position"`
	IsUsedForPromoRules       bool       `json:"is_used_for_promo_rules"`
	IsRequiredInAdminStore    bool       `json:"is_required_in_admin_store"`
	IsUsedInGrid              bool       `json:"is_used_in_grid"`
	IsVisibleInGrid           bool       `json:"is_visible_in_grid"`
	IsFilterableInGrid        bool       `json:"is_filterable_in_grid"`
	SearchWeight              int        `json:"search_weight"`
	AdditionalData            string     `json:"additional_data"`
}
