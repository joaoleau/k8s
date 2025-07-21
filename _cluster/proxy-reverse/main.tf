resource "docker_container" "nginx" {
  name  = "nginx"
  image = "nginx:latest"

  ports {
    internal = 80
    external = 80
  }

  networks_advanced {
    name = "kind"
  }

  volumes {
    host_path      = abspath("${path.module}/nginx.conf")
    container_path = "/etc/nginx/nginx.conf"
  }

  restart = "unless-stopped"
}
