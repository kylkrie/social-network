import { createQuery } from '@tanstack/svelte-query';
import { api } from '$lib/api';
import type { GetUserResponse, GetUserParams } from '$lib/api';

export function useGetUser(username: string, params: GetUserParams = {}) {
  return createQuery<GetUserResponse>({
    queryKey: ['user', username, params],
    queryFn: () => api.get(`/api/users/v1/${username}?${new URLSearchParams(params as any).toString()}`),
  });
}
