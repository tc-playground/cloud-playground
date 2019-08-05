# Terraform

## Introduction

* `Terraform` enables you to safely and predictably create, change, and improve infrastructure.

* `Terraform` is a mechanism for providing `infrastructure as code`.

---

## Resources

* `Resources` are central focus.

    * `Resources` have a `type label` - e.g. `instance`, `NAT Gateway`, etc.

    * `Resources` have a `name label`.

    * `Resources` have `arguments` (inputs) - e.g. `instance_type`.

    * `Resources` have `attributes` (outputs) - e.g. `instance_id`.

---

## Providers

* `Providers` provide integration into a specific underlying (IaaS) API.

    * `Providers` have a `name label` that defines the underlying platform: `aws`, `google`, `oci`, etc.

    * `Providers` have `configuration arguments` (inputs) - e.g. `urls`, `environment variables`, `credentials`, etc.

---

## Data Sources

* `Data sources` are elements that allow you to integrate with external services 

* `Data sources` can fetch data to be used elsewhere within Terraform configuration files. e.g. `fetch a list of instance images`.

* `Data sources` have a `type label` - e.g. `aws_ami`.

* `Data sources` have a `name label` - e.g. `ubuntu`.

* `Data sources` have `query constraint arguments` (inputs) - e.g. `filter`.

* `Data sources` have `attributes` (outputs) - e.g. `ami_id`.


---

## Commands

* `terraform init`

* `terraform validate`

* `terraform plan`

* `terraform apply`

* `terraform show`

* `terraform destroy`


---

## References

* [Terraform Home](https://www.terraform.io/)

* [AWS Provider](https://www.terraform.io/docs/providers/aws/index.html)