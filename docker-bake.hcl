group "default" {
  targets = [
    "state-manager",
    "autooperation",
    "dashboard",
    "proxy",
    "seed-data"
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

target "autooperation" {
  dockerfile = "docker/backend/auto-operation/Dockerfile"
  tags = [
    GET_TAG("autooperation")
  ]
}

target "dashboard" {
  dockerfile = "docker/frontend/dashboard/Dockerfile"
  tags = [
    GET_TAG("dashboard")
  ]
  target = "dashboard-runner"
}

target "proxy" {
  dockerfile = "docker/backend/proxy/Dockerfile"
  tags = [
    GET_TAG("proxy")
  ]
  platforms = [
    "linux/amd64",
    "linux/arm64"
  ]
}

target "seed-data" {
  dockerfile = "docker/backend/seed-data/Dockerfile"
  tags = [
    GET_TAG("seed-data")
  ]
  platforms = [
    "linux/amd64",
    "linux/arm64"
  ]
}
