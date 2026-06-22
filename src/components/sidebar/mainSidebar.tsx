'use client';

import { useState } from 'react';
import { usePathname, useRouter } from 'next/navigation';
import { UserStaffInfo, USER_ROLES } from '@/global/globalAuth';
import { handleLogout, handleMenuClick } from '@/components/sidebar/handler/sidebarHandler';
import MobileToggle from '@/components/sidebar/component/MobileToggle';
import SidebarBackdrop from '@/components/sidebar/component/SidebarBackdrop';
import SidebarHeader from '@/components/sidebar/component/SidebarHeader';
import SidebarMenu from '@/components/sidebar/component/SidebarMenu';
import SidebarFooter from '@/components/sidebar/component/SidebarFooter';

interface SidebarProps {
  user: UserStaffInfo | null;
}

const ALL_MENUS = [
  { 
    name: 'ข้อมูลน้ำ', 
    icon: 'water_drop', 
    children: [
      { name: 'Dashboard', path: '/ecobase/water/dashboard', icon: 'Bar_Chart_4_Bars' },
      { name: 'จัดการข้อมูล', path: '/ecobase/water/manage', icon: 'edit_note' }
    ]
  },
  { 
    name: 'ข้อมูลขยะ', 
    icon: 'delete', 
    children: [
      { name: 'Dashboard', path: '/ecobase/waste/dashboard', icon: 'Bar_Chart_4_Bars' },
      { name: 'จัดการข้อมูล', path: '/ecobase/waste/manage', icon: 'edit_note' }
    ]
  }
];

export default function Sidebar({ user }: SidebarProps) {
  const pathname = usePathname();
  const router = useRouter();
  const [isOpen, setIsOpen] = useState(false);

  const closeSidebar = () => setIsOpen(false);
  const toggleSidebar = () => setIsOpen(true);

  const filteredMenus = ALL_MENUS;

  return (
    <>
      <MobileToggle isOpen={isOpen} onToggle={toggleSidebar} />
      <SidebarBackdrop isOpen={isOpen} onClose={closeSidebar} />

      <aside className={`
        fixed lg:static inset-y-0 left-0 z-40
        w-64 bg-white border-r border-gray-200 flex flex-col h-screen
        transform transition-transform duration-300 ease-in-out
        ${isOpen ? 'translate-x-0' : '-translate-x-full lg:translate-x-0'}
      `}>
        <SidebarHeader user={user} onClose={closeSidebar} />
        <SidebarMenu 
          menus={filteredMenus} 
          currentPath={pathname} 
          onMenuClick={(path) => handleMenuClick(path, router, closeSidebar)} 
        />
        <SidebarFooter onLogout={() => handleLogout(router)} />
      </aside>
    </>
  );
}
