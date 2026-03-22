variable "db_url" {
  type    = string
  default = getenv("DB_URL")
}

variable "test_db" {
  type    = string
  default = getenv("TEST_DB")
}

env "local" {
  
  src = "file://internal/migrations/schema"

  migration {
    dir = "file://internal/migrations/scripts"
    revisions_schema = "atlas_schema_revisions"
  }

  schemas = ["public"]

  dev = var.test_db

  url = var.db_url

  # Formatting: Ensures generated SQL is pretty-printed
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }

  # Safety Guardrails
  lint {
    
    destructive {
      error = false 
    }
    
    incompatible {
      error = false
    }

    naming {
      match   = "^[a-z_]+$"
      message = "Table and column names must be lowercase snake_case"
    }
  }
}