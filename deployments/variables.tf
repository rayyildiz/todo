variable "projectId" {
  sensitive = true
}
variable "region" {
  sensitive = true
}
variable "container_port" {
  default = "4000"
}
variable "name" {
  default = "demoapp"
}
variable "backend_image" {
  type      = string
  sensitive = true
}
variable "clean_image" {
  type      = string
  sensitive = true
}
variable "ui_domain" {
  type      = string
  sensitive = true
}

variable "iam_service_invoker_mail" {
  type      = string
  sensitive = true
}
