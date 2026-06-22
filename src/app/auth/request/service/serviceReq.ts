import { API_BASE_URL, API_ENDPOINTS } from '@/global/globalApi';

export const loginRequest = async (email: string) => {
  try {
    const url = `${API_BASE_URL}${API_ENDPOINTS.AUTH.LOGIN_REQUEST}`;
    
    const res = await fetch(url, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email })
    });

    if (!res.ok) {
      const errorData = await res.json();
      throw new Error(errorData.message || 'เกิดข้อผิดพลาด');
    }
    
    return res.json();
  } catch (error) {
    if (error instanceof TypeError) {
      throw new Error('ไม่สามารถเชื่อมต่อกับเซิร์ฟเวอร์ได้');
    }
    throw error;
  }
};
