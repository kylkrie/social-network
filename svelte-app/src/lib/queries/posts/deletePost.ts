import { createMutation, useQueryClient } from '@tanstack/svelte-query';
import { postsApi } from '$lib/api/posts';
import { QK_POST, QK_POSTS } from './consts';

export function useDeletePost() {
  const queryClient = useQueryClient();

  return createMutation<void, Error, string>({
    mutationFn: (id: string) => postsApi.deletePost(id),
    onSuccess: (_, id) => {
      queryClient.removeQueries({ queryKey: [QK_POST, id] });
      queryClient.invalidateQueries({ queryKey: [QK_POSTS] });
    },
  });
}
