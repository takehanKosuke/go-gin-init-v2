resource "google_sql_database_instance" "go-gin-init-v2-db" {
  name             = var.db
  database_version = "MYSQL_8_0"
  region           = var.region

  settings {
    tier = "db-g1-small"
  }
}

resource "google_sql_user" "go-gin-init-v2-user" {
  instance = google_sql_database_instance.go-gin-init-v2-db.name
  name     = "root"
  password = "password"
}

resource "google_sql_database" "go-gin-init-v2-db" {
  instance = google_sql_database_instance.go-gin-init-v2-db.name
  name     = "go-gin-init-v2-db"
  charset  = "utf8mb4"
}
