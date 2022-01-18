Queries:
- How do the docker-compose, k8s.yml and okteto.yml files relate? I guess I'll understand better the more I learn about kubernetes. Here's what I understand:
  - okteto.yml defines the sync between the kubernetes cluster (which has its configuration defined in k8s.yaml) at okteto cloud and the local filesystem.
  - docker-compose.yaml lets you skip the k8s.yaml manifest but with different commands (```okteto stack deploy --[wait|build]```)

Personal:
- Since the project is small I did not use git submodules because it seemed overkill. 

Details:
- The okteto manifest wasn't included in the repository.