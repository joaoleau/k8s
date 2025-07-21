output "kubectl_set_context" {
  value = "kubectl config set-cluster kind-${kind_cluster.kind.name} --server=https://127.0.0.1:6443"
  description = "Comando kubectl para configurar o contexto"
}
