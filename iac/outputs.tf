output "resource_group_name" {
  description = "Created Azure resource group name."
  value       = azurerm_resource_group.main.name
}

output "aks_cluster_name" {
  description = "Created AKS cluster name."
  value       = azurerm_kubernetes_cluster.main.name
}

output "location" {
  description = "Azure region used for the deployment."
  value       = azurerm_resource_group.main.location
}

output "get_credentials_command" {
  description = "Command to fetch kubeconfig after apply."
  value       = "az aks get-credentials --resource-group ${azurerm_resource_group.main.name} --name ${azurerm_kubernetes_cluster.main.name}"
}
