import { api, cleanUrlParams, parseListPostsResponse } from "$lib/api";
import type { ParsedListPostsResponse } from "../posts";
import type { ListFeedParams } from "./dtos";

const API_PATH = "/feed";

export const feedsApi = {
  listFeedPosts: async (
    params: ListFeedParams = {},
  ): Promise<ParsedListPostsResponse> => {
    const queryString = cleanUrlParams(params);
    const response = await api.get(`${API_PATH}?${queryString}`);
    return parseListPostsResponse(response);
  },
};
