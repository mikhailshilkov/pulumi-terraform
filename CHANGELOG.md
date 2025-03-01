# `@pulumi/terraform` CHANGELOG

This CHANGELOG details important changes made in each version of the
`terraform` provider, and the `@pulumi/terraform` Node.js package.

## v0.18.4 (Unreleased)

- Terraform-based providers can now communicate detailed information about the difference between a resource's desired and actual state during a Pulumi update.
- Add the ability to inject CustomTimeouts into the InstanceDiff during a pulumi update.

## v0.18.3 (Released June 20, 2019)

- Fixed a bug that caused unnecessary changes if the first operation after upgrading a bridged provider was a `pulumi refresh`.
- Fixed a bug that caused maps with keys containing a '.' character to be incorrectly treated as containing nested maps when deserializing Terraform attributes.

### Improvements

- Automatically generate `isInstance` type guards for implementations of `Resource`.
- `TransformJSONDocument` now accepts arrays (in addition to maps).

## v0.18.2 (Released May 28th, 2019)

- Improved the package `README` file to reflect usage of the `@pulumi/terraform`
  package rather than the Terraform bridge.

## v0.18.1 (Released May 16th, 2019)

- Initial release of `@pulumi/terraform` with support for Node.js.
