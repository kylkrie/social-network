import { createMutation, useQueryClient } from '@tanstack/svelte-query';
import { usersApi } from '$lib/api';
import type { UpdateUserRequest } from '$lib/api';
import { QK_CURRENT_USER } from './consts';

export function useUpdateCurrentUser() {
  const queryClient = useQueryClient();

  return createMutation({
    mutationFn: (userData: UpdateUserRequest) => usersApi.updateCurrentUser(userData),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: [QK_CURRENT_USER] });
    },
  });
}
