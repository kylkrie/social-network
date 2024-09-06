import { createMutation, useQueryClient } from '@tanstack/svelte-query';
import { api } from '$lib/api';

export function useDeletePost() {
  const queryClient = useQueryClient();

  return createMutation({
    mutationFn: (id: number) => api.delete(`/posts/v1/${id}`),
    onSuccess: (_, id) => {
      queryClient.invalidateQueries({ queryKey: ['post', id] });
      queryClient.invalidateQueries({ queryKey: ['posts'] });
    },
  });
}
