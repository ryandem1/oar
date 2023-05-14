import type { ModalComponent, ModalSettings } from "@skeletonlabs/skeleton";
import { modalStore } from "@skeletonlabs/skeleton";
import TestsDetailModal from "../components/TestsDetailModal.svelte";

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
		ref: TestsDetailModal,
	};

	const displayModal: ModalSettings = {
		type: 'component',
		title: 'View Tests',
		component: modalComponent,
	};

	modalStore.trigger(displayModal);
}
