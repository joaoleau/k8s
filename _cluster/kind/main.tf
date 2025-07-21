resource "kind_cluster" "kind" {
  name            = "kind"
  node_image      = "kindest/node:v1.32.0"
  kubeconfig_path = pathexpand("~/.kube/config")

  kind_config {
    kind        = "Cluster"
    api_version = "kind.x-k8s.io/v1alpha4"

    node {
      role = "control-plane"
      extra_port_mappings {
        container_port = 6443
        host_port      = 6443
      }
    }

    node {
      role = "worker"
    }

    node {
      role = "worker"
    }
  }
}
