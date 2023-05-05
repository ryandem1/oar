export class OARServiceClient {
	public baseURL: string;
	public testEndpoint: string;

	constructor(baseURL: string) {
		this.baseURL = baseURL;
		if (this.baseURL.endsWith('/')) {
			this.baseURL = this.baseURL.slice(0, -1);
		}

		this.testEndpoint = '/test';
	}

	/*
  addTest will add a new test result via the oar-service. If an error occurs,
  it will be logged to the console and a '-1' testID will be returned
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
}
