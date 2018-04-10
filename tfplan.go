package main

import (
	"github.com/hashicorp/terraform/terraform"
)

type root struct {
	tfplan *terraform.Plan
}

func (r *root) Configure(raw map[string]interface{}) error {
	if _, ok := raw["tfplan"]; !ok {

	}
}
