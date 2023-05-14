import { writable } from 'svelte/store';
import type { Writable } from 'svelte/store';

export const selectedTestIDs: Writable<number[]> = writable([]);
export const refreshTestTable: Writable<boolean> = writable(true);
