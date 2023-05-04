/*
Test represents a single test result from OAR.
 */
type Test = {
  id: bigint,
  summary: string,
  outcome: Outcome,
  analysis: Analysis,
  resolution: Resolution,
  [x: string]: unknown  // Allows for arbitrary properties
}

type TestQuery = {

}
