import { refreshTestTable } from '../stores';
import { OARServiceClient } from '$lib/client';
import { getSelectedTestIDs } from '$lib/table';
import { throwSuccessToast, throwWarningToast } from '$lib/toasts';
import {
	displayConfirmationModal,
	displayEnrichModal,
	displayFilterModal, displaySettingsModal,
	displayViewModal
} from "$lib/modals";

const client = new OARServiceClient();

/*
Handler for the "delete" button on the actions bar
*/
export const onDeleteButtonClick = (): void => {
	const localSelectedTestIDs = getSelectedTestIDs();
	const numSelectedTests = localSelectedTestIDs.length;

	if (numSelectedTests === 0) {
		throwWarningToast('No tests selected to delete!');
		return;
	}

	displayConfirmationModal(
		`Are you sure you would like to delete ${numSelectedTests} tests?`,
		'WARNING, this is not reversible',
		async (r: boolean) => {
			if (r) {
				await client.deleteTests({ ids: localSelectedTestIDs });
				refreshTestTable.set(true);

				throwSuccessToast('Tests deleted successfully!');
			}
		}
	);
};

/*
Handler for the "view" button on the actions bar
*/
export const onViewButtonClick = async (): Promise<void> => {
	const localSelectedTestIDs = getSelectedTestIDs();
	const numSelectedTests = localSelectedTestIDs.length;
	if (numSelectedTests === 0) {
		throwWarningToast('No tests selected to view details of!');
		return;
	}

	displayViewModal();
};

/*
Handler for the "enrich" button on the actions bar
*/
export const onEnrichButtonClick = () => {
	const localSelectedTestIDs = getSelectedTestIDs();
	const numSelectedTests = localSelectedTestIDs.length;
	if (numSelectedTests === 0) {
		throwWarningToast('No tests selected to enrich!');
		return;
	}

	displayEnrichModal();
};

/*
Handler for the "filter" button on the actions bar
*/
export const onFilterButtonClick = () => {
	displayFilterModal();
};


/*
Handler for the "settings" button on the actions bar
*/
export const onSettingsButtonClient = () => {
	displaySettingsModal()
}
