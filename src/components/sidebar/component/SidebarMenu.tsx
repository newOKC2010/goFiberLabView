import { useState } from 'react';

interface MenuItem {
  name: string;
  icon: string;
  path?: string;
  children?: MenuItem[];
}

interface SidebarMenuProps {
  menus: MenuItem[];
  currentPath: string;
  onMenuClick: (path: string) => void;
}

export default function SidebarMenu({ menus, currentPath, onMenuClick }: SidebarMenuProps) {
  const [openMenus, setOpenMenus] = useState<string[]>([]);

  const toggleMenu = (name: string) => {
    setOpenMenus(prev => 
      prev.includes(name) ? prev.filter(m => m !== name) : [...prev, name]
    );
  };

  const isActive = (path?: string) => path === currentPath;
  const isChildActive = (children?: MenuItem[]) => 
    children?.some(child => child.path === currentPath);

  return (
    <nav className="flex-1 p-4 space-y-2">
      {menus.map((menu) => (
        <div key={menu.path || menu.name}>
          <button
            onClick={() => menu.children ? toggleMenu(menu.name) : menu.path && onMenuClick(menu.path)}
            className={`w-full flex items-center gap-3 px-4 py-3 rounded-xl font-bold transition-all duration-200 cursor-pointer
              ${isActive(menu.path) || isChildActive(menu.children)
                ? 'bg-emerald-50 text-emerald-600 shadow-sm'
                : 'text-gray-600 hover:bg-emerald-50/50 hover:scale-105 hover:shadow-md'
            }`}
          >
            <span className="material-symbols-outlined" style={{ fontVariationSettings: "'wght' 700" }}>
              {menu.icon}
            </span>
            <span className="flex-1 text-left">{menu.name}</span>
            {menu.children && (
              <span className="material-symbols-outlined text-sm">
                {openMenus.includes(menu.name) ? 'expand_less' : 'expand_more'}
              </span>
            )}
          </button>

          {menu.children && (
            <div 
              className={`overflow-hidden transition-all duration-300 ease-in-out ${
                openMenus.includes(menu.name) ? 'max-h-96 opacity-100' : 'max-h-0 opacity-0'
              }`}
            >
              <div className="mt-1 ml-4 space-y-1">
                {menu.children.map((child) => (
                  <button
                    key={child.path}
                    onClick={() => child.path && onMenuClick(child.path)}
                    className={`w-full flex items-center gap-3 px-4 py-2 rounded-lg font-bold text-sm transition-all duration-200 cursor-pointer
                      ${isActive(child.path)
                        ? 'bg-emerald-50 text-emerald-600'
                        : 'text-gray-500 hover:bg-emerald-50/50'
                    }`}
                  >
                    <span className="material-symbols-outlined text-sm" style={{ fontVariationSettings: "'wght' 700" }}>
                      {child.icon}
                    </span>
                    <span>{child.name}</span>
                  </button>
                ))}
              </div>
            </div>
          )}
        </div>
      ))}
    </nav>
  );
}
