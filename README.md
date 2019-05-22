# terraform-checkpoint

## Usage

```hcl-terraform
provider "checkpoint" {
  user = "myuser"
  password = "mypassword"
  host = "https://my.checkpoint.example.com"
}

resource "checkpoint_host" "this" {
  count = 10
  name = "host${count.index+1}"
  ip4_address = "192.168.10.${count.index+1}"
}

```
