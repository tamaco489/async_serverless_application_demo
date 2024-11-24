variable "mailer_config" {
  type = map(string)
  default = {
    username  = "hoge@example.com"
    password  = "password"
    to_addr_1 = "fuga1@example.com"
    to_addr_2 = "fuga2@example.com"
  }
}
