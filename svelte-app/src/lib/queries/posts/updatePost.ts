import { createMutation, useQueryClient } from '@tanstack/svelte-query';
import { api } from '$lib/api';

interface UpdatePostData {
  id: number;
  content: string;
}

export function useUpdatePost() {
  const queryClient = useQueryClient();

  return createMutation({
    mutationFn: ({ id, content }: UpdatePostData) => api.put(`/posts/v1/${id}`, { content }),
    onSuccess: (data, variables) => {
      queryClient.invalidateQueries({ queryKey: ['post', variables.id] });
      queryClient.invalidateQueries({ queryKey: ['posts'] });
    },
  });
}
