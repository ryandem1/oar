/*
Test represents a single test result from OAR.
 */
type Test = {
	id: number;
	summary: string;
	outcome: Outcome | string;
	analysis: Analysis | string;
	resolution: Resolution | string;
	[x: string]: unknown; // Allows for arbitrary properties
};

/*
TestQuery represents the query structure for searching for OAR tests.
 */
type TestQuery = {
	ids?: bigint[];
	summaries?: string[];
	outcomes?: Outcome[];
	analyses?: Analysis[];
	resolutions?: Resolution[];
	createdBefore?: Date;
	createdAfter?: Date;
	modifiedBefore?: Date;
	modifiedAfter?: Date;
	docs?: object[];
};

/*
TestQueryResult is the test results and metadata associated with a test query.
 */
type TestQueryResult = {
	count: number;
	tests: Test[];
};
