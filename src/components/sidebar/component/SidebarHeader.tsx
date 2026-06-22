import { UserStaffInfo } from '@/global/globalAuth';

interface SidebarHeaderProps {
  user: UserStaffInfo | null;
  onClose: () => void;
}

export default function SidebarHeader({ user, onClose }: SidebarHeaderProps) {
  return (
    <div className="p-6 border-b border-gray-200 relative">
      {/* Close Button - Mobile Only */}
      <button
        onClick={onClose}
        className="lg:hidden absolute top-4 right-4 w-8 h-8 rounded-full flex items-center justify-center transition-all duration-200 hover:scale-110 text-gray-400 hover:text-gray-600 bg-gray-100 hover:bg-gray-200 cursor-pointer"
      >
        <span
          className="material-symbols-outlined text-sm"
          style={{
            fontVariationSettings: "'wght' 700",
            transition: 'all 0.3s ease'
          }}
        >
          cancel
        </span>
      </button>

      <div className="flex items-center gap-3">
        <span className="material-symbols-outlined text-emerald-500 text-3xl" style={{ fontVariationSettings: "'wght' 700" }}>
          energy_savings_leaf
        </span>
        <h1 className="text-xl font-bold text-gray-800">E-COBASE</h1>
      </div>
      
      {user && (
        <div className="mt-4 space-y-1">
          <p className="text-sm font-bold text-gray-800 truncate">{user.email}</p>
          <div className="flex items-center gap-2">
            <span className="material-symbols-outlined text-gray-500 text-sm" style={{ fontVariationSettings: "'wght' 700" }}>
              badge
            </span>
            <p className="text-xs text-gray-500 font-bold">{user.role}</p>
          </div>
        </div>
      )}
    </div>
  );
}
