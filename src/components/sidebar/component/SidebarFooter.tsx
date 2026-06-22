interface SidebarFooterProps {
  onLogout: () => void;
}

export default function SidebarFooter({ onLogout }: SidebarFooterProps) {
  return (
    <div className="p-4 border-t border-gray-200">
      <button
        onClick={onLogout}
        className="w-full flex items-center gap-3 px-4 py-3 rounded-xl font-bold text-gray-700 hover:bg-red-100 hover:text-red-600 transition-all duration-200 cursor-pointer hover:scale-105 hover:shadow-md"
      >
        <span className="material-symbols-outlined" style={{ fontVariationSettings: "'wght' 700" }}>
          logout
        </span>
        <span>ออกจากระบบ</span>
      </button>
    </div>
  );
}
