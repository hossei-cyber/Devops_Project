variable "project_name" {
  description = "Base name used for Azure resources."
  type        = string
  default     = "webshop-hossay"
}

variable "location" {
  description = "Azure region used for the deployment."
  type        = string
  nullable    = false
}

variable "kubernetes_version" {
  description = "AKS Kubernetes version."
  type        = string
  default     = null
}

variable "node_count" {
  description = "Node count for the AKS default node pool."
  type        = number
  default     = 1
}

variable "vm_size" {
  description = "VM size for the AKS default node pool."
  type        = string
  default = "Standard_B2s_v2"
}
