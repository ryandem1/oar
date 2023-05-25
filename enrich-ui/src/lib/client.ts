import type { EnrichUIError, OARServiceError, Test, TestQuery, TestQueryResult } from "./models";
import { base64Encode } from "./models";
import { oarServiceBaseURL } from "../stores";


export const getOARServiceBaseURL = (): string => {
	let url: string = "";
	const unsubscribe = oarServiceBaseURL.subscribe((baseURL) => {
		url = baseURL;
	});
	unsubscribe();

	return url;
};

/*
The OARServiceClient is the primary way of interacting with the oar-service from
the UI.
 */
export class OARServiceClient {
	public baseURL: string;
	public testEndpoint: string;
	public queryEndpoint: string;
	public testsEndpoint: string;
	public healthEndpoint: string;

	constructor(baseURL: string = getOARServiceBaseURL()) {
		this.baseURL = baseURL;
		if (this.baseURL.endsWith('/')) {
			this.baseURL = this.baseURL.slice(0, -1);
		}

		this.testEndpoint = '/test';
		this.queryEndpoint = '/query';
		this.testsEndpoint = '/tests';
		this.healthEndpoint = '/health';
	}

	/*
  addTest will add a new test result via the oar-service. If an error occurs,
  it will be logged to the console and a '-1' testID will be returned

  @param test - Test result to add
  */
	async addTest(test: Test): Promise<number | OARServiceError | EnrichUIError> {
		const requestOptions = {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify(test)
		};

		return fetch(this.baseURL + this.testEndpoint, requestOptions)
			.then((response) => {
				if (!response.ok) {
					console.error('Error occurred when adding test:', response.json());
					return response.json();
				}
				return response.json();
			})
			.catch((error) => {
				console.error('Error occurred when adding test:', error);
				return { error: error };
			});
	}

	/*
	Health will return the health status of the oar-service. Will return true if the oar-service is
	healthy, it will return False if it is not.
	*/
	async health(): Promise<boolean> {
		const requestOptions = {
			method: 'GET'
		};

		return fetch(this.baseURL + this.healthEndpoint, requestOptions)
			.then((response) => {
				return response.ok
			})
	}

	/*
	query will send a POST to the `/query` endpoint with a testQuery to get a
	base64 encoded query string to use on the other query endpoints

	@param query - TestQuery to encode into a base64 string
	*/
	async query(query: TestQuery): Promise<string | OARServiceError | EnrichUIError> {
		const requestOptions = {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify(query)
		};

		return fetch(this.baseURL + this.queryEndpoint, requestOptions)
			.then((response) => {
				if (!response.ok) {
					console.error('Error occurred when querying:', response.json());
					return response.json();
				}
				return response.json();
			})
			.catch((error) => {
				console.error('Error occurred when querying:', error);
				return { error: error };
			});
	}

	/*
	getTests will return the tests that correspond to a TestQuery

	@param query - TestQuery to return results of. Will be converted into a base64 encoded string, similar to how the
	/query endpoint would.
	@param offset - Offset for query
	@param limit - Results returned limit
	*/
	async getTests(
		query: TestQuery | null = null,
		offset = 0,
		limit = 250
	): Promise<TestQueryResult | OARServiceError | EnrichUIError> {
		const params: Record<string, string> = {
			offset: offset.toString(),
			limit: limit.toString()
		};
		if (query) {
			params['query'] = base64Encode(query);
		}

		return fetch(this.baseURL + this.testsEndpoint + '?' + new URLSearchParams(params))
			.then((response) => {
				const body = response.json();
				if (!response.ok) {
					console.error('Error occurred when getting tests:', body);
					return body;
				}
				return body;
			})
			.then((testQueryResult: TestQueryResult): TestQueryResult => {
				// This will combine the "doc" attribute of each test into the test
				// itself.
				const { tests, ...remainingQuery } = testQueryResult;
				const mergedTests: Test[] = [];
				testQueryResult.tests.forEach((test) => {
					const { doc, ...testWithoutDoc } = test;
					mergedTests.push({ ...testWithoutDoc, ...(doc as object) });
				});
				return { tests: mergedTests, ...remainingQuery };
			})
			.catch((error) => {
				console.error('Error occurred when getting tests:', error);
				return { error: error };
			});
	}

	/*
	enrichTests will right-merge test details to all tests that match a TestQuery

	@param query - TestQuery to return results of. Will be converted into a base64 encoded string, similar to how the
	/query endpoint would.
	@param offset - Offset for query
	@param limit - Results returned limit
	@return - Status code. 304 means no tests were modified, 200 means at least 1 test was modified
	*/
	async enrichTests(
		test: Test,
		query: TestQuery
	): Promise<number | OARServiceError | EnrichUIError> {
		const requestOptions = {
			method: 'PATCH',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify(test)
		};
		const params = {
			query: base64Encode(query)
		};

		return fetch(
			this.baseURL + this.testsEndpoint + '?' + new URLSearchParams(params),
			requestOptions
		)
			.then((response) => {
				if (!response.ok) {
					const body = response.json();
					console.error('Error occurred when enriching tests:', body);
					return body;
				}
				return response.status;
			})
			.catch((error) => {
				console.error('Error occurred when enriching tests:', error);
				return { error: error };
			});
	}

	/*
	deleteTests will delete all tests that match a TestQuery

	@param query - TestQuery to return results of. Will be converted into a base64 encoded string, similar to how the
	/query endpoint would.
	@return - Status code. 304 means no tests were deleted, 200 means at least 1 test was deleted
	*/
	async deleteTests(query: TestQuery): Promise<number | OARServiceError | EnrichUIError> {
		const requestOptions = {
			method: 'DELETE'
		};
		const params = {
			query: base64Encode(query)
		};

		return fetch(
			this.baseURL + this.testsEndpoint + '?' + new URLSearchParams(params),
			requestOptions
		)
			.then((response) => {
				if (!response.ok) {
					console.error('Error occurred when deleting tests:', response.json());
					return response.json();
				}
				return response.status;
			})
			.catch((error) => {
				console.error('Error occurred when deleting tests:', error);
				return { error: error };
			});
	}
}
