import { ReactNode } from 'react';
import TableActions from './TableActions';

interface Column<T> {
  key: string;
  label: string;
  render?: (item: T) => ReactNode;
}

interface TableDesktopProps<T> {
  columns: Column<T>[];
  data: T[];
  onEdit?: (item: T) => void;
  onDelete?: (item: T) => void;
  onToggleStatus?: (item: T) => void;
  getItemId: (item: T) => string | number;
  getItemStatus?: (item: T) => boolean;
  statusButtonText?: string;
}

export default function TableDesktop<T>({
  columns,
  data,
  onEdit,
  onDelete,
  onToggleStatus,
  getItemId,
  getItemStatus,
  statusButtonText
}: TableDesktopProps<T>) {
  const hasActions = onEdit || onDelete || onToggleStatus;

  return (
    <div className="hidden lg:block overflow-x-auto">
      <table className="w-full border-collapse">
        <thead>
          <tr className="bg-blue-50 border-b-2 border-blue-200">
            {columns.map((col) => (
              <th key={col.key} className="px-4 py-3 text-center text-sm font-bold text-gray-700">
                {col.label}
              </th>
            ))}
            {hasActions && (
              <th className="px-4 py-3 text-center text-sm font-bold text-gray-700">จัดการ</th>
            )}
          </tr>
        </thead>
        <tbody>
          {data.map((item) => (
            <tr key={getItemId(item)} className="border-b hover:bg-gray-50 transition-colors">
              {columns.map((col) => (
                <td key={col.key} className="px-4 py-3 text-sm font-bold text-gray-600 text-center">
                  {col.render ? col.render(item) : String((item as any)[col.key])}
                </td>
              ))}
              {hasActions && (
                <td className="px-4 py-3 text-center">
                  <TableActions
                    item={item}
                    onEdit={onEdit}
                    onDelete={onDelete}
                    onToggleStatus={onToggleStatus}
                    getItemStatus={getItemStatus}
                    statusButtonText={statusButtonText}
                  />
                </td>
              )}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
