# Example Consul Terraform Provider, Ctmpl for yaml values and GoLang code consuming it
## Consul
This example expects a consul service running on localhost you can easy bring up one like that:
```shell script
docker run --rm -p 8500:8500 -p 8600:8600/udp --name=consul consul:latest agent -server -bootstrap -ui -client=0.0.0.0
```

## Set KV values in Consul via Terraform 
in ./terraform
```shell script
terraform init
terraform apply
```

## Consul-template
```shell script
consul-template --template "config.yaml.ctmpl:config.yaml"
```
Will render config.yaml with content 
```yaml
ip_set_config:
  ipsets:
    name1: id1
    name2: id2
    name3: id3
    name4: id4
```

## Go
will consume config.yaml and treat `ipsets` as map[string]string
```
name2 -> id2
name3 -> id3
name4 -> id4
name1 -> id1
```
