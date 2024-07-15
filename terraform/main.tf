data "http" "this" {
  url    = "https://swapi.dev/api/films/"
  //url    = "http://127.0.0.1:3000/movies"
  method = "GET"
}

locals {
  titles = [for e in jsondecode(data.http.this.body).results: e.title] 
}

output "all_star_wars_titles" {
  value = local.titles                              
}                                                   
                      
resource "local_file" "all_star_wars_titles" {
  //content = join("\n", jsondecode(data.http.this.body).titles)
  content = join("\n", local.titles)                            
  filename = "${path.module}/all-titles.txt"                    
}
