import type { Test, TestQuery, TestQueryResult } from './models'
import { base64Encode } from "./models";

/*
The OARServiceClient is the primary way of interacting with the oar-service from
the UI.
 */
export class OARServiceClient {
	public baseURL: string;
	public testEndpoint: string;
	public queryEndpoint: string;
	public testsEndpoint: string;

	constructor(baseURL: string) {
		this.baseURL = baseURL;
		if (this.baseURL.endsWith('/')) {
			this.baseURL = this.baseURL.slice(0, -1);
		}

		this.testEndpoint = '/test';
		this.queryEndpoint = '/query';
		this.testsEndpoint = '/tests';
	}

	/*
  addTest will add a new test result via the oar-service. If an error occurs,
  it will be logged to the console and a '-1' testID will be returned

  @param test - Test result to add
  */
	async addTest(test: Test): Promise<number> {
		const requestOptions = {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify(test)
		};

		return fetch(this.baseURL + this.testEndpoint, requestOptions)
			.then((response) => {
				if (!response.ok) {
					console.error('Error occurred when adding test:', response.json());
					return -1;
				}
				return response.json();
			})
			.catch((error) => {
				console.error('Error occurred when adding test:', error);
				return -1;
			});
	}

	/*
	query will send a POST to the `/query` endpoint with a testQuery to get a
	base64 encoded query string to use on the other query endpoints

	@param query - TestQuery to encode into a base64 string
	*/
	async query(query: TestQuery): Promise<string> {
		const requestOptions = {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify(query)
		};

		return fetch(this.baseURL + this.queryEndpoint, requestOptions)
			.then((response) => {
				if (!response.ok) {
					console.error('Error occurred when querying:', response.json());
					return "";
				}
				return response.json();
			})
			.catch((error) => {
				console.error('Error occurred when querying:', error);
				return "";
			});
	}

	/*
	getTests will return the tests that correspond to a TestQuery.

	@param query - TestQuery to return results of. Will be converted into a base64 encoded string, similar to how the
	/query endpoint would.
	@param offset - Offset for query
	@param limit - Results returned limit
	*/
	async getTests(query: TestQuery, offset: number = 0, limit: number = 250): Promise<TestQueryResult> {
		const requestOptions = {
			method: 'GET',
			params: {
				"query": base64Encode(query),
				"offset": offset,
				"limit": limit
			}
		};

		return fetch(this.baseURL + this.testsEndpoint, requestOptions)
			.then((response) => {
				if (!response.ok) {
					console.error('Error occurred when getting tests:', response.json());
					return {"count": 0, "tests": []};
				}
				return response.json();
			})
			.catch((error) => {
				console.error('Error occurred when getting tests:', error);
				return {"count": 0, "tests": []};
			});
	}
}
