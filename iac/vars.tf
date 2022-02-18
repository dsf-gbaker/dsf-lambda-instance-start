## Tags
variable "tag-owner" {
  default = "beerskunk"
  type    = string
}

variable "tag-project-name" {
  default = "dsforgeworks"
  type    = string
}

variable "tag-environment" {
  default = "prod"
  type    = string
}

## Region
variable "region" {
  default = "us-east-1"
  type    = string
}

## Project Variables
variable "filename" {
  default = "main.zip"
  type    = string
}