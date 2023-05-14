import type { ModalSettings, ToastSettings } from "@skeletonlabs/skeleton";
import { modalStore, toastStore } from "@skeletonlabs/skeleton";
import { refreshTestTable } from "../stores";
import { OARServiceClient } from "$lib/client";
import { getSelectedTestIDs } from "$lib/table";
import { isEnrichUIError, isOARServiceError } from "$lib/models";

const client = new OARServiceClient();

/*
Handler for the "delete" button on the actions bar
*/
export const onDeleteButtonClick = (): void => {
  let localSelectedTestIDs = getSelectedTestIDs()
  let numSelectedTests = localSelectedTestIDs.length;

  if (numSelectedTests === 0) {
    const t: ToastSettings = {
      message: "No selected tests to delete!",
      timeout: 3000,
      background: "bg-surface-400"
    };
    toastStore.trigger(t);
    return
  }

  const deleteConfirmationModal: ModalSettings = {
    type: "confirm",
    title: `Are you sure you would like to delete ${numSelectedTests} tests?`,
    body: "WARNING, this is not reversible",
    response: async (r: boolean) => {
      if (r) {
        await client.deleteTests({ids: localSelectedTestIDs});
        refreshTestTable.set(true);

        const t: ToastSettings = {
          message: "Tests deleted successfully",
          timeout: 3000,
          background: "bg-success-400"
        };
        toastStore.trigger(t);
      }
    },
  };

  modalStore.trigger(deleteConfirmationModal);
}

/*
Handler for the "view" button on the actions bar
*/
export const onViewButtonClick = async (): Promise<void> => {
  let result = await client.getTests({ids: getSelectedTestIDs()})
  if (isOARServiceError(result) || isEnrichUIError(result)) {

  }
}
