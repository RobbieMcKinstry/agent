package release

type DatabaseRelease struct {
  version Version
  flavor Flavor
}

func (DatabaseRelease) Eq(DatabaseRelease) bool {
  panic("Not implemented yet.")
}
