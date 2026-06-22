import { ReactNode } from 'react';
import TableActions from './TableActions';

interface Column<T> {
  key: string;
  label: string;
  render?: (item: T) => ReactNode;
}

interface TableMobileProps<T> {
  columns: Column<T>[];
  data: T[];
  onEdit?: (item: T) => void;
  onDelete?: (item: T) => void;
  onToggleStatus?: (item: T) => void;
  getItemId: (item: T) => string | number;
  getItemStatus?: (item: T) => boolean;
  statusButtonText?: string;
}

export default function TableMobile<T>({
  columns,
  data,
  onEdit,
  onDelete,
  onToggleStatus,
  getItemId,
  getItemStatus,
  statusButtonText
}: TableMobileProps<T>) {
  const hasActions = onEdit || onDelete || onToggleStatus;

  return (
    <div className="lg:hidden space-y-4">
      {data.map((item) => (
        <div key={getItemId(item)} className="bg-white rounded-xl shadow-sm border border-gray-200 p-4">
          {columns.map((col) => (
            <div 
              key={col.key} 
              className={`py-2 border-b border-gray-100 last:border-0 ${
                col.key === 'documents' ? 'block' : 'flex justify-between'
              }`}
            >
              <span className="text-sm font-bold text-gray-500 block mb-2">{col.label}:</span>
              <span className={`text-sm font-bold text-gray-700 ${col.key === 'documents' ? 'block' : ''}`}>
                {col.render ? col.render(item) : String((item as any)[col.key])}
              </span>
            </div>
          ))}
          {hasActions && (
            <div className="flex items-center justify-end gap-2 mt-4 pt-3 border-t border-gray-100">
              <TableActions
                item={item}
                onEdit={onEdit}
                onDelete={onDelete}
                onToggleStatus={onToggleStatus}
                getItemStatus={getItemStatus}
                statusButtonText={statusButtonText}
              />
            </div>
          )}
        </div>
      ))}
    </div>
  );
}
