import { API_BASE_URL, API_ENDPOINTS } from '@/global/globalApi';

export const verifyOtp = async (email: string, otp: string) => {
  try {
    const res = await fetch(`${API_BASE_URL}${API_ENDPOINTS.AUTH.LOGIN_VERIFY}`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, otp })
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
