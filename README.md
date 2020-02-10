# Terraform Hiera 5 Provider

This provider implements data sources that can be used to perform hierachical data lookups with Hiera.

This is useful for providing configuration values in an environment with a high level of dimensionality or for making values from an existing Puppet deployment available in Terraform.

It's based on [Terraform hiera provider](https://github.com/ribbybibby/terraform-provider-hiera) and [SmilingNavern's fork](https://github.com/SmilingNavern/terraform-provider-gohiera)

## Goals
* Clean implementation based on [Terraform Plugin SDK](https://www.terraform.io/docs/extend/plugin-sdk.html)
* Clean API implementatation based on [Lyra](https://lyraproj.github.io/)'s [Hiera in golang](https://github.com/lyraproj/hiera)
* Painless migration from [Terraform hiera provider](https://github.com/ribbybibby/terraform-provider-hiera), keeping around some naming and data sources

## Requirements
* [Terraform](https://www.terraform.io/downloads.html) 0.12.x

## Usage

### Configuration
To configure the provider:
```hcl
provider "hiera5" {
  # Optional
  config = "~/hiera.yaml"
  # Optional
  scope {
    environment = "live"
    service     = "api"
    # Complex variables are supported using pdialect
    facts       = "{timezone=>'CET'}"
  }
  # Optional
  merge  = "deep"
}
```

### Data Sources
This provider only implements data sources.

#### Hash
To retrieve a hash:
```hcl
data "hiera5_hash" "aws_tags" {
    key = "aws_tags"
}
```
The following output parameters are returned:
* `id` - matches the key
* `key` - the queried key
* `value` - the hash, represented as a map

Terraform doesn't support nested maps or other more complex data structures. Any keys containing nested elements won't be returned.

#### Array
To retrieve an array:
```hcl
data "hiera5_array" "java_opts" {
    key = "java_opts"
}
```
The following output parameters are returned:
* `id` - matches the key
* `key` - the queried key
* `value` - the array (list)

#### Value
To retrieve any other flat value:
```hcl
data "hiera5" "aws_cloudwatch_enable" {
    key = "aws_cloudwatch_enable"
}
```
The following output parameters are returned:
* `id` - matches the key
* `key` - the queried key
* `value` - the value

All values are returned as strings because Terraform doesn't implement other types like int, float or bool. The values will be implicitly converted into the appropriate type depending on usage.

## Develpment

### Requirements

* [Go](https://golang.org/doc/install) 1.12

### Whishlist
* [ ] Support overriding merge strategy in Data Sources
* [ ] Support overriding scope variables in Data Sources
