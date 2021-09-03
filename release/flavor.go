package release
// A Flavor is a particular kind of DBMS e.g. Postgres or MySQL.
// Each flavor must provide a method connect to it,
// a method to collect knobs, and a method to collect metrics.
type Flavor interface {
  // func InitializeConnection(ConnectionParameters) error
}

struct MySQLFlavor {}
struct PostgresFlavor {}

var _ Flavor = &MySQLFlavor{}
var _ Flavor = &PostgresFlavor{}

// TODO: This API belies the need for the following structs:
// A ConnectionConf struct for connection information.
// Either a struct representing a list of queries, or a mutable
// transaction type which gets passed into the knob/metric collection
// methods.
// Question: Do we want to represent queries opaquely using a
// struct? That way we can test the queries we generate?
