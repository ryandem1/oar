import { describe, expect, it } from 'vitest';
import { OAR_SERVICE_BASE_URL } from '$env/static/private';
import { fakeTests, selectRandomItem } from '$lib/faker';
import { OARServiceClient } from '$lib/client';
import { isError } from '$lib/models';

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
		client.testEndpoint = '/test/bad_response'; // Triggers the mock

		const test = selectRandomItem(fakeTests);
		const response = await client.addTest(test);
		if (typeof response === 'number') {
			throw new Error('Returned valid response!');
		}
		expect(response.error).toBeTruthy();
	});

	it('can handle exceptions without crashing when adding a result', async () => {
		const client = new OARServiceClient(OAR_SERVICE_BASE_URL);
		client.testEndpoint = '/test/exception'; // Triggers the mock

		const test = selectRandomItem(fakeTests);
		const response = await client.addTest(test);
		if (typeof response === 'number') {
			throw new Error('Returned valid response!');
		}
		expect(response.error).toBeTruthy();
	});

	it('can obtain a query string', async () => {
		const client = new OARServiceClient(OAR_SERVICE_BASE_URL);

		const query = {
			ids: [1, 2, 3, 4]
		};
		const queryString = await client.query(query);
		expect(queryString).toBeTruthy();
	});

	it('can handle response errors when querying', async () => {
		const client = new OARServiceClient(OAR_SERVICE_BASE_URL);
		client.queryEndpoint = '/query/bad_response'; // Triggers the mock

		const query = {
			ids: [1, 2, 3, 4]
		};
		const response = await client.query(query);
		if (typeof response === 'string') {
			throw new Error('Returned valid response!');
		}
		expect(response.error).toBeTruthy();
	});

	it('can handle exceptions without crashing when querying', async () => {
		const client = new OARServiceClient(OAR_SERVICE_BASE_URL);
		client.queryEndpoint = '/query/exception'; // Triggers the mock

		const query = {
			ids: [1, 2, 3, 4]
		};
		const response = await client.query(query);
		if (typeof response === 'string') {
			throw new Error('Returned valid response!');
		}
		expect(response.error).toBeTruthy();
	});

	it('can get tests via a query', async () => {
		const client = new OARServiceClient(OAR_SERVICE_BASE_URL);
		client.queryEndpoint = '/tests'; // Triggers the mock

		const query = {
			ids: [1]
		};
		const queryResults = await client.getTests(query);
		if (!isError(queryResults)) {
			expect(queryResults.count).toBe(1);
			expect(queryResults.tests.length).toBe(1);
		}
	});

	it('can handle response errors when getting tests', async () => {
		const client = new OARServiceClient(OAR_SERVICE_BASE_URL);
		client.testsEndpoint = '/tests/bad_response'; // Triggers the mock

		const query = {
			ids: [1]
		};
		const queryResults = await client.getTests(query);

		expect(isError(queryResults)).toBeTruthy();
	});

	it('can handle exceptions when getting tests', async () => {
		const client = new OARServiceClient(OAR_SERVICE_BASE_URL);
		client.testsEndpoint = '/tests/exception'; // Triggers the mock

		const query = {
			ids: [1]
		};
		const queryResults = await client.getTests(query);
		expect(isError(queryResults)).toBeTruthy();
	});

	it('can enrich tests via a query', async () => {
		const client = new OARServiceClient(OAR_SERVICE_BASE_URL);
		client.queryEndpoint = '/tests'; // Triggers the mock

		const query = {
			ids: [1]
		};
		const test = selectRandomItem(fakeTests);
		const response = await client.enrichTests(test, query);

		expect(response).toBe(200);
	});

	it('can handle response errors when getting tests', async () => {
		const client = new OARServiceClient(OAR_SERVICE_BASE_URL);
		client.testsEndpoint = '/tests/bad_response'; // Triggers the mock

		const query = {
			ids: [1]
		};
		const test = selectRandomItem(fakeTests);
		const response = await client.enrichTests(test, query);

		if (typeof response === 'number') {
			throw new Error('Returned valid response!');
		}
		expect(response?.error).toBeTruthy();
	});

	it('can handle exceptions when getting tests', async () => {
		const client = new OARServiceClient(OAR_SERVICE_BASE_URL);
		client.testsEndpoint = '/tests/exception'; // Triggers the mock

		const query = {
			ids: [1]
		};
		const test = selectRandomItem(fakeTests);
		const response = await client.enrichTests(test, query);

		if (typeof response === 'number') {
			throw new Error('Returned valid response!');
		}
		expect(response?.error).toBeTruthy();
	});

	it('can delete tests via a query', async () => {
		const client = new OARServiceClient(OAR_SERVICE_BASE_URL);
		client.queryEndpoint = '/tests'; // Triggers the mock

		const query = {
			ids: [1]
		};
		const response = await client.deleteTests(query);

		expect(response).toBe(200);
	});

	it('can handle response errors when deleting tests', async () => {
		const client = new OARServiceClient(OAR_SERVICE_BASE_URL);
		client.testsEndpoint = '/tests/bad_response'; // Triggers the mock

		const query = {
			ids: [1]
		};
		const response = await client.deleteTests(query);

		if (typeof response === 'number') {
			throw new Error('Returned valid response!');
		}
		expect(response?.error).toBeTruthy();
	});

	it('can handle exceptions when deleting tests', async () => {
		const client = new OARServiceClient(OAR_SERVICE_BASE_URL);
		client.testsEndpoint = '/tests/exception'; // Triggers the mock

		const query = {
			ids: [1]
		};

		const response = await client.deleteTests(query);

		if (typeof response === 'number') {
			throw new Error('Returned valid response!');
		}
		expect(response?.error).toBeTruthy();
	});
});
