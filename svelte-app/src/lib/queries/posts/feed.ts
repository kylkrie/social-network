import { createInfiniteQuery } from "@tanstack/svelte-query";
import type { ListPostsParams, ParsedListPostsResponse } from "$lib/api/posts";
import { QK_POSTS } from "./consts";
import { derived, type Readable } from "svelte/store";
import type { ListPostsQueryResult, ProcessedListPostsData } from "./listPosts";
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

  return derived(query, ($query): ListPostsQueryResult => {
    const processedData: ProcessedListPostsData = {
      posts: $query.data?.pages.flatMap((page) => page.posts) ?? [],
      includes: {
        posts:
          $query.data?.pages.reduce(
            (acc, page) => ({ ...acc, ...page.includes.posts }),
            {},
          ) ?? {},
        users:
          $query.data?.pages.reduce(
            (acc, page) => ({ ...acc, ...page.includes.users }),
            {},
          ) ?? {},

        userInteractions: {
          likedPosts:
            $query.data?.pages.reduce(
              (acc, page) => ({
                ...acc,
                ...page.includes.userInteractions.likedPosts,
              }),
              {},
            ) ?? {},
          bookmarkedPosts:
            $query.data?.pages.reduce(
              (acc, page) => ({
                ...acc,
                ...page.includes.userInteractions.bookmarkedPosts,
              }),
              {},
            ) ?? {},
        },
      },
    };

    return {
      data: processedData,
      query: $query,
    };
  });
}
