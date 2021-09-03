// TODO
// A MetadataStore provdes an opaque store which a client can
// query for information about what knobs and metrics are
// avaiable to a given database release.
// A Database release is a Flavor/Version pair.
type MetadataStore interface {
  // This should be an interface which takes
  // Flavor/Version pair and returns a list of knobs
  // and a list of Metrics.
  Knobs(DatabaseRelease) KnobQuery
  Metrics(DatabaseRelease) MetricQuery
  // Concretely, this can either make an API request to
  // determine which knobs and metrics are needed, OR
  // it can be a hardcoded Map that maps knobs and metrics
  // to function closures?
}
// TODO: Maybe there should be an intermediate layer here.
// Maybe this should return the format/structure of the knobs,
// and another function converts that into a KnobQuery / MetricQuery.

// KnobQuery is a function which knows how to read knobs from
// a database.
type KnobQuery func(context.Context, *sql.Tx) error
// MetricQuery is a function which knows how to read metrics
// from a database.
type MetricQuery func(context.Context, *sql.Tx) error
