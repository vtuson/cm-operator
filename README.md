# cm-operator

This is a proof concept for a chartmuseum operator, that uses a [sidecar](https://github.com/vtuson/cmsidecar).

Once deployed it will clone, pack and publish charts store in a public git

## adding a repo
You can add a repo by creating a Chartmuseum object:
```
apiVersion: "cm.bitnami.com/v1alpha1"
kind: "Chartmuseum"
metadata:
  name: "myrepo"
spec:
  git: "https://github.com/foo/mycharts"
  dependencies:
  - name: bitnami
    url: "https://charts.bitnami.com/bitnami"
    
```
This will make it available in the deployment of chartmuseum as /myrepo/index.yaml

## TODO
* add support for git credential for private repos
