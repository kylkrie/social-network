import {
  api,
  cleanUrlParams,
  parseListPostsResponse,
  type ParsedListPostsResponse,
} from "$lib/api";
import type {
  User,
  GetUserParams,
  GetCurrentUserParams,
  UpdateUserRequest,
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
    const response = await api.get(`${API_PATH}?${queryString}`);
    return response.data;
  },

  updateCurrentUser: async (userData: UpdateUserRequest): Promise<void> => {
    await api.put(API_PATH, userData);
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

    await api.post(`${API_PATH}/pfp`, formData);
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

    await api.post(`${API_PATH}/pfbanner`, formData);
  },
};
