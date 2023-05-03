import type { PlaywrightTestConfig } from '@playwright/test';
import { devices } from "@playwright/test";

const config: PlaywrightTestConfig = {
	// Look for test files in the "tests" directory, relative to this configuration file.
	testDir: 'tests',

	// Run all tests in parallel.
	fullyParallel: true,

	// Fail the build on CI if you accidentally left test.only in the source code.
	forbidOnly: !!process.env.CI,

	// Retry on CI only.
	retries: process.env.CI ? 2 : 0,

	// Opt out of parallel tests on CI.
	workers: process.env.CI ? 1 : undefined,

	// Reporter to use
	reporter: process.env.CI ? 'html' : 'line',

	use: {
		// Base URL to use in actions like `await page.goto('/')`.
		baseURL: 'http://localhost:4173/',

		// Collect trace when retrying the failed test.
		trace: 'on-first-retry',
	},
	// Configure projects for major browsers.
	projects: [
		{
			name: 'chromium',
			use: { ...devices['Desktop Chrome'] },
		},
	],
	// Run your local dev server before starting the tests.
	webServer: {
		command: 'npm run preview',
		url: 'http://localhost:4173/',
		reuseExistingServer: !process.env.CI,
	}
};

export default config;
