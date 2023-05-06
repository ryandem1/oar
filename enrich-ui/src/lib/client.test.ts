import { describe, expect, it } from 'vitest';
import { OAR_SERVICE_BASE_URL } from '$env/static/private';
import { fakeTests, selectRandomItem } from './faker';
import { OARServiceClient } from './client';

describe.concurrent('The oar-service client', () => {
	it('can be initialized', () => {
		new OARServiceClient(OAR_SERVICE_BASE_URL);
	});

	it('can be initialized if base url ends with /', () => {
		new OARServiceClient(OAR_SERVICE_BASE_URL + '/');
	});

	it('can add a test result', async () => {
		const client = new OARServiceClient(OAR_SERVICE_BASE_URL);

		const test = selectRandomItem(fakeTests);
		const testID = await client.addTest(test);
		expect(testID).toBeGreaterThan(0);
	});

	it('can handle response errors when adding a result', async () => {
		const client = new OARServiceClient(OAR_SERVICE_BASE_URL);
		client.testEndpoint = "/test/bad_response"  // Triggers the mock

		const test = selectRandomItem(fakeTests);
		const testID = await client.addTest(test);
		expect(testID).toBe(-1);
	});

	it('can handle exceptions without crashing when adding a result', async () => {
		const client = new OARServiceClient(OAR_SERVICE_BASE_URL);
		client.testEndpoint = "/test/exception"  // Triggers the mock

		const test = selectRandomItem(fakeTests);
		const testID = await client.addTest(test);
		expect(testID).toBe(-1);
	});

	it('can obtain a query string', async() => {
		const client = new OARServiceClient(OAR_SERVICE_BASE_URL);

		const query = {
			"ids": [1, 2, 3, 4]
		}
		const queryString = await client.query(query);
		expect(queryString).toBeTruthy()
	})

	it('can handle response errors when querying', async () => {
		const client = new OARServiceClient(OAR_SERVICE_BASE_URL);
		client.queryEndpoint = "/query/bad_response"  // Triggers the mock

		const query = {
			"ids": [1, 2, 3, 4]
		}
		const queryString = await client.query(query);
		expect(queryString).toBeFalsy()
	});

	it('can handle exceptions without crashing when querying', async () => {
		const client = new OARServiceClient(OAR_SERVICE_BASE_URL);
		client.queryEndpoint = "/query/exception"  // Triggers the mock

		const query = {
			"ids": [1, 2, 3, 4]
		}
		const queryString = await client.query(query);
		expect(queryString).toBeFalsy()
	});
});