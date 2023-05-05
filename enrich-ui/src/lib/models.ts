/*
Test represents a single test result from OAR.
 */
export type Test = {
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
export type TestQuery = {
	ids?: number[];
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
base64Encode will encode a testQuery with base 64. This will only work for
ascii characters. Will return the resulting query string, equivalent to calling
the /query endpoint on the oar-service
 */
export function base64Encode(query: TestQuery): string {
	const jsonString = JSON.stringify(query);
	return btoa(jsonString);
}

/*
TestQueryResult is the test results and metadata associated with a test query.
 */
type TestQueryResult = {
	count: number;
	tests: Test[];
};
