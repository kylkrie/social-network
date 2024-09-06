import { createQuery } from '@tanstack/svelte-query';
import { type GetUserParams, usersApi, type User } from '$lib/api';
import { QK_USER } from './consts';

export function useGetUser(username: string, params: GetUserParams = {}) {
  return createQuery<User>({
    queryKey: [QK_USER, username, params],
    queryFn: () => usersApi.getUser(username, params)
  });
}
