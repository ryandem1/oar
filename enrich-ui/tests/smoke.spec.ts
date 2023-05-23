import { expect, test } from '@playwright/test';

test('navbar displays correct version', async ({ page }) => {
	await page.goto('/');
	const element = await page.getByText('OAR Enrich');
	if (element === null) {
		throw new Error('Could not find header with version in it!');
	}

	const text = await element.textContent();
	expect(text).toContain('0.0'); // Hardcoded for now
});
