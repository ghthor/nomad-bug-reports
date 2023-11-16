job "distinct" {
  constraint {
    operator = "distinct_hosts"
    value    = "true"
  }

  group "group" {
    task "task" {
    }
  }
}
