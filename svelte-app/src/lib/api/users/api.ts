import {
  api,
  cleanUrlParams,
  parseListPostsResponse,
  publicApi,
  type ParsedListPostsResponse,
} from "$lib/api";
import type {
  User,
  GetUserParams,
  GetCurrentUserParams,
  UpdateUserRequest,
  ListFeedParams,
  ListPostsParams,
} from "./dtos";

const API_PATH = "/users";

export const usersApi = {
  getUser: async (
    username: string,
    params: GetUserParams = {},
  ): Promise<User> => {
    const queryString = cleanUrlParams(params);
    const response = await api.get(`${API_PATH}/${username}?${queryString}`);
    return response.data;
  },

  getCurrentUser: async (params: GetCurrentUserParams = {}): Promise<User> => {
    const queryString = cleanUrlParams(params);
    const response = await api.get(`${API_PATH}/me?${queryString}`);
    return response.data;
  },

  updateCurrentUser: async (userData: UpdateUserRequest): Promise<void> => {
    await api.put(`${API_PATH}/me`, userData);
  },

  listFeedPosts: async (
    params: ListFeedParams = {},
  ): Promise<ParsedListPostsResponse> => {
    const queryString = cleanUrlParams(params);

    const response = await api.get(`${API_PATH}/me/feed?${queryString}`);
    return parseListPostsResponse(response);
  },

  listPublicFeedPosts: async (
    params: ListFeedParams = {},
  ): Promise<ParsedListPostsResponse> => {
    const queryString = cleanUrlParams(params);

    const response = await publicApi.get(`/feed?${queryString}`);
    return parseListPostsResponse(response);
  },

  listPosts: async (
    username: string,
    params: ListPostsParams = {},
  ): Promise<ParsedListPostsResponse> => {
    const queryString = cleanUrlParams(params);
    const response = await api.get(
      `${API_PATH}/${username}/posts?${queryString}`,
    );
    return parseListPostsResponse(response);
  },

  getUserLikes: async (
    username: string,
    params: { limit?: number; cursor?: string } = {},
  ): Promise<ParsedListPostsResponse> => {
    const queryString = cleanUrlParams(params);
    const response = await api.get(
      `${API_PATH}/${username}/likes?${queryString}`,
    );
    return parseListPostsResponse(response);
  },

  getUserBookmarks: async (
    username: string,

    params: { limit?: number; cursor?: string } = {},
  ): Promise<ParsedListPostsResponse> => {
    const queryString = cleanUrlParams(params);

    const response = await api.get(
      `${API_PATH}/${username}/bookmarks?${queryString}`,
    );
    return parseListPostsResponse(response);
  },

  uploadProfilePicture: async (file: File): Promise<void> => {
    if (file.size > 5 * 1024 * 1024) {
      throw new Error("File size exceeds maximum limit of 5MB");
    }

    if (file.type !== "image/jpeg" && file.type !== "image/png") {
      throw new Error("Only JPEG and PNG files are allowed");
    }

    const formData = new FormData();

    formData.append("pfp", file);

    await api.post(`${API_PATH}/me/pfp`, formData);
  },

  uploadProfileBanner: async (file: File): Promise<void> => {
    if (file.size > 5 * 1024 * 1024) {
      throw new Error("File size exceeds maximum limit of 5MB");
    }

    if (file.type !== "image/jpeg" && file.type !== "image/png") {
      throw new Error("Only JPEG and PNG files are allowed");
    }

    const formData = new FormData();

    formData.append("banner", file);

    await api.post(`${API_PATH}/me/pfbanner`, formData);
  },
};
