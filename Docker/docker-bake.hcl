
variable "REPO" {
  default = "k33g"
}

variable "TAG" {
  default = "0.0.0"
}

group "default" {
  targets = ["chronicles-of-aethelgard"]
}

target "chronicles-of-aethelgard" {
  context = "."
  platforms = [
    "linux/amd64",
    "linux/arm64"
  ]
  tags = ["${REPO}/chronicles-of-aethelgard:${TAG}"]
}

# docker buildx bake --push --file docker-bake.hcl