import { API_BASE_URL, API_ENDPOINTS } from '@/global/globalApi';
import { RegisterRequest } from '@/app/auth/component/register/utils/types';

export const registerService = async (req: RegisterRequest) => {
  const res = await fetch(`${API_BASE_URL}${API_ENDPOINTS.AUTH.REGISTER}`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(req),
  });

  const data = await res.json();
  if (!res.ok) throw new Error(data.message || 'เกิดข้อผิดพลาด');
  return data;
};
