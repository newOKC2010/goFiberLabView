'use client';

import { ReactNode } from 'react';
import TableLoading from './component/TableLoading';
import TableEmpty from './component/TableEmpty';
import TableDesktop from './component/TableDesktop';
import TableMobile from './component/TableMobile';

interface Column<T> {
  key: string;
  label: string;
  render?: (item: T) => ReactNode;
}

interface TableProps<T> {
  columns: Column<T>[];
  data: T[];
  loading?: boolean;
  onEdit?: (item: T) => void;
  onDelete?: (item: T) => void;
  onToggleStatus?: (item: T) => void;
  getItemId: (item: T) => string | number;
  getItemStatus?: (item: T) => boolean;
  statusButtonText?: string;
}

export default function Table<T>({
  columns,
  data,
  loading = false,
  onEdit,
  onDelete,
  onToggleStatus,
  getItemId,
  getItemStatus,
  statusButtonText
}: TableProps<T>) {
  if (loading) return <TableLoading />;
  if (data.length === 0) return <TableEmpty />;

  return (
    <>
      <TableDesktop
        columns={columns}
        data={data}
        onEdit={onEdit}
        onDelete={onDelete}
        onToggleStatus={onToggleStatus}
        getItemId={getItemId}
        getItemStatus={getItemStatus}
        statusButtonText={statusButtonText}
      />
      <TableMobile
        columns={columns}
        data={data}
        onEdit={onEdit}
        onDelete={onDelete}
        onToggleStatus={onToggleStatus}
        getItemId={getItemId}
        getItemStatus={getItemStatus}
        statusButtonText={statusButtonText}
      />
    </>
  );
}
