# kustomize-interpolate-variables

Kustomize plugin to interpolate variables using $VAR syntax. It also supports using Go templates.

## Usage

Take a look at the example project [here](https://github.com/oscartbeaumont/kustomize-interpolate-variables/tree/main/example).

You can run the example project using the following commands:

```bash
cd example/
kubectl kustomize . --enable-alpha-plugins > output.yaml
# You could replace `output.yaml` with `kubectl apply -f -` to run the configuration against a Kubernetes cluster.
```
