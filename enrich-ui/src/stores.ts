import type { Writable } from "svelte/store";
import { writable } from "svelte/store";
import type { TestQuery } from "$lib/models";

export const selectedTestIDs: Writable<number[]> = writable([]);
export const refreshTestTable: Writable<boolean> = writable(true);
export const testTableQuery: Writable<TestQuery | null> = writable(null);
