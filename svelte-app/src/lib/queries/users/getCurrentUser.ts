import { createQuery } from '@tanstack/svelte-query';
import { type User, type GetCurrentUserParams, usersApi } from '$lib/api';
import { QK_CURRENT_USER } from './consts';

export function useGetCurrentUser(params: GetCurrentUserParams = {}) {
  return createQuery<User>({
    queryKey: [QK_CURRENT_USER, params],
    queryFn: () => usersApi.getCurrentUser(params)
  });
}
