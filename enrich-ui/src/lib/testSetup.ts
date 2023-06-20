import { afterAll, afterEach, beforeAll } from 'vitest';
import { setupServer } from 'msw/node';
import { rest } from 'msw';
import { fakeTests, selectRandomItem } from './faker';
import { base64Encode } from './models';
import { oarServiceBaseURL } from '../stores';

const testOARBaseURL = 'http://localhost:8080';
export const restHandlers = [
	rest.post(testOARBaseURL + '/test', (_, res, ctx) => {
		const test = selectRandomItem(fakeTests);
		return res(ctx.status(200), ctx.json(test.id));
	}),

	rest.post(testOARBaseURL + '/test/bad_response', (_, res, ctx) => {
		return res(ctx.status(400), ctx.json({ error: 'an error occured when creating a test' }));
	}),

	rest.post(testOARBaseURL + '/test/exception', () => {
		throw new Error('Error when creating test');
	}),

	rest.post(testOARBaseURL + '/query', async (req) => {
		req.json().then((data) => {
			base64Encode(data);
		});
	}),

	rest.post(testOARBaseURL + '/query/bad_response', (req, res, ctx) => {
		return res(ctx.status(400), ctx.json({ error: 'an error occured when creating a test' }));
	}),

	rest.post(testOARBaseURL + '/query/exception', () => {
		throw new Error('Error when querying');
	}),

	rest.get(testOARBaseURL + '/tests', (req, res, ctx) => {
		const testQueryResult = { count: 1, tests: [selectRandomItem(fakeTests)] };
		return res(ctx.status(200), ctx.json(testQueryResult));
	}),

	rest.get(testOARBaseURL + '/tests/bad_response', (req, res, ctx) => {
		const testQueryResult = { error: 'an error has occurred when retrieving tests' };
		return res(ctx.status(400), ctx.json(testQueryResult));
	}),

	rest.get(testOARBaseURL + '/tests/exception', () => {
		throw new Error('Error occurred');
	}),

	rest.patch(testOARBaseURL + '/tests', (req, res, ctx) => {
		const patchResponse = 200;
		return res(ctx.status(200), ctx.json(200));
	}),

	rest.patch(testOARBaseURL + '/tests/bad_response', (req, res, ctx) => {
		const patchResponse = { error: 'an error has occurred when updating tests' };
		return res(ctx.status(400), ctx.json(patchResponse));
	}),

	rest.patch(testOARBaseURL + '/tests/exception', () => {
		throw new Error('Error occurred');
	}),

	rest.delete(testOARBaseURL + '/tests', (req, res, ctx) => {
		return res(ctx.status(200), ctx.json(200));
	}),

	rest.delete(testOARBaseURL + '/tests/bad_response', (req, res, ctx) => {
		const deleteResponse = { error: 'an error has occurred when deleting tests' };
		return res(ctx.status(400), ctx.json(deleteResponse));
	}),

	rest.delete(testOARBaseURL + '/tests/exception', () => {
		throw new Error('Error occurred');
	})
];

const server = setupServer(...restHandlers);

// Start server before all tests
beforeAll(() => {
	oarServiceBaseURL.set(testOARBaseURL);
	server.listen({ onUnhandledRequest: 'error' });
});

//  Close server after all tests
afterAll(() => server.close());

// Reset handlers after each test `important for test isolation`
afterEach(() => server.resetHandlers());
