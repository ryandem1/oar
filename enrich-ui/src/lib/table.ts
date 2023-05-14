/*
Contains functions to interact with the main test table
*/

import { selectedTestIDs } from "../stores";

/*
Will return the IDs of the tests that are currently selected in the test table.
*/
export const getSelectedTestIDs = (): number[] => {
  let localSelectedTestIDs: number[] = [];
  const unsubscribe = selectedTestIDs.subscribe(ids => {
    localSelectedTestIDs = ids;
  });
  unsubscribe();

  return localSelectedTestIDs;
}
