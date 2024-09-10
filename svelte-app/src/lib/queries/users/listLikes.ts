import { createInfiniteQuery } from "@tanstack/svelte-query";
import { usersApi } from "$lib/api/users";
import type { ParsedListPostsResponse } from "$lib/api/posts";
import { QK_USER } from "./consts";
import { derived, type Readable } from "svelte/store";
import {
  processListPostsQuery,
  type ListPostsQueryResult,
} from "../posts/listPosts";

export function useUserLikes(
  username: string,
  params: { limit?: number } = {},
): Readable<ListPostsQueryResult> {
  const query = createInfiniteQuery<ParsedListPostsResponse, Error>({
    queryKey: [QK_USER, username, "likes"],
    queryFn: ({ pageParam = undefined }) => {
      return usersApi.getUserLikes(username, {
        ...params,
        cursor: pageParam as string | undefined,
      });
    },
    getNextPageParam: (lastPage) => lastPage.nextCursor || undefined,
    initialPageParam: undefined,
  });

  return derived(query, processListPostsQuery);
}
