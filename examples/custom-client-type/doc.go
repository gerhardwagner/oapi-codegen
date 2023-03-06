package customclienttype

// This is an example of how to add a prefix to the name of the generated Client struct
// See https://github.com/gerhardwagner/oapi-codegen/issues/785 for why this might be necessary

//go:generate go run github.com/gerhardwagner/oapi-codegen/cmd/oapi-codegen -config cfg.yaml api.yaml
