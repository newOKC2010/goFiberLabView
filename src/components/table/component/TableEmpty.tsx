export default function TableEmpty() {
  return (
    <div className="flex items-center justify-center py-16">
      <div className="text-center space-y-3">
        <span className="material-symbols-outlined text-6xl text-gray-300">inbox</span>
        <p className="text-lg font-bold text-gray-500">ไม่พบข้อมูล</p>
        <p className="text-sm font-bold text-gray-400">ยังไม่มีรายการในระบบ</p>
      </div>
    </div>
  );
}
