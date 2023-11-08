group "default" {
  targets = [
    "state-manager"
  ]
}

variable "PREFIX" {
  default = "plarail2023"
}

variable "TAG" {
  default = "latest"
}

function "GET_TAG" {
  params = [image]
  result = "ghcr.io/ueckoken/${PREFIX}-${image}:${TAG}"
}

target "state-manager" {
    dockerfile = "docker/backend/state-manager/Dockerfile"
    tags = [
        GET_TAG("state-manager")
    ]
    platforms = [
        "linux/amd64",
        "linux/arm64"
    ]
}