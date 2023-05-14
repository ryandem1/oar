import type { ToastSettings } from "@skeletonlabs/skeleton";
import { toastStore } from "@skeletonlabs/skeleton";

/*
Will raise a new temporary warning toast with a given message

@param message - Text to display on the toast
*/
export const throwWarningToast = (message: string): void => {
  const t: ToastSettings = {
    message: message,
    timeout: 3000,
    background: "bg-surface-400"
  };
  toastStore.trigger(t);
}

/*
Will raise a new temporary success toast with a given message

@param message - Text to display on the toast
*/
export const throwSuccessToast = (message: string): void => {
  const t: ToastSettings = {
    message: message,
    timeout: 3000,
    background: "bg-success-400"
  };
  toastStore.trigger(t);
}

/*
Will raise a new error toast with a given error message. This toast is not
temporary and will have to get dismissed by the user.

@param message - Text to display on the toast
*/
export const throwFailureToast = (message: string): void => {
  const t: ToastSettings = {
    message: message,
    autohide: false,
    background: "bg-error-400"
  };
  toastStore.trigger(t);
}
