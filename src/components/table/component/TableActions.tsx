interface TableActionsProps<T> {
  item: T;
  onEdit?: (item: T) => void;
  onDelete?: (item: T) => void;
  onToggleStatus?: (item: T) => void;
  getItemStatus?: (item: T) => boolean;
  statusButtonText?: string;
}

export default function TableActions<T>({
  item,
  onEdit,
  onDelete,
  onToggleStatus,
  getItemStatus,
  statusButtonText = 'อนุมัติรายการ'
}: TableActionsProps<T>) {
  const isApproved = getItemStatus?.(item);

  return (
    <div className="flex items-center justify-center gap-2">
      {onToggleStatus && (
        <button
          onClick={() => onToggleStatus(item)}
          className={`px-3 py-1 rounded-lg text-xs font-bold transition-colors cursor-pointer ${
            isApproved
              ? 'bg-red-50 text-red-600 hover:bg-red-100'
              : 'bg-yellow-100 text-yellow-700 hover:bg-yellow-200'
          }`}
        >
          {isApproved ? 'ระงับ' : statusButtonText}
        </button>
      )}
      {onEdit && (
        <button
          onClick={() => onEdit(item)}
          className="p-2 text-blue-600 hover:bg-blue-50 rounded-lg transition-colors cursor-pointer"
          title="แก้ไข"
        >
          <span className="material-symbols-outlined text-xl">edit</span>
        </button>
      )}
      {onDelete && (
        <button
          onClick={() => onDelete(item)}
          className="p-2 text-red-600 hover:bg-red-50 rounded-lg transition-colors cursor-pointer"
          title="ลบ"
        >
          <span className="material-symbols-outlined text-xl">delete</span>
        </button>
      )}
    </div>
  );
}
