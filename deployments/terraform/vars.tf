variable "token" {
  type = string
}

variable "project-name" {
  type = string
  default = "todo-app-dev"
}

variable "environment" {
  type = string
  validation {
    condition = contains(["development", "staging", "production"], self)
  }
  default = "development"
}
