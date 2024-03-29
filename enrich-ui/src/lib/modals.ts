import type { ModalComponent, ModalSettings } from '@skeletonlabs/skeleton';
import { modalStore } from '@skeletonlabs/skeleton';
import TestsDetailModal from '../components/TestsDetailModal.svelte';
import TestsEnrichModal from '../components/TestsEnrichModal.svelte';
import TestsFilterModal from '../components/TestsFilterModal.svelte';
import SettingsModal from '../components/SettingsModal.svelte';

/*
Display confirmation modal will present the standard confirmation modal to the user.

@param title - Title of the modal
@param body - Body text of the modal
@param fn - Function that will handle the user's response to the modal. A true value means the user confirmed
and false means that the user cancelled or clicked out of the modal
*/
export const displayConfirmationModal = (
	title: string,
	body: string,
	fn: (response: boolean) => void
): void => {
	const deleteConfirmationModal: ModalSettings = {
		type: 'confirm',
		title: title,
		body: body,
		response: fn
	};

	modalStore.trigger(deleteConfirmationModal);
};

export const displayViewModal = () => {
	const modalComponent: ModalComponent = {
		// Pass a reference to your custom component
		ref: TestsDetailModal
	};

	const displayModal: ModalSettings = {
		type: 'component',
		title: 'View Tests',
		component: modalComponent
	};

	modalStore.trigger(displayModal);
};

export const displayEnrichModal = () => {
	const modalComponent: ModalComponent = {
		// Pass a reference to your custom component
		ref: TestsEnrichModal
	};

	const displayModal: ModalSettings = {
		type: 'component',
		title: 'Enrich Tests',
		component: modalComponent
	};

	modalStore.trigger(displayModal);
};

export const displayFilterModal = () => {
	const modalComponent: ModalComponent = {
		// Pass a reference to your custom component
		ref: TestsFilterModal
	};

	const displayModal: ModalSettings = {
		type: 'component',
		title: 'Filter Tests',
		component: modalComponent
	};

	modalStore.trigger(displayModal);
};

export const displaySettingsModal = () => {
	const modalComponent: ModalComponent = {
		// Pass a reference to your custom component
		ref: SettingsModal
	};

	const displayModal: ModalSettings = {
		type: 'component',
		title: 'Settings',
		component: modalComponent
	};

	modalStore.trigger(displayModal);
};
