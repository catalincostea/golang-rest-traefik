

1. test
pwd -> superawesome_platform_engineer_test/api

docker build --no-cache -t go-traefik .
<!-- docker build --no-cache -t go-traefik -f Dockerfile-v3 . -->
<!-- docker run -dit -v `pwd`:/app/api -p 3000:3000 --name go-traefik go-traefik -->
docker run -dit -p 3000:3000 --name go-traefik go-traefik
docker exec -it go-traefik bash

go run main.go&
curl localhost:3000/movie

2. 

docker build -t go-traefik .
docker run -d -p 3000:3000 --name go-traefik go-traefik

curl http://localhost:3000/movies

<!-- cleanup -->
docker rm -f go-traefik

3.

docker-compose up -d --build
curl localhost/movies


 <!-- cleanup -->
docker-compose down --remove-orphans
docker-compose rm -f
docker system prune -f
docker volume prune -f
docker network prune -f


4.

docker exec -it go-traefik bash
cd /app/terraform
mv main_local.tf /tmp     # is duplicated for locas tests only


terraform init
terraform apply -auto-approve

Terraform used the selected providers to generate the following execution plan. Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  # local_file.all_star_wars_titles will be created
  + resource "local_file" "all_star_wars_titles" {
      + content              = <<-EOT
            A New Hope
            The Empire Strikes Back
            Return of the Jedi
            The Phantom Menace
            Attack of the Clones
            Revenge of the Sith
        EOT
      + content_base64sha256 = (known after apply)
      + content_base64sha512 = (known after apply)
      + content_md5          = (known after apply)
      + content_sha1         = (known after apply)
      + content_sha256       = (known after apply)
      + content_sha512       = (known after apply)
      + directory_permission = "0777"
      + file_permission      = "0777"
      + filename             = "./all-titles.txt"
      + id                   = (known after apply)
    }

Plan: 1 to add, 0 to change, 0 to destroy.

Changes to Outputs:
  + all_star_wars_titles = [
      + "A New Hope",
      + "The Empire Strikes Back",
      + "Return of the Jedi",
      + "The Phantom Menace",
      + "Attack of the Clones",
      + "Revenge of the Sith",
    ]
local_file.all_star_wars_titles: Creating...
local_file.all_star_wars_titles: Creation complete after 0s [id=cdf09e2bcfc7909223297bc7823c75d0bf7708e1]

Apply complete! Resources: 1 added, 0 changed, 0 destroyed.

Outputs:

all_star_wars_titles = [
  "A New Hope",
  "The Empire Strikes Back",
  "Return of the Jedi",
  "The Phantom Menace",
  "Attack of the Clones",
  "Revenge of the Sith",
]

