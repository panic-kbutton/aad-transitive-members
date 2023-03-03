# aad-transitive-members
hack to work around transitive member retrieval limitations in AzureAD provider in terraform, Open Issue: hashicorp/terraform-provider-azuread#881

Usage: supply group ID in command, returns all transitive members in form {"value":"comma,separated,list,of,members"} for easy terraform integration.
