import { createInfiniteQuery } from "@tanstack/svelte-query";
import type { ParsedListPostsResponse } from "$lib/api/posts";
import { QK_POSTS } from "./consts";
import { derived, type Readable } from "svelte/store";
import { processListPostsQuery, type ListPostsQueryResult } from "./listPosts";
import { feedsApi, type ListFeedParams } from "$lib/api";

export function useListFeed(
  params: ListFeedParams = {},
): Readable<ListPostsQueryResult> {
  const query = createInfiniteQuery<ParsedListPostsResponse, Error>({
    queryKey: [QK_POSTS, "feed"],
    queryFn: ({ pageParam = undefined }) => {
      return feedsApi.listFeedPosts({
        ...params,
        cursor: pageParam as string | undefined,
      });
    },
    getNextPageParam: (lastPage) => lastPage.nextCursor || undefined,
    initialPageParam: undefined,
  });

  return derived(query, processListPostsQuery);
}
