import { expect, test } from '@playwright/test';

test('onboard has title and button', async ({ page }) => {
	await page.goto('/onboard');
	await expect(page.locator('h1')).toHaveText('Welcome to Fusion');
	await expect(page.getByRole('button', { name: 'Get Started' })).toBeVisible();
});

test('onboard button navigates to register form', async ({ page }) => {
	await page.goto('/onboard');
	await page.click('text=Get Started');
	await expect(page.locator('h3')).toHaveText('Create your admin account');
	await expect(page.locator('form')).toBeVisible();
});

test('register form has inputs and button', async ({ page }) => {
	await page.goto('/onboard');
	await page.click('text=Get Started');
	await expect(page.getByLabel('Username')).toBeVisible();
	await expect(page.getByLabel('Password', { exact: true })).toBeVisible();
	await expect(page.getByLabel('Confirm Password')).toBeVisible();
	await expect(page.locator('button[type="submit"]')).toBeVisible();
});

test('register form submits and navigates to final page', async ({ page }) => {
	await page.route('*/**/api/user/register', async (route) => await route.fulfill({ status: 200 }));
	await page.route(
		'*/**/api/auth/login',
		async (route) =>
			await route.fulfill({
				status: 200,
				contentType: 'application/json',
				body: JSON.stringify({ token: '123' }),
				headers: {
					'Set-Cookie': 'refreshToken=321;'
				}
			})
	);

	// Navigate to the register form
	await page.goto('/onboard');
	await page.click('text=Get Started');

	// Fill out the form
	await page.getByLabel('Username').fill('John');
	await page.getByLabel('Password', { exact: true }).fill('Test123123');
	await page.getByLabel('Confirm Password').fill('Test123123');
	await page.click('button[type="submit"]');

	// Check that we navigated to the final page
	await expect(page.locator('h1')).toHaveText('Congratulations!');
	await expect(page.getByRole('button', { name: 'Go to Dashboard' })).toBeVisible();
});
