import { describe, expect, it } from 'vitest';
import { OAR_SERVICE_BASE_URL } from '$env/static/private';
import { fakeTests, selectRandomItem } from './faker';
import { OARServiceClient } from './client';

const client = new OARServiceClient(OAR_SERVICE_BASE_URL);

describe('The oar-service client', () => {
	it('can be initialized', () => {
		new OARServiceClient(OAR_SERVICE_BASE_URL);
	});

	it('can be initialized if base url ends with /', () => {
		new OARServiceClient(OAR_SERVICE_BASE_URL + '/');
	});

	it('can add a test result', async () => {
		const test = selectRandomItem(fakeTests);
		const testID = await client.addTest(test);
		expect(testID).toBeGreaterThan(0);
	});
});
