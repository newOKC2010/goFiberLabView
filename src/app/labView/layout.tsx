'use client';

import { useEffect, useState } from 'react';
import { useRouter, usePathname } from 'next/navigation';
import { checkAuth, UserViewLabInfo } from '@/global/globalAuth';
import Sidebar from '@/components/sidebar/mainSidebar';
import PdpaModal from '@/components/pdpa/PdpaModal';

const PDPA_SESSION_KEY = 'pdpa_acknowledged';

export default function MainLayout({ children }: { children: React.ReactNode }) {
  const router = useRouter();
  const pathname = usePathname();
  const [user, setUser] = useState<UserViewLabInfo | null>(null);
  const [showPdpa, setShowPdpa] = useState(false);

  useEffect(() => {
    const verify = async () => {
      const auth = await checkAuth();

      if (!auth.success) {
        setUser(null);
        router.replace('/auth?error=no_token');
        return;
      }

      setUser(auth.user || null);

      if (!sessionStorage.getItem(PDPA_SESSION_KEY)) {
        setShowPdpa(true);
      }
    };

    verify();
  }, [pathname, router]);

  const handleAcknowledge = () => {
    sessionStorage.setItem(PDPA_SESSION_KEY, '1');
    setShowPdpa(false);
  };

  return (
    <div className="flex h-screen overflow-hidden">
      <Sidebar user={user} />
      <main className="flex-1 overflow-auto bg-gray-50 lg:ml-0">
        <div className="lg:hidden h-16" />
        {children}
      </main>
      {showPdpa && <PdpaModal onAcknowledge={handleAcknowledge} />}
    </div>
  );
}