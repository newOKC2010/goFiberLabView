import { fetchUsers } from '@/app/labView/users/service/serviceGetUsers';
import { updateUserStatus } from '@/app/labView/users/service/serviceUpdateStatus';
import { GetUsersResponse } from '@/app/labView/users/utils/types';

export async function handleGetUsers(
  page: number = 1,
  pageSize: number = 10,
  search: string = ''
): Promise<{ success: boolean; data?: GetUsersResponse; message?: string }> {
  return await fetchUsers(page, pageSize, search);
}

export async function handleUpdateStatus(
  id: number,
  status: boolean
): Promise<{ success: boolean; message?: string }> {
  return await updateUserStatus(id, status);
}
