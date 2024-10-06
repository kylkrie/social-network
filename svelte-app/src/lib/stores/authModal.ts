import { writable } from "svelte/store";

type PostModalState = {
  isOpen: boolean;
  action: string;
};

function createAuthModal() {
  const { subscribe, set } = writable<PostModalState>({
    isOpen: false,
    action: null,
  });

  return {
    subscribe,
    openModal: (action: string) => {
      set({ isOpen: true, action });
    },
    closeModal: () => {
      set({ isOpen: false, action: null });
    },
  };
}

export const authModalStore = createAuthModal();
