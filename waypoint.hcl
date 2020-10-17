project = "example-go-noweb"

app "example-go-noweb" {
  build {
    use "pack" {}
  }

  deploy { 
    use "docker" {}
  }
}
