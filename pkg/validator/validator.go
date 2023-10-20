package validator

import "github.com/go-playground/validator/v10"

// Validator is a singleton that allows validating models.
// Usage: Validator.Struct(s) with s being a struct object.
var Validator = validator.New()
