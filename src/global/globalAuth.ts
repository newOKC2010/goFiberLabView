import { API_BASE_URL, API_ENDPOINTS } from '@/global/globalApi';

export const USER_ROLES = {
  USER: 'user',
  ADMIN: 'admin',
  SUPER_ADMIN: 'super_admin'
} as const;

export interface UserViewLabInfo {
  user_view_lab_id: number;
  cid: string;
  email: string;
  role: string;
}

function parseTimeToSeconds(time: string): number {
  const value = parseInt(time);
  const unit = time.slice(-1);
  
  if (unit === 'm') return value * 60;
  if (unit === 'h') return value * 3600;
  if (unit === 'd') return value * 86400;
  
  return 86400;
}

export class AuthToken {
  private static key = 'auth_token';
  private static getMaxAge = () => {
    const envTime = process.env.NEXT_PUBLIC_TOKEN_EXPIRE || '1d';
    return parseTimeToSeconds(envTime);
  };
  
  static storeToken = (token: string) => {
    document.cookie = `${this.key}=${token}; path=/; max-age=${this.getMaxAge()}`;
  };
  
  static getToken = () => {
    return document.cookie
      .split('; ')
      .find(row => row.startsWith(`${this.key}=`))
      ?.split('=')[1];
  };
  
  static removeToken = () => {
    document.cookie = `${this.key}=; path=/; max-age=0`;
  };
}

export async function checkAuth(): Promise<{ success: boolean; user?: UserViewLabInfo; message?: string }> {
  const token = AuthToken.getToken();
  if (!token) {
    return { success: false, message: 'ไม่พบ token' };
  }

  try {
    const res = await fetch(`${API_BASE_URL}${API_ENDPOINTS.AUTH.STATUS}`, {
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json'
      }
    });

    if (!res.ok) {
      AuthToken.removeToken();
      const data = await res.json();
      return { success: false, message: data.message || 'ไม่ได้รับอนุญาต' };
    }

    const data = await res.json();
    return { success: true, user: data.user };
  } catch {
    return { success: false, message: 'ไม่สามารถเชื่อมต่อเซิร์ฟเวอร์' };
  }
}