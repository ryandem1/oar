import { afterAll, afterEach, beforeAll } from 'vitest';
import { setupServer } from 'msw/node';
import { rest } from 'msw';
import { OAR_SERVICE_BASE_URL } from '$env/static/private';
import { fakeTests, selectRandomItem } from './faker';

export const restHandlers = [
	rest.post(OAR_SERVICE_BASE_URL + '/test', (req, res, ctx) => {
		const test = selectRandomItem(fakeTests);
		return res(ctx.status(200), ctx.json(test.id));
	}),

	rest.post(OAR_SERVICE_BASE_URL + '/test/bad_response', (req, res, ctx) => {
		return res(ctx.status(400), ctx.json({ error: 'an error occured when creating a test' }));
	}),

	rest.post(OAR_SERVICE_BASE_URL + "/test/exception", (req, res, ctx) => {
		throw new Error("Error when creating test")
	})
];

const server = setupServer(...restHandlers);

// Start server before all tests
beforeAll(() => server.listen({ onUnhandledRequest: 'error' }));

//  Close server after all tests
afterAll(() => server.close());

// Reset handlers after each test `important for test isolation`
afterEach(() => server.resetHandlers());
