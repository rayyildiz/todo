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
  default = "todo-backend"
}
variable "image" {
  default   = "rayyildiz/todo"
  sensitive = true
}

variable "ui_domain" {
  type      = string
  sensitive = true
}
