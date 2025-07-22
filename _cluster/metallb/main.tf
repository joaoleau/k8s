resource "helm_release" "metallb" {
  name             = "metallb"
  repository       = "https://metallb.github.io/metallb"
  chart            = "metallb"
  version          = "0.13.10"
  namespace        = "metallb"
  create_namespace = true
}
