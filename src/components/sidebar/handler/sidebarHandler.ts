import { AuthToken } from '@/global/globalAuth';
import { AppRouterInstance } from 'next/dist/shared/lib/app-router-context.shared-runtime';

export const handleLogout = (router: AppRouterInstance) => {
  AuthToken.removeToken();
  // Force reload เพื่อให้แน่ใจว่าหน้าไม่ถูกแคช
  window.location.replace('/auth');
};

export const handleMenuClick = (
  path: string, 
  router: AppRouterInstance, 
  closeSidebar: () => void
) => {
  router.push(path);
  closeSidebar();
};
