
import { api, cleanUrlParams } from '$lib/api';
import type { ListPostsResponse } from '../posts';
import type { ListFeedParams } from './dtos';

const API_PATH = '/feed';

export const feedsApi = {
  listFeedPosts: async (params: ListFeedParams = {}): Promise<ListPostsResponse> => {
    const queryString = cleanUrlParams(params);
    const response = await api.get(`${API_PATH}?${queryString}`);
    return response;
  },
};
