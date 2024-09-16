import {
  type ListRepliesParams,
  type ParsedListPostsResponse,
  postsApi,
  usersApi,
} from "$lib/api";
import { createInfiniteQuery } from "@tanstack/svelte-query";
import type { Readable } from "svelte/motion";
import { derived } from "svelte/store";
import { QK_POSTS } from "./consts";
import { type ListPostsQueryResult, processListPostsQuery } from "./listPosts";

export function useListReplies(
  postId: string,
  params: ListRepliesParams = {},
): Readable<ListPostsQueryResult> {
  const query = createInfiniteQuery<ParsedListPostsResponse, Error>({
    queryKey: [QK_POSTS, "replies", postId],
    queryFn: ({ pageParam = undefined }) => {
      return postsApi.listReplies(postId, {
        ...params,
        cursor: pageParam as string | undefined,
      });
    },
    getNextPageParam: (lastPage) => lastPage.nextCursor || undefined,
    initialPageParam: undefined,
  });

  return derived(query, processListPostsQuery);
}
