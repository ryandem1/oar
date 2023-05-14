import { refreshTestTable } from "../stores";
import { OARServiceClient } from "$lib/client";
import { getSelectedTestIDs } from "$lib/table";
import { isEnrichUIError, isOARServiceError } from "$lib/models";
import { throwSuccessToast, throwWarningToast } from "$lib/toasts";
import { displayConfirmationModal } from "$lib/modals";

const client = new OARServiceClient();

/*
Handler for the "delete" button on the actions bar
*/
export const onDeleteButtonClick = (): void => {
  let localSelectedTestIDs = getSelectedTestIDs()
  let numSelectedTests = localSelectedTestIDs.length;

  if (numSelectedTests === 0) {
    throwWarningToast("No tests selected to delete!")
    return
  }

  displayConfirmationModal(
    `Are you sure you would like to delete ${numSelectedTests} tests?`,
    "WARNING, this is not reversible",
    async (r: boolean) => {
      if (r) {
        await client.deleteTests({ids: localSelectedTestIDs});
        refreshTestTable.set(true);

        throwSuccessToast("Tests deleted successfully!")
      }
    }
  )
}

/*
Handler for the "view" button on the actions bar
*/
export const onViewButtonClick = async (): Promise<void> => {
  let result = await client.getTests({ids: getSelectedTestIDs()})
  if (isOARServiceError(result) || isEnrichUIError(result)) {

  }
}
