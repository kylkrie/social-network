import { createQuery } from '@tanstack/svelte-query';
import { api } from '$lib/api';
import type { GetCurrentUserResponse, GetCurrentUserParams } from '$lib/api';

export function useGetCurrentUser(params: GetCurrentUserParams = {}) {
  return createQuery<GetCurrentUserResponse>({
    queryKey: ['currentUser', params],
    queryFn: () => api.get(`/users/v1?${new URLSearchParams(params as any).toString()}`),
  });
}
