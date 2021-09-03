package cli

type DatabaseFlavor uint8

const (
  MySQL DatabaseFlavor = iota
  Postgres
)

func (flavor DatabaseFlavor) String() string {
  switch flavor {
    case MySQL:
      return "MySQL"
    case Postgres:
      return "Postgres"
  }
}
