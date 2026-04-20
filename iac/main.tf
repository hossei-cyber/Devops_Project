locals {
  resource_group_name = "${var.project_name}-rg"
  aks_cluster_name    = "${var.project_name}-aks"
  dns_prefix          = replace(var.project_name, "-", "")
}

resource "azurerm_resource_group" "main" {
  name     = local.resource_group_name
  location = var.location
}

resource "azurerm_kubernetes_cluster" "main" {
  name                = local.aks_cluster_name
  location            = azurerm_resource_group.main.location
  resource_group_name = azurerm_resource_group.main.name
  dns_prefix          = substr(local.dns_prefix, 0, 20)
  kubernetes_version  = var.kubernetes_version
  sku_tier            = "Free"

  default_node_pool {
    name       = "default"
    node_count = var.node_count
    vm_size    = var.vm_size
  }

  identity {
    type = "SystemAssigned"
  }

  tags = {
    project = var.project_name
    owner   = "student"
  }
}
