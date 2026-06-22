import { fetchUsers } from '@/app/labView/users/service/serviceGetUsers';
import { updateUserStatus } from '@/app/labView/users/service/serviceUpdateStatus';
import { GetUsersResponse } from '@/app/labView/users/utils/types';

export async function handleGetUsers(): Promise<{ success: boolean; data?: GetUsersResponse; message?: string }> {
  return await fetchUsers();
}

export async function handleUpdateStatus(
  id: number,
  status: boolean
): Promise<{ success: boolean; message?: string }> {
  return await updateUserStatus(id, status);
}
