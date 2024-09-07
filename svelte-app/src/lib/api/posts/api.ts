import { api, cleanUrlParams } from '$lib/api';
import type {
  Post,
  CreatePostRequest,
  UpdatePostRequest,
  GetPostParams,
  ListPostsParams,
  ListPostsResponse
} from './dtos';

const API_PATH = '/posts/v1';

export const postsApi = {
  /**
   * Create a new post
   */
  createPost: async (postData: CreatePostRequest): Promise<Post> => {
    const response = await api.post(API_PATH, postData);
    return response.data;
  },

  /**
   * Get a post by ID
   */
  getPost: async (id: number, params: GetPostParams = {}): Promise<Post> => {
    const queryString = cleanUrlParams(params);
    const response = await api.get(`${API_PATH}/${id}?${queryString}`);
    return response.data;
  },

  /**
   * Update an existing post
   */
  updatePost: async (id: number, postData: UpdatePostRequest): Promise<Post> => {
    const response = await api.put(`${API_PATH}/${id}`, postData);
    return response.data;
  },

  /**
   * Delete a post
   */
  deletePost: async (id: number): Promise<void> => {
    await api.delete(`${API_PATH}/${id}`);
  },

  /**
   * List posts
   */
  listPosts: async (params: ListPostsParams = {}): Promise<ListPostsResponse> => {
    const queryString = cleanUrlParams(params);
    const response = await api.get(`${API_PATH}?${queryString}`);
    return response;
  },
};
