import { writable } from "svelte/store";
import type { Post } from "$lib/api/posts/dtos";
import type { User } from "$lib/api/users/dtos";

type PostModalState = {
  isOpen: boolean;
  variant: "normal" | "reply" | "quote";
  post: Post | null;
  user: User | null;
};

function createPostModalStore() {
  const { subscribe, set } = writable<PostModalState>({
    isOpen: false,
    variant: "normal",
    post: null,
    user: null,
  });

  return {
    subscribe,
    openModal: (
      variant: "normal" | "reply" | "quote",
      post: Post | null = null,
      user: User | null = null,
    ) => {
      set({ isOpen: true, variant, post, user });
    },
    closeModal: () => {
      set({ isOpen: false, variant: "normal", post: null, user: null });
    },
  };
}

export const postModalStore = createPostModalStore();
