provider "consul" {
  address    = "127.0.0.1:8500"
  datacenter = "test"
}

resource "consul_key_prefix" "test" {
  path_prefix = "autobahn/"

  subkeys = {
    "ipsets" = <<-EOF
    name1: id1
    name2: id2
    name3: id3
    name4: id4
    EOF
  }
}
