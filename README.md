# Golang Action Docker
A template repository for writing custom GitHub Actions in Golang with Docker.

## Inputs
| Name | Description | Required | Default |
| --- | --- | --- | --- |
| example-input | An example input to use in your action. | true |  |

## Outputs
No outputs.

## External Actions
No external actions.

## Example Usage
```yaml
- name: Golang Action Docker
  uses: owner/repo@latest
  with:
    # An example input to use in your action.
    example-input:
```
