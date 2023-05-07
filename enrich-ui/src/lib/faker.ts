export const fakeTests = [
	{
		id: 20,
		summary: 'Ensures the /metadata endpoint is functional',
		outcome: 'Failed',
		analysis: 'TruePositive',
		resolution: 'NotNeeded',
		created: '2023-05-05T04:30:03.531334Z',
		modified: '2023-05-05T04:30:03.541662Z',
		doc: {
			browsers: ['chrome', 'firefox', 'edge'],
			owner: 'Sandy Cheeks',
			screenshotURL: 'https://some-s3-bucket-that-doesnt-exist.com/714029473432412',
			'test left merge field': 'different value, different type',
			type: 'UI'
		}
	},
	{
		id: 14,
		summary: 'User service load test',
		outcome: 'Passed',
		analysis: 'TrueNegative',
		resolution: 'TicketCreated',
		created: '2023-05-05T04:30:03.458693Z',
		modified: '2023-05-05T04:30:03.458693Z',
		doc: {
			browsers: ['chrome', 'firefox', 'edge'],
			owner: 'Sandy Cheeks',
			screenshotURL: 'https://some-s3-bucket-that-doesnt-exist.com/714029473432412',
			type: 'UI'
		}
	},
	{
		id: 13,
		summary: 'Ensures a bad input returns a correct error message',
		outcome: 'Failed',
		analysis: 'TruePositive',
		resolution: 'KnownIssue',
		created: '2023-05-05T04:30:03.457544Z',
		modified: '2023-05-05T04:30:03.457544Z',
		doc: {
			'latency (ms)': {
				p50: 254.33,
				p75: 332.45,
				p95: 501.99,
				p99: 676.51
			},
			maxRPS: 300,
			owner: 'Squidward Tentacles',
			runtime: '10m',
			samplePayloads: [
				{
					app_id: '47324033',
					status: 'APPROVED'
				},
				{
					app_id: '9948302',
					status: 'REJECTED'
				}
			],
			service: 'application-service',
			type: 'load'
		}
	},
	{
		id: 12,
		summary: 'Ensures that publishing a valid Kafka event gets consumed correctly downstream',
		outcome: 'Passed',
		analysis: 'FalseNegative',
		resolution: 'TestFixed',
		created: '2023-05-05T04:30:03.456243Z',
		modified: '2023-05-05T04:30:03.456243Z',
		doc: {
			'latency (ms)': {
				p50: 254.33,
				p75: 332.45,
				p95: 501.99,
				p99: 676.51
			},
			maxRPS: 300,
			owner: 'Squidward Tentacles',
			runtime: '10m',
			samplePayloads: [
				{
					app_id: '47324033',
					status: 'APPROVED'
				},
				{
					app_id: '9948302',
					status: 'REJECTED'
				}
			],
			service: 'application-service',
			type: 'load'
		}
	},
	{
		id: 11,
		summary: 'Ensures that publishing a valid Kafka event gets consumed correctly downstream',
		outcome: 'Failed',
		analysis: 'TruePositive',
		resolution: 'TestDisabled',
		created: '2023-05-05T04:30:03.454872Z',
		modified: '2023-05-05T04:30:03.454872Z',
		doc: {
			app: 'user-service',
			owner: 'Patrick Star',
			testPayload: {
				accountStatus: 'lock',
				id: 1,
				username: 'someUser48'
			},
			testResponse: {
				responseBody: null,
				responseCode: 200
			},
			type: 'integration'
		}
	},
	{
		id: 10,
		summary: 'Ensures that publishing a valid Kafka event gets consumed correctly downstream',
		outcome: 'Failed',
		analysis: 'FalsePositive',
		resolution: 'TicketCreated',
		created: '2023-05-05T04:30:03.451906Z',
		modified: '2023-05-05T04:30:03.451906Z',
		doc: {
			app: 'user-service',
			owner: 'Patrick Star',
			testPayload: {
				accountStatus: 'lock',
				id: 1,
				username: 'someUser48'
			},
			testResponse: {
				responseBody: null,
				responseCode: 200
			},
			type: 'integration'
		}
	},
	{
		id: 9,
		summary: 'Test user insert query is functional',
		outcome: 'Failed',
		analysis: 'FalsePositive',
		resolution: 'TestDisabled',
		created: '2023-05-05T04:30:03.365358Z',
		modified: '2023-05-05T04:30:03.365358Z',
		doc: {
			'latency (ms)': {
				p50: 254.33,
				p75: 332.45,
				p95: 501.99,
				p99: 676.51
			},
			maxRPS: 300,
			owner: 'Squidward Tentacles',
			runtime: '10m',
			samplePayloads: [
				{
					app_id: '47324033',
					status: 'APPROVED'
				},
				{
					app_id: '9948302',
					status: 'REJECTED'
				}
			],
			service: 'application-service',
			type: 'load'
		}
	},
	{
		id: 8,
		summary: 'Verifies that bad data does not get forwarded downstream',
		outcome: 'Passed',
		analysis: 'NotAnalyzed',
		resolution: 'TestFixed',
		created: '2023-05-05T04:30:03.338032Z',
		modified: '2023-05-05T04:30:03.338032Z',
		doc: {
			'latency (ms)': {
				p50: 254.33,
				p75: 332.45,
				p95: 501.99,
				p99: 676.51
			},
			maxRPS: 300,
			owner: 'Squidward Tentacles',
			runtime: '10m',
			samplePayloads: [
				{
					app_id: '47324033',
					status: 'APPROVED'
				},
				{
					app_id: '9948302',
					status: 'REJECTED'
				}
			],
			service: 'application-service',
			type: 'load'
		}
	},
	{
		id: 7,
		summary: 'Ensures the /metadata endpoint is functional',
		outcome: 'Failed',
		analysis: 'TruePositive',
		resolution: 'NotNeeded',
		created: '2023-05-05T04:30:03.312466Z',
		modified: '2023-05-05T04:30:03.312466Z',
		doc: {
			'latency (ms)': {
				p50: 254.33,
				p75: 332.45,
				p95: 501.99,
				p99: 676.51
			},
			maxRPS: 300,
			owner: 'Squidward Tentacles',
			runtime: '10m',
			samplePayloads: [
				{
					app_id: '47324033',
					status: 'APPROVED'
				},
				{
					app_id: '9948302',
					status: 'REJECTED'
				}
			],
			service: 'application-service',
			type: 'load'
		}
	},
	{
		id: 6,
		summary: 'Ensures a bad input returns a correct error message',
		outcome: 'Passed',
		analysis: 'NotAnalyzed',
		resolution: 'QuickFix',
		created: '2023-05-05T04:30:03.286965Z',
		modified: '2023-05-05T04:30:03.286965Z',
		doc: {
			'latency (ms)': {
				p50: 254.33,
				p75: 332.45,
				p95: 501.99,
				p99: 676.51
			},
			maxRPS: 300,
			owner: 'Squidward Tentacles',
			runtime: '10m',
			samplePayloads: [
				{
					app_id: '47324033',
					status: 'APPROVED'
				},
				{
					app_id: '9948302',
					status: 'REJECTED'
				}
			],
			service: 'application-service',
			type: 'load'
		}
	},
	{
		id: 5,
		summary: 'Ensures the /metadata endpoint is functional',
		outcome: 'Passed',
		analysis: 'TrueNegative',
		resolution: 'TestDisabled',
		created: '2023-05-05T04:30:03.263148Z',
		modified: '2023-05-05T04:30:03.263148Z',
		doc: {
			'latency (ms)': {
				p50: 254.33,
				p75: 332.45,
				p95: 501.99,
				p99: 676.51
			},
			maxRPS: 300,
			owner: 'Squidward Tentacles',
			runtime: '10m',
			samplePayloads: [
				{
					app_id: '47324033',
					status: 'APPROVED'
				},
				{
					app_id: '9948302',
					status: 'REJECTED'
				}
			],
			service: 'application-service',
			type: 'load'
		}
	},
	{
		id: 4,
		summary: 'Ensures the /metadata endpoint is functional',
		outcome: 'Passed',
		analysis: 'NotAnalyzed',
		resolution: 'TicketCreated',
		created: '2023-05-05T04:30:03.137204Z',
		modified: '2023-05-05T04:30:03.166584Z',
		doc: {
			app: 'user-service',
			owner: 'Patrick Star',
			testPayload: {
				accountStatus: 'lock',
				id: 1,
				username: 'someUser48'
			},
			testResponse: {
				responseBody: null,
				responseCode: 200
			},
			type: 'integration'
		}
	},
	{
		id: 1,
		summary: 'Ensures that publishing a valid Kafka event gets consumed correctly downstream',
		outcome: 'Passed',
		analysis: 'NotAnalyzed',
		resolution: 'NotNeeded',
		created: '2023-05-05T04:30:03.005655Z',
		modified: '2023-05-05T04:30:03.005655Z',
		doc: {
			browsers: ['chrome', 'firefox', 'edge'],
			owner: 'Sandy Cheeks',
			screenshotURL: 'https://some-s3-bucket-that-doesnt-exist.com/714029473432412',
			type: 'UI'
		}
	}
];

export function selectRandomItem<T>(items: T[]): T {
	return items[Math.floor(Math.random() * items.length)];
}
