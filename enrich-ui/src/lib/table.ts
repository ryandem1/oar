/*
Contains functions to interact with the main test table
*/

import { selectedTestIDs, testTableFields, testTableQuery } from "../stores";
import type { TestQuery } from "$lib/models";
import { isEnrichUIError, isOARServiceError } from "$lib/models";
import { throwFailureToast } from "$lib/toasts";
import { OARServiceClient } from "$lib/client";
import { tableMapperValues } from "@skeletonlabs/skeleton";

const client = new OARServiceClient();

type TestTable = string[][]; // This is the format that needs to be displayed in the ui

/*
Will return the IDs of the tests that are currently selected in the test table.
*/
export const getSelectedTestIDs = (): number[] => {
	let localSelectedTestIDs: number[] = [];
	const unsubscribe = selectedTestIDs.subscribe((ids) => {
		localSelectedTestIDs = ids;
	});
	unsubscribe();

	return localSelectedTestIDs;
};

/*
Will return the test query that is configured for the main test table
*/
export const getTestQuery = (): TestQuery | null => {
	let localQuery: TestQuery | null = null;
	const unsubscribe = testTableQuery.subscribe((query) => {
		localQuery = query;
	});
	unsubscribe();

	return localQuery;
};

export const getTestTableFields = (): string[] => {
	let localFields: string[] = [];
	const unsubscribe = testTableFields.subscribe((fields) => {
		localFields = fields;
	});
	unsubscribe();

	return localFields;
};

/*
Will retrieve tests from the oar-service and format them like a test table

@param testQuery - Query to send to the API
@param headers - Headers to display on the table
*/
export const getTestTable = async (
	testQuery: TestQuery | null = null,
	headers: string[]
): Promise<TestTable> => {
	const response = await client.getTests(testQuery, 0, 250);
	if (isEnrichUIError(response) || isOARServiceError(response)) {
		throwFailureToast(response.error);
		return [];
	}

	return tableMapperValues(
		response.tests,
		headers.map((f) => f.toLowerCase())
	);
};
