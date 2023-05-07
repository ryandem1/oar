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
export type TestQueryResult = {
	count: number;
	tests: Test[];
};

/*
RequestError is what is returned when an error occurs from the oar-service
*/
export type OARServiceError = {
	error: string;
};

/*
EnrichUIError is what is returned when errors occur from the EnrichUI server.
*/
export type EnrichUIError = {
	error: string;
};

export function isOARServiceError(obj: object): obj is OARServiceError {
	return (<OARServiceError>obj).error !== undefined;
}

export function isEnrichUIError(obj: object): obj is EnrichUIError {
	return (<EnrichUIError>obj).error !== undefined;
}
