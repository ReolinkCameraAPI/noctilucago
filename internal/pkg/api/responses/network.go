package responses

import "github.com/ReolinkCameraAPI/noctilucago/internal/pkg/database/models"

// NetworkProxyResponse is used when creating and updating a proxy setting
// swagger:response networkProxyResponse
type NetworkProxyResponse struct {
	// in: body
	Body *models.Proxy
}

// NetworkProxyArrayResponse is used when returning multiple proxy settings
// swagger:response networkProxyArrayResponse
type NetworkProxyArrayResponse struct {
	// in: body
	Body []*models.Proxy
}

// NetworkProxySchemeResponse is used when returning the supported server schemes
// swagger:response networkProxySchemeResponse
type NetworkProxySchemeResponse struct {
	// in: body
	Body struct {
		Schemes []string `json:"schemes"`
	}
}

// NetworkProtocolResponse is used when returning the supported server protocols
// swagger:response networkProtocolResponse
type NetworkProtocolResponse struct {
	// in: body
	Body struct {
		Protocols []string `json:"protocols"`
	}
}
