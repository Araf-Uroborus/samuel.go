# Kubernetes Cheatsheet

Lista de comandos Kubernetes
![](https://murchie85.github.io/images/kubernetes/kub.jpg)

# Client Configuration

- Ver a configuração do Kubernetes

  `kubectl config view`

- Listar os contextos existentes para clusters Kubernetes

  `kubectl config get-contexts`

# Gerenciar Recursos

- Criar recursos como pods, replicasets, etc.

  `kubectl create -f ./meu-arquivo.yaml`

- Aplicar a configuração a um recurso

  `kubectl apply -f ./meu-arquivo.yaml`

- Iniciar um pod/única instância (com nginx por exmeplo)

  `kubectl run nginx --image=nginx`

- Deletar um recurso (pelo arquivo yaml)

  `kubectl delete -f ./meu-arquivo.yaml`

- Deletar um recurso (pelo nome do recurso)

  `kubectl delete <tipo-de-recurso> <nome-do-recurso>`
  ex:
  `kubectl delete pod nginx`

- Deletar pods baseado em labels (todos os pods com label tipo=servidorweb)

  `kubectl delete pods -l tipo=servidorweb`

# Ver, Encontrar, Listar recursos

- Listar todos os serviços no namespace default

  `kubectl get services`

- Listar todos os pods em todos os namespaces

  `kubectl get pods --all-namespaces`

- Listar todos os pods com visão detalhada (no namespace atual)

  `kubectl get pods -o wide`

- Listar todos os recursos ao mesmo tempo (no namespace atual)

  `kuebctl get all`

- Listar todos os recursos ao mesmo tempo com visão detalhada (no namespace atual)

  `kubectl get all -o wide`

- Listar todos os pods em formato json (ou yaml) (no namespace atual)

  `kubectl get pods -o json`
  `kubectl get pods --o yaml`

- Listar todos os replication controllers (no namespace atual)

  `kubectl get replicationcontroller`
  ou
  `kubectl get rc`

- Listar todos os replicasets (no namespace atual)

  `kubectl get replicaset`
  ou
  `kubectl get rs`

- Listar todos os nodes

  `kubectl get nodes`

- Descrever os detalhes de um recurso (node, pod, svc)

  `kubectl describe nodes meu-node`
  `kubectl describe pod nginx-pod`
  `kubectl describe svc nginx-svc`

- Listar pods por ordem de nome

  `kubectl get pods --sort-by=.metadata.name`

- Listar pods por ordem de contagem de restarts

  `kubectl get pods --sort-by='.status.containerStatuses[0].restartCount'`

- Escalonar ou escalar (Scale) um replicaset chamado 'rs-nginx' para 3

  `kubectl scale --replicas=3 rs/rs-nginx`
  ou
  `kubectl scale --replicas=3 replicaset rs-nginx`

- Escalonar ou escalar (Scale) um recurso especificado no arquivo 'rs-nginx'.yaml para 3

  `kubectl scale --replicas=3 -f rs-nginx.yaml`

- Apenas ver a configuração yaml para um pod nginx

  `kubectl run --image=nginx nginx --dry-run=client -o yaml`

# Criar arquivos yaml

- Arquivo yaml para um pod nginx

  `kubectl run --image=nginx nginx --dry=run=client -o yaml > nginx-pod.yml`

- Arquivo yaml para um deployment nginx com 3 replicas

  `kubectl create deployment --image=nginx nginx --replicas=3 --dry=run=client -o yaml > nginx-pod.yml`

# Kind - Cluster Local com Kind

- Listar containers Docker rodando

  `docker ps`

- Criando um cluster básico com 1 node

  `kind create cluster --name meu-cluster`

- Verificando informações do cluster

  `kubectl cluster-info --context kind-meu-cluster`

- Confirmando que estamos no contexto do nosso cluster kind

  `kubectl config get-contexts`

- Listar clusters do kind

  `kind get clusters`

- Deletando um cluster no kind

  `kind delete cluster --name meu-cluster`

- Criando um cluster baseado em arquivo yaml

  `kind create cluster --name meu-cluster --config meu-arquivo.yml`

- Exemplo de configuração para cluster kind com 3 nodes

  ```
  kind: Cluster
  apiVersion: kind.x-k8s.io/v1alpha4
  nodes:
  - role: control-plane
  - role: worker
  - role: worker
  ```
