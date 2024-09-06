import { createMutation, useQueryClient } from '@tanstack/svelte-query';
import { api } from '$lib/api';

interface CreatePostData {
  content: string;
  conversation_id?: number;
}

export function useCreatePost() {
  const queryClient = useQueryClient();

  return createMutation({
    mutationFn: (postData: CreatePostData) => api.post('/posts/v1', postData),
    onSuccess: () => {
      // Invalidate and refetch
      queryClient.invalidateQueries({ queryKey: ['posts'] });
    },
  });
}
