import { createMutation, useQueryClient } from '@tanstack/svelte-query';
import { postsApi } from '$lib/api/posts';
import type { UpdatePostRequest } from '$lib/api/posts';
import { QK_POST, QK_POSTS } from './consts';

interface UpdatePostMutationVariables {
  id: string;
  postData: UpdatePostRequest;
}

export function useUpdatePost() {
  const queryClient = useQueryClient();

  return createMutation<void, Error, UpdatePostMutationVariables>({
    mutationFn: ({ id, postData }: UpdatePostMutationVariables) => postsApi.updatePost(id, postData),
    onSuccess: (data, variables) => {
      queryClient.setQueryData([QK_POST, variables.id], data);
      queryClient.invalidateQueries({ queryKey: [QK_POSTS] });
    },
  });
}
