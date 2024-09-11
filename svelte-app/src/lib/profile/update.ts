import { usersApi } from "$lib/api";
import type { UpdateUserRequest } from "$lib/api/users/dtos";

export interface ProfileUpdateData extends UpdateUserRequest {
  profilePicture?: File;
  bannerImage?: File;
}

export async function updateProfile(data: ProfileUpdateData): Promise<void> {
  const { profilePicture, bannerImage, ...userUpdateData } = data;

  // Upload profile picture if provided
  if (profilePicture) {
    await usersApi.uploadProfilePicture(profilePicture);
  }

  // Upload banner image if provided
  if (bannerImage) {
    await usersApi.uploadProfileBanner(bannerImage);
  }

  // Update user data
  if (Object.keys(userUpdateData).length > 0) {
    await usersApi.updateCurrentUser(userUpdateData);
  }
}

export function createObjectURL(file: File | null): string | null {
  return file ? URL.createObjectURL(file) : null;
}

export function revokeObjectURL(url: string | null): void {
  if (url) URL.revokeObjectURL(url);
}
