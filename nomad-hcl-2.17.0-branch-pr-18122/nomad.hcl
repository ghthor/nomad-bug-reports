job "distinct" {
  constraint {
    distinct_hosts = true
  }

  group "group" {
    task "task" {
    }
  }
}
