# Terraform Hiera 5 Provider

[![pipeline status](https://gitlab.com/sbitio/terraform-provider-hiera5/badges/master/pipeline.svg)](https://gitlab.com/sbitio/terraform-provider-hiera5/-/commits/master) [![coverage report](https://gitlab.com/sbitio/terraform-provider-hiera5/badges/master/coverage.svg)](https://gitlab.com/sbitio/terraform-provider-hiera5/-/commits/master) [![Go Report Card](https://goreportcard.com/badge/gitlab.com/sbitio/terraform-provider-hiera5)](https://goreportcard.com/report/sbitio/terraform-provider-hiera5)

[Hiera5 provider on Terraform's Registry](https://registry.terraform.io/providers/sbitio/hiera5/latest)

This provider implements data sources that can be used to perform hierachical data lookups with Hiera.

This is useful for providing configuration values in an environment with a high level of dimensionality or for making values from an existing Puppet deployment available in Terraform.

It's based on [Terraform hiera provider](https://github.com/ribbybibby/terraform-provider-hiera) and [SmilingNavern's fork](https://github.com/SmilingNavern/terraform-provider-gohiera)


## Development

Provider development takes place on [source repository (on Gitlab)](https://gitlab.com/sbitio/terraform-provider-hiera5)
