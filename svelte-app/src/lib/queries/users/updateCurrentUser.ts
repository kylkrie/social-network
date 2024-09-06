import { createMutation, useQueryClient } from '@tanstack/svelte-query';
import { api } from '$lib/api';
import type { GetCurrentUserResponse, UpdateUserRequest } from '$lib/api';

export function useUpdateCurrentUser() {
  const queryClient = useQueryClient();

  return createMutation<GetCurrentUserResponse, Error, UpdateUserRequest>({
    mutationFn: (userData: UpdateUserRequest) => api.put('/users/v1', userData),
    onSuccess: (data) => {
      // Update the cache with the new user data
      queryClient.setQueryData(['currentUser'], data);
    },
  });
}
