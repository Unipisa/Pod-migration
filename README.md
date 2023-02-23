# stateful-pod-migration

Repository for Stateful pod migration.

## Local build

### Prerequisites

- [act](https://github.com/nektos/act)
- [docker](https://docs.docker.com/engine/install/ubuntu/)

You can store your secrets on a _secrets.txt_ file or just pass them with the _-s_ flag. The _-j_ flag is used to trigger which job you want to run. 


```Bash
act -j build_and_publish --secret-file ../secrets.txt
```

If you don't want to use act (to run GitHub actions locally) you can just build and push everything using docker.

```Bash
docker build . --tag ghcr.io/leonardopoggiani/virtual-kubelet:latest

docker push ghcr.io/leonardopoggiani/virtual-kubelet:latest
```

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_google"></a> [google](#requirement\_google) | ~> 4.53 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_google"></a> [google](#provider\_google) | ~> 4.53 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [google_compute_firewall.firewall](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_firewall) | resource |
| [google_compute_instance.instance](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_instance) | resource |
| [google_compute_network.network](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_network) | resource |

## Inputs

No inputs.

## Outputs

No outputs.
<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
